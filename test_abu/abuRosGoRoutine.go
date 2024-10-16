package aburos

import (
	"fmt"

	"github.com/abu-lang/goabu/memory"
	"github.com/abu-lang/goabu/parser"
)

type preparedUpdates struct {
	confirm chan bool
	updates []Update
}

func (m *RosExecuter) receiveExternalActions() {
	requests := m.agent.wirePool
	for {
		action := <-requests
		fmt.Printf("%s action received\n", m.localNamespace)
		go m.prepareUpdate(action)
	}
}

func (m *RosExecuter) prepareUpdate(wTask wireTasks) {
	for _, task := range wTask.Tasks {
		for _, res := range task.RemoteResources {
			if !m.memory.Has(res) {
				return
			}
		}
	}
	fmt.Printf("%s parsing...\n", m.localNamespace)
	context, workMem, _ := newEmptyGruleStructures(map[string]memory.Resources{"this": m.memory.GetResources(), "ext": wTask.Resources})
	p := parser.New(m.types, workMem)
	p.ParseExpressions()
	tasks, _ := p.ParseRemoteTasks(wTask.Resources.Types(), wTask.Tasks...)
	var updates []Update
	fmt.Printf("%s evaluating...\n", m.localNamespace)
	for _, task := range tasks {
		m.lockMemory.RLock()
		update, _ := condEvalActions(task.Condition, task.Actions, context, workMem)
		fmt.Printf("task.Condition: %v\n", task.Condition)
		fmt.Printf("task.Actions: %v\n", task.Actions)
		updates = appendNonempty(updates, update)
		m.lockMemory.RUnlock()
	}
	fmt.Printf("%s confirming...\n", m.localNamespace)
	confirm := make(chan bool)
	m.updateReceiver <- preparedUpdates{updates: updates, confirm: confirm}
	confirm <- true
	<-confirm
	fmt.Printf("%s confirm done\n", m.localNamespace)
}

func (m *RosExecuter) startUpdateReceiver() chan<- preparedUpdates {
	res := make(chan preparedUpdates)
	go func(updates <-chan preparedUpdates) {
		var queue []preparedUpdates
		var confirm chan bool = nil
		for {
			fmt.Printf("%s update routine\n", m.localNamespace)
			select {
			case ok := <-confirm:
				fmt.Printf("%s case1\n", m.localNamespace)
				if ok {
					fmt.Printf("%s case11\n", m.localNamespace)
					m.lockPool.Lock()
					m.pool = append(m.pool, queue[0].updates...)
					m.lockPool.Unlock()
					fmt.Printf("%s case111\n", m.localNamespace)
					//m.logger.Info(fmt.Sprintf("Added %d updates to the pool", len(queue[0].updates)),
					//	zap.String("act", "add_updates"),
					//	zapUpdates("updates", queue[0].updates))
				}
				confirm <- ok
				queue = queue[1:]
				fmt.Printf("%s case12\n", m.localNamespace)
			case u := <-updates:
				fmt.Printf("%s case append\n", m.localNamespace)
				queue = append(queue, u)
			}
			if len(queue) == 0 {
				fmt.Printf("%s case2\n", m.localNamespace)

				confirm = nil
			} else {
				fmt.Printf("%s case3\n", m.localNamespace)

				confirm = queue[0].confirm
			}
		}
	}(res)
	return res
}

func (m *RosExecuter) cmdRoutine() {
}
