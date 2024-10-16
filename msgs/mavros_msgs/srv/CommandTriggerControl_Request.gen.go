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

#include <mavros_msgs/srv/command_trigger_control.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/CommandTriggerControl_Request", CommandTriggerControl_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/CommandTriggerControl_Request", CommandTriggerControl_RequestTypeSupport)
}

type CommandTriggerControl_Request struct {
	TriggerEnable bool `yaml:"trigger_enable"`// Trigger enable/disable
	SequenceReset bool `yaml:"sequence_reset"`// Reset the trigger sequence
	TriggerPause bool `yaml:"trigger_pause"`// Pause triggering, but without switching the camera off or retracting it.
}

// NewCommandTriggerControl_Request creates a new CommandTriggerControl_Request with default values.
func NewCommandTriggerControl_Request() *CommandTriggerControl_Request {
	self := CommandTriggerControl_Request{}
	self.SetDefaults()
	return &self
}

func (t *CommandTriggerControl_Request) Clone() *CommandTriggerControl_Request {
	c := &CommandTriggerControl_Request{}
	c.TriggerEnable = t.TriggerEnable
	c.SequenceReset = t.SequenceReset
	c.TriggerPause = t.TriggerPause
	return c
}

func (t *CommandTriggerControl_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CommandTriggerControl_Request) SetDefaults() {
	t.TriggerEnable = false
	t.SequenceReset = false
	t.TriggerPause = false
}

func (t *CommandTriggerControl_Request) GetTypeSupport() types.MessageTypeSupport {
	return CommandTriggerControl_RequestTypeSupport
}

// CommandTriggerControl_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CommandTriggerControl_RequestPublisher struct {
	*rclgo.Publisher
}

// NewCommandTriggerControl_RequestPublisher creates and returns a new publisher for the
// CommandTriggerControl_Request
func NewCommandTriggerControl_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CommandTriggerControl_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, CommandTriggerControl_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CommandTriggerControl_RequestPublisher{pub}, nil
}

func (p *CommandTriggerControl_RequestPublisher) Publish(msg *CommandTriggerControl_Request) error {
	return p.Publisher.Publish(msg)
}

// CommandTriggerControl_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CommandTriggerControl_RequestSubscription struct {
	*rclgo.Subscription
}

// CommandTriggerControl_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a CommandTriggerControl_RequestSubscription.
type CommandTriggerControl_RequestSubscriptionCallback func(msg *CommandTriggerControl_Request, info *rclgo.MessageInfo, err error)

// NewCommandTriggerControl_RequestSubscription creates and returns a new subscription for the
// CommandTriggerControl_Request
func NewCommandTriggerControl_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback CommandTriggerControl_RequestSubscriptionCallback) (*CommandTriggerControl_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg CommandTriggerControl_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CommandTriggerControl_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &CommandTriggerControl_RequestSubscription{sub}, nil
}

func (s *CommandTriggerControl_RequestSubscription) TakeMessage(out *CommandTriggerControl_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneCommandTriggerControl_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCommandTriggerControl_RequestSlice(dst, src []CommandTriggerControl_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CommandTriggerControl_RequestTypeSupport types.MessageTypeSupport = _CommandTriggerControl_RequestTypeSupport{}

type _CommandTriggerControl_RequestTypeSupport struct{}

func (t _CommandTriggerControl_RequestTypeSupport) New() types.Message {
	return NewCommandTriggerControl_Request()
}

func (t _CommandTriggerControl_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__CommandTriggerControl_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__CommandTriggerControl_Request__create())
}

func (t _CommandTriggerControl_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__CommandTriggerControl_Request__destroy((*C.mavros_msgs__srv__CommandTriggerControl_Request)(pointer_to_free))
}

func (t _CommandTriggerControl_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CommandTriggerControl_Request)
	mem := (*C.mavros_msgs__srv__CommandTriggerControl_Request)(dst)
	mem.trigger_enable = C.bool(m.TriggerEnable)
	mem.sequence_reset = C.bool(m.SequenceReset)
	mem.trigger_pause = C.bool(m.TriggerPause)
}

func (t _CommandTriggerControl_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CommandTriggerControl_Request)
	mem := (*C.mavros_msgs__srv__CommandTriggerControl_Request)(ros2_message_buffer)
	m.TriggerEnable = bool(mem.trigger_enable)
	m.SequenceReset = bool(mem.sequence_reset)
	m.TriggerPause = bool(mem.trigger_pause)
}

func (t _CommandTriggerControl_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__CommandTriggerControl_Request())
}

type CCommandTriggerControl_Request = C.mavros_msgs__srv__CommandTriggerControl_Request
type CCommandTriggerControl_Request__Sequence = C.mavros_msgs__srv__CommandTriggerControl_Request__Sequence

func CommandTriggerControl_Request__Sequence_to_Go(goSlice *[]CommandTriggerControl_Request, cSlice CCommandTriggerControl_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CommandTriggerControl_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		CommandTriggerControl_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func CommandTriggerControl_Request__Sequence_to_C(cSlice *CCommandTriggerControl_Request__Sequence, goSlice []CommandTriggerControl_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__CommandTriggerControl_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__CommandTriggerControl_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		CommandTriggerControl_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func CommandTriggerControl_Request__Array_to_Go(goSlice []CommandTriggerControl_Request, cSlice []CCommandTriggerControl_Request) {
	for i := 0; i < len(cSlice); i++ {
		CommandTriggerControl_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CommandTriggerControl_Request__Array_to_C(cSlice []CCommandTriggerControl_Request, goSlice []CommandTriggerControl_Request) {
	for i := 0; i < len(goSlice); i++ {
		CommandTriggerControl_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
