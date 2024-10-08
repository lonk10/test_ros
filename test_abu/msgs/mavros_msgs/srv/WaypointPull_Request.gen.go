// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/srv/waypoint_pull.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/WaypointPull_Request", WaypointPull_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/WaypointPull_Request", WaypointPull_RequestTypeSupport)
}

type WaypointPull_Request struct {
}

// NewWaypointPull_Request creates a new WaypointPull_Request with default values.
func NewWaypointPull_Request() *WaypointPull_Request {
	self := WaypointPull_Request{}
	self.SetDefaults()
	return &self
}

func (t *WaypointPull_Request) Clone() *WaypointPull_Request {
	c := &WaypointPull_Request{}
	return c
}

func (t *WaypointPull_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WaypointPull_Request) SetDefaults() {
}

func (t *WaypointPull_Request) GetTypeSupport() types.MessageTypeSupport {
	return WaypointPull_RequestTypeSupport
}

// WaypointPull_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type WaypointPull_RequestPublisher struct {
	*rclgo.Publisher
}

// NewWaypointPull_RequestPublisher creates and returns a new publisher for the
// WaypointPull_Request
func NewWaypointPull_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*WaypointPull_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, WaypointPull_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &WaypointPull_RequestPublisher{pub}, nil
}

func (p *WaypointPull_RequestPublisher) Publish(msg *WaypointPull_Request) error {
	return p.Publisher.Publish(msg)
}

// WaypointPull_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type WaypointPull_RequestSubscription struct {
	*rclgo.Subscription
}

// WaypointPull_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a WaypointPull_RequestSubscription.
type WaypointPull_RequestSubscriptionCallback func(msg *WaypointPull_Request, info *rclgo.MessageInfo, err error)

// NewWaypointPull_RequestSubscription creates and returns a new subscription for the
// WaypointPull_Request
func NewWaypointPull_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback WaypointPull_RequestSubscriptionCallback) (*WaypointPull_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg WaypointPull_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, WaypointPull_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &WaypointPull_RequestSubscription{sub}, nil
}

func (s *WaypointPull_RequestSubscription) TakeMessage(out *WaypointPull_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneWaypointPull_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWaypointPull_RequestSlice(dst, src []WaypointPull_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WaypointPull_RequestTypeSupport types.MessageTypeSupport = _WaypointPull_RequestTypeSupport{}

type _WaypointPull_RequestTypeSupport struct{}

func (t _WaypointPull_RequestTypeSupport) New() types.Message {
	return NewWaypointPull_Request()
}

func (t _WaypointPull_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__WaypointPull_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__WaypointPull_Request__create())
}

func (t _WaypointPull_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__WaypointPull_Request__destroy((*C.mavros_msgs__srv__WaypointPull_Request)(pointer_to_free))
}

func (t _WaypointPull_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _WaypointPull_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _WaypointPull_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__WaypointPull_Request())
}

type CWaypointPull_Request = C.mavros_msgs__srv__WaypointPull_Request
type CWaypointPull_Request__Sequence = C.mavros_msgs__srv__WaypointPull_Request__Sequence

func WaypointPull_Request__Sequence_to_Go(goSlice *[]WaypointPull_Request, cSlice CWaypointPull_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WaypointPull_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		WaypointPull_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func WaypointPull_Request__Sequence_to_C(cSlice *CWaypointPull_Request__Sequence, goSlice []WaypointPull_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__WaypointPull_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__WaypointPull_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		WaypointPull_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func WaypointPull_Request__Array_to_Go(goSlice []WaypointPull_Request, cSlice []CWaypointPull_Request) {
	for i := 0; i < len(cSlice); i++ {
		WaypointPull_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WaypointPull_Request__Array_to_C(cSlice []CWaypointPull_Request, goSlice []WaypointPull_Request) {
	for i := 0; i < len(goSlice); i++ {
		WaypointPull_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
