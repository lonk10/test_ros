package sub

import (
	"context"
	"fmt"

	geo_msg "test/msgs/geometry_msgs/msg"
	mav_srv "test/msgs/mavros_msgs/srv"

	"github.com/tiiuae/rclgo/pkg/rclgo"

	vbase "test/test_detect_go/vbase"
)

type SubVehicle struct {
	ID             string
	Name           string
	Domain         string
	LocalAgent     *vbase.VehicleBase
	LocalNode      *rclgo.Node
	context        context.Context
	cancelContext  context.CancelFunc
	ServiceClients map[string]*rclgo.Client
	GlobalNode     *rclgo.Node
	Pubs           map[string]*rclgo.Publisher
	Subs           *rclgo.WaitSet
}

func NewSubVehicle(lname string, gname string) (*SubVehicle, error) {
	res := &SubVehicle{
		ID:             lname,
		Name:           "/mavros/" + lname,
		Domain:         gname,
		ServiceClients: make(map[string]*rclgo.Client),
		Pubs:           make(map[string]*rclgo.Publisher),
	}
	var err error

	vb, err := vbase.NewVehicleBase()
	if err != nil {
		return nil, fmt.Errorf("failed to create agent: %v", err)
	}
	res.LocalAgent = vb

	if err := rclgo.Init(nil); err != nil {
		return nil, fmt.Errorf("failed to initialize rclgo: %v", err)
	}

	return res, nil
}

func (s *SubVehicle) Init() error {
	//context init
	spinCtx, cancelCtx := context.WithCancel(context.Background())
	s.context = spinCtx
	s.cancelContext = cancelCtx

	qosProfile := rclgo.NewDefaultServiceQosProfile()

	serviceCtx, err := rclgo.NewContext(0, nil)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("failed to create context: %v", err)
	}

	s.LocalNode, err = serviceCtx.NewNode("arm_sub_client", s.Name)
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}

	s.GlobalNode, err = serviceCtx.NewNode(s.ID, s.Domain)
	if err != nil {
		return fmt.Errorf("failed to create node: %v", err)
	}

	s.LocalAgent.Start(s.Name)

	//arming service client
	client, err := mav_srv.NewCommandBoolClient(s.LocalNode, "cmd/arming", &rclgo.ClientOptions{Qos: qosProfile})
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	s.ServiceClients["cmd/arming"] = client.Client
	//time.Sleep(2 * time.Second)
	//set mode client
	clientSM, err := mav_srv.NewSetModeClient(s.LocalNode, "set_mode", &rclgo.ClientOptions{Qos: qosProfile})
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}
	s.ServiceClients["set_mode"] = clientSM.Client
	//time.Sleep(2 * time.Second)

	//create publisher for movement
	//publisher for velocity
	pubSV, err := geo_msg.NewTwistPublisher(s.LocalNode, "setpoint_velocity/cmd_vel", nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}

	s.LocalAgent.Pubs["setpoint_velocity/cmd_vel"] = pubSV.Publisher

	//publisher for setpos
	pubSP, err := geo_msg.NewPoseStampedPublisher(s.LocalNode, "setpoint_position/local", nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}

	s.LocalAgent.Pubs["setpoint_position/local"] = pubSP.Publisher

	pubBl, err := geo_msg.NewPoseStampedPublisher(s.GlobalNode, "mission", nil)
	if err != nil {
		return fmt.Errorf("failed to create publisher: %v", err)
	}

	s.Pubs["mission"] = pubBl.Publisher
	//time.Sleep(2 * time.Second)
	go serviceCtx.Spin(s.context)
	//
	return nil
}

// sets a velocity vector for the vehicle
// !! won't stop until velocity is set to 0 on all axis
func (s *SubVehicle) SetVelocity(x float64, y float64, z float64) error {
	topic := "setpoint_velocity/cmd_vel"
	pub := s.LocalAgent.Pubs[topic]
	if pub == nil {
		return fmt.Errorf("publisher %s not found", topic)
	}
	msg := geo_msg.Twist{
		Linear: geo_msg.Vector3{X: x, Y: y, Z: z},
	}
	err := pub.Publish(&msg)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}

// sends command to move certain meters on x,y,z from current point
func (s *SubVehicle) SetPoint(x float64, y float64, z float64) error {
	topic := "setpoint_position/local"
	pub := s.LocalAgent.Pubs[topic]
	if pub == nil {
		return fmt.Errorf("publisher %s not found", topic)
	}
	pos := geo_msg.Pose{
		Position: geo_msg.Point{X: x, Y: y, Z: z},
	}
	msg := geo_msg.PoseStamped{
		Pose: pos,
	}
	err := pub.Publish(&msg)
	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}
	return nil
}

func (s *SubVehicle) ArmMotors() error {
	msg := mav_srv.CommandBool_Request{
		Value: true, // Set to "ARMED"
	}
	topic := "cmd/arming"
	client := s.ServiceClients[topic]
	if client == nil {
		return fmt.Errorf("could not find client for service %s", topic)
	}
	fmt.Printf("sending on topic %s\n", topic)
	_, _, err := client.Send(s.context, &msg)
	fmt.Println("bbb")
	if err != nil {
		return fmt.Errorf("failed to call service: %v", err)
	}
	return nil
}

func (s *SubVehicle) SetMode(cmode string) error {
	msg := mav_srv.SetMode_Request{
		CustomMode: cmode, // Set to "ARMED"
	}
	topic := "set_mode"
	client := s.ServiceClients[topic]
	if client == nil {
		return fmt.Errorf("could not find client for service %s", topic)
	}
	_, _, err := client.Send(s.context, &msg)
	if err != nil {
		return fmt.Errorf("failed to call service: %v", err)
	}
	return nil
}
