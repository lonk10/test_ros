// Code generated by rclgo-gen. DO NOT EDIT.

package geometry_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geometry_msgs/msg/point_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geometry_msgs/PointStamped", PointStampedTypeSupport)
	typemap.RegisterMessage("geometry_msgs/msg/PointStamped", PointStampedTypeSupport)
}

type PointStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Point Point `yaml:"point"`
}

// NewPointStamped creates a new PointStamped with default values.
func NewPointStamped() *PointStamped {
	self := PointStamped{}
	self.SetDefaults()
	return &self
}

func (t *PointStamped) Clone() *PointStamped {
	c := &PointStamped{}
	c.Header = *t.Header.Clone()
	c.Point = *t.Point.Clone()
	return c
}

func (t *PointStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PointStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Point.SetDefaults()
}

func (t *PointStamped) GetTypeSupport() types.MessageTypeSupport {
	return PointStampedTypeSupport
}

// PointStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PointStampedPublisher struct {
	*rclgo.Publisher
}

// NewPointStampedPublisher creates and returns a new publisher for the
// PointStamped
func NewPointStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PointStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PointStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PointStampedPublisher{pub}, nil
}

func (p *PointStampedPublisher) Publish(msg *PointStamped) error {
	return p.Publisher.Publish(msg)
}

// PointStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PointStampedSubscription struct {
	*rclgo.Subscription
}

// PointStampedSubscriptionCallback type is used to provide a subscription
// handler function for a PointStampedSubscription.
type PointStampedSubscriptionCallback func(msg *PointStamped, info *rclgo.MessageInfo, err error)

// NewPointStampedSubscription creates and returns a new subscription for the
// PointStamped
func NewPointStampedSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback PointStampedSubscriptionCallback) (*PointStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PointStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PointStampedTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &PointStampedSubscription{sub}, nil
}

func (s *PointStampedSubscription) TakeMessage(out *PointStamped) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePointStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePointStampedSlice(dst, src []PointStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PointStampedTypeSupport types.MessageTypeSupport = _PointStampedTypeSupport{}

type _PointStampedTypeSupport struct{}

func (t _PointStampedTypeSupport) New() types.Message {
	return NewPointStamped()
}

func (t _PointStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geometry_msgs__msg__PointStamped
	return (unsafe.Pointer)(C.geometry_msgs__msg__PointStamped__create())
}

func (t _PointStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geometry_msgs__msg__PointStamped__destroy((*C.geometry_msgs__msg__PointStamped)(pointer_to_free))
}

func (t _PointStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PointStamped)
	mem := (*C.geometry_msgs__msg__PointStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	PointTypeSupport.AsCStruct(unsafe.Pointer(&mem.point), &m.Point)
}

func (t _PointStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PointStamped)
	mem := (*C.geometry_msgs__msg__PointStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	PointTypeSupport.AsGoStruct(&m.Point, unsafe.Pointer(&mem.point))
}

func (t _PointStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geometry_msgs__msg__PointStamped())
}

type CPointStamped = C.geometry_msgs__msg__PointStamped
type CPointStamped__Sequence = C.geometry_msgs__msg__PointStamped__Sequence

func PointStamped__Sequence_to_Go(goSlice *[]PointStamped, cSlice CPointStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PointStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PointStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PointStamped__Sequence_to_C(cSlice *CPointStamped__Sequence, goSlice []PointStamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geometry_msgs__msg__PointStamped)(C.malloc(C.sizeof_struct_geometry_msgs__msg__PointStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PointStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PointStamped__Array_to_Go(goSlice []PointStamped, cSlice []CPointStamped) {
	for i := 0; i < len(cSlice); i++ {
		PointStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PointStamped__Array_to_C(cSlice []CPointStamped, goSlice []PointStamped) {
	for i := 0; i < len(goSlice); i++ {
		PointStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
