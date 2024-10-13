package main

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"syscall"
	aburos "test/msgs/aburos_msgs/msg"

	"github.com/abu-lang/goabu/ecarule"
	"github.com/tiiuae/rclgo/pkg/rclgo"
)

type RosAgent struct {
	topics map[string]string
	subs   *rclgo.WaitSet
	pubs   map[string]*rclgo.Publisher

	node           *rclgo.Node
	generalChannel string

	taskPool chan ecarule.RemoteTask
	wirePool chan wireTasks

	context context.Context
}

func NewRosAgent() (*RosAgent, error) {
	res := &RosAgent{
		topics:   make(map[string]string),
		subs:     nil,
		pubs:     make(map[string]*rclgo.Publisher),
		node:     nil,
		wirePool: make(chan wireTasks),
		taskPool: make(chan ecarule.RemoteTask),
	}

	return res, nil
}

func (r *RosAgent) Start(localname string, remotename string) error {
	if r.node != nil {
		return errors.New("RosAgent is already running")
	}

	err := rclgo.Init(nil)
	if err != nil {
		return fmt.Errorf("failed to initialize rclgo: %v", err)
	}
	//defer rclgo.Uninit()

	serviceCtx, err := rclgo.NewContext(0, nil)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to create context: %v", err)
	}

	r.node, err = serviceCtx.NewNode(localname+"_ros_agent", remotename)
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	//defer r.node.Close()
	r.context, _ = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer cancel()
	ws, err := serviceCtx.NewWaitSet()
	if err != nil {
		return fmt.Errorf("failed to create waitset: %v", err)
	}
	r.subs = ws
	topic := remotename + "/global"
	r.generalChannel = remotename
	err = r.AddGeneralPublisher(topic)
	if err != nil {
		return err
	}
	r.SetGenChannel(topic)
	err = r.AddGeneralSubscriber(topic)
	if err != nil {
		return err
	}
	go r.subs.Run(r.context)
	return nil
}

func (r *RosAgent) SetGenChannel(topic string) {
	//fmt.Printf("Setting general channel %s", topic)
	r.generalChannel = topic
}

func (r *RosAgent) Stop() error {
	if r.node == nil {
		return errors.New("RosAgent is already stopped")
	}
	rclnode, err := rclgo.NewNode("ros_abu_agent", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	defer rclnode.Close()
	r.node.Close()
	r.node = nil
	return nil
}

func (r *RosAgent) getTaskPool() chan ecarule.RemoteTask {
	return r.taskPool
}

func (r *RosAgent) getWirePool() chan wireTasks {
	return r.wirePool
}

func (r *RosAgent) ForAll(payload []byte) error {
	/*
		pub := r.pubs[r.generalChannel]
		if pub == nil {
			return fmt.Errorf("failed to retrieve general channel named: %s", r.generalChannel)
		}
		msg := aburos.AbuBytes{
			Data: payload,
		}
		if err := pub.Publish(&msg); err != nil {
			return fmt.Errorf("failed to publish: %v", err)
		}*/
	r.PublishByte(payload, r.generalChannel)
	fmt.Println("published msg")
	return nil
}

/*
func (r *RosAgent) Publish(rule ecarule.Rule) error {
	pub := r.pubs[rule.Name]
	for _, task := range rule.RemoteTasks {
		msg := aburos.AbuRule{
			Condition:       task.Condition,
			Actions:         task.Actions,
			LocalResources:  task.LocalResources,
			RemoteResources: task.RemoteResources,
		}
		pub.Publish(&msg)
	}
	return nil
}*/

/*
func (r *RosAgent) AddRulePublisher(topic string) error {
	pub, _ := aburos.NewAbuRulePublisher(r.node, topic, nil)
	r.pubs[pub.TopicName] = pub.Publisher
	return nil
}

func (r *RosAgent) AddRuleSubscriber(topic string) error {
	sub, _ := aburos.NewAbuRuleSubscription(r.node, topic, nil, func(msg *aburos.AbuRule, info *rclgo.MessageInfo, err error) {
		if err != nil {
			//m.Logger().Errorf("failed to receive message: %v", err)
			return
		}
		//m.Logger().Infof("Received: %#v", msg)
		task := ecarule.RemoteTask{
			Condition:       msg.Condition,
			Actions:         msg.Actions,
			LocalResources:  msg.LocalResources,
			RemoteResources: msg.RemoteResources,
		}
		r.taskPool <- task

	})
	r.subs.AddSubscriptions(sub.Subscription)
	return nil
}*/

func (r *RosAgent) hasSubscriber(topic string) bool {
	for _, sub := range r.subs.Subscriptions {
		if sub.TopicName == topic {
			return true
		}
	}
	return false
}

func (r *RosAgent) AddGeneralPublisher(topic string) error {
	if r.node == nil {
		return fmt.Errorf("couldn't add publisher, ROS node is null")
	}
	pub, err := aburos.NewAbuBytesPublisher(r.node, topic, nil)
	if err != nil {
		return fmt.Errorf("couldn't create general publisher: %v", err)
	}
	r.pubs[pub.TopicName] = pub.Publisher
	return nil
}

func (r *RosAgent) PublishByte(data []byte, topic string) error {
	pub := r.pubs[topic]
	msg := aburos.AbuBytes{
		Data: data,
	}
	err := pub.Publish(&msg)
	if err != nil {
		return fmt.Errorf("Failed to publish byte: %v", err)
	}
	return nil
}

func (r *RosAgent) AddGeneralSubscriber(topic string) error {
	sub, _ := aburos.NewAbuBytesSubscription(r.node, topic, nil, func(msg *aburos.AbuBytes, info *rclgo.MessageInfo, err error) {
		if err != nil {
			r.node.Logger().Errorf("failed to receive message: %v", err)
			return
		}
		//r.node.Logger().Infof("Received: %#v", msg)
		//m.Logger().Infof("Received: %#v", msg)
		task := msg.Data
		r.node.Logger().Infof("Received data")
		wTask, errs := unmarshalWireTasks(task)
		err = errs
		r.wirePool <- wTask

	})
	r.subs.AddSubscriptions(sub.Subscription)
	//go r.subs.Run(r.context)
	return nil
}
