// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	geometry_msgs_msg "test/msgs/geometry_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/msg/gimbal_device_set_attitude.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/GimbalDeviceSetAttitude", GimbalDeviceSetAttitudeTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/GimbalDeviceSetAttitude", GimbalDeviceSetAttitudeTypeSupport)
}
const (
	GimbalDeviceSetAttitude_FLAGS_RETRACT uint16 = 1// Based on GIMBAL_DEVICE_FLAGS_RETRACT. GIMBAL_DEVICE_FLAGS
	GimbalDeviceSetAttitude_FLAGS_NEUTRAL uint16 = 2// Based on GIMBAL_DEVICE_FLAGS_NEUTRAL
	GimbalDeviceSetAttitude_FLAGS_ROLL_LOCK uint16 = 4// Based on GIMBAL_DEVICE_FLAGS_ROLL_LOCK
	GimbalDeviceSetAttitude_FLAGS_PITCH_LOCK uint16 = 8// Based on GIMBAL_DEVICE_FLAGS_PITCH_LOCK
	GimbalDeviceSetAttitude_FLAGS_YAW_LOCK uint16 = 16// Based on GIMBAL_DEVICE_FLAGS_YAW_LOCK
)

type GimbalDeviceSetAttitude struct {
	TargetSystem uint8 `yaml:"target_system"`// System ID
	TargetComponent uint8 `yaml:"target_component"`// Component ID
	Flags uint16 `yaml:"flags"`// Low level gimbal flags (bitwise) - See GIMBAL_DEVICE_FLAGS
	Q geometry_msgs_msg.Quaternion `yaml:"q"`// Quaternion, x, y, z, w (0 0 0 1 is the null-rotation, the frame is depends on whether the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set)
	AngularVelocityX float32 `yaml:"angular_velocity_x"`// X component of angular velocity, positive is rolling to the right, NaN to be ignored.
	AngularVelocityY float32 `yaml:"angular_velocity_y"`// Y component of angular velocity, positive is pitching up, NaN to be ignored.
	AngularVelocityZ float32 `yaml:"angular_velocity_z"`// Z component of angular velocity, positive is yawing to the right, NaN to be ignored.
}

// NewGimbalDeviceSetAttitude creates a new GimbalDeviceSetAttitude with default values.
func NewGimbalDeviceSetAttitude() *GimbalDeviceSetAttitude {
	self := GimbalDeviceSetAttitude{}
	self.SetDefaults()
	return &self
}

func (t *GimbalDeviceSetAttitude) Clone() *GimbalDeviceSetAttitude {
	c := &GimbalDeviceSetAttitude{}
	c.TargetSystem = t.TargetSystem
	c.TargetComponent = t.TargetComponent
	c.Flags = t.Flags
	c.Q = *t.Q.Clone()
	c.AngularVelocityX = t.AngularVelocityX
	c.AngularVelocityY = t.AngularVelocityY
	c.AngularVelocityZ = t.AngularVelocityZ
	return c
}

func (t *GimbalDeviceSetAttitude) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GimbalDeviceSetAttitude) SetDefaults() {
	t.TargetSystem = 0
	t.TargetComponent = 0
	t.Flags = 0
	t.Q.SetDefaults()
	t.AngularVelocityX = 0
	t.AngularVelocityY = 0
	t.AngularVelocityZ = 0
}

func (t *GimbalDeviceSetAttitude) GetTypeSupport() types.MessageTypeSupport {
	return GimbalDeviceSetAttitudeTypeSupport
}

// GimbalDeviceSetAttitudePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GimbalDeviceSetAttitudePublisher struct {
	*rclgo.Publisher
}

// NewGimbalDeviceSetAttitudePublisher creates and returns a new publisher for the
// GimbalDeviceSetAttitude
func NewGimbalDeviceSetAttitudePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GimbalDeviceSetAttitudePublisher, error) {
	pub, err := node.NewPublisher(topic_name, GimbalDeviceSetAttitudeTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GimbalDeviceSetAttitudePublisher{pub}, nil
}

func (p *GimbalDeviceSetAttitudePublisher) Publish(msg *GimbalDeviceSetAttitude) error {
	return p.Publisher.Publish(msg)
}

// GimbalDeviceSetAttitudeSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GimbalDeviceSetAttitudeSubscription struct {
	*rclgo.Subscription
}

// GimbalDeviceSetAttitudeSubscriptionCallback type is used to provide a subscription
// handler function for a GimbalDeviceSetAttitudeSubscription.
type GimbalDeviceSetAttitudeSubscriptionCallback func(msg *GimbalDeviceSetAttitude, info *rclgo.MessageInfo, err error)

