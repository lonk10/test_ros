// Code generated by rclgo-gen. DO NOT EDIT.

package rcl_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	rcl_interfaces_msg "test/msgs/rcl_interfaces/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rcl_interfaces/srv/set_parameters_atomically.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rcl_interfaces/SetParametersAtomically_Request", SetParametersAtomically_RequestTypeSupport)
	typemap.RegisterMessage("rcl_interfaces/srv/SetParametersAtomically_Request", SetParametersAtomically_RequestTypeSupport)
}

type SetParametersAtomically_Request struct {
	Parameters []rcl_interfaces_msg.Parameter `yaml:"parameters"`// A list of parameters to set atomically.This call will either set all values, or none of the values.
}

// NewSetParametersAtomically_Request creates a new SetParametersAtomically_Request with default values.
func NewSetParametersAtomically_Request() *SetParametersAtomically_Request {
	self := SetParametersAtomically_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetParametersAtomically_Request) Clone() *SetParametersAtomically_Request {
	c := &SetParametersAtomically_Request{}
	if t.Parameters != nil {
		c.Parameters = make([]rcl_interfaces_msg.Parameter, len(t.Parameters))
		rcl_interfaces_msg.CloneParameterSlice(c.Parameters, t.Parameters)
	}
	return c
}

func (t *SetParametersAtomically_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetParametersAtomically_Request) SetDefaults() {
	t.Parameters = nil
}

func (t *SetParametersAtomically_Request) GetTypeSupport() types.MessageTypeSupport {
	return SetParametersAtomically_RequestTypeSupport
}

// SetParametersAtomically_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SetParametersAtomically_RequestPublisher struct {
	*rclgo.Publisher
}

// NewSetParametersAtomically_RequestPublisher creates and returns a new publisher for the
// SetParametersAtomically_Request
func NewSetParametersAtomically_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SetParametersAtomically_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, SetParametersAtomically_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetParametersAtomically_RequestPublisher{pub}, nil
}

func (p *SetParametersAtomically_RequestPublisher) Publish(msg *SetParametersAtomically_Request) error {
	return p.Publisher.Publish(msg)
}

// SetParametersAtomically_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SetParametersAtomically_RequestSubscription struct {
	*rclgo.Subscription
}

// SetParametersAtomically_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a SetParametersAtomically_RequestSubscription.
type SetParametersAtomically_RequestSubscriptionCallback func(msg *SetParametersAtomically_Request, info *rclgo.MessageInfo, err error)

// NewSetParametersAtomically_RequestSubscription creates and returns a new subscription for the
// SetParametersAtomically_Request
func NewSetParametersAtomically_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback SetParametersAtomically_RequestSubscriptionCallback) (*SetParametersAtomically_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SetParametersAtomically_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SetParametersAtomically_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &SetParametersAtomically_RequestSubscription{sub}, nil
}

func (s *SetParametersAtomically_RequestSubscription) TakeMessage(out *SetParametersAtomically_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSetParametersAtomically_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetParametersAtomically_RequestSlice(dst, src []SetParametersAtomically_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetParametersAtomically_RequestTypeSupport types.MessageTypeSupport = _SetParametersAtomically_RequestTypeSupport{}

type _SetParametersAtomically_RequestTypeSupport struct{}

func (t _SetParametersAtomically_RequestTypeSupport) New() types.Message {
	return NewSetParametersAtomically_Request()
}

func (t _SetParametersAtomically_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__SetParametersAtomically_Request
	return (unsafe.Pointer)(C.rcl_interfaces__srv__SetParametersAtomically_Request__create())
}

func (t _SetParametersAtomically_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__SetParametersAtomically_Request__destroy((*C.rcl_interfaces__srv__SetParametersAtomically_Request)(pointer_to_free))
}

func (t _SetParametersAtomically_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetParametersAtomically_Request)
	mem := (*C.rcl_interfaces__srv__SetParametersAtomically_Request)(dst)
	rcl_interfaces_msg.Parameter__Sequence_to_C((*rcl_interfaces_msg.CParameter__Sequence)(unsafe.Pointer(&mem.parameters)), m.Parameters)
}

func (t _SetParametersAtomically_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetParametersAtomically_Request)
	mem := (*C.rcl_interfaces__srv__SetParametersAtomically_Request)(ros2_message_buffer)
	rcl_interfaces_msg.Parameter__Sequence_to_Go(&m.Parameters, *(*rcl_interfaces_msg.CParameter__Sequence)(unsafe.Pointer(&mem.parameters)))
}

func (t _SetParametersAtomically_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__SetParametersAtomically_Request())
}

type CSetParametersAtomically_Request = C.rcl_interfaces__srv__SetParametersAtomically_Request
type CSetParametersAtomically_Request__Sequence = C.rcl_interfaces__srv__SetParametersAtomically_Request__Sequence

func SetParametersAtomically_Request__Sequence_to_Go(goSlice *[]SetParametersAtomically_Request, cSlice CSetParametersAtomically_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetParametersAtomically_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SetParametersAtomically_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SetParametersAtomically_Request__Sequence_to_C(cSlice *CSetParametersAtomically_Request__Sequence, goSlice []SetParametersAtomically_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__SetParametersAtomically_Request)(C.malloc(C.sizeof_struct_rcl_interfaces__srv__SetParametersAtomically_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SetParametersAtomically_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SetParametersAtomically_Request__Array_to_Go(goSlice []SetParametersAtomically_Request, cSlice []CSetParametersAtomically_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetParametersAtomically_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetParametersAtomically_Request__Array_to_C(cSlice []CSetParametersAtomically_Request, goSlice []SetParametersAtomically_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetParametersAtomically_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
