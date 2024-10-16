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

#include <mavros_msgs/srv/command_tol.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/CommandTOL_Response", CommandTOL_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/CommandTOL_Response", CommandTOL_ResponseTypeSupport)
}

type CommandTOL_Response struct {
	Success bool `yaml:"success"`
	Result uint8 `yaml:"result"`
}

// NewCommandTOL_Response creates a new CommandTOL_Response with default values.
func NewCommandTOL_Response() *CommandTOL_Response {
	self := CommandTOL_Response{}
	self.SetDefaults()
	return &self
}

func (t *CommandTOL_Response) Clone() *CommandTOL_Response {
	c := &CommandTOL_Response{}
	c.Success = t.Success
	c.Result = t.Result
	return c
}

func (t *CommandTOL_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *CommandTOL_Response) SetDefaults() {
	t.Success = false
	t.Result = 0
}

func (t *CommandTOL_Response) GetTypeSupport() types.MessageTypeSupport {
	return CommandTOL_ResponseTypeSupport
}

// CommandTOL_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type CommandTOL_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewCommandTOL_ResponsePublisher creates and returns a new publisher for the
// CommandTOL_Response
func NewCommandTOL_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*CommandTOL_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, CommandTOL_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CommandTOL_ResponsePublisher{pub}, nil
}

func (p *CommandTOL_ResponsePublisher) Publish(msg *CommandTOL_Response) error {
	return p.Publisher.Publish(msg)
}

// CommandTOL_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type CommandTOL_ResponseSubscription struct {
	*rclgo.Subscription
}

// CommandTOL_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a CommandTOL_ResponseSubscription.
type CommandTOL_ResponseSubscriptionCallback func(msg *CommandTOL_Response, info *rclgo.MessageInfo, err error)

// NewCommandTOL_ResponseSubscription creates and returns a new subscription for the
// CommandTOL_Response
func NewCommandTOL_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback CommandTOL_ResponseSubscriptionCallback) (*CommandTOL_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg CommandTOL_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, CommandTOL_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &CommandTOL_ResponseSubscription{sub}, nil
}

func (s *CommandTOL_ResponseSubscription) TakeMessage(out *CommandTOL_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneCommandTOL_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneCommandTOL_ResponseSlice(dst, src []CommandTOL_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var CommandTOL_ResponseTypeSupport types.MessageTypeSupport = _CommandTOL_ResponseTypeSupport{}

type _CommandTOL_ResponseTypeSupport struct{}

func (t _CommandTOL_ResponseTypeSupport) New() types.Message {
	return NewCommandTOL_Response()
}

func (t _CommandTOL_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__CommandTOL_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__CommandTOL_Response__create())
}

func (t _CommandTOL_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__CommandTOL_Response__destroy((*C.mavros_msgs__srv__CommandTOL_Response)(pointer_to_free))
}

func (t _CommandTOL_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*CommandTOL_Response)
	mem := (*C.mavros_msgs__srv__CommandTOL_Response)(dst)
	mem.success = C.bool(m.Success)
	mem.result = C.uint8_t(m.Result)
}

func (t _CommandTOL_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*CommandTOL_Response)
	mem := (*C.mavros_msgs__srv__CommandTOL_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	m.Result = uint8(mem.result)
}

func (t _CommandTOL_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__CommandTOL_Response())
}

type CCommandTOL_Response = C.mavros_msgs__srv__CommandTOL_Response
type CCommandTOL_Response__Sequence = C.mavros_msgs__srv__CommandTOL_Response__Sequence

func CommandTOL_Response__Sequence_to_Go(goSlice *[]CommandTOL_Response, cSlice CCommandTOL_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]CommandTOL_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		CommandTOL_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func CommandTOL_Response__Sequence_to_C(cSlice *CCommandTOL_Response__Sequence, goSlice []CommandTOL_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__CommandTOL_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__CommandTOL_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		CommandTOL_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func CommandTOL_Response__Array_to_Go(goSlice []CommandTOL_Response, cSlice []CCommandTOL_Response) {
	for i := 0; i < len(cSlice); i++ {
		CommandTOL_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func CommandTOL_Response__Array_to_C(cSlice []CCommandTOL_Response, goSlice []CommandTOL_Response) {
	for i := 0; i < len(goSlice); i++ {
		CommandTOL_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
