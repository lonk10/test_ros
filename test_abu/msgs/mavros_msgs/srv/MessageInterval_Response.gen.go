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

#include <mavros_msgs/srv/message_interval.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/MessageInterval_Response", MessageInterval_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/MessageInterval_Response", MessageInterval_ResponseTypeSupport)
}

type MessageInterval_Response struct {
	Success bool `yaml:"success"`
}

// NewMessageInterval_Response creates a new MessageInterval_Response with default values.
func NewMessageInterval_Response() *MessageInterval_Response {
	self := MessageInterval_Response{}
	self.SetDefaults()
	return &self
}

func (t *MessageInterval_Response) Clone() *MessageInterval_Response {
	c := &MessageInterval_Response{}
	c.Success = t.Success
	return c
}

func (t *MessageInterval_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MessageInterval_Response) SetDefaults() {
	t.Success = false
}

func (t *MessageInterval_Response) GetTypeSupport() types.MessageTypeSupport {
	return MessageInterval_ResponseTypeSupport
}

// MessageInterval_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MessageInterval_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewMessageInterval_ResponsePublisher creates and returns a new publisher for the
// MessageInterval_Response
func NewMessageInterval_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MessageInterval_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, MessageInterval_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MessageInterval_ResponsePublisher{pub}, nil
}

func (p *MessageInterval_ResponsePublisher) Publish(msg *MessageInterval_Response) error {
	return p.Publisher.Publish(msg)
}

// MessageInterval_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MessageInterval_ResponseSubscription struct {
	*rclgo.Subscription
}

// MessageInterval_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a MessageInterval_ResponseSubscription.
type MessageInterval_ResponseSubscriptionCallback func(msg *MessageInterval_Response, info *rclgo.MessageInfo, err error)

// NewMessageInterval_ResponseSubscription creates and returns a new subscription for the
// MessageInterval_Response
func NewMessageInterval_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MessageInterval_ResponseSubscriptionCallback) (*MessageInterval_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MessageInterval_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MessageInterval_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MessageInterval_ResponseSubscription{sub}, nil
}

func (s *MessageInterval_ResponseSubscription) TakeMessage(out *MessageInterval_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMessageInterval_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMessageInterval_ResponseSlice(dst, src []MessageInterval_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MessageInterval_ResponseTypeSupport types.MessageTypeSupport = _MessageInterval_ResponseTypeSupport{}

type _MessageInterval_ResponseTypeSupport struct{}

func (t _MessageInterval_ResponseTypeSupport) New() types.Message {
	return NewMessageInterval_Response()
}

func (t _MessageInterval_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__MessageInterval_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__MessageInterval_Response__create())
}

func (t _MessageInterval_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__MessageInterval_Response__destroy((*C.mavros_msgs__srv__MessageInterval_Response)(pointer_to_free))
}

func (t _MessageInterval_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MessageInterval_Response)
	mem := (*C.mavros_msgs__srv__MessageInterval_Response)(dst)
	mem.success = C.bool(m.Success)
}

func (t _MessageInterval_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MessageInterval_Response)
	mem := (*C.mavros_msgs__srv__MessageInterval_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
}

func (t _MessageInterval_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__MessageInterval_Response())
}

type CMessageInterval_Response = C.mavros_msgs__srv__MessageInterval_Response
type CMessageInterval_Response__Sequence = C.mavros_msgs__srv__MessageInterval_Response__Sequence

func MessageInterval_Response__Sequence_to_Go(goSlice *[]MessageInterval_Response, cSlice CMessageInterval_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MessageInterval_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MessageInterval_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MessageInterval_Response__Sequence_to_C(cSlice *CMessageInterval_Response__Sequence, goSlice []MessageInterval_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__MessageInterval_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__MessageInterval_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MessageInterval_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MessageInterval_Response__Array_to_Go(goSlice []MessageInterval_Response, cSlice []CMessageInterval_Response) {
	for i := 0; i < len(cSlice); i++ {
		MessageInterval_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MessageInterval_Response__Array_to_C(cSlice []CMessageInterval_Response, goSlice []MessageInterval_Response) {
	for i := 0; i < len(goSlice); i++ {
		MessageInterval_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
