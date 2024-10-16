// Code generated by rclgo-gen. DO NOT EDIT.

package actionlib_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	builtin_interfaces_msg "test/msgs/builtin_interfaces/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <actionlib_msgs/msg/goal_id.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("actionlib_msgs/GoalID", GoalIDTypeSupport)
	typemap.RegisterMessage("actionlib_msgs/msg/GoalID", GoalIDTypeSupport)
}

type GoalID struct {
	Stamp builtin_interfaces_msg.Time `yaml:"stamp"`// The stamp should store the time at which this goal was requested.It is used by an action server when it tries to preempt allgoals that were requested before a certain time
	Id string `yaml:"id"`// The id provides a way to associate feedback andresult message with specific goal requests. The idspecified must be unique.
}

// NewGoalID creates a new GoalID with default values.
func NewGoalID() *GoalID {
	self := GoalID{}
	self.SetDefaults()
	return &self
}

func (t *GoalID) Clone() *GoalID {
	c := &GoalID{}
	c.Stamp = *t.Stamp.Clone()
	c.Id = t.Id
	return c
}

func (t *GoalID) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GoalID) SetDefaults() {
	t.Stamp.SetDefaults()
	t.Id = ""
}

func (t *GoalID) GetTypeSupport() types.MessageTypeSupport {
	return GoalIDTypeSupport
}

// GoalIDPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GoalIDPublisher struct {
	*rclgo.Publisher
}

// NewGoalIDPublisher creates and returns a new publisher for the
// GoalID
func NewGoalIDPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GoalIDPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GoalIDTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GoalIDPublisher{pub}, nil
}

func (p *GoalIDPublisher) Publish(msg *GoalID) error {
	return p.Publisher.Publish(msg)
}

// GoalIDSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GoalIDSubscription struct {
	*rclgo.Subscription
}

// GoalIDSubscriptionCallback type is used to provide a subscription
// handler function for a GoalIDSubscription.
type GoalIDSubscriptionCallback func(msg *GoalID, info *rclgo.MessageInfo, err error)

// NewGoalIDSubscription creates and returns a new subscription for the
// GoalID
func NewGoalIDSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GoalIDSubscriptionCallback) (*GoalIDSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GoalID
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GoalIDTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GoalIDSubscription{sub}, nil
}

func (s *GoalIDSubscription) TakeMessage(out *GoalID) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGoalIDSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGoalIDSlice(dst, src []GoalID) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GoalIDTypeSupport types.MessageTypeSupport = _GoalIDTypeSupport{}

type _GoalIDTypeSupport struct{}

func (t _GoalIDTypeSupport) New() types.Message {
	return NewGoalID()
}

func (t _GoalIDTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.actionlib_msgs__msg__GoalID
	return (unsafe.Pointer)(C.actionlib_msgs__msg__GoalID__create())
}

func (t _GoalIDTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.actionlib_msgs__msg__GoalID__destroy((*C.actionlib_msgs__msg__GoalID)(pointer_to_free))
}

func (t _GoalIDTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GoalID)
	mem := (*C.actionlib_msgs__msg__GoalID)(dst)
	builtin_interfaces_msg.TimeTypeSupport.AsCStruct(unsafe.Pointer(&mem.stamp), &m.Stamp)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.id), m.Id)
}

func (t _GoalIDTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GoalID)
	mem := (*C.actionlib_msgs__msg__GoalID)(ros2_message_buffer)
	builtin_interfaces_msg.TimeTypeSupport.AsGoStruct(&m.Stamp, unsafe.Pointer(&mem.stamp))
	primitives.StringAsGoStruct(&m.Id, unsafe.Pointer(&mem.id))
}

func (t _GoalIDTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__actionlib_msgs__msg__GoalID())
}

type CGoalID = C.actionlib_msgs__msg__GoalID
type CGoalID__Sequence = C.actionlib_msgs__msg__GoalID__Sequence

func GoalID__Sequence_to_Go(goSlice *[]GoalID, cSlice CGoalID__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GoalID, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GoalIDTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GoalID__Sequence_to_C(cSlice *CGoalID__Sequence, goSlice []GoalID) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.actionlib_msgs__msg__GoalID)(C.malloc(C.sizeof_struct_actionlib_msgs__msg__GoalID * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GoalIDTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GoalID__Array_to_Go(goSlice []GoalID, cSlice []CGoalID) {
	for i := 0; i < len(cSlice); i++ {
		GoalIDTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GoalID__Array_to_C(cSlice []CGoalID, goSlice []GoalID) {
	for i := 0; i < len(goSlice); i++ {
		GoalIDTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
