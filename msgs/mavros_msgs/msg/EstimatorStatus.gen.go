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

#include <mavros_msgs/msg/estimator_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/EstimatorStatus", EstimatorStatusTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/EstimatorStatus", EstimatorStatusTypeSupport)
}

type EstimatorStatus struct {
	Header std_msgs_msg.Header `yaml:"header"`
	AttitudeStatusFlag bool `yaml:"attitude_status_flag"`
	VelocityHorizStatusFlag bool `yaml:"velocity_horiz_status_flag"`
	VelocityVertStatusFlag bool `yaml:"velocity_vert_status_flag"`
	PosHorizRelStatusFlag bool `yaml:"pos_horiz_rel_status_flag"`
	PosHorizAbsStatusFlag bool `yaml:"pos_horiz_abs_status_flag"`
	PosVertAbsStatusFlag bool `yaml:"pos_vert_abs_status_flag"`
	PosVertAglStatusFlag bool `yaml:"pos_vert_agl_status_flag"`
	ConstPosModeStatusFlag bool `yaml:"const_pos_mode_status_flag"`
	PredPosHorizRelStatusFlag bool `yaml:"pred_pos_horiz_rel_status_flag"`
	PredPosHorizAbsStatusFlag bool `yaml:"pred_pos_horiz_abs_status_flag"`
	GpsGlitchStatusFlag bool `yaml:"gps_glitch_status_flag"`
	AccelErrorStatusFlag bool `yaml:"accel_error_status_flag"`
}

// NewEstimatorStatus creates a new EstimatorStatus with default values.
func NewEstimatorStatus() *EstimatorStatus {
	self := EstimatorStatus{}
	self.SetDefaults()
	return &self
}

func (t *EstimatorStatus) Clone() *EstimatorStatus {
	c := &EstimatorStatus{}
	c.Header = *t.Header.Clone()
	c.AttitudeStatusFlag = t.AttitudeStatusFlag
	c.VelocityHorizStatusFlag = t.VelocityHorizStatusFlag
	c.VelocityVertStatusFlag = t.VelocityVertStatusFlag
	c.PosHorizRelStatusFlag = t.PosHorizRelStatusFlag
	c.PosHorizAbsStatusFlag = t.PosHorizAbsStatusFlag
	c.PosVertAbsStatusFlag = t.PosVertAbsStatusFlag
	c.PosVertAglStatusFlag = t.PosVertAglStatusFlag
	c.ConstPosModeStatusFlag = t.ConstPosModeStatusFlag
	c.PredPosHorizRelStatusFlag = t.PredPosHorizRelStatusFlag
	c.PredPosHorizAbsStatusFlag = t.PredPosHorizAbsStatusFlag
	c.GpsGlitchStatusFlag = t.GpsGlitchStatusFlag
	c.AccelErrorStatusFlag = t.AccelErrorStatusFlag
	return c
}

func (t *EstimatorStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *EstimatorStatus) SetDefaults() {
	t.Header.SetDefaults()
	t.AttitudeStatusFlag = false
	t.VelocityHorizStatusFlag = false
	t.VelocityVertStatusFlag = false
	t.PosHorizRelStatusFlag = false
	t.PosHorizAbsStatusFlag = false
	t.PosVertAbsStatusFlag = false
	t.PosVertAglStatusFlag = false
	t.ConstPosModeStatusFlag = false
	t.PredPosHorizRelStatusFlag = false
	t.PredPosHorizAbsStatusFlag = false
	t.GpsGlitchStatusFlag = false
	t.AccelErrorStatusFlag = false
}

func (t *EstimatorStatus) GetTypeSupport() types.MessageTypeSupport {
	return EstimatorStatusTypeSupport
}

// EstimatorStatusPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type EstimatorStatusPublisher struct {
	*rclgo.Publisher
}

// NewEstimatorStatusPublisher creates and returns a new publisher for the
// EstimatorStatus
func NewEstimatorStatusPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*EstimatorStatusPublisher, error) {
	pub, err := node.NewPublisher(topic_name, EstimatorStatusTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &EstimatorStatusPublisher{pub}, nil
}

func (p *EstimatorStatusPublisher) Publish(msg *EstimatorStatus) error {
	return p.Publisher.Publish(msg)
}

// EstimatorStatusSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type EstimatorStatusSubscription struct {
	*rclgo.Subscription
}

// EstimatorStatusSubscriptionCallback type is used to provide a subscription
// handler function for a EstimatorStatusSubscription.
type EstimatorStatusSubscriptionCallback func(msg *EstimatorStatus, info *rclgo.MessageInfo, err error)

