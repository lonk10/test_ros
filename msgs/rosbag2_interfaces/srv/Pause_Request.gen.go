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
	typemap.RegisterMessage("rosbag2_interfaces/Pause_Request", Pause_RequestTypeSupport)
	typemap.RegisterMessage("rosbag2_interfaces/srv/Pause_Request", Pause_RequestTypeSupport)
}

type Pause_Request struct {
}

// NewPause_Request creates a new Pause_Request with default values.
func NewPause_Request() *Pause_Request {
	self := Pause_Request{}
	self.SetDefaults()
	return &self
}

func (t *Pause_Request) Clone() *Pause_Request {
	c := &Pause_Request{}
	return c
}

func (t *Pause_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Pause_Request) SetDefaults() {
}

func (t *Pause_Request) GetTypeSupport() types.MessageTypeSupport {
	return Pause_RequestTypeSupport
}

// Pause_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type Pause_RequestPublisher struct {
	*rclgo.Publisher
}

// NewPause_RequestPublisher creates and returns a new publisher for the
// Pause_Request
func NewPause_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*Pause_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, Pause_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Pause_RequestPublisher{pub}, nil
}

func (p *Pause_RequestPublisher) Publish(msg *Pause_Request) error {
	return p.Publisher.Publish(msg)
}

// Pause_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type Pause_RequestSubscription struct {
	*rclgo.Subscription
}

// Pause_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a Pause_RequestSubscription.
type Pause_RequestSubscriptionCallback func(msg *Pause_Request, info *rclgo.MessageInfo, err error)

// NewPause_RequestSubscription creates and returns a new subscription for the
// Pause_Request
func NewPause_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback Pause_RequestSubscriptionCallback) (*Pause_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Pause_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, Pause_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &Pause_RequestSubscription{sub}, nil
}

func (s *Pause_RequestSubscription) TakeMessage(out *Pause_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// ClonePause_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func ClonePause_RequestSlice(dst, src []Pause_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var Pause_RequestTypeSupport types.MessageTypeSupport = _Pause_RequestTypeSupport{}

type _Pause_RequestTypeSupport struct{}

func (t _Pause_RequestTypeSupport) New() types.Message {
	return NewPause_Request()
}

func (t _Pause_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.rosbag2_interfaces__srv__Pause_Request
	return (unsafe.Pointer)(C.rosbag2_interfaces__srv__Pause_Request__create())
}

func (t _Pause_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.rosbag2_interfaces__srv__Pause_Request__destroy((*C.rosbag2_interfaces__srv__Pause_Request)(pointer_to_free))
}

func (t _Pause_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	
}

func (t _Pause_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	
}

func (t _Pause_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__rosbag2_interfaces__srv__Pause_Request())
}

type CPause_Request = C.rosbag2_interfaces__srv__Pause_Request
type CPause_Request__Sequence = C.rosbag2_interfaces__srv__Pause_Request__Sequence

func Pause_Request__Sequence_to_Go(goSlice *[]Pause_Request, cSlice CPause_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Pause_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		Pause_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Pause_Request__Sequence_to_C(cSlice *CPause_Request__Sequence, goSlice []Pause_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.rosbag2_interfaces__srv__Pause_Request)(C.malloc(C.sizeof_struct_rosbag2_interfaces__srv__Pause_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		Pause_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Pause_Request__Array_to_Go(goSlice []Pause_Request, cSlice []CPause_Request) {
	for i := 0; i < len(cSlice); i++ {
		Pause_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Pause_Request__Array_to_C(cSlice []CPause_Request, goSlice []Pause_Request) {
	for i := 0; i < len(goSlice); i++ {
		Pause_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
