// Code generated by rclgo-gen. DO NOT EDIT.

package geographic_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <geographic_msgs/msg/key_value.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("geographic_msgs/KeyValue", KeyValueTypeSupport)
	typemap.RegisterMessage("geographic_msgs/msg/KeyValue", KeyValueTypeSupport)
}

type KeyValue struct {
	Key string `yaml:"key"`// tag label
	Value string `yaml:"value"`// corresponding value
}

// NewKeyValue creates a new KeyValue with default values.
func NewKeyValue() *KeyValue {
	self := KeyValue{}
	self.SetDefaults()
	return &self
}

func (t *KeyValue) Clone() *KeyValue {
	c := &KeyValue{}
	c.Key = t.Key
	c.Value = t.Value
	return c
}

func (t *KeyValue) CloneMsg() types.Message {
	return t.Clone()
}

func (t *KeyValue) SetDefaults() {
	t.Key = ""
	t.Value = ""
}

func (t *KeyValue) GetTypeSupport() types.MessageTypeSupport {
	return KeyValueTypeSupport
}

// KeyValuePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type KeyValuePublisher struct {
	*rclgo.Publisher
}

// NewKeyValuePublisher creates and returns a new publisher for the
// KeyValue
func NewKeyValuePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*KeyValuePublisher, error) {
	pub, err := node.NewPublisher(topic_name, KeyValueTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &KeyValuePublisher{pub}, nil
}

func (p *KeyValuePublisher) Publish(msg *KeyValue) error {
	return p.Publisher.Publish(msg)
}

// KeyValueSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type KeyValueSubscription struct {
	*rclgo.Subscription
}

// KeyValueSubscriptionCallback type is used to provide a subscription
// handler function for a KeyValueSubscription.
type KeyValueSubscriptionCallback func(msg *KeyValue, info *rclgo.MessageInfo, err error)

// NewKeyValueSubscription creates and returns a new subscription for the
// KeyValue
func NewKeyValueSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback KeyValueSubscriptionCallback) (*KeyValueSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg KeyValue
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, KeyValueTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &KeyValueSubscription{sub}, nil
}

func (s *KeyValueSubscription) TakeMessage(out *KeyValue) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneKeyValueSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneKeyValueSlice(dst, src []KeyValue) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var KeyValueTypeSupport types.MessageTypeSupport = _KeyValueTypeSupport{}

type _KeyValueTypeSupport struct{}

func (t _KeyValueTypeSupport) New() types.Message {
	return NewKeyValue()
}

func (t _KeyValueTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.geographic_msgs__msg__KeyValue
	return (unsafe.Pointer)(C.geographic_msgs__msg__KeyValue__create())
}

func (t _KeyValueTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.geographic_msgs__msg__KeyValue__destroy((*C.geographic_msgs__msg__KeyValue)(pointer_to_free))
}

func (t _KeyValueTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*KeyValue)
	mem := (*C.geographic_msgs__msg__KeyValue)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.key), m.Key)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.value), m.Value)
}

func (t _KeyValueTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*KeyValue)
	mem := (*C.geographic_msgs__msg__KeyValue)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Key, unsafe.Pointer(&mem.key))
	primitives.StringAsGoStruct(&m.Value, unsafe.Pointer(&mem.value))
}

func (t _KeyValueTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__geographic_msgs__msg__KeyValue())
}

type CKeyValue = C.geographic_msgs__msg__KeyValue
type CKeyValue__Sequence = C.geographic_msgs__msg__KeyValue__Sequence

func KeyValue__Sequence_to_Go(goSlice *[]KeyValue, cSlice CKeyValue__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]KeyValue, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		KeyValueTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func KeyValue__Sequence_to_C(cSlice *CKeyValue__Sequence, goSlice []KeyValue) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.geographic_msgs__msg__KeyValue)(C.malloc(C.sizeof_struct_geographic_msgs__msg__KeyValue * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		KeyValueTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func KeyValue__Array_to_Go(goSlice []KeyValue, cSlice []CKeyValue) {
	for i := 0; i < len(cSlice); i++ {
		KeyValueTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func KeyValue__Array_to_C(cSlice []CKeyValue, goSlice []KeyValue) {
	for i := 0; i < len(goSlice); i++ {
		KeyValueTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
