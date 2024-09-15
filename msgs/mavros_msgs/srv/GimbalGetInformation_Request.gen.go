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

#include <mavros_msgs/srv/gimbal_get_information.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/GimbalGetInformation_Request", GimbalGetInformation_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/GimbalGetInformation_Request", GimbalGetInformation_RequestTypeSupport)
}

type GimbalGetInformation_Request struct {
}

// NewGimbalGetInformation_Request creates a new GimbalGetInformation_Request with default values.
func NewGimbalGetInformation_Request() *GimbalGetInformation_Request {
	self := GimbalGetInformation_Request{}
	self.SetDefaults()
	return &self
}

func (t *GimbalGetInformation_Request) Clone() *GimbalGetInformation_Request {
	c := &GimbalGetInformation_Request{}
	return c
}

func (t *GimbalGetInformation_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GimbalGetInformation_Request) SetDefaults() {
}

func (t *GimbalGetInformation_Request) GetTypeSupport() types.MessageTypeSupport {
	return GimbalGetInformation_RequestTypeSupport
}

// GimbalGetInformation_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GimbalGetInformation_RequestPublisher struct {
	*rclgo.Publisher
}

// NewGimbalGetInformation_RequestPublisher creates and returns a new publisher for the
// GimbalGetInformation_Request
func NewGimbalGetInformation_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GimbalGetInformation_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GimbalGetInformation_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GimbalGetInformation_RequestPublisher{pub}, nil
}

func (p *GimbalGetInformation_RequestPublisher) Publish(msg *GimbalGetInformation_Request) error {
	return p.Publisher.Publish(msg)
}

// GimbalGetInformation_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GimbalGetInformation_RequestSubscription struct {
	*rclgo.Subscription
}

// GimbalGetInformation_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a GimbalGetInformation_RequestSubscription.
type GimbalGetInformation_RequestSubscriptionCallback func(msg *GimbalGetInformation_Request, info *rclgo.MessageInfo, err error)

// NewGimbalGetInformation_RequestSubscription creates and returns a new subscription for the
// GimbalGetInformation_Request
func NewGimbalGetInformation_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GimbalGetInformation_RequestSubscriptionCallback) (*GimbalGetInformation_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GimbalGetInformation_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GimbalGetInformation_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GimbalGetInformation_RequestSubscription{sub}, nil
}

func (s *GimbalGetInformation_RequestSubscription) TakeMessage(out *GimbalGetInformation_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGimbalGetInformation_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGimbalGetInformation_RequestSlice(dst, src []GimbalGetInformation_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GimbalGetInformation_RequestTypeSupport types.MessageTypeSupport = _GimbalGetInformation_RequestTypeSupport{}

type _GimbalGetInformation_RequestTypeSupport struct{}

func (t _GimbalGetInformation_RequestTypeSupport) New() types.Message {
	return NewGimbalGetInformation_Request()
}

func (t _GimbalGetInformation_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__GimbalGetInformation_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__GimbalGetInformation_Request__create())
}

func (t _GimbalGetInformation_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__GimbalGetInformation_Request__destroy((*C.mavros_msgs__srv__GimbalGetInformation_Request)(pointer_to_free))
}

func (t _GimbalGetInformation_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _GimbalGetInformation_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _GimbalGetInformation_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__GimbalGetInformation_Request())
}

type CGimbalGetInformation_Request = C.mavros_msgs__srv__GimbalGetInformation_Request
type CGimbalGetInformation_Request__Sequence = C.mavros_msgs__srv__GimbalGetInformation_Request__Sequence

func GimbalGetInformation_Request__Sequence_to_Go(goSlice *[]GimbalGetInformation_Request, cSlice CGimbalGetInformation_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GimbalGetInformation_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GimbalGetInformation_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GimbalGetInformation_Request__Sequence_to_C(cSlice *CGimbalGetInformation_Request__Sequence, goSlice []GimbalGetInformation_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__GimbalGetInformation_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__GimbalGetInformation_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GimbalGetInformation_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GimbalGetInformation_Request__Array_to_Go(goSlice []GimbalGetInformation_Request, cSlice []CGimbalGetInformation_Request) {
	for i := 0; i < len(cSlice); i++ {
		GimbalGetInformation_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GimbalGetInformation_Request__Array_to_C(cSlice []CGimbalGetInformation_Request, goSlice []GimbalGetInformation_Request) {
	for i := 0; i < len(goSlice); i++ {
		GimbalGetInformation_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
