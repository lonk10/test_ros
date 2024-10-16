// Code generated by rclgo-gen. DO NOT EDIT.

package map_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <map_msgs/srv/set_map_projections.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("map_msgs/SetMapProjections_Request", SetMapProjections_RequestTypeSupport)
	typemap.RegisterMessage("map_msgs/srv/SetMapProjections_Request", SetMapProjections_RequestTypeSupport)
}

type SetMapProjections_Request struct {
}

// NewSetMapProjections_Request creates a new SetMapProjections_Request with default values.
func NewSetMapProjections_Request() *SetMapProjections_Request {
	self := SetMapProjections_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetMapProjections_Request) Clone() *SetMapProjections_Request {
	c := &SetMapProjections_Request{}
	return c
}

func (t *SetMapProjections_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetMapProjections_Request) SetDefaults() {
}

func (t *SetMapProjections_Request) GetTypeSupport() types.MessageTypeSupport {
	return SetMapProjections_RequestTypeSupport
}

// SetMapProjections_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SetMapProjections_RequestPublisher struct {
	*rclgo.Publisher
}

// NewSetMapProjections_RequestPublisher creates and returns a new publisher for the
// SetMapProjections_Request
func NewSetMapProjections_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SetMapProjections_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, SetMapProjections_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetMapProjections_RequestPublisher{pub}, nil
}

func (p *SetMapProjections_RequestPublisher) Publish(msg *SetMapProjections_Request) error {
	return p.Publisher.Publish(msg)
}

// SetMapProjections_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SetMapProjections_RequestSubscription struct {
	*rclgo.Subscription
}

// SetMapProjections_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a SetMapProjections_RequestSubscription.
type SetMapProjections_RequestSubscriptionCallback func(msg *SetMapProjections_Request, info *rclgo.MessageInfo, err error)

// NewSetMapProjections_RequestSubscription creates and returns a new subscription for the
// SetMapProjections_Request
func NewSetMapProjections_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback SetMapProjections_RequestSubscriptionCallback) (*SetMapProjections_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SetMapProjections_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SetMapProjections_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &SetMapProjections_RequestSubscription{sub}, nil
}

func (s *SetMapProjections_RequestSubscription) TakeMessage(out *SetMapProjections_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSetMapProjections_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetMapProjections_RequestSlice(dst, src []SetMapProjections_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetMapProjections_RequestTypeSupport types.MessageTypeSupport = _SetMapProjections_RequestTypeSupport{}

type _SetMapProjections_RequestTypeSupport struct{}

func (t _SetMapProjections_RequestTypeSupport) New() types.Message {
	return NewSetMapProjections_Request()
}

func (t _SetMapProjections_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.map_msgs__srv__SetMapProjections_Request
	return (unsafe.Pointer)(C.map_msgs__srv__SetMapProjections_Request__create())
}

func (t _SetMapProjections_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.map_msgs__srv__SetMapProjections_Request__destroy((*C.map_msgs__srv__SetMapProjections_Request)(pointer_to_free))
}

func (t _SetMapProjections_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _SetMapProjections_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _SetMapProjections_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__map_msgs__srv__SetMapProjections_Request())
}

type CSetMapProjections_Request = C.map_msgs__srv__SetMapProjections_Request
type CSetMapProjections_Request__Sequence = C.map_msgs__srv__SetMapProjections_Request__Sequence

func SetMapProjections_Request__Sequence_to_Go(goSlice *[]SetMapProjections_Request, cSlice CSetMapProjections_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetMapProjections_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SetMapProjections_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SetMapProjections_Request__Sequence_to_C(cSlice *CSetMapProjections_Request__Sequence, goSlice []SetMapProjections_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.map_msgs__srv__SetMapProjections_Request)(C.malloc(C.sizeof_struct_map_msgs__srv__SetMapProjections_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SetMapProjections_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SetMapProjections_Request__Array_to_Go(goSlice []SetMapProjections_Request, cSlice []CSetMapProjections_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetMapProjections_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetMapProjections_Request__Array_to_C(cSlice []CSetMapProjections_Request, goSlice []SetMapProjections_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetMapProjections_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
