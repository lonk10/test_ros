package aburos

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"

	"github.com/abu-lang/goabu/ecarule"
	"github.com/abu-lang/goabu/memory"
	"github.com/abu-lang/goabu/parser"
	"github.com/abu-lang/goabu/stringset"
	"github.com/hyperjumptech/grule-rule-engine/ast"
)

type RosExecuter struct {
	memory     memory.ResourceController
	lockMemory sync.RWMutex

	workingMemory *ast.WorkingMemory

	agent           *RosAgent
	dataContext     ast.IDataContext
	lexerParserPool sync.Pool
	types           map[string]string

	ruleLibrary map[string]ecarule.RuleDict
	lockRules   sync.Mutex

	pool     []Update
	lockPool sync.Mutex

	localNamespace  string
	remoteNamespace string

	//logger *zap.Logger

	//routines
	updateReceiver chan<- preparedUpdates
}

func NewRosExecuter(mem memory.ResourceController, rules []string, agt *RosAgent, localname string, remotename string) (*RosExecuter, error) {
	exec := &RosExecuter{
		memory:          mem.Copy(),
		ruleLibrary:     make(map[string]ecarule.RuleDict),
		agent:           agt,
		localNamespace:  localname,
		remoteNamespace: remotename,
	}
	// resources checks
	if exec.memory.HasDuplicates() {

		return nil, errors.New("multiple resources have the same name")
	}
	err := validNames(exec.memory.ResourceNames())
	if err != nil {
		return nil, err
	}
	// set types
	exec.types = exec.memory.Types()
	exec.dataContext, exec.workingMemory, err = newEmptyGruleStructures(map[string]memory.Resources{"this": exec.memory.GetResources()})
	if err != nil {
		return nil, err
	}
	// set lexer/parser pool
	exec.lexerParserPool = sync.Pool{
		New: func() interface{} {
			return parser.New(exec.types, exec.workingMemory, &exec.lockMemory)
		},
	}
	err = exec.StartAgent()
	if err != nil {
		return nil, err
	}
	// add rules
	err = exec.AddRules(rules...)
	if err != nil {
		return nil, err
	}

	//exec.logger, err = zapCfg.Build()
	exec.updateReceiver = exec.startUpdateReceiver()

	return exec, nil
}

func (m *RosExecuter) StartAgent() error {
	err := m.agent.Start(m.localNamespace, m.remoteNamespace)
	if err != nil {
		return fmt.Errorf("couldn't start agent: %v", err)
	}
	go m.receiveExternalActions()
	return nil
}

func (m *RosExecuter) AddRules(rules ...string) error {
	//fmt.Printf("Adding rules")
	if len(rules) == 0 {
		return nil
	}
	m.lockRules.Lock()
	defer m.lockRules.Unlock()
	parser := m.lexerParserPool.Get().(ecarule.Parser)
	defer m.lexerParserPool.Put(parser)
	parsedRules, _ := parser.Parse(rules...)
	/*err := m.createPubSubs()
	if err != nil {
		return err
	}*/
	if len(parsedRules) == 1 {
		return m.addRuleAux(parsedRules[0])
	}
	return addList(parsedRules, m.addRuleAux)
}

func (m *RosExecuter) TakeState() (memory.Resources, []Update) {
	//m.coordinator.requestWrite(false)
	fmt.Printf("%s lock in take state...\n", m.localNamespace)
	m.lockMemory.RLock()
	memCopy := m.memory.Copy().GetResources()
	m.lockMemory.RUnlock()
	lock := make(chan bool)
	fmt.Printf("%s preparing update in take state...\n", m.localNamespace)
	m.updateReceiver <- preparedUpdates{confirm: lock}
	fmt.Printf("%s sent update in take state...\n", m.localNamespace)
	lock <- false // no updates are added
	poolCopy := make([]Update, 0, len(m.pool))
	for _, update := range m.pool {
		var updateCopy Update = make([]Assignment, len(update))
		copy(updateCopy, update)
		poolCopy = append(poolCopy, updateCopy)
	}
	fmt.Printf("%s lock something?...\n", m.localNamespace)
	<-lock
	//m.coordinator.closeWrite()
	fmt.Printf("%s take state finish...\n", m.localNamespace)
	return memCopy, poolCopy
}

