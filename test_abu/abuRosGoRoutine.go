package main

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
		fmt.Printf("action received\n")
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
	context, workMem, _ := newEmptyGruleStructures(map[string]memory.Resources{"this": m.memory.GetResources(), "ext": wTask.Resources})
	p := parser.New(m.types, workMem)
	p.ParseExpressions()
	tasks, _ := p.ParseRemoteTasks(wTask.Resources.Types(), wTask.Tasks...)
	var updates []Update
	for _, task := range tasks {
		m.lockMemory.RLock()
		update, _ := condEvalActions(task.Condition, task.Actions, context, workMem)
		updates = appendNonempty(updates, update)
		m.lockMemory.RUnlock()
	}
	confirm := make(chan bool)
	m.updateReceiver <- preparedUpdates{updates: updates, confirm: confirm}
}

func (m *RosExecuter) startUpdateReceiver() chan<- preparedUpdates {
	res := make(chan preparedUpdates)
	go func(updates <-chan preparedUpdates) {
		var queue []preparedUpdates
		var confirm chan bool = nil
		for {
			select {
			case ok := <-confirm:
				if ok {
					m.lockPool.Lock()
					m.pool = append(m.pool, queue[0].updates...)
					m.lockPool.Unlock()
					//m.logger.Info(fmt.Sprintf("Added %d updates to the pool", len(queue[0].updates)),
					//	zap.String("act", "add_updates"),
					//	zapUpdates("updates", queue[0].updates))
				}
				confirm <- ok
				queue = queue[1:]
			case u := <-updates:
				queue = append(queue, u)
			}
			if len(queue) == 0 {
				confirm = nil
			} else {
				confirm = queue[0].confirm
			}
		}
	}(res)
	return res
}
