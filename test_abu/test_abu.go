package main

import (
	"fmt"
	"time"

	"github.com/abu-lang/goabu/memory"
)

type Resources struct {
	Bool    map[string]bool
	Integer map[string]int64
	Float   map[string]float64
	Text    map[string]string
	Time    map[string]time.Time
	Other   map[string]interface{}
}

func run() error {
	mem := memory.MakeResources()
	mem.Integer["foo"] = 1
	mem.Text["bar"] = "octocat"

	localRule := `rule MyLocalRule on foo bar for "octocat" == bar do foo = foo * 2, bar = "gopher"`

	globalRule := `rule MyGlobalRule on foo for all foo >= ext.foo do foo = ext.foo + foo`

	agent, _ := NewRosAgent()

	executer, err := NewRosExecuter(mem, []string{localRule}, agent, "dev1", "/aburos")
	if err != nil {
		return fmt.Errorf("failed to create executer: %v", err)
	}
	fmt.Println("Created first executer")

	mem2 := memory.MakeResources()
	mem2.Integer["foo"] = 1
	mem2.Float["baz"] = 3.14

	agent2, _ := NewRosAgent()

	executer2, err := NewRosExecuter(mem2, []string{globalRule}, agent2, "dev2", "/aburos")
	if err != nil {
		return fmt.Errorf("failed to create executer: %v", err)
	}
	fmt.Println("Created second executer")

	executer2.Input("foo = 4, baz = 2.72")
	fmt.Println("Input done")

	executer.Exec()
	executer.Exec()
	executer2.Exec()
	executer.Exec()

	state, _ := executer.TakeState()
	fmt.Println("foo =", state.Integer["foo"])
	fmt.Println("bar =", state.Text["bar"])
	state2, _ := executer2.TakeState()
	fmt.Println("foo =", state2.Integer["foo"])
	fmt.Println("baz =", state2.Float["baz"])

	/*
		for ind, sub := range executer.agent.subs.Subscriptions {
			fmt.Printf("Subscription %d: %s \n", ind, sub.TopicName)
			pc, _ := sub.GetPublisherCount()
			sc, _ := agent.pubs[sub.TopicName].GetSubscriptionCount()
			fmt.Printf("Publishers on %s: %d\n", sub.TopicName, pc)
			fmt.Printf("Subscribers on %s: %d\n", sub.TopicName, sc)
			data, info, _ := sub.TakeSerializedMessage()
			msg, _ := rclgo.Deserialize(data, aburos_msgs_msg.AbuBytesTypeSupport)
			m := msg.(*aburos_msgs_msg.AbuBytes)
			w, _ := unmarshalWireTasks(m.Data)
			fmt.Printf("Rule %s\n%s\n", w.Tasks[0].Condition, w.Tasks[0].Actions[0])
			fmt.Printf("Timestamp %s\n", info.ReceivedTimestamp)
			agent.node.Logger()
		}*/

	w := <-agent.wirePool
	fmt.Printf("Rule %s\n%s\n", w.Tasks[0].Condition, w.Tasks[0].Actions[0])

	return nil
}

func main() {
	err := run()
	fmt.Printf("err: %v\n", err)
}
