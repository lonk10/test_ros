// Code generated by rclgo-gen. DO NOT EDIT.

package turtlesim_action
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/action/rotate_absolute.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/RotateAbsolute_Feedback", RotateAbsolute_FeedbackTypeSupport)
	typemap.RegisterMessage("turtlesim/action/RotateAbsolute_Feedback", RotateAbsolute_FeedbackTypeSupport)
}

type RotateAbsolute_Feedback struct {
	Remaining float32 `yaml:"remaining"`// The remaining rotation in radians
}

// NewRotateAbsolute_Feedback creates a new RotateAbsolute_Feedback with default values.
func NewRotateAbsolute_Feedback() *RotateAbsolute_Feedback {
	self := RotateAbsolute_Feedback{}
	self.SetDefaults()
	return &self
}

func (t *RotateAbsolute_Feedback) Clone() *RotateAbsolute_Feedback {
	c := &RotateAbsolute_Feedback{}
	c.Remaining = t.Remaining
	return c
}

func (t *RotateAbsolute_Feedback) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RotateAbsolute_Feedback) SetDefaults() {
	t.Remaining = 0
}

func (t *RotateAbsolute_Feedback) GetTypeSupport() types.MessageTypeSupport {
	return RotateAbsolute_FeedbackTypeSupport
}

// RotateAbsolute_FeedbackPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type RotateAbsolute_FeedbackPublisher struct {
	*rclgo.Publisher
}

// NewRotateAbsolute_FeedbackPublisher creates and returns a new publisher for the
// RotateAbsolute_Feedback
func NewRotateAbsolute_FeedbackPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*RotateAbsolute_FeedbackPublisher, error) {
	pub, err := node.NewPublisher(topic_name, RotateAbsolute_FeedbackTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &RotateAbsolute_FeedbackPublisher{pub}, nil
}

func (p *RotateAbsolute_FeedbackPublisher) Publish(msg *RotateAbsolute_Feedback) error {
	return p.Publisher.Publish(msg)
}

// RotateAbsolute_FeedbackSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type RotateAbsolute_FeedbackSubscription struct {
	*rclgo.Subscription
}

// RotateAbsolute_FeedbackSubscriptionCallback type is used to provide a subscription
// handler function for a RotateAbsolute_FeedbackSubscription.
type RotateAbsolute_FeedbackSubscriptionCallback func(msg *RotateAbsolute_Feedback, info *rclgo.MessageInfo, err error)

// NewRotateAbsolute_FeedbackSubscription creates and returns a new subscription for the
// RotateAbsolute_Feedback
func NewRotateAbsolute_FeedbackSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback RotateAbsolute_FeedbackSubscriptionCallback) (*RotateAbsolute_FeedbackSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg RotateAbsolute_Feedback
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, RotateAbsolute_FeedbackTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &RotateAbsolute_FeedbackSubscription{sub}, nil
}

func (s *RotateAbsolute_FeedbackSubscription) TakeMessage(out *RotateAbsolute_Feedback) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneRotateAbsolute_FeedbackSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRotateAbsolute_FeedbackSlice(dst, src []RotateAbsolute_Feedback) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RotateAbsolute_FeedbackTypeSupport types.MessageTypeSupport = _RotateAbsolute_FeedbackTypeSupport{}

type _RotateAbsolute_FeedbackTypeSupport struct{}

func (t _RotateAbsolute_FeedbackTypeSupport) New() types.Message {
	return NewRotateAbsolute_Feedback()
}

func (t _RotateAbsolute_FeedbackTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__action__RotateAbsolute_Feedback
	return (unsafe.Pointer)(C.turtlesim__action__RotateAbsolute_Feedback__create())
}

func (t _RotateAbsolute_FeedbackTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__action__RotateAbsolute_Feedback__destroy((*C.turtlesim__action__RotateAbsolute_Feedback)(pointer_to_free))
}

func (t _RotateAbsolute_FeedbackTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RotateAbsolute_Feedback)
	mem := (*C.turtlesim__action__RotateAbsolute_Feedback)(dst)
	mem.remaining = C.float(m.Remaining)
}

func (t _RotateAbsolute_FeedbackTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RotateAbsolute_Feedback)
	mem := (*C.turtlesim__action__RotateAbsolute_Feedback)(ros2_message_buffer)
	m.Remaining = float32(mem.remaining)
}

func (t _RotateAbsolute_FeedbackTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__action__RotateAbsolute_Feedback())
}

type CRotateAbsolute_Feedback = C.turtlesim__action__RotateAbsolute_Feedback
type CRotateAbsolute_Feedback__Sequence = C.turtlesim__action__RotateAbsolute_Feedback__Sequence

func RotateAbsolute_Feedback__Sequence_to_Go(goSlice *[]RotateAbsolute_Feedback, cSlice CRotateAbsolute_Feedback__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RotateAbsolute_Feedback, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		RotateAbsolute_FeedbackTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func RotateAbsolute_Feedback__Sequence_to_C(cSlice *CRotateAbsolute_Feedback__Sequence, goSlice []RotateAbsolute_Feedback) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.turtlesim__action__RotateAbsolute_Feedback)(C.malloc(C.sizeof_struct_turtlesim__action__RotateAbsolute_Feedback * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		RotateAbsolute_FeedbackTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func RotateAbsolute_Feedback__Array_to_Go(goSlice []RotateAbsolute_Feedback, cSlice []CRotateAbsolute_Feedback) {
	for i := 0; i < len(cSlice); i++ {
		RotateAbsolute_FeedbackTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RotateAbsolute_Feedback__Array_to_C(cSlice []CRotateAbsolute_Feedback, goSlice []RotateAbsolute_Feedback) {
	for i := 0; i < len(goSlice); i++ {
		RotateAbsolute_FeedbackTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
