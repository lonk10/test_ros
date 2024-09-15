package main

import (
	"context"
	"fmt"
	"os"

	msgs "test/msgs/mavros_msgs/srv"

	"github.com/tiiuae/rclgo/pkg/rclgo"
)

func run() error {
	spinCtx, cancelSpin := context.WithCancel(context.Background())
	defer cancelSpin()

	ctx, err := rclgo.NewContext(0, nil)
	if err != nil {
		return fmt.Errorf("error while creating context: %v", err)
	}

	node, err := ctx.NewNode("setmode_node", "/test")
	if err != nil {
		return fmt.Errorf("error while creating node: %v", err)
	}

	_, err = msgs.NewSetModeService(node, "set_mode", nil, func(info *rclgo.ServiceInfo, req *msgs.SetMode_Request, sender msgs.SetModeServiceResponseSender) {
		res := msgs.NewSetMode_Response()
		res.ModeSent = true
		sender.SendResponse(res)
	})

	if err != nil {
		return fmt.Errorf("error occured starting service: %v", err)
	}

	go ctx.Spin(spinCtx)

	select {}
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
