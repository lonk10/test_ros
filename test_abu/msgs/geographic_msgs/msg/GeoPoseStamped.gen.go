// Code generated by rclgo-gen. DO NOT EDIT.

package geographic_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geographic_msgs/msg/geo_pose_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geographic_msgs/GeoPoseStamped", GeoPoseStampedTypeSupport)
	typemap.RegisterMessage("geographic_msgs/msg/GeoPoseStamped", GeoPoseStampedTypeSupport)
}

type GeoPoseStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Pose GeoPose `yaml:"pose"`
}

// NewGeoPoseStamped creates a new GeoPoseStamped with default values.
func NewGeoPoseStamped() *GeoPoseStamped {
	self := GeoPoseStamped{}
	self.SetDefaults()
	return &self
}

func (t *GeoPoseStamped) Clone() *GeoPoseStamped {
	c := &GeoPoseStamped{}
	c.Header = *t.Header.Clone()
	c.Pose = *t.Pose.Clone()
	return c
}

func (t *GeoPoseStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GeoPoseStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Pose.SetDefaults()
}

func (t *GeoPoseStamped) GetTypeSupport() types.MessageTypeSupport {
	return GeoPoseStampedTypeSupport
}

// GeoPoseStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GeoPoseStampedPublisher struct {
	*rclgo.Publisher
}

// NewGeoPoseStampedPublisher creates and returns a new publisher for the
// GeoPoseStamped
func NewGeoPoseStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GeoPoseStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GeoPoseStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GeoPoseStampedPublisher{pub}, nil
}

func (p *GeoPoseStampedPublisher) Publish(msg *GeoPoseStamped) error {
	return p.Publisher.Publish(msg)
}

// GeoPoseStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GeoPoseStampedSubscription struct {
	*rclgo.Subscription
}

// GeoPoseStampedSubscriptionCallback type is used to provide a subscription
// handler function for a GeoPoseStampedSubscription.
type GeoPoseStampedSubscriptionCallback func(msg *GeoPoseStamped, info *rclgo.MessageInfo, err error)

// NewGeoPoseStampedSubscription creates and returns a new subscription for the
// GeoPoseStamped
func NewGeoPoseStampedSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GeoPoseStampedSubscriptionCallback) (*GeoPoseStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GeoPoseStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GeoPoseStampedTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GeoPoseStampedSubscription{sub}, nil
}

func (s *GeoPoseStampedSubscription) TakeMessage(out *GeoPoseStamped) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGeoPoseStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGeoPoseStampedSlice(dst, src []GeoPoseStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GeoPoseStampedTypeSupport types.MessageTypeSupport = _GeoPoseStampedTypeSupport{}

type _GeoPoseStampedTypeSupport struct{}

func (t _GeoPoseStampedTypeSupport) New() types.Message {
	return NewGeoPoseStamped()
}

func (t _GeoPoseStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geographic_msgs__msg__GeoPoseStamped
	return (unsafe.Pointer)(C.geographic_msgs__msg__GeoPoseStamped__create())
}

func (t _GeoPoseStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geographic_msgs__msg__GeoPoseStamped__destroy((*C.geographic_msgs__msg__GeoPoseStamped)(pointer_to_free))
}

func (t _GeoPoseStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GeoPoseStamped)
	mem := (*C.geographic_msgs__msg__GeoPoseStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	GeoPoseTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
}

func (t _GeoPoseStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GeoPoseStamped)
	mem := (*C.geographic_msgs__msg__GeoPoseStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	GeoPoseTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
}

func (t _GeoPoseStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geographic_msgs__msg__GeoPoseStamped())
}

type CGeoPoseStamped = C.geographic_msgs__msg__GeoPoseStamped
type CGeoPoseStamped__Sequence = C.geographic_msgs__msg__GeoPoseStamped__Sequence

func GeoPoseStamped__Sequence_to_Go(goSlice *[]GeoPoseStamped, cSlice CGeoPoseStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GeoPoseStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GeoPoseStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GeoPoseStamped__Sequence_to_C(cSlice *CGeoPoseStamped__Sequence, goSlice []GeoPoseStamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geographic_msgs__msg__GeoPoseStamped)(C.malloc(C.sizeof_struct_geographic_msgs__msg__GeoPoseStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GeoPoseStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GeoPoseStamped__Array_to_Go(goSlice []GeoPoseStamped, cSlice []CGeoPoseStamped) {
	for i := 0; i < len(cSlice); i++ {
		GeoPoseStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GeoPoseStamped__Array_to_C(cSlice []CGeoPoseStamped, goSlice []GeoPoseStamped) {
	for i := 0; i < len(goSlice); i++ {
		GeoPoseStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