func (m *RosExecuter) addRuleAux(rule ecarule.Rule) error {
	if m.hasRuleAux(rule.Name) {
		return fmt.Errorf("there is already a rule named %s", rule.Name)
	}
	for _, evt := range rule.Events {
		if m.ruleLibrary[evt] == nil {
			m.ruleLibrary[evt] = ecarule.MakeRuleDict()
		}
		m.ruleLibrary[evt].Insert(&rule)
	}
	//m.logger.Debug("Introduced new rule", zap.String("act", "add_rule"), zap.String("obj", rule.Name))
	return nil
}

func (m *RosExecuter) hasRuleAux(name string) bool {
	for _, d := range m.ruleLibrary {
		if d.Has(name) {
			return true
		}
	}
	return false
}

func (m *RosExecuter) createPubSubs() error {
	//fmt.Printf("Creating pubsubs")
	topic := m.remoteNamespace + "/global"
	err := m.agent.AddGeneralPublisher(topic)
	if err != nil {
		return err
	}
	m.agent.SetGenChannel(topic)
	err = m.agent.AddGeneralSubscriber(topic)
	if err != nil {
		return err
	}
	return nil
}

/***
* support functions
***/

// newEmptyGruleStructures creates a clean working memory and data context containing
// the [memory.Resources] from resources as struct instances referenced by the map keys.
func newEmptyGruleStructures(resources map[string]memory.Resources) (ast.IDataContext, *ast.WorkingMemory, error) {
	dataContext := ast.NewDataContext()
	kbName := "dummy"
	for name, rs := range resources {
		rs := rs
		kbName += "_" + name
		err := dataContext.Add(name, &(rs))
		if err != nil {
			return dataContext, nil, err
		}
	}
	version := "0.0.0"
	knowledgeBase := &ast.KnowledgeBase{
		Name:          kbName,
		Version:       version,
		RuleEntries:   make(map[string]*ast.RuleEntry),
		WorkingMemory: ast.NewWorkingMemory(kbName, version),
	}
	err := dataContext.Add("DEFUNC",
		makeBuiltinFunctions(
			knowledgeBase,
			knowledgeBase.WorkingMemory,
			dataContext,
		))
	if err != nil {
		return dataContext, nil, err
	}
	knowledgeBase.InitializeContext(dataContext)
	return dataContext, knowledgeBase.WorkingMemory, nil
}

// validNames verifies if the arguments can be used as valid identifiers in GoAbU rules.
func validNames(names []string) error {
	if len(names) == 0 {
		return errors.New("no resource specified")
	}
	results := parser.ValidateIdentifiers(names...)
	for i, n := range names {
		if !results[i] {
			return fmt.Errorf(`invalid resource name: "%s"`, n)
		}
	}
	return nil
}

func (m *RosExecuter) Exec() {
	fmt.Printf("%s Exec func start...\n", m.localNamespace)
	m.lockPool.Lock()
	if len(m.pool) == 0 {
		fmt.Printf("%s Exec pool empty...\n", m.localNamespace)
		m.lockPool.Unlock()
		return
	}
	fmt.Printf("%s Exec getting update...\n", m.localNamespace)
	update, index := m.chooseUpdate()
	m.lockPool.Unlock()
	workingSet := stringset.Make()
	fmt.Printf("%s Exec getting actions...\n", m.localNamespace)
	for _, action := range update {
		workingSet.Insert(action.Resource)
	}
	m.lockPool.Lock()
	fmt.Printf("%s Exec removing update...\n", m.localNamespace)
	m.removeUpdate(index)
	m.lockPool.Unlock()
	m.lockMemory.Lock()
	var modified stringset.Set
	fmt.Printf("%s Exec applying update...\n", m.localNamespace)
	modified = m.applyUpdate(update, false)
	fmt.Printf("%s Exec signaling modification...\n", m.localNamespace)
	m.signalModified(modified)
	fmt.Printf("%s Exec discovery...\n", m.localNamespace)
	m.discovery(modified)
	fmt.Printf("%s Exec end...\n", m.localNamespace)
}

func (m *RosExecuter) Input(actions string) error {

	parsed, err := m.parseActions(actions)
	if err != nil {
		return err
	}
	workingSet := stringset.Make()
	for _, p := range parsed {
		workingSet.Insert(p.Resource)
	}
	//m.coordinator.requestWrite(m.HasOptimisticInput())
	//defer m.coordinator.closeWrite()
	//m.coordinator.fixWorkingSetWrite(workingSet)
	m.lockMemory.RLock()
	update, err := evalActions(parsed, m.dataContext, m.workingMemory)
	if err != nil {
	}
	m.lockMemory.RUnlock()
	//m.logger.Info("Input: "+actions, zap.String("act", "input"), zapUpdate("update", update))
	m.lockMemory.Lock()
	//fmt.Println("Locky locks ok")
	m.discovery(m.applyUpdate(update, true))
	//m.logger.Debug("Processed input", zap.String("act", "input"))
	//m.logger.Sync()
	return nil
}

