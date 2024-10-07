package vbase

import (
	"context"
	"errors"
	"fmt"
	"os/signal"
	"syscall"
	geographic_msgs_msg "test/msgs/geographic_msgs/msg"
	sensor_msgs_msg "test/msgs/sensor_msgs/msg"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

type VehicleBase struct {
	node           *rclgo.Node
	Subs           *rclgo.WaitSet
	Pubs           map[string]*rclgo.Publisher
	generalChannel string
	context        context.Context
	ServiceClients map[string]*rclgo.Client
	Position       geographic_msgs_msg.GeoPoseStamped
}

func NewVehicleBase() (*VehicleBase, error) {
	res := &VehicleBase{
		Pubs: make(map[string]*rclgo.Publisher),
	}

	return res, nil
}

func (v *VehicleBase) GetNode() *rclgo.Node {
	return v.node
}

func (v *VehicleBase) Start(name string) error {
	if v.node != nil {
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

	v.node, err = serviceCtx.NewNode("ros_agent", name)
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	//defer r.node.Close()
	v.context, _ = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer cancel()
	ws, err := serviceCtx.NewWaitSet()
	if err != nil {
		return fmt.Errorf("failed to create waitset: %v", err)
	}
	//defer ws.Close()
	sub, err := v.addGPSSubscriber()
	if err != nil {
		return fmt.Errorf("failed to create GPS sub: %v", err)
	}
	ws.AddSubscriptions(sub)
	v.Subs = ws
	go v.Subs.Run(v.context)
	//go serviceCtx.Spin(v.context)
	return nil
}

func (v *VehicleBase) GetPublisher(topic string) (*rclgo.Publisher, error) {
	res := v.Pubs[topic]
	if res == nil {
		return nil, fmt.Errorf("Publisher for topic %s not found", topic)
	}
	return res, nil
}

func (v *VehicleBase) AddSubscriber(sub *rclgo.Subscription) {
	v.Subs.AddSubscriptions(sub)
}

func (v *VehicleBase) GetSubscriber(topic string) (*rclgo.Subscription, error) {

	for _, sub := range v.Subs.Subscriptions {
		if sub.TopicName == topic {
			return sub, nil
		}
	}
	return nil, fmt.Errorf("publisher for topic %s not found", topic)
}

func (v *VehicleBase) addGPSSubscriber() (*rclgo.Subscription, error) {
	qos := rclgo.QosProfile{
		Reliability: rclgo.ReliabilityBestEffort,
		Durability:  rclgo.DurabilityVolatile,
	}
	opts := rclgo.SubscriptionOptions{
		Qos: qos,
	}
	sub, err := sensor_msgs_msg.NewNavSatFixSubscription(
		v.node, v.node.Namespace()+"/global_position/global", &opts, //incompatible qos, need to fix
		func(msg *sensor_msgs_msg.NavSatFix, info *rclgo.MessageInfo, err error) {
			if err != nil {
				v.node.Logger().Errorf("failed to receive message: %v", err)
				return
			}
			pos := geographic_msgs_msg.GeoPoseStamped{
				Pose: geographic_msgs_msg.GeoPose{
					Position: geographic_msgs_msg.GeoPoint{
						Latitude:  msg.Latitude,
						Longitude: msg.Longitude,
						Altitude:  msg.Altitude,
					},
				},
			}
			//v.node.Logger().Infof("Received: %#v", msg)
			v.Position = pos
			//fmt.Printf("v.Position: %v\n", v.Position)
		})
	if err != nil {
		return nil, fmt.Errorf("failed to create subscription: %v", err)
	}
	fmt.Printf("Sub topic: %s on node: %s\n", sub.Subscription.TopicName, v.node.FullyQualifiedName())
	return sub.Subscription, nil
}

func (v *VehicleBase) GetPosition() geographic_msgs_msg.GeoPoseStamped {
	return v.Position
}
