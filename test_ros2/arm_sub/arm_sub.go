package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	std_msgs "test/msgs/std_msgs/msg"

	arm "test/test_ros2/arm_copter"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

func run() error {
	rclArgs, _, err := rclgo.ParseArgs(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to parse ROS args: %v", err)
	}

	if err := rclgo.Init(rclArgs); err != nil {
		return fmt.Errorf("failed to initialize rclgo: %v", err)
	}
	defer rclgo.Uninit()

	node, err := rclgo.NewNode("arm_sub_client", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	defer node.Close()

	sub, err := std_msgs.NewBoolSubscription(
		node,
		"/abu/start",
		nil,
		func(msg *std_msgs.Bool, info *rclgo.MessageInfo, err error) {
			if err != nil {
				node.Logger().Errorf("failed to receive message: %v", err)
				return
			}
			node.Logger().Infof("Received: %#v", msg)
			if msg.Data {
				node.Logger().Infof("Starting copter...")
				err = arm.StartDrone("GUIDED", 10, 600)
				if err != nil {
					node.Logger().Errorf("failed to start copter: %v", err)
					return
				}
			}
		},
	)
	if err != nil {
		return fmt.Errorf("failed to create subscriber: %v", err)
	}
	defer sub.Close()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	ws, err := rclgo.NewWaitSet()
	if err != nil {
		return fmt.Errorf("failed to create waitset: %v", err)
	}
	defer ws.Close()
	ws.AddSubscriptions(sub.Subscription)
	return ws.Run(ctx)
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
