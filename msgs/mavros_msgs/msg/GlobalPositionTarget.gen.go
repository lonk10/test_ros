// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	geometry_msgs_msg "test/msgs/geometry_msgs/msg"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/msg/global_position_target.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/GlobalPositionTarget", GlobalPositionTargetTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/GlobalPositionTarget", GlobalPositionTargetTypeSupport)
}
const (
	GlobalPositionTarget_FRAME_GLOBAL_INT uint8 = 5
	GlobalPositionTarget_FRAME_GLOBAL_REL_ALT uint8 = 6
	GlobalPositionTarget_FRAME_GLOBAL_TERRAIN_ALT uint8 = 11
	GlobalPositionTarget_IGNORE_LATITUDE uint16 = 1// Position ignore flags
	GlobalPositionTarget_IGNORE_LONGITUDE uint16 = 2
	GlobalPositionTarget_IGNORE_ALTITUDE uint16 = 4
	GlobalPositionTarget_IGNORE_VX uint16 = 8// Velocity vector ignore flags
	GlobalPositionTarget_IGNORE_VY uint16 = 16
	GlobalPositionTarget_IGNORE_VZ uint16 = 32
	GlobalPositionTarget_IGNORE_AFX uint16 = 64// Acceleration/Force vector ignore flags
	GlobalPositionTarget_IGNORE_AFY uint16 = 128
	GlobalPositionTarget_IGNORE_AFZ uint16 = 256
	GlobalPositionTarget_FORCE uint16 = 512// Force in af vector flag
	GlobalPositionTarget_IGNORE_YAW uint16 = 1024
	GlobalPositionTarget_IGNORE_YAW_RATE uint16 = 2048
)

type GlobalPositionTarget struct {
	Header std_msgs_msg.Header `yaml:"header"`
	CoordinateFrame uint8 `yaml:"coordinate_frame"`
	TypeMask uint16 `yaml:"type_mask"`
	Latitude float64 `yaml:"latitude"`
	Longitude float64 `yaml:"longitude"`
	Altitude float32 `yaml:"altitude"`// in meters, AMSL or above terrain
	Velocity geometry_msgs_msg.Vector3 `yaml:"velocity"`
	AccelerationOrForce geometry_msgs_msg.Vector3 `yaml:"acceleration_or_force"`
	Yaw float32 `yaml:"yaw"`
	YawRate float32 `yaml:"yaw_rate"`
}

// NewGlobalPositionTarget creates a new GlobalPositionTarget with default values.
func NewGlobalPositionTarget() *GlobalPositionTarget {
	self := GlobalPositionTarget{}
	self.SetDefaults()
	return &self
}

func (t *GlobalPositionTarget) Clone() *GlobalPositionTarget {
	c := &GlobalPositionTarget{}
	c.Header = *t.Header.Clone()
	c.CoordinateFrame = t.CoordinateFrame
	c.TypeMask = t.TypeMask
	c.Latitude = t.Latitude
	c.Longitude = t.Longitude
	c.Altitude = t.Altitude
	c.Velocity = *t.Velocity.Clone()
	c.AccelerationOrForce = *t.AccelerationOrForce.Clone()
	c.Yaw = t.Yaw
	c.YawRate = t.YawRate
	return c
}

func (t *GlobalPositionTarget) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GlobalPositionTarget) SetDefaults() {
	t.Header.SetDefaults()
	t.CoordinateFrame = 0
	t.TypeMask = 0
	t.Latitude = 0
	t.Longitude = 0
	t.Altitude = 0
	t.Velocity.SetDefaults()
	t.AccelerationOrForce.SetDefaults()
	t.Yaw = 0
	t.YawRate = 0
}

func (t *GlobalPositionTarget) GetTypeSupport() types.MessageTypeSupport {
	return GlobalPositionTargetTypeSupport
}

// GlobalPositionTargetPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GlobalPositionTargetPublisher struct {
	*rclgo.Publisher
}

// NewGlobalPositionTargetPublisher creates and returns a new publisher for the
// GlobalPositionTarget
func NewGlobalPositionTargetPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GlobalPositionTargetPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GlobalPositionTargetTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GlobalPositionTargetPublisher{pub}, nil
}

func (p *GlobalPositionTargetPublisher) Publish(msg *GlobalPositionTarget) error {
	return p.Publisher.Publish(msg)
}

// GlobalPositionTargetSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GlobalPositionTargetSubscription struct {
	*rclgo.Subscription
}

