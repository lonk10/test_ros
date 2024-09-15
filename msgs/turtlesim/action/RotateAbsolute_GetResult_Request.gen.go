// Code generated by rclgo-gen. DO NOT EDIT.

package turtlesim_action
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	unique_identifier_msgs_msg "test/msgs/unique_identifier_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/action/rotate_absolute.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/RotateAbsolute_GetResult_Request", RotateAbsolute_GetResult_RequestTypeSupport)
	typemap.RegisterMessage("turtlesim/action/RotateAbsolute_GetResult_Request", RotateAbsolute_GetResult_RequestTypeSupport)
}

type RotateAbsolute_GetResult_Request struct {
	GoalID unique_identifier_msgs_msg.UUID `yaml:"goal_id"`
}

// NewRotateAbsolute_GetResult_Request creates a new RotateAbsolute_GetResult_Request with default values.
func NewRotateAbsolute_GetResult_Request() *RotateAbsolute_GetResult_Request {
	self := RotateAbsolute_GetResult_Request{}
	self.SetDefaults()
	return &self
}

func (t *RotateAbsolute_GetResult_Request) Clone() *RotateAbsolute_GetResult_Request {
	c := &RotateAbsolute_GetResult_Request{}
	c.GoalID = *t.GoalID.Clone()
	return c
}

func (t *RotateAbsolute_GetResult_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RotateAbsolute_GetResult_Request) SetDefaults() {
	t.GoalID.SetDefaults()
}

func (t *RotateAbsolute_GetResult_Request) GetTypeSupport() types.MessageTypeSupport {
	return RotateAbsolute_GetResult_RequestTypeSupport
}
func (t *RotateAbsolute_GetResult_Request) GetGoalID() *types.GoalID {
	return (*types.GoalID)(&t.GoalID.Uuid)
}

func (t *RotateAbsolute_GetResult_Request) SetGoalID(id *types.GoalID) {
	t.GoalID.Uuid = *id
}

// RotateAbsolute_GetResult_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type RotateAbsolute_GetResult_RequestPublisher struct {
	*rclgo.Publisher
}

// NewRotateAbsolute_GetResult_RequestPublisher creates and returns a new publisher for the
// RotateAbsolute_GetResult_Request
func NewRotateAbsolute_GetResult_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*RotateAbsolute_GetResult_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, RotateAbsolute_GetResult_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &RotateAbsolute_GetResult_RequestPublisher{pub}, nil
}

func (p *RotateAbsolute_GetResult_RequestPublisher) Publish(msg *RotateAbsolute_GetResult_Request) error {
	return p.Publisher.Publish(msg)
}

// RotateAbsolute_GetResult_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type RotateAbsolute_GetResult_RequestSubscription struct {
	*rclgo.Subscription
}

// RotateAbsolute_GetResult_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a RotateAbsolute_GetResult_RequestSubscription.
type RotateAbsolute_GetResult_RequestSubscriptionCallback func(msg *RotateAbsolute_GetResult_Request, info *rclgo.MessageInfo, err error)

// NewRotateAbsolute_GetResult_RequestSubscription creates and returns a new subscription for the
// RotateAbsolute_GetResult_Request
func NewRotateAbsolute_GetResult_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback RotateAbsolute_GetResult_RequestSubscriptionCallback) (*RotateAbsolute_GetResult_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg RotateAbsolute_GetResult_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, RotateAbsolute_GetResult_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &RotateAbsolute_GetResult_RequestSubscription{sub}, nil
}

func (s *RotateAbsolute_GetResult_RequestSubscription) TakeMessage(out *RotateAbsolute_GetResult_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneRotateAbsolute_GetResult_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRotateAbsolute_GetResult_RequestSlice(dst, src []RotateAbsolute_GetResult_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RotateAbsolute_GetResult_RequestTypeSupport types.MessageTypeSupport = _RotateAbsolute_GetResult_RequestTypeSupport{}

type _RotateAbsolute_GetResult_RequestTypeSupport struct{}

func (t _RotateAbsolute_GetResult_RequestTypeSupport) New() types.Message {
	return NewRotateAbsolute_GetResult_Request()
}

func (t _RotateAbsolute_GetResult_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__action__RotateAbsolute_GetResult_Request
	return (unsafe.Pointer)(C.turtlesim__action__RotateAbsolute_GetResult_Request__create())
}

func (t _RotateAbsolute_GetResult_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__action__RotateAbsolute_GetResult_Request__destroy((*C.turtlesim__action__RotateAbsolute_GetResult_Request)(pointer_to_free))
}

func (t _RotateAbsolute_GetResult_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RotateAbsolute_GetResult_Request)
	mem := (*C.turtlesim__action__RotateAbsolute_GetResult_Request)(dst)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.goal_id), &m.GoalID)
}

func (t _RotateAbsolute_GetResult_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RotateAbsolute_GetResult_Request)
	mem := (*C.turtlesim__action__RotateAbsolute_GetResult_Request)(ros2_message_buffer)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.GoalID, unsafe.Pointer(&mem.goal_id))
}

func (t _RotateAbsolute_GetResult_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__action__RotateAbsolute_GetResult_Request())
}

type CRotateAbsolute_GetResult_Request = C.turtlesim__action__RotateAbsolute_GetResult_Request
type CRotateAbsolute_GetResult_Request__Sequence = C.turtlesim__action__RotateAbsolute_GetResult_Request__Sequence

func RotateAbsolute_GetResult_Request__Sequence_to_Go(goSlice *[]RotateAbsolute_GetResult_Request, cSlice CRotateAbsolute_GetResult_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RotateAbsolute_GetResult_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		RotateAbsolute_GetResult_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func RotateAbsolute_GetResult_Request__Sequence_to_C(cSlice *CRotateAbsolute_GetResult_Request__Sequence, goSlice []RotateAbsolute_GetResult_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.turtlesim__action__RotateAbsolute_GetResult_Request)(C.malloc(C.sizeof_struct_turtlesim__action__RotateAbsolute_GetResult_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		RotateAbsolute_GetResult_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func RotateAbsolute_GetResult_Request__Array_to_Go(goSlice []RotateAbsolute_GetResult_Request, cSlice []CRotateAbsolute_GetResult_Request) {
	for i := 0; i < len(cSlice); i++ {
		RotateAbsolute_GetResult_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RotateAbsolute_GetResult_Request__Array_to_C(cSlice []CRotateAbsolute_GetResult_Request, goSlice []RotateAbsolute_GetResult_Request) {
	for i := 0; i < len(goSlice); i++ {
		RotateAbsolute_GetResult_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