func (m *RosExecuter) parseActions(actions string) ([]ecarule.Action, error) {
	parser := m.lexerParserPool.Get().(ecarule.Parser)
	defer m.lexerParserPool.Put(parser)
	res, errs := parser.ParseActions(actions)
	if len(errs) > 0 {
		return nil, errs[0]
	}

	return res, nil
}

func (m *RosExecuter) discovery(modified stringset.Set) {
	fmt.Println("Entered discovery")
	updates, wire := m.triggeredActions(modified)
	m.lockMemory.Unlock()
	ok := make(chan bool)
	log.Printf("About to use channel %v", m.updateReceiver)
	m.updateReceiver <- preparedUpdates{updates: updates, confirm: ok}
	ok <- true
	<-ok // ?
	if len(wire.Tasks) > 0 {
		payload, err := marshalWireTasks(wire)
		if err != nil {
		}
		tentatives := 0
		for {
			err = m.agent.ForAll(payload)
			if err == nil {
				break
			}
			//fmt.Printf("Couldn't publish: %v", err)
			tentatives++
			if tentatives%10 == 0 {
			}
		}
	}
}

func (m *RosExecuter) signalModified(modified stringset.Set) {
	for r := range modified {
		m.memory.Modified(r)
	}
}

func (m *RosExecuter) triggeredActions(modified stringset.Set) ([]Update, wireTasks) {
	var newpool []Update
	var wTask wireTasks
	rules := m.activeRules(modified)
	localResources := stringset.Make()
	for _, rule := range rules {
		fmt.Printf("%s triggered rule %s\n", m.localNamespace, rule.Name)
		for _, task := range rule.LocalTasks {
			tActions, err := condEvalActions(task.Condition, task.Actions, m.dataContext, m.workingMemory)
			if err != nil {
			}
			fmt.Printf("%s appending task %v\n", m.localNamespace, tActions)
			newpool = appendNonempty(newpool, tActions)
		}
		for _, task := range rule.RemoteTasks {
			wTask.Tasks = append(wTask.Tasks, task)
			localResources.Add(stringset.Make(task.LocalResources...))
		}
	}
	wTask.Resources = m.memory.Extract(localResources.Slice())
	return newpool, wTask
}

func (m *RosExecuter) activeRules(modified stringset.Set) ecarule.RuleDict {
	res := ecarule.MakeRuleDict()
	m.lockRules.Lock()
	for resource := range modified {
		res.Add(m.ruleLibrary[resource])
	}
	m.lockRules.Unlock()
	return res
}

func (m *RosExecuter) chooseUpdate() (Update, int) {
	// TODO: implement other strategies
	return m.pool[0], 0
}

func (m *RosExecuter) removeUpdate(index int) {
	m.pool = append(m.pool[:index], m.pool[index+1:len(m.pool)]...)
}

func (m *RosExecuter) applyUpdate(update Update, input bool) stringset.Set {
	modified := stringset.Make()
	for _, action := range update {
		variable := action.variable
		variable = m.workingMemory.AddVariable(variable)
		currentVal, err := variable.Evaluate(m.dataContext, m.workingMemory)
		if err != nil {
		}
		fmt.Printf("%s evaluating %v to %v\n", m.localNamespace, currentVal, action.Value)
		if reflect.DeepEqual(currentVal, action.Value) {
			continue
		}
		ltype := currentVal.Type()
		rtype := action.Value.Type()
		if !rtype.AssignableTo(ltype) {
		} else {
			err := variable.Assign(action.Value, m.dataContext, m.workingMemory)
			if err != nil {
			}
			modified.Insert(action.Resource)
			if input {
				m.memory.Modified(action.Resource)
			}
		}
	}
	return modified
}

func addList[T any](objs []T, add func(T) error) error {
	var fstErr error
	failed := ""
	for i, obj := range objs {
		err := add(obj)
		if err != nil {
			failed += strconv.Itoa(i) + ", "
			if fstErr == nil {
				fstErr = err
			}
		}
	}
	if fstErr != nil {
		return fmt.Errorf("could not add elements with indexes %s as %s", failed[:len(failed)-2], fstErr.Error())
	}
	return nil
}