// NewGimbalDeviceSetAttitudeSubscription creates and returns a new subscription for the
// GimbalDeviceSetAttitude
func NewGimbalDeviceSetAttitudeSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GimbalDeviceSetAttitudeSubscriptionCallback) (*GimbalDeviceSetAttitudeSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GimbalDeviceSetAttitude
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GimbalDeviceSetAttitudeTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GimbalDeviceSetAttitudeSubscription{sub}, nil
}

func (s *GimbalDeviceSetAttitudeSubscription) TakeMessage(out *GimbalDeviceSetAttitude) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGimbalDeviceSetAttitudeSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGimbalDeviceSetAttitudeSlice(dst, src []GimbalDeviceSetAttitude) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GimbalDeviceSetAttitudeTypeSupport types.MessageTypeSupport = _GimbalDeviceSetAttitudeTypeSupport{}

type _GimbalDeviceSetAttitudeTypeSupport struct{}

func (t _GimbalDeviceSetAttitudeTypeSupport) New() types.Message {
	return NewGimbalDeviceSetAttitude()
}

func (t _GimbalDeviceSetAttitudeTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__GimbalDeviceSetAttitude
	return (unsafe.Pointer)(C.mavros_msgs__msg__GimbalDeviceSetAttitude__create())
}

func (t _GimbalDeviceSetAttitudeTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__GimbalDeviceSetAttitude__destroy((*C.mavros_msgs__msg__GimbalDeviceSetAttitude)(pointer_to_free))
}

func (t _GimbalDeviceSetAttitudeTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GimbalDeviceSetAttitude)
	mem := (*C.mavros_msgs__msg__GimbalDeviceSetAttitude)(dst)
	mem.target_system = C.uint8_t(m.TargetSystem)
	mem.target_component = C.uint8_t(m.TargetComponent)
	mem.flags = C.uint16_t(m.Flags)
	geometry_msgs_msg.QuaternionTypeSupport.AsCStruct(unsafe.Pointer(&mem.q), &m.Q)
	mem.angular_velocity_x = C.float(m.AngularVelocityX)
	mem.angular_velocity_y = C.float(m.AngularVelocityY)
	mem.angular_velocity_z = C.float(m.AngularVelocityZ)
}

func (t _GimbalDeviceSetAttitudeTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GimbalDeviceSetAttitude)
	mem := (*C.mavros_msgs__msg__GimbalDeviceSetAttitude)(ros2_message_buffer)
	m.TargetSystem = uint8(mem.target_system)
	m.TargetComponent = uint8(mem.target_component)
	m.Flags = uint16(mem.flags)
	geometry_msgs_msg.QuaternionTypeSupport.AsGoStruct(&m.Q, unsafe.Pointer(&mem.q))
	m.AngularVelocityX = float32(mem.angular_velocity_x)
	m.AngularVelocityY = float32(mem.angular_velocity_y)
	m.AngularVelocityZ = float32(mem.angular_velocity_z)
}

func (t _GimbalDeviceSetAttitudeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__GimbalDeviceSetAttitude())
}

type CGimbalDeviceSetAttitude = C.mavros_msgs__msg__GimbalDeviceSetAttitude
type CGimbalDeviceSetAttitude__Sequence = C.mavros_msgs__msg__GimbalDeviceSetAttitude__Sequence

func GimbalDeviceSetAttitude__Sequence_to_Go(goSlice *[]GimbalDeviceSetAttitude, cSlice CGimbalDeviceSetAttitude__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GimbalDeviceSetAttitude, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GimbalDeviceSetAttitudeTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GimbalDeviceSetAttitude__Sequence_to_C(cSlice *CGimbalDeviceSetAttitude__Sequence, goSlice []GimbalDeviceSetAttitude) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__GimbalDeviceSetAttitude)(C.malloc(C.sizeof_struct_mavros_msgs__msg__GimbalDeviceSetAttitude * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GimbalDeviceSetAttitudeTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GimbalDeviceSetAttitude__Array_to_Go(goSlice []GimbalDeviceSetAttitude, cSlice []CGimbalDeviceSetAttitude) {
	for i := 0; i < len(cSlice); i++ {
		GimbalDeviceSetAttitudeTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GimbalDeviceSetAttitude__Array_to_C(cSlice []CGimbalDeviceSetAttitude, goSlice []GimbalDeviceSetAttitude) {
	for i := 0; i < len(goSlice); i++ {
		GimbalDeviceSetAttitudeTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
