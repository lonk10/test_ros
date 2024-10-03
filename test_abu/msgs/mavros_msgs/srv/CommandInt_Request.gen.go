// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/srv/command_int.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/CommandInt_Request", CommandInt_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/CommandInt_Request", CommandInt_RequestTypeSupport)
}

type CommandInt_Request struct {
	Broadcast bool `yaml:"broadcast"`// send this command in broadcast mode
	Frame uint8 `yaml:"frame"`
	Command uint16 `yaml:"command"`
	Current uint8 `yaml:"current"`
	Autocontinue uint8 `yaml:"autocontinue"`
	Param1 float32 `yaml:"param1"`
	Param2 float32 `yaml:"param2"`
	Param3 float32 `yaml:"param3"`
	Param4 float32 `yaml:"param4"`
	X int32 `yaml:"x"`// latitude in deg * 1E7 or local x * 1E4 m
	Y int32 `yaml:"y"`// longitude in deg * 1E7 or local y * 1E4 m
	Z float32 `yaml:"z"`// altitude
}

// NewCommandInt_Request creates a new CommandInt_Request with default values.
func NewCommandInt_Request() *CommandInt_Request {
	self := CommandInt_Request{}
	self.SetDefaults()
	return &self
}

func (t *CommandInt_Request) Clone() *CommandInt_Request {
	c := &CommandInt_Request{}
	c.Broadcast = t.Broadcast
	c.Frame = t.Frame
	c.Command = t.Command
	c.Current = t.Current
	c.Autocontinue = t.Autocontinue
	c.Param1 = t.Param1
	c.Param2 = t.Param2
	c.Param3 = t.Param3
	c.Param4 = t.Param4
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	return c
}

func (t *CommandInt_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CommandInt_Request) SetDefaults() {
	t.Broadcast = false
	t.Frame = 0
	t.Command = 0
	t.Current = 0
	t.Autocontinue = 0
	t.Param1 = 0
	t.Param2 = 0
	t.Param3 = 0
	t.Param4 = 0
	t.X = 0
	t.Y = 0
	t.Z = 0
}

func (t *CommandInt_Request) GetTypeSupport() types.MessageTypeSupport {
	return CommandInt_RequestTypeSupport
}

// CommandInt_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CommandInt_RequestPublisher struct {
	*rclgo.Publisher
}

// NewCommandInt_RequestPublisher creates and returns a new publisher for the
// CommandInt_Request
func NewCommandInt_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CommandInt_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, CommandInt_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CommandInt_RequestPublisher{pub}, nil
}

func (p *CommandInt_RequestPublisher) Publish(msg *CommandInt_Request) error {
	return p.Publisher.Publish(msg)
}

// CommandInt_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CommandInt_RequestSubscription struct {
	*rclgo.Subscription
}

// CommandInt_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a CommandInt_RequestSubscription.
type CommandInt_RequestSubscriptionCallback func(msg *CommandInt_Request, info *rclgo.MessageInfo, err error)

// NewCommandInt_RequestSubscription creates and returns a new subscription for the
// CommandInt_Request
func NewCommandInt_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback CommandInt_RequestSubscriptionCallback) (*CommandInt_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg CommandInt_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CommandInt_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &CommandInt_RequestSubscription{sub}, nil
}

func (s *CommandInt_RequestSubscription) TakeMessage(out *CommandInt_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneCommandInt_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCommandInt_RequestSlice(dst, src []CommandInt_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CommandInt_RequestTypeSupport types.MessageTypeSupport = _CommandInt_RequestTypeSupport{}

type _CommandInt_RequestTypeSupport struct{}

func (t _CommandInt_RequestTypeSupport) New() types.Message {
	return NewCommandInt_Request()
}

func (t _CommandInt_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__CommandInt_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__CommandInt_Request__create())
}

func (t _CommandInt_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__CommandInt_Request__destroy((*C.mavros_msgs__srv__CommandInt_Request)(pointer_to_free))
}

func (t _CommandInt_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CommandInt_Request)
	mem := (*C.mavros_msgs__srv__CommandInt_Request)(dst)
	mem.broadcast = C.bool(m.Broadcast)
	mem.frame = C.uint8_t(m.Frame)
	mem.command = C.uint16_t(m.Command)
	mem.current = C.uint8_t(m.Current)
	mem.autocontinue = C.uint8_t(m.Autocontinue)
	mem.param1 = C.float(m.Param1)
	mem.param2 = C.float(m.Param2)
	mem.param3 = C.float(m.Param3)
	mem.param4 = C.float(m.Param4)
	mem.x = C.int32_t(m.X)
	mem.y = C.int32_t(m.Y)
	mem.z = C.float(m.Z)
}

func (t _CommandInt_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CommandInt_Request)
	mem := (*C.mavros_msgs__srv__CommandInt_Request)(ros2_message_buffer)
	m.Broadcast = bool(mem.broadcast)
	m.Frame = uint8(mem.frame)
	m.Command = uint16(mem.command)
	m.Current = uint8(mem.current)
	m.Autocontinue = uint8(mem.autocontinue)
	m.Param1 = float32(mem.param1)
	m.Param2 = float32(mem.param2)
	m.Param3 = float32(mem.param3)
	m.Param4 = float32(mem.param4)
	m.X = int32(mem.x)
	m.Y = int32(mem.y)
	m.Z = float32(mem.z)
}

func (t _CommandInt_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__CommandInt_Request())
}

type CCommandInt_Request = C.mavros_msgs__srv__CommandInt_Request
type CCommandInt_Request__Sequence = C.mavros_msgs__srv__CommandInt_Request__Sequence

func CommandInt_Request__Sequence_to_Go(goSlice *[]CommandInt_Request, cSlice CCommandInt_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CommandInt_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		CommandInt_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func CommandInt_Request__Sequence_to_C(cSlice *CCommandInt_Request__Sequence, goSlice []CommandInt_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__CommandInt_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__CommandInt_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		CommandInt_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func CommandInt_Request__Array_to_Go(goSlice []CommandInt_Request, cSlice []CCommandInt_Request) {
	for i := 0; i < len(cSlice); i++ {
		CommandInt_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CommandInt_Request__Array_to_C(cSlice []CCommandInt_Request, goSlice []CommandInt_Request) {
	for i := 0; i < len(goSlice); i++ {
		CommandInt_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
