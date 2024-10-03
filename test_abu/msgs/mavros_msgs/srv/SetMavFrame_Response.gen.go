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

#include <mavros_msgs/srv/set_mav_frame.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/SetMavFrame_Response", SetMavFrame_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/SetMavFrame_Response", SetMavFrame_ResponseTypeSupport)
}

type SetMavFrame_Response struct {
	Success bool `yaml:"success"`
}

// NewSetMavFrame_Response creates a new SetMavFrame_Response with default values.
func NewSetMavFrame_Response() *SetMavFrame_Response {
	self := SetMavFrame_Response{}
	self.SetDefaults()
	return &self
}

func (t *SetMavFrame_Response) Clone() *SetMavFrame_Response {
	c := &SetMavFrame_Response{}
	c.Success = t.Success
	return c
}

func (t *SetMavFrame_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetMavFrame_Response) SetDefaults() {
	t.Success = false
}

func (t *SetMavFrame_Response) GetTypeSupport() types.MessageTypeSupport {
	return SetMavFrame_ResponseTypeSupport
}

// SetMavFrame_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SetMavFrame_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewSetMavFrame_ResponsePublisher creates and returns a new publisher for the
// SetMavFrame_Response
func NewSetMavFrame_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SetMavFrame_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, SetMavFrame_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetMavFrame_ResponsePublisher{pub}, nil
}

func (p *SetMavFrame_ResponsePublisher) Publish(msg *SetMavFrame_Response) error {
	return p.Publisher.Publish(msg)
}

// SetMavFrame_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SetMavFrame_ResponseSubscription struct {
	*rclgo.Subscription
}

// SetMavFrame_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a SetMavFrame_ResponseSubscription.
type SetMavFrame_ResponseSubscriptionCallback func(msg *SetMavFrame_Response, info *rclgo.MessageInfo, err error)

// NewSetMavFrame_ResponseSubscription creates and returns a new subscription for the
// SetMavFrame_Response
func NewSetMavFrame_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback SetMavFrame_ResponseSubscriptionCallback) (*SetMavFrame_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SetMavFrame_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SetMavFrame_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &SetMavFrame_ResponseSubscription{sub}, nil
}

func (s *SetMavFrame_ResponseSubscription) TakeMessage(out *SetMavFrame_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSetMavFrame_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetMavFrame_ResponseSlice(dst, src []SetMavFrame_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetMavFrame_ResponseTypeSupport types.MessageTypeSupport = _SetMavFrame_ResponseTypeSupport{}

type _SetMavFrame_ResponseTypeSupport struct{}

func (t _SetMavFrame_ResponseTypeSupport) New() types.Message {
	return NewSetMavFrame_Response()
}

func (t _SetMavFrame_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__SetMavFrame_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__SetMavFrame_Response__create())
}

func (t _SetMavFrame_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__SetMavFrame_Response__destroy((*C.mavros_msgs__srv__SetMavFrame_Response)(pointer_to_free))
}

func (t _SetMavFrame_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetMavFrame_Response)
	mem := (*C.mavros_msgs__srv__SetMavFrame_Response)(dst)
	mem.success = C.bool(m.Success)
}

func (t _SetMavFrame_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetMavFrame_Response)
	mem := (*C.mavros_msgs__srv__SetMavFrame_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
}

func (t _SetMavFrame_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__SetMavFrame_Response())
}

type CSetMavFrame_Response = C.mavros_msgs__srv__SetMavFrame_Response
type CSetMavFrame_Response__Sequence = C.mavros_msgs__srv__SetMavFrame_Response__Sequence

func SetMavFrame_Response__Sequence_to_Go(goSlice *[]SetMavFrame_Response, cSlice CSetMavFrame_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetMavFrame_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SetMavFrame_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SetMavFrame_Response__Sequence_to_C(cSlice *CSetMavFrame_Response__Sequence, goSlice []SetMavFrame_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__SetMavFrame_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__SetMavFrame_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SetMavFrame_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SetMavFrame_Response__Array_to_Go(goSlice []SetMavFrame_Response, cSlice []CSetMavFrame_Response) {
	for i := 0; i < len(cSlice); i++ {
		SetMavFrame_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetMavFrame_Response__Array_to_C(cSlice []CSetMavFrame_Response, goSlice []SetMavFrame_Response) {
	for i := 0; i < len(goSlice); i++ {
		SetMavFrame_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
