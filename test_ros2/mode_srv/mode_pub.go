package main

import (
	"fmt"
	"os"
	"time"

	"github.com/tiiuae/rclgo/pkg/rclgo"

	mavros "test/msgs/mavros_msgs/srv"
)

func run2() error {
	// Initialize ROS 2 context
	rclArgs, _, err := rclgo.ParseArgs(os.Args[1:])
	if err != nil {
		return fmt.Errorf("failed to parse ROS args: %v", err)
	}

	if err := rclgo.Init(rclArgs); err != nil {
		return fmt.Errorf("failed to initialize rclgo: %v", err)
	}
	defer rclgo.Uninit()

	// Create a ROS 2 node
	node, err := rclgo.NewNode("mode_change_node", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}

	// Create a publisher to the /mavros/set_mode topic
	publisher, err := node.NewPublisher("/mavros/set_mode", mavros.SetMode_RequestTypeSupport, nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}
	defer publisher.Close()

	// Wait a few seconds to ensure everything is connected
	time.Sleep(2 * time.Second)

	// Construct the mode change message
	modeMsg := mavros.SetMode_Request{
		BaseMode:   mavros.SetMode_Request_MAV_MODE_GUIDED_ARMED, // Leave as 0 for custom mode
		CustomMode: "",                                           // Set to "GUIDED"
	}

	// Publish the message to change mode to GUIDED
	fmt.Println("Changing mode to GUIDED...")
	err = publisher.Publish(&modeMsg)
	if err != nil {
		return fmt.Errorf("failed to publish: %v", err)
	}

	// Keep the node running for a short time to ensure the message is sent
	time.Sleep(5 * time.Second)

	return nil
}

func main2() {
	if err := run2(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
