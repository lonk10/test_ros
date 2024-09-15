// Code generated by rclgo-gen. DO NOT EDIT.

package pcl_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <pcl_msgs/msg/point_indices.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("pcl_msgs/PointIndices", PointIndicesTypeSupport)
	typemap.RegisterMessage("pcl_msgs/msg/PointIndices", PointIndicesTypeSupport)
}

type PointIndices struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Indices []int32 `yaml:"indices"`
}

// NewPointIndices creates a new PointIndices with default values.
func NewPointIndices() *PointIndices {
	self := PointIndices{}
	self.SetDefaults()
	return &self
}

func (t *PointIndices) Clone() *PointIndices {
	c := &PointIndices{}
	c.Header = *t.Header.Clone()
	if t.Indices != nil {
		c.Indices = make([]int32, len(t.Indices))
		copy(c.Indices, t.Indices)
	}
	return c
}

func (t *PointIndices) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PointIndices) SetDefaults() {
	t.Header.SetDefaults()
	t.Indices = nil
}

func (t *PointIndices) GetTypeSupport() types.MessageTypeSupport {
	return PointIndicesTypeSupport
}

// PointIndicesPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PointIndicesPublisher struct {
	*rclgo.Publisher
}

// NewPointIndicesPublisher creates and returns a new publisher for the
// PointIndices
func NewPointIndicesPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PointIndicesPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PointIndicesTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PointIndicesPublisher{pub}, nil
}

func (p *PointIndicesPublisher) Publish(msg *PointIndices) error {
	return p.Publisher.Publish(msg)
}

// PointIndicesSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PointIndicesSubscription struct {
	*rclgo.Subscription
}

// PointIndicesSubscriptionCallback type is used to provide a subscription
// handler function for a PointIndicesSubscription.
type PointIndicesSubscriptionCallback func(msg *PointIndices, info *rclgo.MessageInfo, err error)

// NewPointIndicesSubscription creates and returns a new subscription for the
// PointIndices
func NewPointIndicesSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback PointIndicesSubscriptionCallback) (*PointIndicesSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PointIndices
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PointIndicesTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &PointIndicesSubscription{sub}, nil
}

func (s *PointIndicesSubscription) TakeMessage(out *PointIndices) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePointIndicesSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePointIndicesSlice(dst, src []PointIndices) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PointIndicesTypeSupport types.MessageTypeSupport = _PointIndicesTypeSupport{}

type _PointIndicesTypeSupport struct{}

func (t _PointIndicesTypeSupport) New() types.Message {
	return NewPointIndices()
}

func (t _PointIndicesTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.pcl_msgs__msg__PointIndices
	return (unsafe.Pointer)(C.pcl_msgs__msg__PointIndices__create())
}

func (t _PointIndicesTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.pcl_msgs__msg__PointIndices__destroy((*C.pcl_msgs__msg__PointIndices)(pointer_to_free))
}

func (t _PointIndicesTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*PointIndices)
	mem := (*C.pcl_msgs__msg__PointIndices)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.Int32__Sequence_to_C((*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.indices)), m.Indices)
}

func (t _PointIndicesTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*PointIndices)
	mem := (*C.pcl_msgs__msg__PointIndices)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.Int32__Sequence_to_Go(&m.Indices, *(*primitives.CInt32__Sequence)(unsafe.Pointer(&mem.indices)))
}

func (t _PointIndicesTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__pcl_msgs__msg__PointIndices())
}

type CPointIndices = C.pcl_msgs__msg__PointIndices
type CPointIndices__Sequence = C.pcl_msgs__msg__PointIndices__Sequence

func PointIndices__Sequence_to_Go(goSlice *[]PointIndices, cSlice CPointIndices__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PointIndices, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PointIndicesTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PointIndices__Sequence_to_C(cSlice *CPointIndices__Sequence, goSlice []PointIndices) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.pcl_msgs__msg__PointIndices)(C.malloc(C.sizeof_struct_pcl_msgs__msg__PointIndices * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PointIndicesTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PointIndices__Array_to_Go(goSlice []PointIndices, cSlice []CPointIndices) {
	for i := 0; i < len(cSlice); i++ {
		PointIndicesTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PointIndices__Array_to_C(cSlice []CPointIndices, goSlice []PointIndices) {
	for i := 0; i < len(goSlice); i++ {
		PointIndicesTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