// NewEstimatorStatusSubscription creates and returns a new subscription for the
// EstimatorStatus
func NewEstimatorStatusSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback EstimatorStatusSubscriptionCallback) (*EstimatorStatusSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg EstimatorStatus
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, EstimatorStatusTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &EstimatorStatusSubscription{sub}, nil
}

func (s *EstimatorStatusSubscription) TakeMessage(out *EstimatorStatus) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneEstimatorStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneEstimatorStatusSlice(dst, src []EstimatorStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var EstimatorStatusTypeSupport types.MessageTypeSupport = _EstimatorStatusTypeSupport{}

type _EstimatorStatusTypeSupport struct{}

func (t _EstimatorStatusTypeSupport) New() types.Message {
	return NewEstimatorStatus()
}

func (t _EstimatorStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__EstimatorStatus
	return (unsafe.Pointer)(C.mavros_msgs__msg__EstimatorStatus__create())
}

func (t _EstimatorStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__EstimatorStatus__destroy((*C.mavros_msgs__msg__EstimatorStatus)(pointer_to_free))
}

func (t _EstimatorStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*EstimatorStatus)
	mem := (*C.mavros_msgs__msg__EstimatorStatus)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.attitude_status_flag = C.bool(m.AttitudeStatusFlag)
	mem.velocity_horiz_status_flag = C.bool(m.VelocityHorizStatusFlag)
	mem.velocity_vert_status_flag = C.bool(m.VelocityVertStatusFlag)
	mem.pos_horiz_rel_status_flag = C.bool(m.PosHorizRelStatusFlag)
	mem.pos_horiz_abs_status_flag = C.bool(m.PosHorizAbsStatusFlag)
	mem.pos_vert_abs_status_flag = C.bool(m.PosVertAbsStatusFlag)
	mem.pos_vert_agl_status_flag = C.bool(m.PosVertAglStatusFlag)
	mem.const_pos_mode_status_flag = C.bool(m.ConstPosModeStatusFlag)
	mem.pred_pos_horiz_rel_status_flag = C.bool(m.PredPosHorizRelStatusFlag)
	mem.pred_pos_horiz_abs_status_flag = C.bool(m.PredPosHorizAbsStatusFlag)
	mem.gps_glitch_status_flag = C.bool(m.GpsGlitchStatusFlag)
	mem.accel_error_status_flag = C.bool(m.AccelErrorStatusFlag)
}

func (t _EstimatorStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*EstimatorStatus)
	mem := (*C.mavros_msgs__msg__EstimatorStatus)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.AttitudeStatusFlag = bool(mem.attitude_status_flag)
	m.VelocityHorizStatusFlag = bool(mem.velocity_horiz_status_flag)
	m.VelocityVertStatusFlag = bool(mem.velocity_vert_status_flag)
	m.PosHorizRelStatusFlag = bool(mem.pos_horiz_rel_status_flag)
	m.PosHorizAbsStatusFlag = bool(mem.pos_horiz_abs_status_flag)
	m.PosVertAbsStatusFlag = bool(mem.pos_vert_abs_status_flag)
	m.PosVertAglStatusFlag = bool(mem.pos_vert_agl_status_flag)
	m.ConstPosModeStatusFlag = bool(mem.const_pos_mode_status_flag)
	m.PredPosHorizRelStatusFlag = bool(mem.pred_pos_horiz_rel_status_flag)
	m.PredPosHorizAbsStatusFlag = bool(mem.pred_pos_horiz_abs_status_flag)
	m.GpsGlitchStatusFlag = bool(mem.gps_glitch_status_flag)
	m.AccelErrorStatusFlag = bool(mem.accel_error_status_flag)
}

func (t _EstimatorStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__EstimatorStatus())
}

type CEstimatorStatus = C.mavros_msgs__msg__EstimatorStatus
type CEstimatorStatus__Sequence = C.mavros_msgs__msg__EstimatorStatus__Sequence

func EstimatorStatus__Sequence_to_Go(goSlice *[]EstimatorStatus, cSlice CEstimatorStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]EstimatorStatus, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		EstimatorStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func EstimatorStatus__Sequence_to_C(cSlice *CEstimatorStatus__Sequence, goSlice []EstimatorStatus) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__EstimatorStatus)(C.malloc(C.sizeof_struct_mavros_msgs__msg__EstimatorStatus * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		EstimatorStatusTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func EstimatorStatus__Array_to_Go(goSlice []EstimatorStatus, cSlice []CEstimatorStatus) {
	for i := 0; i < len(cSlice); i++ {
		EstimatorStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func EstimatorStatus__Array_to_C(cSlice []CEstimatorStatus, goSlice []EstimatorStatus) {
	for i := 0; i < len(goSlice); i++ {
		EstimatorStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
