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

#include <mavros_msgs/srv/command_bool.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/CommandBool_Request", CommandBool_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/CommandBool_Request", CommandBool_RequestTypeSupport)
}

type CommandBool_Request struct {
	Value bool `yaml:"value"`
}

// NewCommandBool_Request creates a new CommandBool_Request with default values.
func NewCommandBool_Request() *CommandBool_Request {
	self := CommandBool_Request{}
	self.SetDefaults()
	return &self
}

func (t *CommandBool_Request) Clone() *CommandBool_Request {
	c := &CommandBool_Request{}
	c.Value = t.Value
	return c
}

func (t *CommandBool_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CommandBool_Request) SetDefaults() {
	t.Value = false
}

func (t *CommandBool_Request) GetTypeSupport() types.MessageTypeSupport {
	return CommandBool_RequestTypeSupport
}

// CommandBool_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CommandBool_RequestPublisher struct {
	*rclgo.Publisher
}

// NewCommandBool_RequestPublisher creates and returns a new publisher for the
// CommandBool_Request
func NewCommandBool_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CommandBool_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, CommandBool_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CommandBool_RequestPublisher{pub}, nil
}

func (p *CommandBool_RequestPublisher) Publish(msg *CommandBool_Request) error {
	return p.Publisher.Publish(msg)
}

// CommandBool_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CommandBool_RequestSubscription struct {
	*rclgo.Subscription
}

// CommandBool_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a CommandBool_RequestSubscription.
type CommandBool_RequestSubscriptionCallback func(msg *CommandBool_Request, info *rclgo.MessageInfo, err error)

// NewCommandBool_RequestSubscription creates and returns a new subscription for the
// CommandBool_Request
func NewCommandBool_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback CommandBool_RequestSubscriptionCallback) (*CommandBool_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg CommandBool_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CommandBool_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &CommandBool_RequestSubscription{sub}, nil
}

func (s *CommandBool_RequestSubscription) TakeMessage(out *CommandBool_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneCommandBool_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCommandBool_RequestSlice(dst, src []CommandBool_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CommandBool_RequestTypeSupport types.MessageTypeSupport = _CommandBool_RequestTypeSupport{}

type _CommandBool_RequestTypeSupport struct{}

func (t _CommandBool_RequestTypeSupport) New() types.Message {
	return NewCommandBool_Request()
}

func (t _CommandBool_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__CommandBool_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__CommandBool_Request__create())
}

func (t _CommandBool_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__CommandBool_Request__destroy((*C.mavros_msgs__srv__CommandBool_Request)(pointer_to_free))
}

func (t _CommandBool_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CommandBool_Request)
	mem := (*C.mavros_msgs__srv__CommandBool_Request)(dst)
	mem.value = C.bool(m.Value)
}

func (t _CommandBool_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CommandBool_Request)
	mem := (*C.mavros_msgs__srv__CommandBool_Request)(ros2_message_buffer)
	m.Value = bool(mem.value)
}

func (t _CommandBool_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__CommandBool_Request())
}

type CCommandBool_Request = C.mavros_msgs__srv__CommandBool_Request
type CCommandBool_Request__Sequence = C.mavros_msgs__srv__CommandBool_Request__Sequence

func CommandBool_Request__Sequence_to_Go(goSlice *[]CommandBool_Request, cSlice CCommandBool_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CommandBool_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		CommandBool_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func CommandBool_Request__Sequence_to_C(cSlice *CCommandBool_Request__Sequence, goSlice []CommandBool_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__CommandBool_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__CommandBool_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		CommandBool_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func CommandBool_Request__Array_to_Go(goSlice []CommandBool_Request, cSlice []CCommandBool_Request) {
	for i := 0; i < len(cSlice); i++ {
		CommandBool_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CommandBool_Request__Array_to_C(cSlice []CCommandBool_Request, goSlice []CommandBool_Request) {
	for i := 0; i < len(goSlice); i++ {
		CommandBool_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
