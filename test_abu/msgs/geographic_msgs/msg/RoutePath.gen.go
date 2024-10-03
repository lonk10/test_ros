// Code generated by rclgo-gen. DO NOT EDIT.

package geographic_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "test/msgs/std_msgs/msg"
	unique_identifier_msgs_msg "test/msgs/unique_identifier_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geographic_msgs/msg/route_path.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geographic_msgs/RoutePath", RoutePathTypeSupport)
	typemap.RegisterMessage("geographic_msgs/msg/RoutePath", RoutePathTypeSupport)
}

type RoutePath struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Network unique_identifier_msgs_msg.UUID `yaml:"network"`// Route network containing this path
	Segments []unique_identifier_msgs_msg.UUID `yaml:"segments"`// Sequence of RouteSegment IDs
	Props []KeyValue `yaml:"props"`// Key/value properties
}

// NewRoutePath creates a new RoutePath with default values.
func NewRoutePath() *RoutePath {
	self := RoutePath{}
	self.SetDefaults()
	return &self
}

func (t *RoutePath) Clone() *RoutePath {
	c := &RoutePath{}
	c.Header = *t.Header.Clone()
	c.Network = *t.Network.Clone()
	if t.Segments != nil {
		c.Segments = make([]unique_identifier_msgs_msg.UUID, len(t.Segments))
		unique_identifier_msgs_msg.CloneUUIDSlice(c.Segments, t.Segments)
	}
	if t.Props != nil {
		c.Props = make([]KeyValue, len(t.Props))
		CloneKeyValueSlice(c.Props, t.Props)
	}
	return c
}

func (t *RoutePath) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RoutePath) SetDefaults() {
	t.Header.SetDefaults()
	t.Network.SetDefaults()
	t.Segments = nil
	t.Props = nil
}

func (t *RoutePath) GetTypeSupport() types.MessageTypeSupport {
	return RoutePathTypeSupport
}

// RoutePathPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type RoutePathPublisher struct {
	*rclgo.Publisher
}

// NewRoutePathPublisher creates and returns a new publisher for the
// RoutePath
func NewRoutePathPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*RoutePathPublisher, error) {
	pub, err := node.NewPublisher(topic_name, RoutePathTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &RoutePathPublisher{pub}, nil
}

func (p *RoutePathPublisher) Publish(msg *RoutePath) error {
	return p.Publisher.Publish(msg)
}

// RoutePathSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type RoutePathSubscription struct {
	*rclgo.Subscription
}

// RoutePathSubscriptionCallback type is used to provide a subscription
// handler function for a RoutePathSubscription.
type RoutePathSubscriptionCallback func(msg *RoutePath, info *rclgo.MessageInfo, err error)

// NewRoutePathSubscription creates and returns a new subscription for the
// RoutePath
func NewRoutePathSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback RoutePathSubscriptionCallback) (*RoutePathSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg RoutePath
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, RoutePathTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &RoutePathSubscription{sub}, nil
}

func (s *RoutePathSubscription) TakeMessage(out *RoutePath) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneRoutePathSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRoutePathSlice(dst, src []RoutePath) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RoutePathTypeSupport types.MessageTypeSupport = _RoutePathTypeSupport{}

type _RoutePathTypeSupport struct{}

func (t _RoutePathTypeSupport) New() types.Message {
	return NewRoutePath()
}

func (t _RoutePathTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geographic_msgs__msg__RoutePath
	return (unsafe.Pointer)(C.geographic_msgs__msg__RoutePath__create())
}

func (t _RoutePathTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geographic_msgs__msg__RoutePath__destroy((*C.geographic_msgs__msg__RoutePath)(pointer_to_free))
}

func (t _RoutePathTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RoutePath)
	mem := (*C.geographic_msgs__msg__RoutePath)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.network), &m.Network)
	unique_identifier_msgs_msg.UUID__Sequence_to_C((*unique_identifier_msgs_msg.CUUID__Sequence)(unsafe.Pointer(&mem.segments)), m.Segments)
	KeyValue__Sequence_to_C(&mem.props, m.Props)
}

func (t _RoutePathTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RoutePath)
	mem := (*C.geographic_msgs__msg__RoutePath)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.Network, unsafe.Pointer(&mem.network))
	unique_identifier_msgs_msg.UUID__Sequence_to_Go(&m.Segments, *(*unique_identifier_msgs_msg.CUUID__Sequence)(unsafe.Pointer(&mem.segments)))
	KeyValue__Sequence_to_Go(&m.Props, mem.props)
}

func (t _RoutePathTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geographic_msgs__msg__RoutePath())
}

type CRoutePath = C.geographic_msgs__msg__RoutePath
type CRoutePath__Sequence = C.geographic_msgs__msg__RoutePath__Sequence

func RoutePath__Sequence_to_Go(goSlice *[]RoutePath, cSlice CRoutePath__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RoutePath, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		RoutePathTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func RoutePath__Sequence_to_C(cSlice *CRoutePath__Sequence, goSlice []RoutePath) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geographic_msgs__msg__RoutePath)(C.malloc(C.sizeof_struct_geographic_msgs__msg__RoutePath * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		RoutePathTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func RoutePath__Array_to_Go(goSlice []RoutePath, cSlice []CRoutePath) {
	for i := 0; i < len(cSlice); i++ {
		RoutePathTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RoutePath__Array_to_C(cSlice []CRoutePath, goSlice []RoutePath) {
	for i := 0; i < len(goSlice); i++ {
		RoutePathTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
