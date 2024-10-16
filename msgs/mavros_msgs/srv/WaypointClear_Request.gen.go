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

#include <mavros_msgs/srv/waypoint_clear.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/WaypointClear_Request", WaypointClear_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/WaypointClear_Request", WaypointClear_RequestTypeSupport)
}

type WaypointClear_Request struct {
}

// NewWaypointClear_Request creates a new WaypointClear_Request with default values.
func NewWaypointClear_Request() *WaypointClear_Request {
	self := WaypointClear_Request{}
	self.SetDefaults()
	return &self
}

func (t *WaypointClear_Request) Clone() *WaypointClear_Request {
	c := &WaypointClear_Request{}
	return c
}

func (t *WaypointClear_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WaypointClear_Request) SetDefaults() {
}

func (t *WaypointClear_Request) GetTypeSupport() types.MessageTypeSupport {
	return WaypointClear_RequestTypeSupport
}

// WaypointClear_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type WaypointClear_RequestPublisher struct {
	*rclgo.Publisher
}

// NewWaypointClear_RequestPublisher creates and returns a new publisher for the
// WaypointClear_Request
func NewWaypointClear_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*WaypointClear_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, WaypointClear_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &WaypointClear_RequestPublisher{pub}, nil
}

func (p *WaypointClear_RequestPublisher) Publish(msg *WaypointClear_Request) error {
	return p.Publisher.Publish(msg)
}

// WaypointClear_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type WaypointClear_RequestSubscription struct {
	*rclgo.Subscription
}

// WaypointClear_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a WaypointClear_RequestSubscription.
type WaypointClear_RequestSubscriptionCallback func(msg *WaypointClear_Request, info *rclgo.MessageInfo, err error)

// NewWaypointClear_RequestSubscription creates and returns a new subscription for the
// WaypointClear_Request
func NewWaypointClear_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback WaypointClear_RequestSubscriptionCallback) (*WaypointClear_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg WaypointClear_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, WaypointClear_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &WaypointClear_RequestSubscription{sub}, nil
}

func (s *WaypointClear_RequestSubscription) TakeMessage(out *WaypointClear_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneWaypointClear_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWaypointClear_RequestSlice(dst, src []WaypointClear_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WaypointClear_RequestTypeSupport types.MessageTypeSupport = _WaypointClear_RequestTypeSupport{}

type _WaypointClear_RequestTypeSupport struct{}

func (t _WaypointClear_RequestTypeSupport) New() types.Message {
	return NewWaypointClear_Request()
}

func (t _WaypointClear_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__WaypointClear_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__WaypointClear_Request__create())
}

func (t _WaypointClear_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__WaypointClear_Request__destroy((*C.mavros_msgs__srv__WaypointClear_Request)(pointer_to_free))
}

func (t _WaypointClear_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _WaypointClear_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _WaypointClear_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__WaypointClear_Request())
}

type CWaypointClear_Request = C.mavros_msgs__srv__WaypointClear_Request
type CWaypointClear_Request__Sequence = C.mavros_msgs__srv__WaypointClear_Request__Sequence

func WaypointClear_Request__Sequence_to_Go(goSlice *[]WaypointClear_Request, cSlice CWaypointClear_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WaypointClear_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		WaypointClear_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func WaypointClear_Request__Sequence_to_C(cSlice *CWaypointClear_Request__Sequence, goSlice []WaypointClear_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__WaypointClear_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__WaypointClear_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		WaypointClear_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func WaypointClear_Request__Array_to_Go(goSlice []WaypointClear_Request, cSlice []CWaypointClear_Request) {
	for i := 0; i < len(cSlice); i++ {
		WaypointClear_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WaypointClear_Request__Array_to_C(cSlice []CWaypointClear_Request, goSlice []WaypointClear_Request) {
	for i := 0; i < len(goSlice); i++ {
		WaypointClear_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