// GlobalPositionTargetSubscriptionCallback type is used to provide a subscription
// handler function for a GlobalPositionTargetSubscription.
type GlobalPositionTargetSubscriptionCallback func(msg *GlobalPositionTarget, info *rclgo.MessageInfo, err error)

// NewGlobalPositionTargetSubscription creates and returns a new subscription for the
// GlobalPositionTarget
func NewGlobalPositionTargetSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GlobalPositionTargetSubscriptionCallback) (*GlobalPositionTargetSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GlobalPositionTarget
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GlobalPositionTargetTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GlobalPositionTargetSubscription{sub}, nil
}

func (s *GlobalPositionTargetSubscription) TakeMessage(out *GlobalPositionTarget) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGlobalPositionTargetSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGlobalPositionTargetSlice(dst, src []GlobalPositionTarget) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GlobalPositionTargetTypeSupport types.MessageTypeSupport = _GlobalPositionTargetTypeSupport{}

type _GlobalPositionTargetTypeSupport struct{}

func (t _GlobalPositionTargetTypeSupport) New() types.Message {
	return NewGlobalPositionTarget()
}

func (t _GlobalPositionTargetTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__GlobalPositionTarget
	return (unsafe.Pointer)(C.mavros_msgs__msg__GlobalPositionTarget__create())
}

func (t _GlobalPositionTargetTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__GlobalPositionTarget__destroy((*C.mavros_msgs__msg__GlobalPositionTarget)(pointer_to_free))
}

func (t _GlobalPositionTargetTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GlobalPositionTarget)
	mem := (*C.mavros_msgs__msg__GlobalPositionTarget)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.coordinate_frame = C.uint8_t(m.CoordinateFrame)
	mem.type_mask = C.uint16_t(m.TypeMask)
	mem.latitude = C.double(m.Latitude)
	mem.longitude = C.double(m.Longitude)
	mem.altitude = C.float(m.Altitude)
	geometry_msgs_msg.Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.velocity), &m.Velocity)
	geometry_msgs_msg.Vector3TypeSupport.AsCStruct(unsafe.Pointer(&mem.acceleration_or_force), &m.AccelerationOrForce)
	mem.yaw = C.float(m.Yaw)
	mem.yaw_rate = C.float(m.YawRate)
}

func (t _GlobalPositionTargetTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GlobalPositionTarget)
	mem := (*C.mavros_msgs__msg__GlobalPositionTarget)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.CoordinateFrame = uint8(mem.coordinate_frame)
	m.TypeMask = uint16(mem.type_mask)
	m.Latitude = float64(mem.latitude)
	m.Longitude = float64(mem.longitude)
	m.Altitude = float32(mem.altitude)
	geometry_msgs_msg.Vector3TypeSupport.AsGoStruct(&m.Velocity, unsafe.Pointer(&mem.velocity))
	geometry_msgs_msg.Vector3TypeSupport.AsGoStruct(&m.AccelerationOrForce, unsafe.Pointer(&mem.acceleration_or_force))
	m.Yaw = float32(mem.yaw)
	m.YawRate = float32(mem.yaw_rate)
}

func (t _GlobalPositionTargetTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__GlobalPositionTarget())
}

type CGlobalPositionTarget = C.mavros_msgs__msg__GlobalPositionTarget
type CGlobalPositionTarget__Sequence = C.mavros_msgs__msg__GlobalPositionTarget__Sequence

func GlobalPositionTarget__Sequence_to_Go(goSlice *[]GlobalPositionTarget, cSlice CGlobalPositionTarget__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GlobalPositionTarget, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GlobalPositionTargetTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GlobalPositionTarget__Sequence_to_C(cSlice *CGlobalPositionTarget__Sequence, goSlice []GlobalPositionTarget) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__GlobalPositionTarget)(C.malloc(C.sizeof_struct_mavros_msgs__msg__GlobalPositionTarget * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GlobalPositionTargetTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GlobalPositionTarget__Array_to_Go(goSlice []GlobalPositionTarget, cSlice []CGlobalPositionTarget) {
	for i := 0; i < len(cSlice); i++ {
		GlobalPositionTargetTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GlobalPositionTarget__Array_to_C(cSlice []CGlobalPositionTarget, goSlice []GlobalPositionTarget) {
	for i := 0; i < len(goSlice); i++ {
		GlobalPositionTargetTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
