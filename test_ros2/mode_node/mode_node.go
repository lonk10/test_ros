package main

import (
	"context"
	"fmt"
	"log"
	"time"

	mavros "test/msgs/mavros_msgs/srv"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

func main() {
	// Create ROS2 context
	ctx := context.Background()
	rosCtx, err := rclgo.NewContext(0, nil)
	if err != nil {
		log.Fatalf("Failed to create ROS2 context: %v", err)
	}
	defer rosCtx.Close()

	// Create a ROS2 node
	fmt.Println("a")
	node, err := rosCtx.NewNode("arm_node", "")
	if err != nil {
		log.Fatalf("Failed to create ROS2 node: %v", err)
	}
	defer node.Close()

	qosProfile := rclgo.NewDefaultServiceQosProfile()
	// Create a client for the "arming" service
	fmt.Println("b")
	client, err := mavros.NewSetModeClient(node, "/mavros/set_mode", &rclgo.ClientOptions{Qos: qosProfile})
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	log.Println("Service is available, proceeding with the request.")

	// Create a service request to arm the drone
	req := mavros.NewSetMode_Request()
	req.BaseMode = 216

	fmt.Println("c")
	// Call the arming service and wait for the response
	res, _, err := client.Send(ctx, req)
	if err != nil {
		log.Fatalf("Failed to send: %v", err)
	}

	fmt.Println("d")
	// Check if the drone was successfully armed
	if res.ModeSent {
		log.Println("Drone armed successfully!")
	} else {
		log.Printf("Failed to arm the drone, result code: %d\n", res.ModeSent)
	}

	fmt.Println("e")
	// Keep the node alive for a while to observe results
	time.Sleep(5 * time.Second)
}
