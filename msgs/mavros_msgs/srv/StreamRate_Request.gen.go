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

#include <mavros_msgs/srv/stream_rate.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/StreamRate_Request", StreamRate_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/StreamRate_Request", StreamRate_RequestTypeSupport)
}
const (
	StreamRate_Request_STREAM_ALL uint8 = 0
	StreamRate_Request_STREAM_RAW_SENSORS uint8 = 1
	StreamRate_Request_STREAM_EXTENDED_STATUS uint8 = 2
	StreamRate_Request_STREAM_RC_CHANNELS uint8 = 3
	StreamRate_Request_STREAM_RAW_CONTROLLER uint8 = 4
	StreamRate_Request_STREAM_POSITION uint8 = 6
	StreamRate_Request_STREAM_EXTRA1 uint8 = 10
	StreamRate_Request_STREAM_EXTRA2 uint8 = 11
	StreamRate_Request_STREAM_EXTRA3 uint8 = 12
)

type StreamRate_Request struct {
	StreamId uint8 `yaml:"stream_id"`
	MessageRate uint16 `yaml:"message_rate"`
	OnOff bool `yaml:"on_off"`
}

// NewStreamRate_Request creates a new StreamRate_Request with default values.
func NewStreamRate_Request() *StreamRate_Request {
	self := StreamRate_Request{}
	self.SetDefaults()
	return &self
}

func (t *StreamRate_Request) Clone() *StreamRate_Request {
	c := &StreamRate_Request{}
	c.StreamId = t.StreamId
	c.MessageRate = t.MessageRate
	c.OnOff = t.OnOff
	return c
}

func (t *StreamRate_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *StreamRate_Request) SetDefaults() {
	t.StreamId = 0
	t.MessageRate = 0
	t.OnOff = false
}

func (t *StreamRate_Request) GetTypeSupport() types.MessageTypeSupport {
	return StreamRate_RequestTypeSupport
}

// StreamRate_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type StreamRate_RequestPublisher struct {
	*rclgo.Publisher
}

// NewStreamRate_RequestPublisher creates and returns a new publisher for the
// StreamRate_Request
func NewStreamRate_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*StreamRate_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, StreamRate_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &StreamRate_RequestPublisher{pub}, nil
}

func (p *StreamRate_RequestPublisher) Publish(msg *StreamRate_Request) error {
	return p.Publisher.Publish(msg)
}

// StreamRate_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type StreamRate_RequestSubscription struct {
	*rclgo.Subscription
}

// StreamRate_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a StreamRate_RequestSubscription.
type StreamRate_RequestSubscriptionCallback func(msg *StreamRate_Request, info *rclgo.MessageInfo, err error)

// NewStreamRate_RequestSubscription creates and returns a new subscription for the
// StreamRate_Request
func NewStreamRate_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback StreamRate_RequestSubscriptionCallback) (*StreamRate_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg StreamRate_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, StreamRate_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &StreamRate_RequestSubscription{sub}, nil
}

func (s *StreamRate_RequestSubscription) TakeMessage(out *StreamRate_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneStreamRate_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneStreamRate_RequestSlice(dst, src []StreamRate_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var StreamRate_RequestTypeSupport types.MessageTypeSupport = _StreamRate_RequestTypeSupport{}

type _StreamRate_RequestTypeSupport struct{}

func (t _StreamRate_RequestTypeSupport) New() types.Message {
	return NewStreamRate_Request()
}

func (t _StreamRate_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__StreamRate_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__StreamRate_Request__create())
}

func (t _StreamRate_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__StreamRate_Request__destroy((*C.mavros_msgs__srv__StreamRate_Request)(pointer_to_free))
}

func (t _StreamRate_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*StreamRate_Request)
	mem := (*C.mavros_msgs__srv__StreamRate_Request)(dst)
	mem.stream_id = C.uint8_t(m.StreamId)
	mem.message_rate = C.uint16_t(m.MessageRate)
	mem.on_off = C.bool(m.OnOff)
}

func (t _StreamRate_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*StreamRate_Request)
	mem := (*C.mavros_msgs__srv__StreamRate_Request)(ros2_message_buffer)
	m.StreamId = uint8(mem.stream_id)
	m.MessageRate = uint16(mem.message_rate)
	m.OnOff = bool(mem.on_off)
}

func (t _StreamRate_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__StreamRate_Request())
}

type CStreamRate_Request = C.mavros_msgs__srv__StreamRate_Request
type CStreamRate_Request__Sequence = C.mavros_msgs__srv__StreamRate_Request__Sequence

func StreamRate_Request__Sequence_to_Go(goSlice *[]StreamRate_Request, cSlice CStreamRate_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]StreamRate_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		StreamRate_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func StreamRate_Request__Sequence_to_C(cSlice *CStreamRate_Request__Sequence, goSlice []StreamRate_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__StreamRate_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__StreamRate_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		StreamRate_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func StreamRate_Request__Array_to_Go(goSlice []StreamRate_Request, cSlice []CStreamRate_Request) {
	for i := 0; i < len(cSlice); i++ {
		StreamRate_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func StreamRate_Request__Array_to_C(cSlice []CStreamRate_Request, goSlice []StreamRate_Request) {
	for i := 0; i < len(goSlice); i++ {
		StreamRate_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
