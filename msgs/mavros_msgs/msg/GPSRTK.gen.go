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

#include <mavros_msgs/msg/gpsrtk.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/GPSRTK", GPSRTKTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/GPSRTK", GPSRTKTypeSupport)
}

type GPSRTK struct {
	Header std_msgs_msg.Header `yaml:"header"`
	RtkReceiverId uint8 `yaml:"rtk_receiver_id"`// Identification of connected RTK receiver.
	Wn int16 `yaml:"wn"`// GPS Week Number of last baseline.
	Tow uint32 `yaml:"tow"`// [ms] GPS Time of Week of last baseline.
	RtkHealth uint8 `yaml:"rtk_health"`// GPS-specific health report for RTK data.
	RtkRate uint8 `yaml:"rtk_rate"`// [Hz] Rate of baseline messages being received by GPS.
	Nsats uint8 `yaml:"nsats"`// Current number of sats used for RTK calculation.
	BaselineA int32 `yaml:"baseline_a"`// [mm] Current baseline in ECEF x or NED north component, depends on header.frame_id.
	BaselineB int32 `yaml:"baseline_b"`// [mm] Current baseline in ECEF y or NED east component, depends on header.frame_id.
	BaselineC int32 `yaml:"baseline_c"`// [mm] Current baseline in ECEF z or NED down component, depends on header.frame_id.
	Accuracy uint32 `yaml:"accuracy"`// Current estimate of baseline accuracy.
	IarNumHypotheses int32 `yaml:"iar_num_hypotheses"`// Current number of integer ambiguity hypotheses.
}

// NewGPSRTK creates a new GPSRTK with default values.
func NewGPSRTK() *GPSRTK {
	self := GPSRTK{}
	self.SetDefaults()
	return &self
}

func (t *GPSRTK) Clone() *GPSRTK {
	c := &GPSRTK{}
	c.Header = *t.Header.Clone()
	c.RtkReceiverId = t.RtkReceiverId
	c.Wn = t.Wn
	c.Tow = t.Tow
	c.RtkHealth = t.RtkHealth
	c.RtkRate = t.RtkRate
	c.Nsats = t.Nsats
	c.BaselineA = t.BaselineA
	c.BaselineB = t.BaselineB
	c.BaselineC = t.BaselineC
	c.Accuracy = t.Accuracy
	c.IarNumHypotheses = t.IarNumHypotheses
	return c
}

func (t *GPSRTK) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GPSRTK) SetDefaults() {
	t.Header.SetDefaults()
	t.RtkReceiverId = 0
	t.Wn = 0
	t.Tow = 0
	t.RtkHealth = 0
	t.RtkRate = 0
	t.Nsats = 0
	t.BaselineA = 0
	t.BaselineB = 0
	t.BaselineC = 0
	t.Accuracy = 0
	t.IarNumHypotheses = 0
}

func (t *GPSRTK) GetTypeSupport() types.MessageTypeSupport {
	return GPSRTKTypeSupport
}

// GPSRTKPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GPSRTKPublisher struct {
	*rclgo.Publisher
}

// NewGPSRTKPublisher creates and returns a new publisher for the
// GPSRTK
func NewGPSRTKPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GPSRTKPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GPSRTKTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GPSRTKPublisher{pub}, nil
}

func (p *GPSRTKPublisher) Publish(msg *GPSRTK) error {
	return p.Publisher.Publish(msg)
}

// GPSRTKSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GPSRTKSubscription struct {
	*rclgo.Subscription
}

// GPSRTKSubscriptionCallback type is used to provide a subscription
// handler function for a GPSRTKSubscription.
type GPSRTKSubscriptionCallback func(msg *GPSRTK, info *rclgo.MessageInfo, err error)

// NewGPSRTKSubscription creates and returns a new subscription for the
// GPSRTK
func NewGPSRTKSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GPSRTKSubscriptionCallback) (*GPSRTKSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GPSRTK
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GPSRTKTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GPSRTKSubscription{sub}, nil
}

func (s *GPSRTKSubscription) TakeMessage(out *GPSRTK) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGPSRTKSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGPSRTKSlice(dst, src []GPSRTK) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GPSRTKTypeSupport types.MessageTypeSupport = _GPSRTKTypeSupport{}

type _GPSRTKTypeSupport struct{}

func (t _GPSRTKTypeSupport) New() types.Message {
	return NewGPSRTK()
}

func (t _GPSRTKTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__GPSRTK
	return (unsafe.Pointer)(C.mavros_msgs__msg__GPSRTK__create())
}

func (t _GPSRTKTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__GPSRTK__destroy((*C.mavros_msgs__msg__GPSRTK)(pointer_to_free))
}

func (t _GPSRTKTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GPSRTK)
	mem := (*C.mavros_msgs__msg__GPSRTK)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.rtk_receiver_id = C.uint8_t(m.RtkReceiverId)
	mem.wn = C.int16_t(m.Wn)
	mem.tow = C.uint32_t(m.Tow)
	mem.rtk_health = C.uint8_t(m.RtkHealth)
	mem.rtk_rate = C.uint8_t(m.RtkRate)
	mem.nsats = C.uint8_t(m.Nsats)
	mem.baseline_a = C.int32_t(m.BaselineA)
	mem.baseline_b = C.int32_t(m.BaselineB)
	mem.baseline_c = C.int32_t(m.BaselineC)
	mem.accuracy = C.uint32_t(m.Accuracy)
	mem.iar_num_hypotheses = C.int32_t(m.IarNumHypotheses)
}

func (t _GPSRTKTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GPSRTK)
	mem := (*C.mavros_msgs__msg__GPSRTK)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.RtkReceiverId = uint8(mem.rtk_receiver_id)
	m.Wn = int16(mem.wn)
	m.Tow = uint32(mem.tow)
	m.RtkHealth = uint8(mem.rtk_health)
	m.RtkRate = uint8(mem.rtk_rate)
	m.Nsats = uint8(mem.nsats)
	m.BaselineA = int32(mem.baseline_a)
	m.BaselineB = int32(mem.baseline_b)
	m.BaselineC = int32(mem.baseline_c)
	m.Accuracy = uint32(mem.accuracy)
	m.IarNumHypotheses = int32(mem.iar_num_hypotheses)
}

func (t _GPSRTKTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__GPSRTK())
}

type CGPSRTK = C.mavros_msgs__msg__GPSRTK
type CGPSRTK__Sequence = C.mavros_msgs__msg__GPSRTK__Sequence

func GPSRTK__Sequence_to_Go(goSlice *[]GPSRTK, cSlice CGPSRTK__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GPSRTK, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GPSRTKTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GPSRTK__Sequence_to_C(cSlice *CGPSRTK__Sequence, goSlice []GPSRTK) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__GPSRTK)(C.malloc(C.sizeof_struct_mavros_msgs__msg__GPSRTK * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GPSRTKTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GPSRTK__Array_to_Go(goSlice []GPSRTK, cSlice []CGPSRTK) {
	for i := 0; i < len(cSlice); i++ {
		GPSRTKTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GPSRTK__Array_to_C(cSlice []CGPSRTK, goSlice []GPSRTK) {
	for i := 0; i < len(goSlice); i++ {
		GPSRTKTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
