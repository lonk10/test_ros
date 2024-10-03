// Code generated by rclgo-gen. DO NOT EDIT.

package turtlesim_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <turtlesim/srv/teleport_absolute.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("turtlesim/TeleportAbsolute_Request", TeleportAbsolute_RequestTypeSupport)
	typemap.RegisterMessage("turtlesim/srv/TeleportAbsolute_Request", TeleportAbsolute_RequestTypeSupport)
}

type TeleportAbsolute_Request struct {
	X float32 `yaml:"x"`
	Y float32 `yaml:"y"`
	Theta float32 `yaml:"theta"`
}

// NewTeleportAbsolute_Request creates a new TeleportAbsolute_Request with default values.
func NewTeleportAbsolute_Request() *TeleportAbsolute_Request {
	self := TeleportAbsolute_Request{}
	self.SetDefaults()
	return &self
}

func (t *TeleportAbsolute_Request) Clone() *TeleportAbsolute_Request {
	c := &TeleportAbsolute_Request{}
	c.X = t.X
	c.Y = t.Y
	c.Theta = t.Theta
	return c
}

func (t *TeleportAbsolute_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TeleportAbsolute_Request) SetDefaults() {
	t.X = 0
	t.Y = 0
	t.Theta = 0
}

func (t *TeleportAbsolute_Request) GetTypeSupport() types.MessageTypeSupport {
	return TeleportAbsolute_RequestTypeSupport
}

// TeleportAbsolute_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TeleportAbsolute_RequestPublisher struct {
	*rclgo.Publisher
}

// NewTeleportAbsolute_RequestPublisher creates and returns a new publisher for the
// TeleportAbsolute_Request
func NewTeleportAbsolute_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TeleportAbsolute_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, TeleportAbsolute_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TeleportAbsolute_RequestPublisher{pub}, nil
}

func (p *TeleportAbsolute_RequestPublisher) Publish(msg *TeleportAbsolute_Request) error {
	return p.Publisher.Publish(msg)
}

// TeleportAbsolute_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TeleportAbsolute_RequestSubscription struct {
	*rclgo.Subscription
}

// TeleportAbsolute_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a TeleportAbsolute_RequestSubscription.
type TeleportAbsolute_RequestSubscriptionCallback func(msg *TeleportAbsolute_Request, info *rclgo.MessageInfo, err error)

// NewTeleportAbsolute_RequestSubscription creates and returns a new subscription for the
// TeleportAbsolute_Request
func NewTeleportAbsolute_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback TeleportAbsolute_RequestSubscriptionCallback) (*TeleportAbsolute_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg TeleportAbsolute_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TeleportAbsolute_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &TeleportAbsolute_RequestSubscription{sub}, nil
}

func (s *TeleportAbsolute_RequestSubscription) TakeMessage(out *TeleportAbsolute_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTeleportAbsolute_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTeleportAbsolute_RequestSlice(dst, src []TeleportAbsolute_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TeleportAbsolute_RequestTypeSupport types.MessageTypeSupport = _TeleportAbsolute_RequestTypeSupport{}

type _TeleportAbsolute_RequestTypeSupport struct{}

func (t _TeleportAbsolute_RequestTypeSupport) New() types.Message {
	return NewTeleportAbsolute_Request()
}

func (t _TeleportAbsolute_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.turtlesim__srv__TeleportAbsolute_Request
	return (unsafe.Pointer)(C.turtlesim__srv__TeleportAbsolute_Request__create())
}

func (t _TeleportAbsolute_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.turtlesim__srv__TeleportAbsolute_Request__destroy((*C.turtlesim__srv__TeleportAbsolute_Request)(pointer_to_free))
}

func (t _TeleportAbsolute_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TeleportAbsolute_Request)
	mem := (*C.turtlesim__srv__TeleportAbsolute_Request)(dst)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.theta = C.float(m.Theta)
}

func (t _TeleportAbsolute_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TeleportAbsolute_Request)
	mem := (*C.turtlesim__srv__TeleportAbsolute_Request)(ros2_message_buffer)
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Theta = float32(mem.theta)
}

func (t _TeleportAbsolute_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__turtlesim__srv__TeleportAbsolute_Request())
}

type CTeleportAbsolute_Request = C.turtlesim__srv__TeleportAbsolute_Request
type CTeleportAbsolute_Request__Sequence = C.turtlesim__srv__TeleportAbsolute_Request__Sequence

func TeleportAbsolute_Request__Sequence_to_Go(goSlice *[]TeleportAbsolute_Request, cSlice CTeleportAbsolute_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TeleportAbsolute_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		TeleportAbsolute_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func TeleportAbsolute_Request__Sequence_to_C(cSlice *CTeleportAbsolute_Request__Sequence, goSlice []TeleportAbsolute_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.turtlesim__srv__TeleportAbsolute_Request)(C.malloc(C.sizeof_struct_turtlesim__srv__TeleportAbsolute_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		TeleportAbsolute_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func TeleportAbsolute_Request__Array_to_Go(goSlice []TeleportAbsolute_Request, cSlice []CTeleportAbsolute_Request) {
	for i := 0; i < len(cSlice); i++ {
		TeleportAbsolute_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TeleportAbsolute_Request__Array_to_C(cSlice []CTeleportAbsolute_Request, goSlice []TeleportAbsolute_Request) {
	for i := 0; i < len(goSlice); i++ {
		TeleportAbsolute_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
