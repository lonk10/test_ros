// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/msg/rc_in.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/RCIn", RCInTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/RCIn", RCInTypeSupport)
}

type RCIn struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Rssi uint8 `yaml:"rssi"`
	Channels []uint16 `yaml:"channels"`
}

// NewRCIn creates a new RCIn with default values.
func NewRCIn() *RCIn {
	self := RCIn{}
	self.SetDefaults()
	return &self
}

func (t *RCIn) Clone() *RCIn {
	c := &RCIn{}
	c.Header = *t.Header.Clone()
	c.Rssi = t.Rssi
	if t.Channels != nil {
		c.Channels = make([]uint16, len(t.Channels))
		copy(c.Channels, t.Channels)
	}
	return c
}

func (t *RCIn) CloneMsg() types.Message {
	return t.Clone()
}

func (t *RCIn) SetDefaults() {
	t.Header.SetDefaults()
	t.Rssi = 0
	t.Channels = nil
}

func (t *RCIn) GetTypeSupport() types.MessageTypeSupport {
	return RCInTypeSupport
}

// RCInPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type RCInPublisher struct {
	*rclgo.Publisher
}

// NewRCInPublisher creates and returns a new publisher for the
// RCIn
func NewRCInPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*RCInPublisher, error) {
	pub, err := node.NewPublisher(topic_name, RCInTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &RCInPublisher{pub}, nil
}

func (p *RCInPublisher) Publish(msg *RCIn) error {
	return p.Publisher.Publish(msg)
}

// RCInSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type RCInSubscription struct {
	*rclgo.Subscription
}

// RCInSubscriptionCallback type is used to provide a subscription
// handler function for a RCInSubscription.
type RCInSubscriptionCallback func(msg *RCIn, info *rclgo.MessageInfo, err error)

// NewRCInSubscription creates and returns a new subscription for the
// RCIn
func NewRCInSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback RCInSubscriptionCallback) (*RCInSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg RCIn
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, RCInTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &RCInSubscription{sub}, nil
}

func (s *RCInSubscription) TakeMessage(out *RCIn) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneRCInSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneRCInSlice(dst, src []RCIn) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var RCInTypeSupport types.MessageTypeSupport = _RCInTypeSupport{}

type _RCInTypeSupport struct{}

func (t _RCInTypeSupport) New() types.Message {
	return NewRCIn()
}

func (t _RCInTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__RCIn
	return (unsafe.Pointer)(C.mavros_msgs__msg__RCIn__create())
}

func (t _RCInTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__RCIn__destroy((*C.mavros_msgs__msg__RCIn)(pointer_to_free))
}

func (t _RCInTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*RCIn)
	mem := (*C.mavros_msgs__msg__RCIn)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.rssi = C.uint8_t(m.Rssi)
	primitives.Uint16__Sequence_to_C((*primitives.CUint16__Sequence)(unsafe.Pointer(&mem.channels)), m.Channels)
}

func (t _RCInTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*RCIn)
	mem := (*C.mavros_msgs__msg__RCIn)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Rssi = uint8(mem.rssi)
	primitives.Uint16__Sequence_to_Go(&m.Channels, *(*primitives.CUint16__Sequence)(unsafe.Pointer(&mem.channels)))
}

func (t _RCInTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__RCIn())
}

type CRCIn = C.mavros_msgs__msg__RCIn
type CRCIn__Sequence = C.mavros_msgs__msg__RCIn__Sequence

func RCIn__Sequence_to_Go(goSlice *[]RCIn, cSlice CRCIn__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]RCIn, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		RCInTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func RCIn__Sequence_to_C(cSlice *CRCIn__Sequence, goSlice []RCIn) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__RCIn)(C.malloc(C.sizeof_struct_mavros_msgs__msg__RCIn * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		RCInTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func RCIn__Array_to_Go(goSlice []RCIn, cSlice []CRCIn) {
	for i := 0; i < len(cSlice); i++ {
		RCInTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func RCIn__Array_to_C(cSlice []CRCIn, goSlice []RCIn) {
	for i := 0; i < len(goSlice); i++ {
		RCInTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
