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

#include <mavros_msgs/msg/esc_telemetry_item.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/ESCTelemetryItem", ESCTelemetryItemTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/ESCTelemetryItem", ESCTelemetryItemTypeSupport)
}

type ESCTelemetryItem struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Temperature float32 `yaml:"temperature"`// deg C
	Voltage float32 `yaml:"voltage"`// V
	Current float32 `yaml:"current"`// A
	Totalcurrent float32 `yaml:"totalcurrent"`// Ah
	Rpm int32 `yaml:"rpm"`// 1/min
	Count uint16 `yaml:"count"`// count of telemetry packets
}

// NewESCTelemetryItem creates a new ESCTelemetryItem with default values.
func NewESCTelemetryItem() *ESCTelemetryItem {
	self := ESCTelemetryItem{}
	self.SetDefaults()
	return &self
}

func (t *ESCTelemetryItem) Clone() *ESCTelemetryItem {
	c := &ESCTelemetryItem{}
	c.Header = *t.Header.Clone()
	c.Temperature = t.Temperature
	c.Voltage = t.Voltage
	c.Current = t.Current
	c.Totalcurrent = t.Totalcurrent
	c.Rpm = t.Rpm
	c.Count = t.Count
	return c
}

func (t *ESCTelemetryItem) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ESCTelemetryItem) SetDefaults() {
	t.Header.SetDefaults()
	t.Temperature = 0
	t.Voltage = 0
	t.Current = 0
	t.Totalcurrent = 0
	t.Rpm = 0
	t.Count = 0
}

func (t *ESCTelemetryItem) GetTypeSupport() types.MessageTypeSupport {
	return ESCTelemetryItemTypeSupport
}

// ESCTelemetryItemPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ESCTelemetryItemPublisher struct {
	*rclgo.Publisher
}

// NewESCTelemetryItemPublisher creates and returns a new publisher for the
// ESCTelemetryItem
func NewESCTelemetryItemPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ESCTelemetryItemPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ESCTelemetryItemTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ESCTelemetryItemPublisher{pub}, nil
}

func (p *ESCTelemetryItemPublisher) Publish(msg *ESCTelemetryItem) error {
	return p.Publisher.Publish(msg)
}

// ESCTelemetryItemSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ESCTelemetryItemSubscription struct {
	*rclgo.Subscription
}

// ESCTelemetryItemSubscriptionCallback type is used to provide a subscription
// handler function for a ESCTelemetryItemSubscription.
type ESCTelemetryItemSubscriptionCallback func(msg *ESCTelemetryItem, info *rclgo.MessageInfo, err error)

// NewESCTelemetryItemSubscription creates and returns a new subscription for the
// ESCTelemetryItem
func NewESCTelemetryItemSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ESCTelemetryItemSubscriptionCallback) (*ESCTelemetryItemSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ESCTelemetryItem
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ESCTelemetryItemTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ESCTelemetryItemSubscription{sub}, nil
}

func (s *ESCTelemetryItemSubscription) TakeMessage(out *ESCTelemetryItem) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneESCTelemetryItemSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneESCTelemetryItemSlice(dst, src []ESCTelemetryItem) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ESCTelemetryItemTypeSupport types.MessageTypeSupport = _ESCTelemetryItemTypeSupport{}

type _ESCTelemetryItemTypeSupport struct{}

func (t _ESCTelemetryItemTypeSupport) New() types.Message {
	return NewESCTelemetryItem()
}

func (t _ESCTelemetryItemTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__ESCTelemetryItem
	return (unsafe.Pointer)(C.mavros_msgs__msg__ESCTelemetryItem__create())
}

func (t _ESCTelemetryItemTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__ESCTelemetryItem__destroy((*C.mavros_msgs__msg__ESCTelemetryItem)(pointer_to_free))
}

func (t _ESCTelemetryItemTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ESCTelemetryItem)
	mem := (*C.mavros_msgs__msg__ESCTelemetryItem)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.temperature = C.float(m.Temperature)
	mem.voltage = C.float(m.Voltage)
	mem.current = C.float(m.Current)
	mem.totalcurrent = C.float(m.Totalcurrent)
	mem.rpm = C.int32_t(m.Rpm)
	mem.count = C.uint16_t(m.Count)
}

func (t _ESCTelemetryItemTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ESCTelemetryItem)
	mem := (*C.mavros_msgs__msg__ESCTelemetryItem)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Temperature = float32(mem.temperature)
	m.Voltage = float32(mem.voltage)
	m.Current = float32(mem.current)
	m.Totalcurrent = float32(mem.totalcurrent)
	m.Rpm = int32(mem.rpm)
	m.Count = uint16(mem.count)
}

func (t _ESCTelemetryItemTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__ESCTelemetryItem())
}

type CESCTelemetryItem = C.mavros_msgs__msg__ESCTelemetryItem
type CESCTelemetryItem__Sequence = C.mavros_msgs__msg__ESCTelemetryItem__Sequence

func ESCTelemetryItem__Sequence_to_Go(goSlice *[]ESCTelemetryItem, cSlice CESCTelemetryItem__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ESCTelemetryItem, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ESCTelemetryItemTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ESCTelemetryItem__Sequence_to_C(cSlice *CESCTelemetryItem__Sequence, goSlice []ESCTelemetryItem) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__ESCTelemetryItem)(C.malloc(C.sizeof_struct_mavros_msgs__msg__ESCTelemetryItem * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ESCTelemetryItemTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ESCTelemetryItem__Array_to_Go(goSlice []ESCTelemetryItem, cSlice []CESCTelemetryItem) {
	for i := 0; i < len(cSlice); i++ {
		ESCTelemetryItemTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ESCTelemetryItem__Array_to_C(cSlice []CESCTelemetryItem, goSlice []ESCTelemetryItem) {
	for i := 0; i < len(goSlice); i++ {
		ESCTelemetryItemTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
