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

#include <mavros_msgs/msg/gimbal_device_attitude_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/GimbalDeviceAttitudeStatus", GimbalDeviceAttitudeStatusTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/GimbalDeviceAttitudeStatus", GimbalDeviceAttitudeStatusTypeSupport)
}
const (
	GimbalDeviceAttitudeStatus_FLAGS_RETRACT uint16 = 1// Set to retracted safe position (no stabilization), takes presedence over all other flags.. GIMBAL_DEVICE_FLAGS
	GimbalDeviceAttitudeStatus_FLAGS_NEUTRAL uint16 = 2// Set to neutral/default position, taking precedence over all other flags except RETRACT. Neutral is commonly forward-facing and horizontal (pitch=yaw=0) but may be any orientation.
	GimbalDeviceAttitudeStatus_FLAGS_ROLL_LOCK uint16 = 4// Lock roll angle to absolute angle relative to horizon (not relative to drone). This is generally the default with a stabilizing gimbal.
	GimbalDeviceAttitudeStatus_FLAGS_PITCH_LOCK uint16 = 8// Lock pitch angle to absolute angle relative to horizon (not relative to drone). This is generally the default.
	GimbalDeviceAttitudeStatus_FLAGS_YAW_LOCK uint16 = 16// Lock yaw angle to absolute angle relative to North (not relative to drone). If this flag is set, the quaternion is in the Earth frame with the x-axis pointing North (yaw absolute). If this flag is not set, the quaternion frame is in the Earth frame rotated so that the x-axis is pointing forward (yaw relative to vehicle).
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_AT_ROLL_LIMIT uint32 = 1// Gimbal device is limited by hardware roll limit.. GIMBAL_DEVICE_ERROR_FLAGS
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_AT_PITCH_LIMIT uint32 = 2// Gimbal device is limited by hardware pitch limit.
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_AT_YAW_LIMIT uint32 = 4// Gimbal device is limited by hardware yaw limit.
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_ENCODER_ERROR uint32 = 8// There is an error with the gimbal encoders.
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_POWER_ERROR uint32 = 16// There is an error with the gimbal power source.
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_MOTOR_ERROR uint32 = 32// There is an error with the gimbal motor's.
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_SOFTWARE_ERROR uint32 = 64// There is an error with the gimbal's software.
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_COMMS_ERROR uint32 = 128// There is an error with the gimbal's communication.
	GimbalDeviceAttitudeStatus_ERROR_FLAGS_CALIBRATION_RUNNING uint32 = 256// Gimbal is currently calibrating.
)

type GimbalDeviceAttitudeStatus struct {
	Header std_msgs_msg.Header `yaml:"header"`
	TargetSystem uint8 `yaml:"target_system"`// System ID
	TargetComponent uint8 `yaml:"target_component"`// Component ID
	Flags uint16 `yaml:"flags"`// Current gimbal flags set (bitwise) - See GIMBAL_DEVICE_FLAGS
	Q geometry_msgs_msg.Quaternion `yaml:"q"`// Quaternion, x, y, z, w (0 0 0 1 is the null-rotation, the frame is depends on whether the flag GIMBAL_DEVICE_FLAGS_YAW_LOCK is set)
	AngularVelocityX float32 `yaml:"angular_velocity_x"`// X component of angular velocity (NaN if unknown)
	AngularVelocityY float32 `yaml:"angular_velocity_y"`// Y component of angular velocity (NaN if unknown)
	AngularVelocityZ float32 `yaml:"angular_velocity_z"`// Z component of angular velocity (NaN if unknown)
	FailureFlags uint32 `yaml:"failure_flags"`// Failure flags (0 for no failure) (bitwise) - See GIMBAL_DEVICE_ERROR_FLAGS
}

// NewGimbalDeviceAttitudeStatus creates a new GimbalDeviceAttitudeStatus with default values.
func NewGimbalDeviceAttitudeStatus() *GimbalDeviceAttitudeStatus {
	self := GimbalDeviceAttitudeStatus{}
	self.SetDefaults()
	return &self
}

func (t *GimbalDeviceAttitudeStatus) Clone() *GimbalDeviceAttitudeStatus {
	c := &GimbalDeviceAttitudeStatus{}
	c.Header = *t.Header.Clone()
	c.TargetSystem = t.TargetSystem
	c.TargetComponent = t.TargetComponent
	c.Flags = t.Flags
	c.Q = *t.Q.Clone()
	c.AngularVelocityX = t.AngularVelocityX
	c.AngularVelocityY = t.AngularVelocityY
	c.AngularVelocityZ = t.AngularVelocityZ
	c.FailureFlags = t.FailureFlags
	return c
}

func (t *GimbalDeviceAttitudeStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GimbalDeviceAttitudeStatus) SetDefaults() {
	t.Header.SetDefaults()
	t.TargetSystem = 0
	t.TargetComponent = 0
	t.Flags = 0
	t.Q.SetDefaults()
	t.AngularVelocityX = 0
	t.AngularVelocityY = 0
	t.AngularVelocityZ = 0
	t.FailureFlags = 0
}

func (t *GimbalDeviceAttitudeStatus) GetTypeSupport() types.MessageTypeSupport {
	return GimbalDeviceAttitudeStatusTypeSupport
}

// GimbalDeviceAttitudeStatusPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GimbalDeviceAttitudeStatusPublisher struct {
	*rclgo.Publisher
}

