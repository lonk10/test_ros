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

#include <mavros_msgs/srv/waypoint_push.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/WaypointPush_Response", WaypointPush_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/WaypointPush_Response", WaypointPush_ResponseTypeSupport)
}

type WaypointPush_Response struct {
	Success bool `yaml:"success"`
	WpTransfered uint32 `yaml:"wp_transfered"`
}

// NewWaypointPush_Response creates a new WaypointPush_Response with default values.
func NewWaypointPush_Response() *WaypointPush_Response {
	self := WaypointPush_Response{}
	self.SetDefaults()
	return &self
}

func (t *WaypointPush_Response) Clone() *WaypointPush_Response {
	c := &WaypointPush_Response{}
	c.Success = t.Success
	c.WpTransfered = t.WpTransfered
	return c
}

func (t *WaypointPush_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *WaypointPush_Response) SetDefaults() {
	t.Success = false
	t.WpTransfered = 0
}

func (t *WaypointPush_Response) GetTypeSupport() types.MessageTypeSupport {
	return WaypointPush_ResponseTypeSupport
}

// WaypointPush_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type WaypointPush_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewWaypointPush_ResponsePublisher creates and returns a new publisher for the
// WaypointPush_Response
func NewWaypointPush_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*WaypointPush_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, WaypointPush_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &WaypointPush_ResponsePublisher{pub}, nil
}

func (p *WaypointPush_ResponsePublisher) Publish(msg *WaypointPush_Response) error {
	return p.Publisher.Publish(msg)
}

// WaypointPush_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type WaypointPush_ResponseSubscription struct {
	*rclgo.Subscription
}

// WaypointPush_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a WaypointPush_ResponseSubscription.
type WaypointPush_ResponseSubscriptionCallback func(msg *WaypointPush_Response, info *rclgo.MessageInfo, err error)

// NewWaypointPush_ResponseSubscription creates and returns a new subscription for the
// WaypointPush_Response
func NewWaypointPush_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback WaypointPush_ResponseSubscriptionCallback) (*WaypointPush_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg WaypointPush_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, WaypointPush_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &WaypointPush_ResponseSubscription{sub}, nil
}

func (s *WaypointPush_ResponseSubscription) TakeMessage(out *WaypointPush_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneWaypointPush_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneWaypointPush_ResponseSlice(dst, src []WaypointPush_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var WaypointPush_ResponseTypeSupport types.MessageTypeSupport = _WaypointPush_ResponseTypeSupport{}

type _WaypointPush_ResponseTypeSupport struct{}

func (t _WaypointPush_ResponseTypeSupport) New() types.Message {
	return NewWaypointPush_Response()
}

func (t _WaypointPush_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__WaypointPush_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__WaypointPush_Response__create())
}

func (t _WaypointPush_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__WaypointPush_Response__destroy((*C.mavros_msgs__srv__WaypointPush_Response)(pointer_to_free))
}

func (t _WaypointPush_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*WaypointPush_Response)
	mem := (*C.mavros_msgs__srv__WaypointPush_Response)(dst)
	mem.success = C.bool(m.Success)
	mem.wp_transfered = C.uint32_t(m.WpTransfered)
}

func (t _WaypointPush_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*WaypointPush_Response)
	mem := (*C.mavros_msgs__srv__WaypointPush_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	m.WpTransfered = uint32(mem.wp_transfered)
}

func (t _WaypointPush_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__WaypointPush_Response())
}

type CWaypointPush_Response = C.mavros_msgs__srv__WaypointPush_Response
type CWaypointPush_Response__Sequence = C.mavros_msgs__srv__WaypointPush_Response__Sequence

func WaypointPush_Response__Sequence_to_Go(goSlice *[]WaypointPush_Response, cSlice CWaypointPush_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]WaypointPush_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		WaypointPush_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func WaypointPush_Response__Sequence_to_C(cSlice *CWaypointPush_Response__Sequence, goSlice []WaypointPush_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__WaypointPush_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__WaypointPush_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		WaypointPush_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func WaypointPush_Response__Array_to_Go(goSlice []WaypointPush_Response, cSlice []CWaypointPush_Response) {
	for i := 0; i < len(cSlice); i++ {
		WaypointPush_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func WaypointPush_Response__Array_to_C(cSlice []CWaypointPush_Response, goSlice []WaypointPush_Response) {
	for i := 0; i < len(goSlice); i++ {
		WaypointPush_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
