// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/msg/waypoint_list.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/WaypointList", WaypointListTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/WaypointList", WaypointListTypeSupport)
}

type WaypointList struct {
	CurrentSeq uint16 `yaml:"current_seq"`
	Waypoints []Waypoint `yaml:"waypoints"`
}

// NewWaypointList creates a new WaypointList with default values.
func NewWaypointList() *WaypointList {
	self := WaypointList{}
	self.SetDefaults()
	return &self
}

func (t *WaypointList) Clone() *WaypointList {
	c := &WaypointList{}
	c.CurrentSeq = t.CurrentSeq
	if t.Waypoints != nil {
		c.Waypoints = make([]Waypoint, len(t.Waypoints))
		CloneWaypointSlice(c.Waypoints, t.Waypoints)
	}
	return c
}

func (t *WaypointList) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WaypointList) SetDefaults() {
	t.CurrentSeq = 0
	t.Waypoints = nil
}

func (t *WaypointList) GetTypeSupport() types.MessageTypeSupport {
	return WaypointListTypeSupport
}

// WaypointListPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type WaypointListPublisher struct {
	*rclgo.Publisher
}

// NewWaypointListPublisher creates and returns a new publisher for the
// WaypointList
func NewWaypointListPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*WaypointListPublisher, error) {
	pub, err := node.NewPublisher(topic_name, WaypointListTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &WaypointListPublisher{pub}, nil
}

func (p *WaypointListPublisher) Publish(msg *WaypointList) error {
	return p.Publisher.Publish(msg)
}

// WaypointListSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type WaypointListSubscription struct {
	*rclgo.Subscription
}

// WaypointListSubscriptionCallback type is used to provide a subscription
// handler function for a WaypointListSubscription.
type WaypointListSubscriptionCallback func(msg *WaypointList, info *rclgo.MessageInfo, err error)

// NewWaypointListSubscription creates and returns a new subscription for the
// WaypointList
func NewWaypointListSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback WaypointListSubscriptionCallback) (*WaypointListSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg WaypointList
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, WaypointListTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &WaypointListSubscription{sub}, nil
}

func (s *WaypointListSubscription) TakeMessage(out *WaypointList) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneWaypointListSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWaypointListSlice(dst, src []WaypointList) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WaypointListTypeSupport types.MessageTypeSupport = _WaypointListTypeSupport{}

type _WaypointListTypeSupport struct{}

func (t _WaypointListTypeSupport) New() types.Message {
	return NewWaypointList()
}

func (t _WaypointListTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__WaypointList
	return (unsafe.Pointer)(C.mavros_msgs__msg__WaypointList__create())
}

func (t _WaypointListTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__WaypointList__destroy((*C.mavros_msgs__msg__WaypointList)(pointer_to_free))
}

func (t _WaypointListTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*WaypointList)
	mem := (*C.mavros_msgs__msg__WaypointList)(dst)
	mem.current_seq = C.uint16_t(m.CurrentSeq)
	Waypoint__Sequence_to_C(&mem.waypoints, m.Waypoints)
}

func (t _WaypointListTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*WaypointList)
	mem := (*C.mavros_msgs__msg__WaypointList)(ros2_message_buffer)
	m.CurrentSeq = uint16(mem.current_seq)
	Waypoint__Sequence_to_Go(&m.Waypoints, mem.waypoints)
}

func (t _WaypointListTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__WaypointList())
}

type CWaypointList = C.mavros_msgs__msg__WaypointList
type CWaypointList__Sequence = C.mavros_msgs__msg__WaypointList__Sequence

func WaypointList__Sequence_to_Go(goSlice *[]WaypointList, cSlice CWaypointList__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WaypointList, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		WaypointListTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func WaypointList__Sequence_to_C(cSlice *CWaypointList__Sequence, goSlice []WaypointList) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__WaypointList)(C.malloc(C.sizeof_struct_mavros_msgs__msg__WaypointList * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		WaypointListTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func WaypointList__Array_to_Go(goSlice []WaypointList, cSlice []CWaypointList) {
	for i := 0; i < len(cSlice); i++ {
		WaypointListTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WaypointList__Array_to_C(cSlice []CWaypointList, goSlice []WaypointList) {
	for i := 0; i < len(goSlice); i++ {
		WaypointListTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
