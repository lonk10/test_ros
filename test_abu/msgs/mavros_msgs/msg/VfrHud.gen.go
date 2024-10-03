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

#include <mavros_msgs/msg/vfr_hud.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/VfrHud", VfrHudTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/VfrHud", VfrHudTypeSupport)
}

type VfrHud struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Airspeed float32 `yaml:"airspeed"`// m/s
	Groundspeed float32 `yaml:"groundspeed"`// m/s
	Heading int16 `yaml:"heading"`// degrees 0..360
	Throttle float32 `yaml:"throttle"`// normalized to 0.0..1.0
	Altitude float32 `yaml:"altitude"`// MSL
	Climb float32 `yaml:"climb"`// current climb rate m/s
}

// NewVfrHud creates a new VfrHud with default values.
func NewVfrHud() *VfrHud {
	self := VfrHud{}
	self.SetDefaults()
	return &self
}

func (t *VfrHud) Clone() *VfrHud {
	c := &VfrHud{}
	c.Header = *t.Header.Clone()
	c.Airspeed = t.Airspeed
	c.Groundspeed = t.Groundspeed
	c.Heading = t.Heading
	c.Throttle = t.Throttle
	c.Altitude = t.Altitude
	c.Climb = t.Climb
	return c
}

func (t *VfrHud) CloneMsg() types.Message {
	return t.Clone()
}

func (t *VfrHud) SetDefaults() {
	t.Header.SetDefaults()
	t.Airspeed = 0
	t.Groundspeed = 0
	t.Heading = 0
	t.Throttle = 0
	t.Altitude = 0
	t.Climb = 0
}

func (t *VfrHud) GetTypeSupport() types.MessageTypeSupport {
	return VfrHudTypeSupport
}

// VfrHudPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type VfrHudPublisher struct {
	*rclgo.Publisher
}

// NewVfrHudPublisher creates and returns a new publisher for the
// VfrHud
func NewVfrHudPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*VfrHudPublisher, error) {
	pub, err := node.NewPublisher(topic_name, VfrHudTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &VfrHudPublisher{pub}, nil
}

func (p *VfrHudPublisher) Publish(msg *VfrHud) error {
	return p.Publisher.Publish(msg)
}

// VfrHudSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type VfrHudSubscription struct {
	*rclgo.Subscription
}

// VfrHudSubscriptionCallback type is used to provide a subscription
// handler function for a VfrHudSubscription.
type VfrHudSubscriptionCallback func(msg *VfrHud, info *rclgo.MessageInfo, err error)

// NewVfrHudSubscription creates and returns a new subscription for the
// VfrHud
func NewVfrHudSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback VfrHudSubscriptionCallback) (*VfrHudSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg VfrHud
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, VfrHudTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &VfrHudSubscription{sub}, nil
}

func (s *VfrHudSubscription) TakeMessage(out *VfrHud) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneVfrHudSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneVfrHudSlice(dst, src []VfrHud) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var VfrHudTypeSupport types.MessageTypeSupport = _VfrHudTypeSupport{}

type _VfrHudTypeSupport struct{}

func (t _VfrHudTypeSupport) New() types.Message {
	return NewVfrHud()
}

func (t _VfrHudTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__VfrHud
	return (unsafe.Pointer)(C.mavros_msgs__msg__VfrHud__create())
}

func (t _VfrHudTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__VfrHud__destroy((*C.mavros_msgs__msg__VfrHud)(pointer_to_free))
}

func (t _VfrHudTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*VfrHud)
	mem := (*C.mavros_msgs__msg__VfrHud)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.airspeed = C.float(m.Airspeed)
	mem.groundspeed = C.float(m.Groundspeed)
	mem.heading = C.int16_t(m.Heading)
	mem.throttle = C.float(m.Throttle)
	mem.altitude = C.float(m.Altitude)
	mem.climb = C.float(m.Climb)
}

func (t _VfrHudTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*VfrHud)
	mem := (*C.mavros_msgs__msg__VfrHud)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Airspeed = float32(mem.airspeed)
	m.Groundspeed = float32(mem.groundspeed)
	m.Heading = int16(mem.heading)
	m.Throttle = float32(mem.throttle)
	m.Altitude = float32(mem.altitude)
	m.Climb = float32(mem.climb)
}

func (t _VfrHudTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__VfrHud())
}

type CVfrHud = C.mavros_msgs__msg__VfrHud
type CVfrHud__Sequence = C.mavros_msgs__msg__VfrHud__Sequence

func VfrHud__Sequence_to_Go(goSlice *[]VfrHud, cSlice CVfrHud__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]VfrHud, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		VfrHudTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func VfrHud__Sequence_to_C(cSlice *CVfrHud__Sequence, goSlice []VfrHud) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__VfrHud)(C.malloc(C.sizeof_struct_mavros_msgs__msg__VfrHud * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		VfrHudTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func VfrHud__Array_to_Go(goSlice []VfrHud, cSlice []CVfrHud) {
	for i := 0; i < len(cSlice); i++ {
		VfrHudTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func VfrHud__Array_to_C(cSlice []CVfrHud, goSlice []VfrHud) {
	for i := 0; i < len(goSlice); i++ {
		VfrHudTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
