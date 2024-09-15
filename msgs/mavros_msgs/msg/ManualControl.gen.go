// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/msg/manual_control.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/ManualControl", ManualControlTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/ManualControl", ManualControlTypeSupport)
}

type ManualControl struct {
	Header std_msgs_msg.Header `yaml:"header"`// Manual Control state
	X float32 `yaml:"x"`
	Y float32 `yaml:"y"`
	Z float32 `yaml:"z"`
	R float32 `yaml:"r"`
	Buttons uint16 `yaml:"buttons"`
}

// NewManualControl creates a new ManualControl with default values.
func NewManualControl() *ManualControl {
	self := ManualControl{}
	self.SetDefaults()
	return &self
}

func (t *ManualControl) Clone() *ManualControl {
	c := &ManualControl{}
	c.Header = *t.Header.Clone()
	c.X = t.X
	c.Y = t.Y
	c.Z = t.Z
	c.R = t.R
	c.Buttons = t.Buttons
	return c
}

func (t *ManualControl) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ManualControl) SetDefaults() {
	t.Header.SetDefaults()
	t.X = 0
	t.Y = 0
	t.Z = 0
	t.R = 0
	t.Buttons = 0
}

func (t *ManualControl) GetTypeSupport() types.MessageTypeSupport {
	return ManualControlTypeSupport
}

// ManualControlPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ManualControlPublisher struct {
	*rclgo.Publisher
}

// NewManualControlPublisher creates and returns a new publisher for the
// ManualControl
func NewManualControlPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ManualControlPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ManualControlTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ManualControlPublisher{pub}, nil
}

func (p *ManualControlPublisher) Publish(msg *ManualControl) error {
	return p.Publisher.Publish(msg)
}

// ManualControlSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ManualControlSubscription struct {
	*rclgo.Subscription
}

// ManualControlSubscriptionCallback type is used to provide a subscription
// handler function for a ManualControlSubscription.
type ManualControlSubscriptionCallback func(msg *ManualControl, info *rclgo.MessageInfo, err error)

// NewManualControlSubscription creates and returns a new subscription for the
// ManualControl
func NewManualControlSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ManualControlSubscriptionCallback) (*ManualControlSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ManualControl
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ManualControlTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ManualControlSubscription{sub}, nil
}

func (s *ManualControlSubscription) TakeMessage(out *ManualControl) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneManualControlSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneManualControlSlice(dst, src []ManualControl) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ManualControlTypeSupport types.MessageTypeSupport = _ManualControlTypeSupport{}

type _ManualControlTypeSupport struct{}

func (t _ManualControlTypeSupport) New() types.Message {
	return NewManualControl()
}

func (t _ManualControlTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__ManualControl
	return (unsafe.Pointer)(C.mavros_msgs__msg__ManualControl__create())
}

func (t _ManualControlTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__ManualControl__destroy((*C.mavros_msgs__msg__ManualControl)(pointer_to_free))
}

func (t _ManualControlTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ManualControl)
	mem := (*C.mavros_msgs__msg__ManualControl)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.x = C.float(m.X)
	mem.y = C.float(m.Y)
	mem.z = C.float(m.Z)
	mem.r = C.float(m.R)
	mem.buttons = C.uint16_t(m.Buttons)
}

func (t _ManualControlTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ManualControl)
	mem := (*C.mavros_msgs__msg__ManualControl)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.X = float32(mem.x)
	m.Y = float32(mem.y)
	m.Z = float32(mem.z)
	m.R = float32(mem.r)
	m.Buttons = uint16(mem.buttons)
}

func (t _ManualControlTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__ManualControl())
}

type CManualControl = C.mavros_msgs__msg__ManualControl
type CManualControl__Sequence = C.mavros_msgs__msg__ManualControl__Sequence

func ManualControl__Sequence_to_Go(goSlice *[]ManualControl, cSlice CManualControl__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ManualControl, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ManualControlTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ManualControl__Sequence_to_C(cSlice *CManualControl__Sequence, goSlice []ManualControl) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__ManualControl)(C.malloc(C.sizeof_struct_mavros_msgs__msg__ManualControl * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ManualControlTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ManualControl__Array_to_Go(goSlice []ManualControl, cSlice []CManualControl) {
	for i := 0; i < len(cSlice); i++ {
		ManualControlTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ManualControl__Array_to_C(cSlice []CManualControl, goSlice []ManualControl) {
	for i := 0; i < len(goSlice); i++ {
		ManualControlTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
