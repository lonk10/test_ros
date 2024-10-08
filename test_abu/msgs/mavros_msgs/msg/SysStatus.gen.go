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

#include <mavros_msgs/msg/sys_status.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/SysStatus", SysStatusTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/SysStatus", SysStatusTypeSupport)
}

type SysStatus struct {
	Header std_msgs_msg.Header `yaml:"header"`
	SensorsPresent uint32 `yaml:"sensors_present"`
	SensorsEnabled uint32 `yaml:"sensors_enabled"`
	SensorsHealth uint32 `yaml:"sensors_health"`
	Load uint16 `yaml:"load"`
	VoltageBattery uint16 `yaml:"voltage_battery"`
	CurrentBattery int16 `yaml:"current_battery"`
	BatteryRemaining int8 `yaml:"battery_remaining"`
	DropRateComm uint16 `yaml:"drop_rate_comm"`
	ErrorsComm uint16 `yaml:"errors_comm"`
	ErrorsCount1 uint16 `yaml:"errors_count1"`
	ErrorsCount2 uint16 `yaml:"errors_count2"`
	ErrorsCount3 uint16 `yaml:"errors_count3"`
	ErrorsCount4 uint16 `yaml:"errors_count4"`
}

// NewSysStatus creates a new SysStatus with default values.
func NewSysStatus() *SysStatus {
	self := SysStatus{}
	self.SetDefaults()
	return &self
}

func (t *SysStatus) Clone() *SysStatus {
	c := &SysStatus{}
	c.Header = *t.Header.Clone()
	c.SensorsPresent = t.SensorsPresent
	c.SensorsEnabled = t.SensorsEnabled
	c.SensorsHealth = t.SensorsHealth
	c.Load = t.Load
	c.VoltageBattery = t.VoltageBattery
	c.CurrentBattery = t.CurrentBattery
	c.BatteryRemaining = t.BatteryRemaining
	c.DropRateComm = t.DropRateComm
	c.ErrorsComm = t.ErrorsComm
	c.ErrorsCount1 = t.ErrorsCount1
	c.ErrorsCount2 = t.ErrorsCount2
	c.ErrorsCount3 = t.ErrorsCount3
	c.ErrorsCount4 = t.ErrorsCount4
	return c
}

func (t *SysStatus) CloneMsg() types.Message {
	return t.Clone()
}

func (t *SysStatus) SetDefaults() {
	t.Header.SetDefaults()
	t.SensorsPresent = 0
	t.SensorsEnabled = 0
	t.SensorsHealth = 0
	t.Load = 0
	t.VoltageBattery = 0
	t.CurrentBattery = 0
	t.BatteryRemaining = 0
	t.DropRateComm = 0
	t.ErrorsComm = 0
	t.ErrorsCount1 = 0
	t.ErrorsCount2 = 0
	t.ErrorsCount3 = 0
	t.ErrorsCount4 = 0
}

func (t *SysStatus) GetTypeSupport() types.MessageTypeSupport {
	return SysStatusTypeSupport
}

// SysStatusPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type SysStatusPublisher struct {
	*rclgo.Publisher
}

// NewSysStatusPublisher creates and returns a new publisher for the
// SysStatus
func NewSysStatusPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*SysStatusPublisher, error) {
	pub, err := node.NewPublisher(topic_name, SysStatusTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SysStatusPublisher{pub}, nil
}

func (p *SysStatusPublisher) Publish(msg *SysStatus) error {
	return p.Publisher.Publish(msg)
}

// SysStatusSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type SysStatusSubscription struct {
	*rclgo.Subscription
}

// SysStatusSubscriptionCallback type is used to provide a subscription
// handler function for a SysStatusSubscription.
type SysStatusSubscriptionCallback func(msg *SysStatus, info *rclgo.MessageInfo, err error)

// NewSysStatusSubscription creates and returns a new subscription for the
// SysStatus
func NewSysStatusSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback SysStatusSubscriptionCallback) (*SysStatusSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg SysStatus
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, SysStatusTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &SysStatusSubscription{sub}, nil
}

func (s *SysStatusSubscription) TakeMessage(out *SysStatus) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneSysStatusSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneSysStatusSlice(dst, src []SysStatus) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var SysStatusTypeSupport types.MessageTypeSupport = _SysStatusTypeSupport{}

type _SysStatusTypeSupport struct{}

func (t _SysStatusTypeSupport) New() types.Message {
	return NewSysStatus()
}

func (t _SysStatusTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__SysStatus
	return (unsafe.Pointer)(C.mavros_msgs__msg__SysStatus__create())
}

func (t _SysStatusTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__SysStatus__destroy((*C.mavros_msgs__msg__SysStatus)(pointer_to_free))
}

func (t _SysStatusTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*SysStatus)
	mem := (*C.mavros_msgs__msg__SysStatus)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.sensors_present = C.uint32_t(m.SensorsPresent)
	mem.sensors_enabled = C.uint32_t(m.SensorsEnabled)
	mem.sensors_health = C.uint32_t(m.SensorsHealth)
	mem.load = C.uint16_t(m.Load)
	mem.voltage_battery = C.uint16_t(m.VoltageBattery)
	mem.current_battery = C.int16_t(m.CurrentBattery)
	mem.battery_remaining = C.int8_t(m.BatteryRemaining)
	mem.drop_rate_comm = C.uint16_t(m.DropRateComm)
	mem.errors_comm = C.uint16_t(m.ErrorsComm)
	mem.errors_count1 = C.uint16_t(m.ErrorsCount1)
	mem.errors_count2 = C.uint16_t(m.ErrorsCount2)
	mem.errors_count3 = C.uint16_t(m.ErrorsCount3)
	mem.errors_count4 = C.uint16_t(m.ErrorsCount4)
}

func (t _SysStatusTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*SysStatus)
	mem := (*C.mavros_msgs__msg__SysStatus)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.SensorsPresent = uint32(mem.sensors_present)
	m.SensorsEnabled = uint32(mem.sensors_enabled)
	m.SensorsHealth = uint32(mem.sensors_health)
	m.Load = uint16(mem.load)
	m.VoltageBattery = uint16(mem.voltage_battery)
	m.CurrentBattery = int16(mem.current_battery)
	m.BatteryRemaining = int8(mem.battery_remaining)
	m.DropRateComm = uint16(mem.drop_rate_comm)
	m.ErrorsComm = uint16(mem.errors_comm)
	m.ErrorsCount1 = uint16(mem.errors_count1)
	m.ErrorsCount2 = uint16(mem.errors_count2)
	m.ErrorsCount3 = uint16(mem.errors_count3)
	m.ErrorsCount4 = uint16(mem.errors_count4)
}

func (t _SysStatusTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__SysStatus())
}

type CSysStatus = C.mavros_msgs__msg__SysStatus
type CSysStatus__Sequence = C.mavros_msgs__msg__SysStatus__Sequence

func SysStatus__Sequence_to_Go(goSlice *[]SysStatus, cSlice CSysStatus__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]SysStatus, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		SysStatusTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func SysStatus__Sequence_to_C(cSlice *CSysStatus__Sequence, goSlice []SysStatus) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__SysStatus)(C.malloc(C.sizeof_struct_mavros_msgs__msg__SysStatus * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		SysStatusTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func SysStatus__Array_to_Go(goSlice []SysStatus, cSlice []CSysStatus) {
	for i := 0; i < len(cSlice); i++ {
		SysStatusTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func SysStatus__Array_to_C(cSlice []CSysStatus, goSlice []SysStatus) {
	for i := 0; i < len(goSlice); i++ {
		SysStatusTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
