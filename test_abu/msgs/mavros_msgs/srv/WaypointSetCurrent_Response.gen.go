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

#include <mavros_msgs/srv/waypoint_set_current.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/WaypointSetCurrent_Response", WaypointSetCurrent_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/WaypointSetCurrent_Response", WaypointSetCurrent_ResponseTypeSupport)
}

type WaypointSetCurrent_Response struct {
	Success bool `yaml:"success"`
}

// NewWaypointSetCurrent_Response creates a new WaypointSetCurrent_Response with default values.
func NewWaypointSetCurrent_Response() *WaypointSetCurrent_Response {
	self := WaypointSetCurrent_Response{}
	self.SetDefaults()
	return &self
}

func (t *WaypointSetCurrent_Response) Clone() *WaypointSetCurrent_Response {
	c := &WaypointSetCurrent_Response{}
	c.Success = t.Success
	return c
}

func (t *WaypointSetCurrent_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WaypointSetCurrent_Response) SetDefaults() {
	t.Success = false
}

func (t *WaypointSetCurrent_Response) GetTypeSupport() types.MessageTypeSupport {
	return WaypointSetCurrent_ResponseTypeSupport
}

// WaypointSetCurrent_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type WaypointSetCurrent_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewWaypointSetCurrent_ResponsePublisher creates and returns a new publisher for the
// WaypointSetCurrent_Response
func NewWaypointSetCurrent_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*WaypointSetCurrent_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, WaypointSetCurrent_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &WaypointSetCurrent_ResponsePublisher{pub}, nil
}

func (p *WaypointSetCurrent_ResponsePublisher) Publish(msg *WaypointSetCurrent_Response) error {
	return p.Publisher.Publish(msg)
}

// WaypointSetCurrent_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type WaypointSetCurrent_ResponseSubscription struct {
	*rclgo.Subscription
}

// WaypointSetCurrent_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a WaypointSetCurrent_ResponseSubscription.
type WaypointSetCurrent_ResponseSubscriptionCallback func(msg *WaypointSetCurrent_Response, info *rclgo.MessageInfo, err error)

// NewWaypointSetCurrent_ResponseSubscription creates and returns a new subscription for the
// WaypointSetCurrent_Response
func NewWaypointSetCurrent_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback WaypointSetCurrent_ResponseSubscriptionCallback) (*WaypointSetCurrent_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg WaypointSetCurrent_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, WaypointSetCurrent_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &WaypointSetCurrent_ResponseSubscription{sub}, nil
}

func (s *WaypointSetCurrent_ResponseSubscription) TakeMessage(out *WaypointSetCurrent_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneWaypointSetCurrent_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWaypointSetCurrent_ResponseSlice(dst, src []WaypointSetCurrent_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WaypointSetCurrent_ResponseTypeSupport types.MessageTypeSupport = _WaypointSetCurrent_ResponseTypeSupport{}

type _WaypointSetCurrent_ResponseTypeSupport struct{}

func (t _WaypointSetCurrent_ResponseTypeSupport) New() types.Message {
	return NewWaypointSetCurrent_Response()
}

func (t _WaypointSetCurrent_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__WaypointSetCurrent_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__WaypointSetCurrent_Response__create())
}

func (t _WaypointSetCurrent_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__WaypointSetCurrent_Response__destroy((*C.mavros_msgs__srv__WaypointSetCurrent_Response)(pointer_to_free))
}

func (t _WaypointSetCurrent_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*WaypointSetCurrent_Response)
	mem := (*C.mavros_msgs__srv__WaypointSetCurrent_Response)(dst)
	mem.success = C.bool(m.Success)
}

func (t _WaypointSetCurrent_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*WaypointSetCurrent_Response)
	mem := (*C.mavros_msgs__srv__WaypointSetCurrent_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
}

func (t _WaypointSetCurrent_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__WaypointSetCurrent_Response())
}

type CWaypointSetCurrent_Response = C.mavros_msgs__srv__WaypointSetCurrent_Response
type CWaypointSetCurrent_Response__Sequence = C.mavros_msgs__srv__WaypointSetCurrent_Response__Sequence

func WaypointSetCurrent_Response__Sequence_to_Go(goSlice *[]WaypointSetCurrent_Response, cSlice CWaypointSetCurrent_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WaypointSetCurrent_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		WaypointSetCurrent_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func WaypointSetCurrent_Response__Sequence_to_C(cSlice *CWaypointSetCurrent_Response__Sequence, goSlice []WaypointSetCurrent_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__WaypointSetCurrent_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__WaypointSetCurrent_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		WaypointSetCurrent_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func WaypointSetCurrent_Response__Array_to_Go(goSlice []WaypointSetCurrent_Response, cSlice []CWaypointSetCurrent_Response) {
	for i := 0; i < len(cSlice); i++ {
		WaypointSetCurrent_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WaypointSetCurrent_Response__Array_to_C(cSlice []CWaypointSetCurrent_Response, goSlice []WaypointSetCurrent_Response) {
	for i := 0; i < len(goSlice); i++ {
		WaypointSetCurrent_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
