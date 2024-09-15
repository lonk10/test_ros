package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	mavros "test/msgs/mavros_msgs/srv"

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

	node, err := rclgo.NewNode("publisher", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}
	defer node.Close()

	pub, err := mavros.NewCommandBool_RequestPublisher(node, "/abu/start", nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}
	defer pub.Close()

	msg := mavros.CommandBool_Request{
		Value: true, // Set to "ARMED"
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	node.Logger().Infof("Publishing: %#v", msg)
	if err := pub.Publish(&msg); err != nil {
		return fmt.Errorf("failed to publish: %v", err)
	}
	select {
	case <-ctx.Done():
		return nil
	case <-time.After(time.Second):
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
