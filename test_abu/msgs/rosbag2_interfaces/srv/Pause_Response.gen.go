// Code generated by rclgo-gen. DO NOT EDIT.

package rosbag2_interfaces_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <rosbag2_interfaces/srv/pause.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rosbag2_interfaces/Pause_Response", Pause_ResponseTypeSupport)
	typemap.RegisterMessage("rosbag2_interfaces/srv/Pause_Response", Pause_ResponseTypeSupport)
}

type Pause_Response struct {
}

// NewPause_Response creates a new Pause_Response with default values.
func NewPause_Response() *Pause_Response {
	self := Pause_Response{}
	self.SetDefaults()
	return &self
}

func (t *Pause_Response) Clone() *Pause_Response {
	c := &Pause_Response{}
	return c
}

func (t *Pause_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Pause_Response) SetDefaults() {
}

func (t *Pause_Response) GetTypeSupport() types.MessageTypeSupport {
	return Pause_ResponseTypeSupport
}

// Pause_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Pause_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewPause_ResponsePublisher creates and returns a new publisher for the
// Pause_Response
func NewPause_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Pause_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, Pause_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Pause_ResponsePublisher{pub}, nil
}

func (p *Pause_ResponsePublisher) Publish(msg *Pause_Response) error {
	return p.Publisher.Publish(msg)
}

// Pause_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Pause_ResponseSubscription struct {
	*rclgo.Subscription
}

// Pause_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a Pause_ResponseSubscription.
type Pause_ResponseSubscriptionCallback func(msg *Pause_Response, info *rclgo.MessageInfo, err error)

// NewPause_ResponseSubscription creates and returns a new subscription for the
// Pause_Response
func NewPause_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback Pause_ResponseSubscriptionCallback) (*Pause_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Pause_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Pause_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &Pause_ResponseSubscription{sub}, nil
}

func (s *Pause_ResponseSubscription) TakeMessage(out *Pause_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePause_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePause_ResponseSlice(dst, src []Pause_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Pause_ResponseTypeSupport types.MessageTypeSupport = _Pause_ResponseTypeSupport{}

type _Pause_ResponseTypeSupport struct{}

func (t _Pause_ResponseTypeSupport) New() types.Message {
	return NewPause_Response()
}

func (t _Pause_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rosbag2_interfaces__srv__Pause_Response
	return (unsafe.Pointer)(C.rosbag2_interfaces__srv__Pause_Response__create())
}

func (t _Pause_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rosbag2_interfaces__srv__Pause_Response__destroy((*C.rosbag2_interfaces__srv__Pause_Response)(pointer_to_free))
}

func (t _Pause_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _Pause_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _Pause_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rosbag2_interfaces__srv__Pause_Response())
}

type CPause_Response = C.rosbag2_interfaces__srv__Pause_Response
type CPause_Response__Sequence = C.rosbag2_interfaces__srv__Pause_Response__Sequence

func Pause_Response__Sequence_to_Go(goSlice *[]Pause_Response, cSlice CPause_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Pause_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Pause_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Pause_Response__Sequence_to_C(cSlice *CPause_Response__Sequence, goSlice []Pause_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rosbag2_interfaces__srv__Pause_Response)(C.malloc(C.sizeof_struct_rosbag2_interfaces__srv__Pause_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Pause_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Pause_Response__Array_to_Go(goSlice []Pause_Response, cSlice []CPause_Response) {
	for i := 0; i < len(cSlice); i++ {
		Pause_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Pause_Response__Array_to_C(cSlice []CPause_Response, goSlice []Pause_Response) {
	for i := 0; i < len(goSlice); i++ {
		Pause_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
