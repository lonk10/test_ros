package main

import (
	"context"
	"fmt"
	"os"

	mavros "test/msgs/mavros_msgs/srv"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

//go:generate go run github.com/tiiuae/rclgo/cmd/rclgo-gen generate -d msgs --include-go-package-deps ./... --cgo-flags-path ""

func run() error {
	spinCtx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()
	qosProfile := rclgo.NewDefaultServiceQosProfile()

	serviceCtx, err := rclgo.NewContext(0, nil)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error while creating context: %v", err)
	}

	node, err := serviceCtx.NewNode("client_node", "/test")
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error while creating node: %v", err)
	}

	client, err := node.NewClient("set_mode", mavros.SetModeTypeSupport, &rclgo.ClientOptions{Qos: qosProfile})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error while creating client: %v", err)
	}

	go serviceCtx.Spin(spinCtx)

	req := mavros.NewSetMode_Request()
	req.BaseMode = 216

	res, _, err := client.Send(spinCtx, req)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error while sending: %v", err)
	}
	fmt.Println(res)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
