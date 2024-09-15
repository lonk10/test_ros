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

#include <mavros_msgs/srv/param_pull.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/ParamPull_Response", ParamPull_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/ParamPull_Response", ParamPull_ResponseTypeSupport)
}

type ParamPull_Response struct {
	Success bool `yaml:"success"`
	ParamReceived uint32 `yaml:"param_received"`
}

// NewParamPull_Response creates a new ParamPull_Response with default values.
func NewParamPull_Response() *ParamPull_Response {
	self := ParamPull_Response{}
	self.SetDefaults()
	return &self
}

func (t *ParamPull_Response) Clone() *ParamPull_Response {
	c := &ParamPull_Response{}
	c.Success = t.Success
	c.ParamReceived = t.ParamReceived
	return c
}

func (t *ParamPull_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ParamPull_Response) SetDefaults() {
	t.Success = false
	t.ParamReceived = 0
}

func (t *ParamPull_Response) GetTypeSupport() types.MessageTypeSupport {
	return ParamPull_ResponseTypeSupport
}

// ParamPull_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ParamPull_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewParamPull_ResponsePublisher creates and returns a new publisher for the
// ParamPull_Response
func NewParamPull_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ParamPull_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, ParamPull_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ParamPull_ResponsePublisher{pub}, nil
}

func (p *ParamPull_ResponsePublisher) Publish(msg *ParamPull_Response) error {
	return p.Publisher.Publish(msg)
}

// ParamPull_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ParamPull_ResponseSubscription struct {
	*rclgo.Subscription
}

// ParamPull_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a ParamPull_ResponseSubscription.
type ParamPull_ResponseSubscriptionCallback func(msg *ParamPull_Response, info *rclgo.MessageInfo, err error)

// NewParamPull_ResponseSubscription creates and returns a new subscription for the
// ParamPull_Response
func NewParamPull_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ParamPull_ResponseSubscriptionCallback) (*ParamPull_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ParamPull_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ParamPull_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ParamPull_ResponseSubscription{sub}, nil
}

func (s *ParamPull_ResponseSubscription) TakeMessage(out *ParamPull_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneParamPull_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneParamPull_ResponseSlice(dst, src []ParamPull_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ParamPull_ResponseTypeSupport types.MessageTypeSupport = _ParamPull_ResponseTypeSupport{}

type _ParamPull_ResponseTypeSupport struct{}

func (t _ParamPull_ResponseTypeSupport) New() types.Message {
	return NewParamPull_Response()
}

func (t _ParamPull_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__ParamPull_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__ParamPull_Response__create())
}

func (t _ParamPull_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__ParamPull_Response__destroy((*C.mavros_msgs__srv__ParamPull_Response)(pointer_to_free))
}

func (t _ParamPull_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ParamPull_Response)
	mem := (*C.mavros_msgs__srv__ParamPull_Response)(dst)
	mem.success = C.bool(m.Success)
	mem.param_received = C.uint32_t(m.ParamReceived)
}

func (t _ParamPull_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ParamPull_Response)
	mem := (*C.mavros_msgs__srv__ParamPull_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	m.ParamReceived = uint32(mem.param_received)
}

func (t _ParamPull_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__ParamPull_Response())
}

type CParamPull_Response = C.mavros_msgs__srv__ParamPull_Response
type CParamPull_Response__Sequence = C.mavros_msgs__srv__ParamPull_Response__Sequence

func ParamPull_Response__Sequence_to_Go(goSlice *[]ParamPull_Response, cSlice CParamPull_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ParamPull_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ParamPull_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ParamPull_Response__Sequence_to_C(cSlice *CParamPull_Response__Sequence, goSlice []ParamPull_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__ParamPull_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__ParamPull_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ParamPull_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ParamPull_Response__Array_to_Go(goSlice []ParamPull_Response, cSlice []CParamPull_Response) {
	for i := 0; i < len(cSlice); i++ {
		ParamPull_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ParamPull_Response__Array_to_C(cSlice []CParamPull_Response, goSlice []ParamPull_Response) {
	for i := 0; i < len(goSlice); i++ {
		ParamPull_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
