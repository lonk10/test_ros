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

#include <mavros_msgs/msg/status_text.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/StatusText", StatusTextTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/StatusText", StatusTextTypeSupport)
}
const (
	StatusText_EMERGENCY uint8 = 0// Severity levels
	StatusText_ALERT uint8 = 1
	StatusText_CRITICAL uint8 = 2
	StatusText_ERROR uint8 = 3
	StatusText_WARNING uint8 = 4
	StatusText_NOTICE uint8 = 5
	StatusText_INFO uint8 = 6
	StatusText_DEBUG uint8 = 7
)

type StatusText struct {
	Header std_msgs_msg.Header `yaml:"header"`// Fields
	Severity uint8 `yaml:"severity"`
	Text string `yaml:"text"`
}

// NewStatusText creates a new StatusText with default values.
func NewStatusText() *StatusText {
	self := StatusText{}
	self.SetDefaults()
	return &self
}

func (t *StatusText) Clone() *StatusText {
	c := &StatusText{}
	c.Header = *t.Header.Clone()
	c.Severity = t.Severity
	c.Text = t.Text
	return c
}

func (t *StatusText) CloneMsg() types.Message {
	return t.Clone()
}

func (t *StatusText) SetDefaults() {
	t.Header.SetDefaults()
	t.Severity = 0
	t.Text = ""
}

func (t *StatusText) GetTypeSupport() types.MessageTypeSupport {
	return StatusTextTypeSupport
}

// StatusTextPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type StatusTextPublisher struct {
	*rclgo.Publisher
}

// NewStatusTextPublisher creates and returns a new publisher for the
// StatusText
func NewStatusTextPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*StatusTextPublisher, error) {
	pub, err := node.NewPublisher(topic_name, StatusTextTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &StatusTextPublisher{pub}, nil
}

func (p *StatusTextPublisher) Publish(msg *StatusText) error {
	return p.Publisher.Publish(msg)
}

// StatusTextSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type StatusTextSubscription struct {
	*rclgo.Subscription
}

// StatusTextSubscriptionCallback type is used to provide a subscription
// handler function for a StatusTextSubscription.
type StatusTextSubscriptionCallback func(msg *StatusText, info *rclgo.MessageInfo, err error)

// NewStatusTextSubscription creates and returns a new subscription for the
// StatusText
func NewStatusTextSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback StatusTextSubscriptionCallback) (*StatusTextSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg StatusText
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, StatusTextTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &StatusTextSubscription{sub}, nil
}

func (s *StatusTextSubscription) TakeMessage(out *StatusText) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneStatusTextSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneStatusTextSlice(dst, src []StatusText) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var StatusTextTypeSupport types.MessageTypeSupport = _StatusTextTypeSupport{}

type _StatusTextTypeSupport struct{}

func (t _StatusTextTypeSupport) New() types.Message {
	return NewStatusText()
}

func (t _StatusTextTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__StatusText
	return (unsafe.Pointer)(C.mavros_msgs__msg__StatusText__create())
}

func (t _StatusTextTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__StatusText__destroy((*C.mavros_msgs__msg__StatusText)(pointer_to_free))
}

func (t _StatusTextTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*StatusText)
	mem := (*C.mavros_msgs__msg__StatusText)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.severity = C.uint8_t(m.Severity)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.text), m.Text)
}

func (t _StatusTextTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*StatusText)
	mem := (*C.mavros_msgs__msg__StatusText)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Severity = uint8(mem.severity)
	primitives.StringAsGoStruct(&m.Text, unsafe.Pointer(&mem.text))
}

func (t _StatusTextTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__StatusText())
}

type CStatusText = C.mavros_msgs__msg__StatusText
type CStatusText__Sequence = C.mavros_msgs__msg__StatusText__Sequence

func StatusText__Sequence_to_Go(goSlice *[]StatusText, cSlice CStatusText__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]StatusText, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		StatusTextTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func StatusText__Sequence_to_C(cSlice *CStatusText__Sequence, goSlice []StatusText) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__StatusText)(C.malloc(C.sizeof_struct_mavros_msgs__msg__StatusText * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		StatusTextTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func StatusText__Array_to_Go(goSlice []StatusText, cSlice []CStatusText) {
	for i := 0; i < len(cSlice); i++ {
		StatusTextTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func StatusText__Array_to_C(cSlice []CStatusText, goSlice []StatusText) {
	for i := 0; i < len(goSlice); i++ {
		StatusTextTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
