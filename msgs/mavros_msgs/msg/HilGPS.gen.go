// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	geographic_msgs_msg "test/msgs/geographic_msgs/msg"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/msg/hil_gps.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/HilGPS", HilGPSTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/HilGPS", HilGPSTypeSupport)
}

type HilGPS struct {
	Header std_msgs_msg.Header `yaml:"header"`
	FixType uint8 `yaml:"fix_type"`
	Geo geographic_msgs_msg.GeoPoint `yaml:"geo"`
	Eph uint16 `yaml:"eph"`
	Epv uint16 `yaml:"epv"`
	Vel uint16 `yaml:"vel"`
	Vn int16 `yaml:"vn"`
	Ve int16 `yaml:"ve"`
	Vd int16 `yaml:"vd"`
	Cog uint16 `yaml:"cog"`
	SatellitesVisible uint8 `yaml:"satellites_visible"`
}

// NewHilGPS creates a new HilGPS with default values.
func NewHilGPS() *HilGPS {
	self := HilGPS{}
	self.SetDefaults()
	return &self
}

func (t *HilGPS) Clone() *HilGPS {
	c := &HilGPS{}
	c.Header = *t.Header.Clone()
	c.FixType = t.FixType
	c.Geo = *t.Geo.Clone()
	c.Eph = t.Eph
	c.Epv = t.Epv
	c.Vel = t.Vel
	c.Vn = t.Vn
	c.Ve = t.Ve
	c.Vd = t.Vd
	c.Cog = t.Cog
	c.SatellitesVisible = t.SatellitesVisible
	return c
}

func (t *HilGPS) CloneMsg() types.Message {
	return t.Clone()
}

func (t *HilGPS) SetDefaults() {
	t.Header.SetDefaults()
	t.FixType = 0
	t.Geo.SetDefaults()
	t.Eph = 0
	t.Epv = 0
	t.Vel = 0
	t.Vn = 0
	t.Ve = 0
	t.Vd = 0
	t.Cog = 0
	t.SatellitesVisible = 0
}

func (t *HilGPS) GetTypeSupport() types.MessageTypeSupport {
	return HilGPSTypeSupport
}

// HilGPSPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type HilGPSPublisher struct {
	*rclgo.Publisher
}

// NewHilGPSPublisher creates and returns a new publisher for the
// HilGPS
func NewHilGPSPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*HilGPSPublisher, error) {
	pub, err := node.NewPublisher(topic_name, HilGPSTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &HilGPSPublisher{pub}, nil
}

func (p *HilGPSPublisher) Publish(msg *HilGPS) error {
	return p.Publisher.Publish(msg)
}

// HilGPSSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type HilGPSSubscription struct {
	*rclgo.Subscription
}

// HilGPSSubscriptionCallback type is used to provide a subscription
// handler function for a HilGPSSubscription.
type HilGPSSubscriptionCallback func(msg *HilGPS, info *rclgo.MessageInfo, err error)

// NewHilGPSSubscription creates and returns a new subscription for the
// HilGPS
func NewHilGPSSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback HilGPSSubscriptionCallback) (*HilGPSSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg HilGPS
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, HilGPSTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &HilGPSSubscription{sub}, nil
}

func (s *HilGPSSubscription) TakeMessage(out *HilGPS) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneHilGPSSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneHilGPSSlice(dst, src []HilGPS) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var HilGPSTypeSupport types.MessageTypeSupport = _HilGPSTypeSupport{}

type _HilGPSTypeSupport struct{}

func (t _HilGPSTypeSupport) New() types.Message {
	return NewHilGPS()
}

func (t _HilGPSTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__HilGPS
	return (unsafe.Pointer)(C.mavros_msgs__msg__HilGPS__create())
}

func (t _HilGPSTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__HilGPS__destroy((*C.mavros_msgs__msg__HilGPS)(pointer_to_free))
}

func (t _HilGPSTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*HilGPS)
	mem := (*C.mavros_msgs__msg__HilGPS)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.fix_type = C.uint8_t(m.FixType)
	geographic_msgs_msg.GeoPointTypeSupport.AsCStruct(unsafe.Pointer(&mem.geo), &m.Geo)
	mem.eph = C.uint16_t(m.Eph)
	mem.epv = C.uint16_t(m.Epv)
	mem.vel = C.uint16_t(m.Vel)
	mem.vn = C.int16_t(m.Vn)
	mem.ve = C.int16_t(m.Ve)
	mem.vd = C.int16_t(m.Vd)
	mem.cog = C.uint16_t(m.Cog)
	mem.satellites_visible = C.uint8_t(m.SatellitesVisible)
}

func (t _HilGPSTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*HilGPS)
	mem := (*C.mavros_msgs__msg__HilGPS)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.FixType = uint8(mem.fix_type)
	geographic_msgs_msg.GeoPointTypeSupport.AsGoStruct(&m.Geo, unsafe.Pointer(&mem.geo))
	m.Eph = uint16(mem.eph)
	m.Epv = uint16(mem.epv)
	m.Vel = uint16(mem.vel)
	m.Vn = int16(mem.vn)
	m.Ve = int16(mem.ve)
	m.Vd = int16(mem.vd)
	m.Cog = uint16(mem.cog)
	m.SatellitesVisible = uint8(mem.satellites_visible)
}

func (t _HilGPSTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__HilGPS())
}

type CHilGPS = C.mavros_msgs__msg__HilGPS
type CHilGPS__Sequence = C.mavros_msgs__msg__HilGPS__Sequence

func HilGPS__Sequence_to_Go(goSlice *[]HilGPS, cSlice CHilGPS__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]HilGPS, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		HilGPSTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func HilGPS__Sequence_to_C(cSlice *CHilGPS__Sequence, goSlice []HilGPS) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__HilGPS)(C.malloc(C.sizeof_struct_mavros_msgs__msg__HilGPS * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		HilGPSTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func HilGPS__Array_to_Go(goSlice []HilGPS, cSlice []CHilGPS) {
	for i := 0; i < len(cSlice); i++ {
		HilGPSTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func HilGPS__Array_to_C(cSlice []CHilGPS, goSlice []HilGPS) {
	for i := 0; i < len(goSlice); i++ {
		HilGPSTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
