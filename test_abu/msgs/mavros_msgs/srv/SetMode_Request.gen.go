// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/srv/set_mode.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/SetMode_Request", SetMode_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/SetMode_Request", SetMode_RequestTypeSupport)
}
const (
	SetMode_Request_MAV_MODE_PREFLIGHT uint8 = 0// basic modes from MAV_MODE
	SetMode_Request_MAV_MODE_STABILIZE_DISARMED uint8 = 80
	SetMode_Request_MAV_MODE_STABILIZE_ARMED uint8 = 208
	SetMode_Request_MAV_MODE_MANUAL_DISARMED uint8 = 64
	SetMode_Request_MAV_MODE_MANUAL_ARMED uint8 = 192
	SetMode_Request_MAV_MODE_GUIDED_DISARMED uint8 = 88
	SetMode_Request_MAV_MODE_GUIDED_ARMED uint8 = 216
	SetMode_Request_MAV_MODE_AUTO_DISARMED uint8 = 92
	SetMode_Request_MAV_MODE_AUTO_ARMED uint8 = 220
	SetMode_Request_MAV_MODE_TEST_DISARMED uint8 = 66
	SetMode_Request_MAV_MODE_TEST_ARMED uint8 = 194
)

type SetMode_Request struct {
	BaseMode uint8 `yaml:"base_mode"`// filled by MAV_MODE enum value or 0 if custom_mode != ''
	CustomMode string `yaml:"custom_mode"`// string mode representation or integer
}

// NewSetMode_Request creates a new SetMode_Request with default values.
func NewSetMode_Request() *SetMode_Request {
	self := SetMode_Request{}
	self.SetDefaults()
	return &self
}

func (t *SetMode_Request) Clone() *SetMode_Request {
	c := &SetMode_Request{}
	c.BaseMode = t.BaseMode
	c.CustomMode = t.CustomMode
	return c
}

func (t *SetMode_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SetMode_Request) SetDefaults() {
	t.BaseMode = 0
	t.CustomMode = ""
}

func (t *SetMode_Request) GetTypeSupport() types.MessageTypeSupport {
	return SetMode_RequestTypeSupport
}

// SetMode_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SetMode_RequestPublisher struct {
	*rclgo.Publisher
}

// NewSetMode_RequestPublisher creates and returns a new publisher for the
// SetMode_Request
func NewSetMode_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SetMode_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, SetMode_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetMode_RequestPublisher{pub}, nil
}

func (p *SetMode_RequestPublisher) Publish(msg *SetMode_Request) error {
	return p.Publisher.Publish(msg)
}

// SetMode_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SetMode_RequestSubscription struct {
	*rclgo.Subscription
}

// SetMode_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a SetMode_RequestSubscription.
type SetMode_RequestSubscriptionCallback func(msg *SetMode_Request, info *rclgo.MessageInfo, err error)

// NewSetMode_RequestSubscription creates and returns a new subscription for the
// SetMode_Request
func NewSetMode_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback SetMode_RequestSubscriptionCallback) (*SetMode_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SetMode_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SetMode_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &SetMode_RequestSubscription{sub}, nil
}

func (s *SetMode_RequestSubscription) TakeMessage(out *SetMode_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSetMode_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSetMode_RequestSlice(dst, src []SetMode_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SetMode_RequestTypeSupport types.MessageTypeSupport = _SetMode_RequestTypeSupport{}

type _SetMode_RequestTypeSupport struct{}

func (t _SetMode_RequestTypeSupport) New() types.Message {
	return NewSetMode_Request()
}

func (t _SetMode_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__SetMode_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__SetMode_Request__create())
}

func (t _SetMode_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__SetMode_Request__destroy((*C.mavros_msgs__srv__SetMode_Request)(pointer_to_free))
}

func (t _SetMode_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SetMode_Request)
	mem := (*C.mavros_msgs__srv__SetMode_Request)(dst)
	mem.base_mode = C.uint8_t(m.BaseMode)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.custom_mode), m.CustomMode)
}

func (t _SetMode_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SetMode_Request)
	mem := (*C.mavros_msgs__srv__SetMode_Request)(ros2_message_buffer)
	m.BaseMode = uint8(mem.base_mode)
	primitives.StringAsGoStruct(&m.CustomMode, unsafe.Pointer(&mem.custom_mode))
}

func (t _SetMode_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__SetMode_Request())
}

type CSetMode_Request = C.mavros_msgs__srv__SetMode_Request
type CSetMode_Request__Sequence = C.mavros_msgs__srv__SetMode_Request__Sequence

func SetMode_Request__Sequence_to_Go(goSlice *[]SetMode_Request, cSlice CSetMode_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SetMode_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SetMode_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SetMode_Request__Sequence_to_C(cSlice *CSetMode_Request__Sequence, goSlice []SetMode_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__SetMode_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__SetMode_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SetMode_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SetMode_Request__Array_to_Go(goSlice []SetMode_Request, cSlice []CSetMode_Request) {
	for i := 0; i < len(cSlice); i++ {
		SetMode_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SetMode_Request__Array_to_C(cSlice []CSetMode_Request, goSlice []SetMode_Request) {
	for i := 0; i < len(goSlice); i++ {
		SetMode_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
