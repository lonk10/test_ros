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

#include <mavros_msgs/srv/mount_configure.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/MountConfigure_Response", MountConfigure_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/MountConfigure_Response", MountConfigure_ResponseTypeSupport)
}

type MountConfigure_Response struct {
	Success bool `yaml:"success"`
	Result uint8 `yaml:"result"`// raw result returned by COMMAND_ACK
}

// NewMountConfigure_Response creates a new MountConfigure_Response with default values.
func NewMountConfigure_Response() *MountConfigure_Response {
	self := MountConfigure_Response{}
	self.SetDefaults()
	return &self
}

func (t *MountConfigure_Response) Clone() *MountConfigure_Response {
	c := &MountConfigure_Response{}
	c.Success = t.Success
	c.Result = t.Result
	return c
}

func (t *MountConfigure_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MountConfigure_Response) SetDefaults() {
	t.Success = false
	t.Result = 0
}

func (t *MountConfigure_Response) GetTypeSupport() types.MessageTypeSupport {
	return MountConfigure_ResponseTypeSupport
}

// MountConfigure_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MountConfigure_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewMountConfigure_ResponsePublisher creates and returns a new publisher for the
// MountConfigure_Response
func NewMountConfigure_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MountConfigure_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, MountConfigure_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MountConfigure_ResponsePublisher{pub}, nil
}

func (p *MountConfigure_ResponsePublisher) Publish(msg *MountConfigure_Response) error {
	return p.Publisher.Publish(msg)
}

// MountConfigure_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MountConfigure_ResponseSubscription struct {
	*rclgo.Subscription
}

// MountConfigure_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a MountConfigure_ResponseSubscription.
type MountConfigure_ResponseSubscriptionCallback func(msg *MountConfigure_Response, info *rclgo.MessageInfo, err error)

// NewMountConfigure_ResponseSubscription creates and returns a new subscription for the
// MountConfigure_Response
func NewMountConfigure_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MountConfigure_ResponseSubscriptionCallback) (*MountConfigure_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MountConfigure_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MountConfigure_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MountConfigure_ResponseSubscription{sub}, nil
}

func (s *MountConfigure_ResponseSubscription) TakeMessage(out *MountConfigure_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMountConfigure_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMountConfigure_ResponseSlice(dst, src []MountConfigure_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MountConfigure_ResponseTypeSupport types.MessageTypeSupport = _MountConfigure_ResponseTypeSupport{}

type _MountConfigure_ResponseTypeSupport struct{}

func (t _MountConfigure_ResponseTypeSupport) New() types.Message {
	return NewMountConfigure_Response()
}

func (t _MountConfigure_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__MountConfigure_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__MountConfigure_Response__create())
}

func (t _MountConfigure_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__MountConfigure_Response__destroy((*C.mavros_msgs__srv__MountConfigure_Response)(pointer_to_free))
}

func (t _MountConfigure_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MountConfigure_Response)
	mem := (*C.mavros_msgs__srv__MountConfigure_Response)(dst)
	mem.success = C.bool(m.Success)
	mem.result = C.uint8_t(m.Result)
}

func (t _MountConfigure_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MountConfigure_Response)
	mem := (*C.mavros_msgs__srv__MountConfigure_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	m.Result = uint8(mem.result)
}

func (t _MountConfigure_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__MountConfigure_Response())
}

type CMountConfigure_Response = C.mavros_msgs__srv__MountConfigure_Response
type CMountConfigure_Response__Sequence = C.mavros_msgs__srv__MountConfigure_Response__Sequence

func MountConfigure_Response__Sequence_to_Go(goSlice *[]MountConfigure_Response, cSlice CMountConfigure_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MountConfigure_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MountConfigure_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MountConfigure_Response__Sequence_to_C(cSlice *CMountConfigure_Response__Sequence, goSlice []MountConfigure_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__MountConfigure_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__MountConfigure_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MountConfigure_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MountConfigure_Response__Array_to_Go(goSlice []MountConfigure_Response, cSlice []CMountConfigure_Response) {
	for i := 0; i < len(cSlice); i++ {
		MountConfigure_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MountConfigure_Response__Array_to_C(cSlice []CMountConfigure_Response, goSlice []MountConfigure_Response) {
	for i := 0; i < len(goSlice); i++ {
		MountConfigure_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
