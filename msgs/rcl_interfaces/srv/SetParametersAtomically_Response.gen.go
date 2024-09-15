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
	typemap.RegisterMessage("rcl_interfaces/SetParametersAtomically_Response", SetParametersAtomically_ResponseTypeSupport)
	typemap.RegisterMessage("rcl_interfaces/srv/SetParametersAtomically_Response", SetParametersAtomically_ResponseTypeSupport)
}

type SetParametersAtomically_Response struct {
	Result rcl_interfaces_msg.SetParametersResult `yaml:"result"`// Indicates whether setting all of the parameters succeeded or not and why.
}

// NewSetParametersAtomically_Response creates a new SetParametersAtomically_Response with default values.
func NewSetParametersAtomically_Response() *SetParametersAtomically_Response {
	self := SetParametersAtomically_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetParametersAtomically_Response) Clone() *SetParametersAtomically_Response {
	c := &SetParametersAtomically_Response{}
	c.Result = *t.Result.Clone()
	return c
}

func (t *SetParametersAtomically_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetParametersAtomically_Response) SetDefaults() {
	t.Result.SetDefaults()
}

func (t *SetParametersAtomically_Response) GetTypeSupport() types.MessageTypeSupport {
	return SetParametersAtomically_ResponseTypeSupport
}

// SetParametersAtomically_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SetParametersAtomically_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewSetParametersAtomically_ResponsePublisher creates and returns a new publisher for the
// SetParametersAtomically_Response
func NewSetParametersAtomically_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SetParametersAtomically_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, SetParametersAtomically_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetParametersAtomically_ResponsePublisher{pub}, nil
}

func (p *SetParametersAtomically_ResponsePublisher) Publish(msg *SetParametersAtomically_Response) error {
	return p.Publisher.Publish(msg)
}

// SetParametersAtomically_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SetParametersAtomically_ResponseSubscription struct {
	*rclgo.Subscription
}

// SetParametersAtomically_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a SetParametersAtomically_ResponseSubscription.
type SetParametersAtomically_ResponseSubscriptionCallback func(msg *SetParametersAtomically_Response, info *rclgo.MessageInfo, err error)

// NewSetParametersAtomically_ResponseSubscription creates and returns a new subscription for the
// SetParametersAtomically_Response
func NewSetParametersAtomically_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback SetParametersAtomically_ResponseSubscriptionCallback) (*SetParametersAtomically_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SetParametersAtomically_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SetParametersAtomically_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &SetParametersAtomically_ResponseSubscription{sub}, nil
}

func (s *SetParametersAtomically_ResponseSubscription) TakeMessage(out *SetParametersAtomically_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSetParametersAtomically_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetParametersAtomically_ResponseSlice(dst, src []SetParametersAtomically_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetParametersAtomically_ResponseTypeSupport types.MessageTypeSupport = _SetParametersAtomically_ResponseTypeSupport{}

type _SetParametersAtomically_ResponseTypeSupport struct{}

func (t _SetParametersAtomically_ResponseTypeSupport) New() types.Message {
	return NewSetParametersAtomically_Response()
}

func (t _SetParametersAtomically_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rcl_interfaces__srv__SetParametersAtomically_Response
	return (unsafe.Pointer)(C.rcl_interfaces__srv__SetParametersAtomically_Response__create())
}

func (t _SetParametersAtomically_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rcl_interfaces__srv__SetParametersAtomically_Response__destroy((*C.rcl_interfaces__srv__SetParametersAtomically_Response)(pointer_to_free))
}

func (t _SetParametersAtomically_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetParametersAtomically_Response)
	mem := (*C.rcl_interfaces__srv__SetParametersAtomically_Response)(dst)
	rcl_interfaces_msg.SetParametersResultTypeSupport.AsCStruct(unsafe.Pointer(&mem.result), &m.Result)
}

func (t _SetParametersAtomically_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetParametersAtomically_Response)
	mem := (*C.rcl_interfaces__srv__SetParametersAtomically_Response)(ros2_message_buffer)
	rcl_interfaces_msg.SetParametersResultTypeSupport.AsGoStruct(&m.Result, unsafe.Pointer(&mem.result))
}

func (t _SetParametersAtomically_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rcl_interfaces__srv__SetParametersAtomically_Response())
}

type CSetParametersAtomically_Response = C.rcl_interfaces__srv__SetParametersAtomically_Response
type CSetParametersAtomically_Response__Sequence = C.rcl_interfaces__srv__SetParametersAtomically_Response__Sequence

func SetParametersAtomically_Response__Sequence_to_Go(goSlice *[]SetParametersAtomically_Response, cSlice CSetParametersAtomically_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetParametersAtomically_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SetParametersAtomically_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SetParametersAtomically_Response__Sequence_to_C(cSlice *CSetParametersAtomically_Response__Sequence, goSlice []SetParametersAtomically_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rcl_interfaces__srv__SetParametersAtomically_Response)(C.malloc(C.sizeof_struct_rcl_interfaces__srv__SetParametersAtomically_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SetParametersAtomically_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SetParametersAtomically_Response__Array_to_Go(goSlice []SetParametersAtomically_Response, cSlice []CSetParametersAtomically_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetParametersAtomically_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetParametersAtomically_Response__Array_to_C(cSlice []CSetParametersAtomically_Response, goSlice []SetParametersAtomically_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetParametersAtomically_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
