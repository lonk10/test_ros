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

#include <geographic_msgs/msg/geo_pose_with_covariance_stamped.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geographic_msgs/GeoPoseWithCovarianceStamped", GeoPoseWithCovarianceStampedTypeSupport)
	typemap.RegisterMessage("geographic_msgs/msg/GeoPoseWithCovarianceStamped", GeoPoseWithCovarianceStampedTypeSupport)
}

type GeoPoseWithCovarianceStamped struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Pose GeoPoseWithCovariance `yaml:"pose"`
}

// NewGeoPoseWithCovarianceStamped creates a new GeoPoseWithCovarianceStamped with default values.
func NewGeoPoseWithCovarianceStamped() *GeoPoseWithCovarianceStamped {
	self := GeoPoseWithCovarianceStamped{}
	self.SetDefaults()
	return &self
}

func (t *GeoPoseWithCovarianceStamped) Clone() *GeoPoseWithCovarianceStamped {
	c := &GeoPoseWithCovarianceStamped{}
	c.Header = *t.Header.Clone()
	c.Pose = *t.Pose.Clone()
	return c
}

func (t *GeoPoseWithCovarianceStamped) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GeoPoseWithCovarianceStamped) SetDefaults() {
	t.Header.SetDefaults()
	t.Pose.SetDefaults()
}

func (t *GeoPoseWithCovarianceStamped) GetTypeSupport() types.MessageTypeSupport {
	return GeoPoseWithCovarianceStampedTypeSupport
}

// GeoPoseWithCovarianceStampedPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GeoPoseWithCovarianceStampedPublisher struct {
	*rclgo.Publisher
}

// NewGeoPoseWithCovarianceStampedPublisher creates and returns a new publisher for the
// GeoPoseWithCovarianceStamped
func NewGeoPoseWithCovarianceStampedPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GeoPoseWithCovarianceStampedPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GeoPoseWithCovarianceStampedTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GeoPoseWithCovarianceStampedPublisher{pub}, nil
}

func (p *GeoPoseWithCovarianceStampedPublisher) Publish(msg *GeoPoseWithCovarianceStamped) error {
	return p.Publisher.Publish(msg)
}

// GeoPoseWithCovarianceStampedSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GeoPoseWithCovarianceStampedSubscription struct {
	*rclgo.Subscription
}

// GeoPoseWithCovarianceStampedSubscriptionCallback type is used to provide a subscription
// handler function for a GeoPoseWithCovarianceStampedSubscription.
type GeoPoseWithCovarianceStampedSubscriptionCallback func(msg *GeoPoseWithCovarianceStamped, info *rclgo.MessageInfo, err error)

// NewGeoPoseWithCovarianceStampedSubscription creates and returns a new subscription for the
// GeoPoseWithCovarianceStamped
func NewGeoPoseWithCovarianceStampedSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GeoPoseWithCovarianceStampedSubscriptionCallback) (*GeoPoseWithCovarianceStampedSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GeoPoseWithCovarianceStamped
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GeoPoseWithCovarianceStampedTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GeoPoseWithCovarianceStampedSubscription{sub}, nil
}

func (s *GeoPoseWithCovarianceStampedSubscription) TakeMessage(out *GeoPoseWithCovarianceStamped) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGeoPoseWithCovarianceStampedSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGeoPoseWithCovarianceStampedSlice(dst, src []GeoPoseWithCovarianceStamped) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GeoPoseWithCovarianceStampedTypeSupport types.MessageTypeSupport = _GeoPoseWithCovarianceStampedTypeSupport{}

type _GeoPoseWithCovarianceStampedTypeSupport struct{}

func (t _GeoPoseWithCovarianceStampedTypeSupport) New() types.Message {
	return NewGeoPoseWithCovarianceStamped()
}

func (t _GeoPoseWithCovarianceStampedTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geographic_msgs__msg__GeoPoseWithCovarianceStamped
	return (unsafe.Pointer)(C.geographic_msgs__msg__GeoPoseWithCovarianceStamped__create())
}

func (t _GeoPoseWithCovarianceStampedTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geographic_msgs__msg__GeoPoseWithCovarianceStamped__destroy((*C.geographic_msgs__msg__GeoPoseWithCovarianceStamped)(pointer_to_free))
}

func (t _GeoPoseWithCovarianceStampedTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GeoPoseWithCovarianceStamped)
	mem := (*C.geographic_msgs__msg__GeoPoseWithCovarianceStamped)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	GeoPoseWithCovarianceTypeSupport.AsCStruct(unsafe.Pointer(&mem.pose), &m.Pose)
}

func (t _GeoPoseWithCovarianceStampedTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GeoPoseWithCovarianceStamped)
	mem := (*C.geographic_msgs__msg__GeoPoseWithCovarianceStamped)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	GeoPoseWithCovarianceTypeSupport.AsGoStruct(&m.Pose, unsafe.Pointer(&mem.pose))
}

func (t _GeoPoseWithCovarianceStampedTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geographic_msgs__msg__GeoPoseWithCovarianceStamped())
}

type CGeoPoseWithCovarianceStamped = C.geographic_msgs__msg__GeoPoseWithCovarianceStamped
type CGeoPoseWithCovarianceStamped__Sequence = C.geographic_msgs__msg__GeoPoseWithCovarianceStamped__Sequence

func GeoPoseWithCovarianceStamped__Sequence_to_Go(goSlice *[]GeoPoseWithCovarianceStamped, cSlice CGeoPoseWithCovarianceStamped__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GeoPoseWithCovarianceStamped, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GeoPoseWithCovarianceStampedTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GeoPoseWithCovarianceStamped__Sequence_to_C(cSlice *CGeoPoseWithCovarianceStamped__Sequence, goSlice []GeoPoseWithCovarianceStamped) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geographic_msgs__msg__GeoPoseWithCovarianceStamped)(C.malloc(C.sizeof_struct_geographic_msgs__msg__GeoPoseWithCovarianceStamped * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GeoPoseWithCovarianceStampedTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GeoPoseWithCovarianceStamped__Array_to_Go(goSlice []GeoPoseWithCovarianceStamped, cSlice []CGeoPoseWithCovarianceStamped) {
	for i := 0; i < len(cSlice); i++ {
		GeoPoseWithCovarianceStampedTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GeoPoseWithCovarianceStamped__Array_to_C(cSlice []CGeoPoseWithCovarianceStamped, goSlice []GeoPoseWithCovarianceStamped) {
	for i := 0; i < len(goSlice); i++ {
		GeoPoseWithCovarianceStampedTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
