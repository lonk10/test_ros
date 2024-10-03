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

#include <mavros_msgs/msg/optical_flow_rad.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/OpticalFlowRad", OpticalFlowRadTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/OpticalFlowRad", OpticalFlowRadTypeSupport)
}

type OpticalFlowRad struct {
	Header std_msgs_msg.Header `yaml:"header"`
	IntegrationTimeUs uint32 `yaml:"integration_time_us"`
	IntegratedX float32 `yaml:"integrated_x"`
	IntegratedY float32 `yaml:"integrated_y"`
	IntegratedXgyro float32 `yaml:"integrated_xgyro"`
	IntegratedYgyro float32 `yaml:"integrated_ygyro"`
	IntegratedZgyro float32 `yaml:"integrated_zgyro"`
	Temperature int16 `yaml:"temperature"`
	Quality uint8 `yaml:"quality"`
	TimeDeltaDistanceUs uint32 `yaml:"time_delta_distance_us"`
	Distance float32 `yaml:"distance"`
}

// NewOpticalFlowRad creates a new OpticalFlowRad with default values.
func NewOpticalFlowRad() *OpticalFlowRad {
	self := OpticalFlowRad{}
	self.SetDefaults()
	return &self
}

func (t *OpticalFlowRad) Clone() *OpticalFlowRad {
	c := &OpticalFlowRad{}
	c.Header = *t.Header.Clone()
	c.IntegrationTimeUs = t.IntegrationTimeUs
	c.IntegratedX = t.IntegratedX
	c.IntegratedY = t.IntegratedY
	c.IntegratedXgyro = t.IntegratedXgyro
	c.IntegratedYgyro = t.IntegratedYgyro
	c.IntegratedZgyro = t.IntegratedZgyro
	c.Temperature = t.Temperature
	c.Quality = t.Quality
	c.TimeDeltaDistanceUs = t.TimeDeltaDistanceUs
	c.Distance = t.Distance
	return c
}

func (t *OpticalFlowRad) CloneMsg() types.Message {
	return t.Clone()
}

func (t *OpticalFlowRad) SetDefaults() {
	t.Header.SetDefaults()
	t.IntegrationTimeUs = 0
	t.IntegratedX = 0
	t.IntegratedY = 0
	t.IntegratedXgyro = 0
	t.IntegratedYgyro = 0
	t.IntegratedZgyro = 0
	t.Temperature = 0
	t.Quality = 0
	t.TimeDeltaDistanceUs = 0
	t.Distance = 0
}

func (t *OpticalFlowRad) GetTypeSupport() types.MessageTypeSupport {
	return OpticalFlowRadTypeSupport
}

// OpticalFlowRadPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type OpticalFlowRadPublisher struct {
	*rclgo.Publisher
}

// NewOpticalFlowRadPublisher creates and returns a new publisher for the
// OpticalFlowRad
func NewOpticalFlowRadPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*OpticalFlowRadPublisher, error) {
	pub, err := node.NewPublisher(topic_name, OpticalFlowRadTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &OpticalFlowRadPublisher{pub}, nil
}

func (p *OpticalFlowRadPublisher) Publish(msg *OpticalFlowRad) error {
	return p.Publisher.Publish(msg)
}

// OpticalFlowRadSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type OpticalFlowRadSubscription struct {
	*rclgo.Subscription
}

// OpticalFlowRadSubscriptionCallback type is used to provide a subscription
// handler function for a OpticalFlowRadSubscription.
type OpticalFlowRadSubscriptionCallback func(msg *OpticalFlowRad, info *rclgo.MessageInfo, err error)

// NewOpticalFlowRadSubscription creates and returns a new subscription for the
// OpticalFlowRad
func NewOpticalFlowRadSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback OpticalFlowRadSubscriptionCallback) (*OpticalFlowRadSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg OpticalFlowRad
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, OpticalFlowRadTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &OpticalFlowRadSubscription{sub}, nil
}

func (s *OpticalFlowRadSubscription) TakeMessage(out *OpticalFlowRad) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneOpticalFlowRadSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneOpticalFlowRadSlice(dst, src []OpticalFlowRad) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var OpticalFlowRadTypeSupport types.MessageTypeSupport = _OpticalFlowRadTypeSupport{}

type _OpticalFlowRadTypeSupport struct{}

func (t _OpticalFlowRadTypeSupport) New() types.Message {
	return NewOpticalFlowRad()
}

func (t _OpticalFlowRadTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__OpticalFlowRad
	return (unsafe.Pointer)(C.mavros_msgs__msg__OpticalFlowRad__create())
}

func (t _OpticalFlowRadTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__OpticalFlowRad__destroy((*C.mavros_msgs__msg__OpticalFlowRad)(pointer_to_free))
}

func (t _OpticalFlowRadTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*OpticalFlowRad)
	mem := (*C.mavros_msgs__msg__OpticalFlowRad)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.integration_time_us = C.uint32_t(m.IntegrationTimeUs)
	mem.integrated_x = C.float(m.IntegratedX)
	mem.integrated_y = C.float(m.IntegratedY)
	mem.integrated_xgyro = C.float(m.IntegratedXgyro)
	mem.integrated_ygyro = C.float(m.IntegratedYgyro)
	mem.integrated_zgyro = C.float(m.IntegratedZgyro)
	mem.temperature = C.int16_t(m.Temperature)
	mem.quality = C.uint8_t(m.Quality)
	mem.time_delta_distance_us = C.uint32_t(m.TimeDeltaDistanceUs)
	mem.distance = C.float(m.Distance)
}

func (t _OpticalFlowRadTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*OpticalFlowRad)
	mem := (*C.mavros_msgs__msg__OpticalFlowRad)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.IntegrationTimeUs = uint32(mem.integration_time_us)
	m.IntegratedX = float32(mem.integrated_x)
	m.IntegratedY = float32(mem.integrated_y)
	m.IntegratedXgyro = float32(mem.integrated_xgyro)
	m.IntegratedYgyro = float32(mem.integrated_ygyro)
	m.IntegratedZgyro = float32(mem.integrated_zgyro)
	m.Temperature = int16(mem.temperature)
	m.Quality = uint8(mem.quality)
	m.TimeDeltaDistanceUs = uint32(mem.time_delta_distance_us)
	m.Distance = float32(mem.distance)
}

func (t _OpticalFlowRadTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__OpticalFlowRad())
}

type COpticalFlowRad = C.mavros_msgs__msg__OpticalFlowRad
type COpticalFlowRad__Sequence = C.mavros_msgs__msg__OpticalFlowRad__Sequence

func OpticalFlowRad__Sequence_to_Go(goSlice *[]OpticalFlowRad, cSlice COpticalFlowRad__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]OpticalFlowRad, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		OpticalFlowRadTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func OpticalFlowRad__Sequence_to_C(cSlice *COpticalFlowRad__Sequence, goSlice []OpticalFlowRad) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__OpticalFlowRad)(C.malloc(C.sizeof_struct_mavros_msgs__msg__OpticalFlowRad * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		OpticalFlowRadTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func OpticalFlowRad__Array_to_Go(goSlice []OpticalFlowRad, cSlice []COpticalFlowRad) {
	for i := 0; i < len(cSlice); i++ {
		OpticalFlowRadTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func OpticalFlowRad__Array_to_C(cSlice []COpticalFlowRad, goSlice []OpticalFlowRad) {
	for i := 0; i < len(goSlice); i++ {
		OpticalFlowRadTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
