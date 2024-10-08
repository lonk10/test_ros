// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/srv/mount_configure.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/MountConfigure_Request", MountConfigure_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/MountConfigure_Request", MountConfigure_RequestTypeSupport)
}
const (
	MountConfigure_Request_MODE_RETRACT uint8 = 0// MAV_MOUNT_MODE
	MountConfigure_Request_MODE_NEUTRAL uint8 = 1
	MountConfigure_Request_MODE_MAVLINK_TARGETING uint8 = 2
	MountConfigure_Request_MODE_RC_TARGETING uint8 = 3
	MountConfigure_Request_MODE_GPS_POINT uint8 = 4
	MountConfigure_Request_INPUT_ANGLE_BODY_FRAME uint8 = 0// MOUNT_INPUT
	MountConfigure_Request_INPUT_ANGULAR_RATE uint8 = 1
	MountConfigure_Request_INPUT_ANGLE_ABSOLUTE_FRAME uint8 = 2
)

type MountConfigure_Request struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Mode uint8 `yaml:"mode"`// See enum MAV_MOUNT_MODE.
	StabilizeRoll bool `yaml:"stabilize_roll"`// stabilize roll? (1 = yes, 0 = no)
	StabilizePitch bool `yaml:"stabilize_pitch"`// stabilize pitch? (1 = yes, 0 = no)
	StabilizeYaw bool `yaml:"stabilize_yaw"`// stabilize yaw? (1 = yes, 0 = no)
	RollInput uint8 `yaml:"roll_input"`// roll input (See enum MOUNT_INPUT)
	PitchInput uint8 `yaml:"pitch_input"`// pitch input (See enum MOUNT_INPUT)
	YawInput uint8 `yaml:"yaw_input"`// yaw input (See enum MOUNT_INPUT)
}

// NewMountConfigure_Request creates a new MountConfigure_Request with default values.
func NewMountConfigure_Request() *MountConfigure_Request {
	self := MountConfigure_Request{}
	self.SetDefaults()
	return &self
}

func (t *MountConfigure_Request) Clone() *MountConfigure_Request {
	c := &MountConfigure_Request{}
	c.Header = *t.Header.Clone()
	c.Mode = t.Mode
	c.StabilizeRoll = t.StabilizeRoll
	c.StabilizePitch = t.StabilizePitch
	c.StabilizeYaw = t.StabilizeYaw
	c.RollInput = t.RollInput
	c.PitchInput = t.PitchInput
	c.YawInput = t.YawInput
	return c
}

func (t *MountConfigure_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MountConfigure_Request) SetDefaults() {
	t.Header.SetDefaults()
	t.Mode = 0
	t.StabilizeRoll = false
	t.StabilizePitch = false
	t.StabilizeYaw = false
	t.RollInput = 0
	t.PitchInput = 0
	t.YawInput = 0
}

func (t *MountConfigure_Request) GetTypeSupport() types.MessageTypeSupport {
	return MountConfigure_RequestTypeSupport
}

// MountConfigure_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MountConfigure_RequestPublisher struct {
	*rclgo.Publisher
}

// NewMountConfigure_RequestPublisher creates and returns a new publisher for the
// MountConfigure_Request
func NewMountConfigure_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MountConfigure_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, MountConfigure_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MountConfigure_RequestPublisher{pub}, nil
}

func (p *MountConfigure_RequestPublisher) Publish(msg *MountConfigure_Request) error {
	return p.Publisher.Publish(msg)
}

// MountConfigure_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MountConfigure_RequestSubscription struct {
	*rclgo.Subscription
}

// MountConfigure_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a MountConfigure_RequestSubscription.
type MountConfigure_RequestSubscriptionCallback func(msg *MountConfigure_Request, info *rclgo.MessageInfo, err error)

// NewMountConfigure_RequestSubscription creates and returns a new subscription for the
// MountConfigure_Request
func NewMountConfigure_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MountConfigure_RequestSubscriptionCallback) (*MountConfigure_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MountConfigure_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MountConfigure_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MountConfigure_RequestSubscription{sub}, nil
}

func (s *MountConfigure_RequestSubscription) TakeMessage(out *MountConfigure_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMountConfigure_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMountConfigure_RequestSlice(dst, src []MountConfigure_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MountConfigure_RequestTypeSupport types.MessageTypeSupport = _MountConfigure_RequestTypeSupport{}

type _MountConfigure_RequestTypeSupport struct{}

func (t _MountConfigure_RequestTypeSupport) New() types.Message {
	return NewMountConfigure_Request()
}

func (t _MountConfigure_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__MountConfigure_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__MountConfigure_Request__create())
}

func (t _MountConfigure_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__MountConfigure_Request__destroy((*C.mavros_msgs__srv__MountConfigure_Request)(pointer_to_free))
}

func (t _MountConfigure_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MountConfigure_Request)
	mem := (*C.mavros_msgs__srv__MountConfigure_Request)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.mode = C.uint8_t(m.Mode)
	mem.stabilize_roll = C.bool(m.StabilizeRoll)
	mem.stabilize_pitch = C.bool(m.StabilizePitch)
	mem.stabilize_yaw = C.bool(m.StabilizeYaw)
	mem.roll_input = C.uint8_t(m.RollInput)
	mem.pitch_input = C.uint8_t(m.PitchInput)
	mem.yaw_input = C.uint8_t(m.YawInput)
}

func (t _MountConfigure_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MountConfigure_Request)
	mem := (*C.mavros_msgs__srv__MountConfigure_Request)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Mode = uint8(mem.mode)
	m.StabilizeRoll = bool(mem.stabilize_roll)
	m.StabilizePitch = bool(mem.stabilize_pitch)
	m.StabilizeYaw = bool(mem.stabilize_yaw)
	m.RollInput = uint8(mem.roll_input)
	m.PitchInput = uint8(mem.pitch_input)
	m.YawInput = uint8(mem.yaw_input)
}

func (t _MountConfigure_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__MountConfigure_Request())
}

type CMountConfigure_Request = C.mavros_msgs__srv__MountConfigure_Request
type CMountConfigure_Request__Sequence = C.mavros_msgs__srv__MountConfigure_Request__Sequence

func MountConfigure_Request__Sequence_to_Go(goSlice *[]MountConfigure_Request, cSlice CMountConfigure_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MountConfigure_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MountConfigure_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MountConfigure_Request__Sequence_to_C(cSlice *CMountConfigure_Request__Sequence, goSlice []MountConfigure_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__MountConfigure_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__MountConfigure_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MountConfigure_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MountConfigure_Request__Array_to_Go(goSlice []MountConfigure_Request, cSlice []CMountConfigure_Request) {
	for i := 0; i < len(cSlice); i++ {
		MountConfigure_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MountConfigure_Request__Array_to_C(cSlice []CMountConfigure_Request, goSlice []MountConfigure_Request) {
	for i := 0; i < len(goSlice); i++ {
		MountConfigure_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
