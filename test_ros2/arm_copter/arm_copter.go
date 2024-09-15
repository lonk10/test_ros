package arm_copter

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/tiiuae/rclgo/pkg/rclgo"

	mav_msg "test/msgs/mavros_msgs/msg"
	mav_srv "test/msgs/mavros_msgs/srv"
)

func StartDrone(mode string, altitude float32, newalt float32) error {
	// Initialize ROS 2 context
	spinCtx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	qosProfile := rclgo.NewDefaultServiceQosProfile()

	serviceCtx, err := rclgo.NewContext(0, nil)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to create context: %v", err)
	}

	// Create a ROS 2 node
	node, err := serviceCtx.NewNode("arming_node", "")
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}

	err = ArmMotors(node, serviceCtx, spinCtx, qosProfile)
	if err != nil {
		return fmt.Errorf("failed to arm motors: %v", err)
	}
	err = SetMode(mode, node, serviceCtx, spinCtx, qosProfile)
	if err != nil {
		return fmt.Errorf("failed to set mode: %v", err)
	}
	err = TakeOff(altitude, node, serviceCtx, spinCtx, qosProfile)
	if err != nil {
		return fmt.Errorf("failed to change altitude: %v", err)
	}
	err = ChangeAltitude(newalt, node, serviceCtx, spinCtx, qosProfile)
	if err != nil {
		return fmt.Errorf("failed to change altitude: %v", err)
	}

	return nil
}

func ArmMotors(node *rclgo.Node, serviceCtx *rclgo.Context, spinCtx context.Context, qosProfile rclgo.QosProfile) error {

	// Create a publisher to the /mavros/cmd/arming topic
	client, err := mav_srv.NewCommandBoolClient(node, "/mavros/cmd/arming", &rclgo.ClientOptions{Qos: qosProfile})
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	go serviceCtx.Spin(spinCtx)

	// Wait a few seconds to ensure everything is connected
	time.Sleep(2 * time.Second)

	// Construct the mode change message
	modeMsg := mav_srv.CommandBool_Request{
		Value: true, // Set to "ARMED"
	}

	// Publish the message to change mode to ARMED
	fmt.Println("Arming motors...")
	_, _, err = client.Send(spinCtx, &modeMsg)
	if err != nil {
		return fmt.Errorf("failed to call: %v", err)
	}
	fmt.Println("Done.")

	// Keep the node running for a short time to ensure the message is sent
	time.Sleep(5 * time.Second)

	return nil
}

func SetMode(mode string, node *rclgo.Node, serviceCtx *rclgo.Context, spinCtx context.Context, qosProfile rclgo.QosProfile) error {
	// Create a publisher to the /mavros/cmd/arming topic
	client, err := mav_srv.NewSetModeClient(node, "/mavros/set_mode", &rclgo.ClientOptions{Qos: qosProfile})
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	go serviceCtx.Spin(spinCtx)

	// Wait a few seconds to ensure everything is connected
	time.Sleep(2 * time.Second)

	// Construct the mode change message
	modeMsg := mav_srv.SetMode_Request{
		CustomMode: mode, // Set to "ARMED"
	}

	// Publish the message to change mode to ARMED
	fmt.Println("Changing mode to GUIDED...")
	_, _, err = client.Send(spinCtx, &modeMsg)
	if err != nil {
		return fmt.Errorf("failed to call: %v", err)
	}
	fmt.Println("Done.")

	// Keep the node running for a short time to ensure the message is sent
	time.Sleep(5 * time.Second)

	return nil
}

func TakeOff(altitude float32, node *rclgo.Node, serviceCtx *rclgo.Context, spinCtx context.Context, qosProfile rclgo.QosProfile) error {
	client, err := mav_srv.NewCommandTOLClient(node, "/mavros/cmd/takeoff", &rclgo.ClientOptions{Qos: qosProfile})
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	defer client.Close()

	go serviceCtx.Spin(spinCtx)

	// Wait a few seconds to ensure everything is connected
	time.Sleep(2 * time.Second)

	// Construct the mode change message
	modeMsg := mav_srv.CommandTOL_Request{
		Altitude: altitude,
	}

	// Publish the message to change mode to ARMED
	fmt.Println("Taking off...")
	_, _, err = client.Send(spinCtx, &modeMsg)
	if err != nil {
		return fmt.Errorf("failed to call: %v", err)
	}
	fmt.Println("Done.")

	// Keep the node running for a short time to ensure the message is sent
	time.Sleep(5 * time.Second)

	return nil
}

func ChangeAltitude(altitude float32, node *rclgo.Node, serviceCtx *rclgo.Context, spinCtx context.Context, qosProfile rclgo.QosProfile) error {
	pub, err := mav_msg.NewGlobalPositionTargetPublisher(node, "/mavros/setpoint_raw/global", nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}
	defer pub.Close()

	msg := mav_msg.GlobalPositionTarget{
		Altitude: altitude, // Set to "ARMED"
	}

	_, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	node.Logger().Infof("Publishing: %#v", msg)
	if err := pub.Publish(&msg); err != nil {
		return fmt.Errorf("failed to publish: %v", err)
	}
	node.Logger().Infof("Published.")
	return nil
}

func main() {
	if err := StartDrone("GUIDED", 4, 50); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