// NewGimbalDeviceAttitudeStatusPublisher creates and returns a new publisher for the
// GimbalDeviceAttitudeStatus
func NewGimbalDeviceAttitudeStatusPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GimbalDeviceAttitudeStatusPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GimbalDeviceAttitudeStatusTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GimbalDeviceAttitudeStatusPublisher{pub}, nil
}

func (p *GimbalDeviceAttitudeStatusPublisher) Publish(msg *GimbalDeviceAttitudeStatus) error {
	return p.Publisher.Publish(msg)
}

// GimbalDeviceAttitudeStatusSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GimbalDeviceAttitudeStatusSubscription struct {
	*rclgo.Subscription
}

// GimbalDeviceAttitudeStatusSubscriptionCallback type is used to provide a subscription
// handler function for a GimbalDeviceAttitudeStatusSubscription.
type GimbalDeviceAttitudeStatusSubscriptionCallback func(msg *GimbalDeviceAttitudeStatus, info *rclgo.MessageInfo, err error)

// NewGimbalDeviceAttitudeStatusSubscription creates and returns a new subscription for the
// GimbalDeviceAttitudeStatus
func NewGimbalDeviceAttitudeStatusSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GimbalDeviceAttitudeStatusSubscriptionCallback) (*GimbalDeviceAttitudeStatusSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GimbalDeviceAttitudeStatus
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GimbalDeviceAttitudeStatusTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GimbalDeviceAttitudeStatusSubscription{sub}, nil
}

func (s *GimbalDeviceAttitudeStatusSubscription) TakeMessage(out *GimbalDeviceAttitudeStatus) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGimbalDeviceAttitudeStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGimbalDeviceAttitudeStatusSlice(dst, src []GimbalDeviceAttitudeStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GimbalDeviceAttitudeStatusTypeSupport types.MessageTypeSupport = _GimbalDeviceAttitudeStatusTypeSupport{}

type _GimbalDeviceAttitudeStatusTypeSupport struct{}

func (t _GimbalDeviceAttitudeStatusTypeSupport) New() types.Message {
	return NewGimbalDeviceAttitudeStatus()
}

func (t _GimbalDeviceAttitudeStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__GimbalDeviceAttitudeStatus
	return (unsafe.Pointer)(C.mavros_msgs__msg__GimbalDeviceAttitudeStatus__create())
}

func (t _GimbalDeviceAttitudeStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__GimbalDeviceAttitudeStatus__destroy((*C.mavros_msgs__msg__GimbalDeviceAttitudeStatus)(pointer_to_free))
}

func (t _GimbalDeviceAttitudeStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GimbalDeviceAttitudeStatus)
	mem := (*C.mavros_msgs__msg__GimbalDeviceAttitudeStatus)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.target_system = C.uint8_t(m.TargetSystem)
	mem.target_component = C.uint8_t(m.TargetComponent)
	mem.flags = C.uint16_t(m.Flags)
	geometry_msgs_msg.QuaternionTypeSupport.AsCStruct(unsafe.Pointer(&mem.q), &m.Q)
	mem.angular_velocity_x = C.float(m.AngularVelocityX)
	mem.angular_velocity_y = C.float(m.AngularVelocityY)
	mem.angular_velocity_z = C.float(m.AngularVelocityZ)
	mem.failure_flags = C.uint32_t(m.FailureFlags)
}

func (t _GimbalDeviceAttitudeStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GimbalDeviceAttitudeStatus)
	mem := (*C.mavros_msgs__msg__GimbalDeviceAttitudeStatus)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.TargetSystem = uint8(mem.target_system)
	m.TargetComponent = uint8(mem.target_component)
	m.Flags = uint16(mem.flags)
	geometry_msgs_msg.QuaternionTypeSupport.AsGoStruct(&m.Q, unsafe.Pointer(&mem.q))
	m.AngularVelocityX = float32(mem.angular_velocity_x)
	m.AngularVelocityY = float32(mem.angular_velocity_y)
	m.AngularVelocityZ = float32(mem.angular_velocity_z)
	m.FailureFlags = uint32(mem.failure_flags)
}

func (t _GimbalDeviceAttitudeStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__GimbalDeviceAttitudeStatus())
}

type CGimbalDeviceAttitudeStatus = C.mavros_msgs__msg__GimbalDeviceAttitudeStatus
type CGimbalDeviceAttitudeStatus__Sequence = C.mavros_msgs__msg__GimbalDeviceAttitudeStatus__Sequence

func GimbalDeviceAttitudeStatus__Sequence_to_Go(goSlice *[]GimbalDeviceAttitudeStatus, cSlice CGimbalDeviceAttitudeStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GimbalDeviceAttitudeStatus, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GimbalDeviceAttitudeStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GimbalDeviceAttitudeStatus__Sequence_to_C(cSlice *CGimbalDeviceAttitudeStatus__Sequence, goSlice []GimbalDeviceAttitudeStatus) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__GimbalDeviceAttitudeStatus)(C.malloc(C.sizeof_struct_mavros_msgs__msg__GimbalDeviceAttitudeStatus * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GimbalDeviceAttitudeStatusTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GimbalDeviceAttitudeStatus__Array_to_Go(goSlice []GimbalDeviceAttitudeStatus, cSlice []CGimbalDeviceAttitudeStatus) {
	for i := 0; i < len(cSlice); i++ {
		GimbalDeviceAttitudeStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GimbalDeviceAttitudeStatus__Array_to_C(cSlice []CGimbalDeviceAttitudeStatus, goSlice []GimbalDeviceAttitudeStatus) {
	for i := 0; i < len(goSlice); i++ {
		GimbalDeviceAttitudeStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
