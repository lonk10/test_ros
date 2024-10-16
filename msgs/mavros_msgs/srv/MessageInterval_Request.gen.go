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
	typemap.RegisterMessage("mavros_msgs/MessageInterval_Request", MessageInterval_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/MessageInterval_Request", MessageInterval_RequestTypeSupport)
}

type MessageInterval_Request struct {
	MessageId uint32 `yaml:"message_id"`
	MessageRate float32 `yaml:"message_rate"`
}

// NewMessageInterval_Request creates a new MessageInterval_Request with default values.
func NewMessageInterval_Request() *MessageInterval_Request {
	self := MessageInterval_Request{}
	self.SetDefaults()
	return &self
}

func (t *MessageInterval_Request) Clone() *MessageInterval_Request {
	c := &MessageInterval_Request{}
	c.MessageId = t.MessageId
	c.MessageRate = t.MessageRate
	return c
}

func (t *MessageInterval_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MessageInterval_Request) SetDefaults() {
	t.MessageId = 0
	t.MessageRate = 0
}

func (t *MessageInterval_Request) GetTypeSupport() types.MessageTypeSupport {
	return MessageInterval_RequestTypeSupport
}

// MessageInterval_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MessageInterval_RequestPublisher struct {
	*rclgo.Publisher
}

// NewMessageInterval_RequestPublisher creates and returns a new publisher for the
// MessageInterval_Request
func NewMessageInterval_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MessageInterval_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, MessageInterval_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MessageInterval_RequestPublisher{pub}, nil
}

func (p *MessageInterval_RequestPublisher) Publish(msg *MessageInterval_Request) error {
	return p.Publisher.Publish(msg)
}

// MessageInterval_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MessageInterval_RequestSubscription struct {
	*rclgo.Subscription
}

// MessageInterval_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a MessageInterval_RequestSubscription.
type MessageInterval_RequestSubscriptionCallback func(msg *MessageInterval_Request, info *rclgo.MessageInfo, err error)

// NewMessageInterval_RequestSubscription creates and returns a new subscription for the
// MessageInterval_Request
func NewMessageInterval_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MessageInterval_RequestSubscriptionCallback) (*MessageInterval_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MessageInterval_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MessageInterval_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MessageInterval_RequestSubscription{sub}, nil
}

func (s *MessageInterval_RequestSubscription) TakeMessage(out *MessageInterval_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMessageInterval_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMessageInterval_RequestSlice(dst, src []MessageInterval_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MessageInterval_RequestTypeSupport types.MessageTypeSupport = _MessageInterval_RequestTypeSupport{}

type _MessageInterval_RequestTypeSupport struct{}

func (t _MessageInterval_RequestTypeSupport) New() types.Message {
	return NewMessageInterval_Request()
}

func (t _MessageInterval_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__MessageInterval_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__MessageInterval_Request__create())
}

func (t _MessageInterval_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__MessageInterval_Request__destroy((*C.mavros_msgs__srv__MessageInterval_Request)(pointer_to_free))
}

func (t _MessageInterval_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MessageInterval_Request)
	mem := (*C.mavros_msgs__srv__MessageInterval_Request)(dst)
	mem.message_id = C.uint32_t(m.MessageId)
	mem.message_rate = C.float(m.MessageRate)
}

func (t _MessageInterval_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MessageInterval_Request)
	mem := (*C.mavros_msgs__srv__MessageInterval_Request)(ros2_message_buffer)
	m.MessageId = uint32(mem.message_id)
	m.MessageRate = float32(mem.message_rate)
}

func (t _MessageInterval_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__MessageInterval_Request())
}

type CMessageInterval_Request = C.mavros_msgs__srv__MessageInterval_Request
type CMessageInterval_Request__Sequence = C.mavros_msgs__srv__MessageInterval_Request__Sequence

func MessageInterval_Request__Sequence_to_Go(goSlice *[]MessageInterval_Request, cSlice CMessageInterval_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MessageInterval_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MessageInterval_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MessageInterval_Request__Sequence_to_C(cSlice *CMessageInterval_Request__Sequence, goSlice []MessageInterval_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__MessageInterval_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__MessageInterval_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MessageInterval_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MessageInterval_Request__Array_to_Go(goSlice []MessageInterval_Request, cSlice []CMessageInterval_Request) {
	for i := 0; i < len(cSlice); i++ {
		MessageInterval_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MessageInterval_Request__Array_to_C(cSlice []CMessageInterval_Request, goSlice []MessageInterval_Request) {
	for i := 0; i < len(goSlice); i++ {
		MessageInterval_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
