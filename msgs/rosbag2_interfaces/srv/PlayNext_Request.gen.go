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

#include <rosbag2_interfaces/srv/play_next.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("rosbag2_interfaces/PlayNext_Request", PlayNext_RequestTypeSupport)
	typemap.RegisterMessage("rosbag2_interfaces/srv/PlayNext_Request", PlayNext_RequestTypeSupport)
}

type PlayNext_Request struct {
}

// NewPlayNext_Request creates a new PlayNext_Request with default values.
func NewPlayNext_Request() *PlayNext_Request {
	self := PlayNext_Request{}
	self.SetDefaults()
	return &self
}

func (t *PlayNext_Request) Clone() *PlayNext_Request {
	c := &PlayNext_Request{}
	return c
}

func (t *PlayNext_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *PlayNext_Request) SetDefaults() {
}

func (t *PlayNext_Request) GetTypeSupport() types.MessageTypeSupport {
	return PlayNext_RequestTypeSupport
}

// PlayNext_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type PlayNext_RequestPublisher struct {
	*rclgo.Publisher
}

// NewPlayNext_RequestPublisher creates and returns a new publisher for the
// PlayNext_Request
func NewPlayNext_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*PlayNext_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, PlayNext_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PlayNext_RequestPublisher{pub}, nil
}

func (p *PlayNext_RequestPublisher) Publish(msg *PlayNext_Request) error {
	return p.Publisher.Publish(msg)
}

// PlayNext_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type PlayNext_RequestSubscription struct {
	*rclgo.Subscription
}

// PlayNext_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a PlayNext_RequestSubscription.
type PlayNext_RequestSubscriptionCallback func(msg *PlayNext_Request, info *rclgo.MessageInfo, err error)

// NewPlayNext_RequestSubscription creates and returns a new subscription for the
// PlayNext_Request
func NewPlayNext_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback PlayNext_RequestSubscriptionCallback) (*PlayNext_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg PlayNext_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, PlayNext_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &PlayNext_RequestSubscription{sub}, nil
}

func (s *PlayNext_RequestSubscription) TakeMessage(out *PlayNext_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePlayNext_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePlayNext_RequestSlice(dst, src []PlayNext_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var PlayNext_RequestTypeSupport types.MessageTypeSupport = _PlayNext_RequestTypeSupport{}

type _PlayNext_RequestTypeSupport struct{}

func (t _PlayNext_RequestTypeSupport) New() types.Message {
	return NewPlayNext_Request()
}

func (t _PlayNext_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rosbag2_interfaces__srv__PlayNext_Request
	return (unsafe.Pointer)(C.rosbag2_interfaces__srv__PlayNext_Request__create())
}

func (t _PlayNext_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rosbag2_interfaces__srv__PlayNext_Request__destroy((*C.rosbag2_interfaces__srv__PlayNext_Request)(pointer_to_free))
}

func (t _PlayNext_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _PlayNext_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _PlayNext_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rosbag2_interfaces__srv__PlayNext_Request())
}

type CPlayNext_Request = C.rosbag2_interfaces__srv__PlayNext_Request
type CPlayNext_Request__Sequence = C.rosbag2_interfaces__srv__PlayNext_Request__Sequence

func PlayNext_Request__Sequence_to_Go(goSlice *[]PlayNext_Request, cSlice CPlayNext_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]PlayNext_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		PlayNext_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func PlayNext_Request__Sequence_to_C(cSlice *CPlayNext_Request__Sequence, goSlice []PlayNext_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rosbag2_interfaces__srv__PlayNext_Request)(C.malloc(C.sizeof_struct_rosbag2_interfaces__srv__PlayNext_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		PlayNext_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func PlayNext_Request__Array_to_Go(goSlice []PlayNext_Request, cSlice []CPlayNext_Request) {
	for i := 0; i < len(cSlice); i++ {
		PlayNext_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func PlayNext_Request__Array_to_C(cSlice []CPlayNext_Request, goSlice []PlayNext_Request) {
	for i := 0; i < len(goSlice); i++ {
		PlayNext_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
