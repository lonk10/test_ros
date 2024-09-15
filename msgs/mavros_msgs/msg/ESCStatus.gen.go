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

#include <mavros_msgs/msg/esc_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/ESCStatus", ESCStatusTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/ESCStatus", ESCStatusTypeSupport)
}

type ESCStatus struct {
	Header std_msgs_msg.Header `yaml:"header"`
	EscStatus []ESCStatusItem `yaml:"esc_status"`
}

// NewESCStatus creates a new ESCStatus with default values.
func NewESCStatus() *ESCStatus {
	self := ESCStatus{}
	self.SetDefaults()
	return &self
}

func (t *ESCStatus) Clone() *ESCStatus {
	c := &ESCStatus{}
	c.Header = *t.Header.Clone()
	if t.EscStatus != nil {
		c.EscStatus = make([]ESCStatusItem, len(t.EscStatus))
		CloneESCStatusItemSlice(c.EscStatus, t.EscStatus)
	}
	return c
}

func (t *ESCStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ESCStatus) SetDefaults() {
	t.Header.SetDefaults()
	t.EscStatus = nil
}

func (t *ESCStatus) GetTypeSupport() types.MessageTypeSupport {
	return ESCStatusTypeSupport
}

// ESCStatusPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ESCStatusPublisher struct {
	*rclgo.Publisher
}

// NewESCStatusPublisher creates and returns a new publisher for the
// ESCStatus
func NewESCStatusPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ESCStatusPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ESCStatusTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ESCStatusPublisher{pub}, nil
}

func (p *ESCStatusPublisher) Publish(msg *ESCStatus) error {
	return p.Publisher.Publish(msg)
}

// ESCStatusSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ESCStatusSubscription struct {
	*rclgo.Subscription
}

// ESCStatusSubscriptionCallback type is used to provide a subscription
// handler function for a ESCStatusSubscription.
type ESCStatusSubscriptionCallback func(msg *ESCStatus, info *rclgo.MessageInfo, err error)

// NewESCStatusSubscription creates and returns a new subscription for the
// ESCStatus
func NewESCStatusSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ESCStatusSubscriptionCallback) (*ESCStatusSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ESCStatus
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ESCStatusTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ESCStatusSubscription{sub}, nil
}

func (s *ESCStatusSubscription) TakeMessage(out *ESCStatus) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneESCStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneESCStatusSlice(dst, src []ESCStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ESCStatusTypeSupport types.MessageTypeSupport = _ESCStatusTypeSupport{}

type _ESCStatusTypeSupport struct{}

func (t _ESCStatusTypeSupport) New() types.Message {
	return NewESCStatus()
}

func (t _ESCStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__ESCStatus
	return (unsafe.Pointer)(C.mavros_msgs__msg__ESCStatus__create())
}

func (t _ESCStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__ESCStatus__destroy((*C.mavros_msgs__msg__ESCStatus)(pointer_to_free))
}

func (t _ESCStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ESCStatus)
	mem := (*C.mavros_msgs__msg__ESCStatus)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	ESCStatusItem__Sequence_to_C(&mem.esc_status, m.EscStatus)
}

func (t _ESCStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ESCStatus)
	mem := (*C.mavros_msgs__msg__ESCStatus)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	ESCStatusItem__Sequence_to_Go(&m.EscStatus, mem.esc_status)
}

func (t _ESCStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__ESCStatus())
}

type CESCStatus = C.mavros_msgs__msg__ESCStatus
type CESCStatus__Sequence = C.mavros_msgs__msg__ESCStatus__Sequence

func ESCStatus__Sequence_to_Go(goSlice *[]ESCStatus, cSlice CESCStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ESCStatus, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ESCStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ESCStatus__Sequence_to_C(cSlice *CESCStatus__Sequence, goSlice []ESCStatus) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__ESCStatus)(C.malloc(C.sizeof_struct_mavros_msgs__msg__ESCStatus * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ESCStatusTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ESCStatus__Array_to_Go(goSlice []ESCStatus, cSlice []CESCStatus) {
	for i := 0; i < len(cSlice); i++ {
		ESCStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ESCStatus__Array_to_C(cSlice []CESCStatus, goSlice []ESCStatus) {
	for i := 0; i < len(goSlice); i++ {
		ESCStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
