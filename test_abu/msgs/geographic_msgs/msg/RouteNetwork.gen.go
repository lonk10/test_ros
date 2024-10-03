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

#include <geographic_msgs/msg/route_network.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geographic_msgs/RouteNetwork", RouteNetworkTypeSupport)
	typemap.RegisterMessage("geographic_msgs/msg/RouteNetwork", RouteNetworkTypeSupport)
}

type RouteNetwork struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Id unique_identifier_msgs_msg.UUID `yaml:"id"`// This route network identifier
	Bounds BoundingBox `yaml:"bounds"`// 2D bounding box for network
	Points []WayPoint `yaml:"points"`// Way points in this network
	Segments []RouteSegment `yaml:"segments"`// Directed edges of this network
	Props []KeyValue `yaml:"props"`// Network key/value properties
}

// NewRouteNetwork creates a new RouteNetwork with default values.
func NewRouteNetwork() *RouteNetwork {
	self := RouteNetwork{}
	self.SetDefaults()
	return &self
}

func (t *RouteNetwork) Clone() *RouteNetwork {
	c := &RouteNetwork{}
	c.Header = *t.Header.Clone()
	c.Id = *t.Id.Clone()
	c.Bounds = *t.Bounds.Clone()
	if t.Points != nil {
		c.Points = make([]WayPoint, len(t.Points))
		CloneWayPointSlice(c.Points, t.Points)
	}
	if t.Segments != nil {
		c.Segments = make([]RouteSegment, len(t.Segments))
		CloneRouteSegmentSlice(c.Segments, t.Segments)
	}
	if t.Props != nil {
		c.Props = make([]KeyValue, len(t.Props))
		CloneKeyValueSlice(c.Props, t.Props)
	}
	return c
}

func (t *RouteNetwork) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RouteNetwork) SetDefaults() {
	t.Header.SetDefaults()
	t.Id.SetDefaults()
	t.Bounds.SetDefaults()
	t.Points = nil
	t.Segments = nil
	t.Props = nil
}

func (t *RouteNetwork) GetTypeSupport() types.MessageTypeSupport {
	return RouteNetworkTypeSupport
}

// RouteNetworkPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type RouteNetworkPublisher struct {
	*rclgo.Publisher
}

// NewRouteNetworkPublisher creates and returns a new publisher for the
// RouteNetwork
func NewRouteNetworkPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*RouteNetworkPublisher, error) {
	pub, err := node.NewPublisher(topic_name, RouteNetworkTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &RouteNetworkPublisher{pub}, nil
}

func (p *RouteNetworkPublisher) Publish(msg *RouteNetwork) error {
	return p.Publisher.Publish(msg)
}

// RouteNetworkSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type RouteNetworkSubscription struct {
	*rclgo.Subscription
}

// RouteNetworkSubscriptionCallback type is used to provide a subscription
// handler function for a RouteNetworkSubscription.
type RouteNetworkSubscriptionCallback func(msg *RouteNetwork, info *rclgo.MessageInfo, err error)

// NewRouteNetworkSubscription creates and returns a new subscription for the
// RouteNetwork
func NewRouteNetworkSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback RouteNetworkSubscriptionCallback) (*RouteNetworkSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg RouteNetwork
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, RouteNetworkTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &RouteNetworkSubscription{sub}, nil
}

func (s *RouteNetworkSubscription) TakeMessage(out *RouteNetwork) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneRouteNetworkSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRouteNetworkSlice(dst, src []RouteNetwork) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RouteNetworkTypeSupport types.MessageTypeSupport = _RouteNetworkTypeSupport{}

type _RouteNetworkTypeSupport struct{}

func (t _RouteNetworkTypeSupport) New() types.Message {
	return NewRouteNetwork()
}

func (t _RouteNetworkTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geographic_msgs__msg__RouteNetwork
	return (unsafe.Pointer)(C.geographic_msgs__msg__RouteNetwork__create())
}

func (t _RouteNetworkTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geographic_msgs__msg__RouteNetwork__destroy((*C.geographic_msgs__msg__RouteNetwork)(pointer_to_free))
}

func (t _RouteNetworkTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RouteNetwork)
	mem := (*C.geographic_msgs__msg__RouteNetwork)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	unique_identifier_msgs_msg.UUIDTypeSupport.AsCStruct(unsafe.Pointer(&mem.id), &m.Id)
	BoundingBoxTypeSupport.AsCStruct(unsafe.Pointer(&mem.bounds), &m.Bounds)
	WayPoint__Sequence_to_C(&mem.points, m.Points)
	RouteSegment__Sequence_to_C(&mem.segments, m.Segments)
	KeyValue__Sequence_to_C(&mem.props, m.Props)
}

func (t _RouteNetworkTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RouteNetwork)
	mem := (*C.geographic_msgs__msg__RouteNetwork)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	unique_identifier_msgs_msg.UUIDTypeSupport.AsGoStruct(&m.Id, unsafe.Pointer(&mem.id))
	BoundingBoxTypeSupport.AsGoStruct(&m.Bounds, unsafe.Pointer(&mem.bounds))
	WayPoint__Sequence_to_Go(&m.Points, mem.points)
	RouteSegment__Sequence_to_Go(&m.Segments, mem.segments)
	KeyValue__Sequence_to_Go(&m.Props, mem.props)
}

func (t _RouteNetworkTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geographic_msgs__msg__RouteNetwork())
}

type CRouteNetwork = C.geographic_msgs__msg__RouteNetwork
type CRouteNetwork__Sequence = C.geographic_msgs__msg__RouteNetwork__Sequence

func RouteNetwork__Sequence_to_Go(goSlice *[]RouteNetwork, cSlice CRouteNetwork__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RouteNetwork, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		RouteNetworkTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func RouteNetwork__Sequence_to_C(cSlice *CRouteNetwork__Sequence, goSlice []RouteNetwork) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geographic_msgs__msg__RouteNetwork)(C.malloc(C.sizeof_struct_geographic_msgs__msg__RouteNetwork * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		RouteNetworkTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func RouteNetwork__Array_to_Go(goSlice []RouteNetwork, cSlice []CRouteNetwork) {
	for i := 0; i < len(cSlice); i++ {
		RouteNetworkTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RouteNetwork__Array_to_C(cSlice []CRouteNetwork, goSlice []RouteNetwork) {
	for i := 0; i < len(goSlice); i++ {
		RouteNetworkTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
