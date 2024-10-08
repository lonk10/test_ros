// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	builtin_interfaces_msg "test/msgs/builtin_interfaces/msg"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/msg/adsb_vehicle.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/ADSBVehicle", ADSBVehicleTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/ADSBVehicle", ADSBVehicleTypeSupport)
}
const (
	ADSBVehicle_ALT_PRESSURE_QNH uint8 = 0// Altitude reported from a Baro source using QNH reference. [[[cog:import mavros_cogmavros_cog.idl_decl_enum('ADSB_ALTITUDE_TYPE', 'ALT_')mavros_cog.idl_decl_enum('ADSB_EMITTER_TYPE', 'EMITTER_')mavros_cog.idl_decl_enum('ADSB_FLAGS', 'FLAG_', 16)]]]ADSB_ALTITUDE_TYPE
	ADSBVehicle_ALT_GEOMETRIC uint8 = 1// Altitude reported from a GNSS source
	ADSBVehicle_EMITTER_NO_INFO uint8 = 0// ADSB_EMITTER_TYPE
	ADSBVehicle_EMITTER_LIGHT uint8 = 1
	ADSBVehicle_EMITTER_SMALL uint8 = 2
	ADSBVehicle_EMITTER_LARGE uint8 = 3
	ADSBVehicle_EMITTER_HIGH_VORTEX_LARGE uint8 = 4
	ADSBVehicle_EMITTER_HEAVY uint8 = 5
	ADSBVehicle_EMITTER_HIGHLY_MANUV uint8 = 6
	ADSBVehicle_EMITTER_ROTOCRAFT uint8 = 7
	ADSBVehicle_EMITTER_UNASSIGNED uint8 = 8
	ADSBVehicle_EMITTER_GLIDER uint8 = 9
	ADSBVehicle_EMITTER_LIGHTER_AIR uint8 = 10
	ADSBVehicle_EMITTER_PARACHUTE uint8 = 11
	ADSBVehicle_EMITTER_ULTRA_LIGHT uint8 = 12
	ADSBVehicle_EMITTER_UNASSIGNED2 uint8 = 13
	ADSBVehicle_EMITTER_UAV uint8 = 14
	ADSBVehicle_EMITTER_SPACE uint8 = 15
	ADSBVehicle_EMITTER_UNASSGINED3 uint8 = 16
	ADSBVehicle_EMITTER_EMERGENCY_SURFACE uint8 = 17
	ADSBVehicle_EMITTER_SERVICE_SURFACE uint8 = 18
	ADSBVehicle_EMITTER_POINT_OBSTACLE uint8 = 19
	ADSBVehicle_FLAG_VALID_COORDS uint16 = 1// ADSB_FLAGS
	ADSBVehicle_FLAG_VALID_ALTITUDE uint16 = 2
	ADSBVehicle_FLAG_VALID_HEADING uint16 = 4
	ADSBVehicle_FLAG_VALID_VELOCITY uint16 = 8
	ADSBVehicle_FLAG_VALID_CALLSIGN uint16 = 16
	ADSBVehicle_FLAG_VALID_SQUAWK uint16 = 32
	ADSBVehicle_FLAG_SIMULATED uint16 = 64
	ADSBVehicle_FLAG_VERTICAL_VELOCITY_VALID uint16 = 128
	ADSBVehicle_FLAG_BARO_VALID uint16 = 256
	ADSBVehicle_FLAG_SOURCE_UAT uint16 = 32768
)

type ADSBVehicle struct {
	Header std_msgs_msg.Header `yaml:"header"`
	IcaoAddress uint32 `yaml:"icao_address"`
	Callsign string `yaml:"callsign"`
	Latitude float64 `yaml:"latitude"`
	Longitude float64 `yaml:"longitude"`
	Altitude float32 `yaml:"altitude"`// AMSL
	Heading float32 `yaml:"heading"`// deg [0..360)
	HorVelocity float32 `yaml:"hor_velocity"`// m/s
	VerVelocity float32 `yaml:"ver_velocity"`// m/s
	AltitudeType uint8 `yaml:"altitude_type"`// Type from ADSB_ALTITUDE_TYPE enum
	EmitterType uint8 `yaml:"emitter_type"`// Type from ADSB_EMITTER_TYPE enum
	Tslc builtin_interfaces_msg.Duration `yaml:"tslc"`// Duration from last communication, seconds [0..255]
	Flags uint16 `yaml:"flags"`// ADSB_FLAGS bit field
	Squawk uint16 `yaml:"squawk"`// Squawk code
}

// NewADSBVehicle creates a new ADSBVehicle with default values.
func NewADSBVehicle() *ADSBVehicle {
	self := ADSBVehicle{}
	self.SetDefaults()
	return &self
}

func (t *ADSBVehicle) Clone() *ADSBVehicle {
	c := &ADSBVehicle{}
	c.Header = *t.Header.Clone()
	c.IcaoAddress = t.IcaoAddress
	c.Callsign = t.Callsign
	c.Latitude = t.Latitude
	c.Longitude = t.Longitude
	c.Altitude = t.Altitude
	c.Heading = t.Heading
	c.HorVelocity = t.HorVelocity
	c.VerVelocity = t.VerVelocity
	c.AltitudeType = t.AltitudeType
	c.EmitterType = t.EmitterType
	c.Tslc = *t.Tslc.Clone()
	c.Flags = t.Flags
	c.Squawk = t.Squawk
	return c
}

func (t *ADSBVehicle) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ADSBVehicle) SetDefaults() {
	t.Header.SetDefaults()
	t.IcaoAddress = 0
	t.Callsign = ""
	t.Latitude = 0
	t.Longitude = 0
	t.Altitude = 0
	t.Heading = 0
	t.HorVelocity = 0
	t.VerVelocity = 0
	t.AltitudeType = 0
	t.EmitterType = 0
	t.Tslc.SetDefaults()
	t.Flags = 0
	t.Squawk = 0
}

func (t *ADSBVehicle) GetTypeSupport() types.MessageTypeSupport {
	return ADSBVehicleTypeSupport
}

// ADSBVehiclePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ADSBVehiclePublisher struct {
	*rclgo.Publisher
}

// NewADSBVehiclePublisher creates and returns a new publisher for the
// ADSBVehicle
func NewADSBVehiclePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ADSBVehiclePublisher, error) {
	pub, err := node.NewPublisher(topic_name, ADSBVehicleTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ADSBVehiclePublisher{pub}, nil
}

func (p *ADSBVehiclePublisher) Publish(msg *ADSBVehicle) error {
	return p.Publisher.Publish(msg)
}

// ADSBVehicleSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ADSBVehicleSubscription struct {
	*rclgo.Subscription
}

// ADSBVehicleSubscriptionCallback type is used to provide a subscription
// handler function for a ADSBVehicleSubscription.
type ADSBVehicleSubscriptionCallback func(msg *ADSBVehicle, info *rclgo.MessageInfo, err error)

// NewADSBVehicleSubscription creates and returns a new subscription for the
// ADSBVehicle
func NewADSBVehicleSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ADSBVehicleSubscriptionCallback) (*ADSBVehicleSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ADSBVehicle
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ADSBVehicleTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ADSBVehicleSubscription{sub}, nil
}

func (s *ADSBVehicleSubscription) TakeMessage(out *ADSBVehicle) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneADSBVehicleSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneADSBVehicleSlice(dst, src []ADSBVehicle) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ADSBVehicleTypeSupport types.MessageTypeSupport = _ADSBVehicleTypeSupport{}

type _ADSBVehicleTypeSupport struct{}

func (t _ADSBVehicleTypeSupport) New() types.Message {
	return NewADSBVehicle()
}

func (t _ADSBVehicleTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__ADSBVehicle
	return (unsafe.Pointer)(C.mavros_msgs__msg__ADSBVehicle__create())
}

func (t _ADSBVehicleTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__ADSBVehicle__destroy((*C.mavros_msgs__msg__ADSBVehicle)(pointer_to_free))
}

func (t _ADSBVehicleTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ADSBVehicle)
	mem := (*C.mavros_msgs__msg__ADSBVehicle)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.icao_address = C.uint32_t(m.IcaoAddress)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.callsign), m.Callsign)
	mem.latitude = C.double(m.Latitude)
	mem.longitude = C.double(m.Longitude)
	mem.altitude = C.float(m.Altitude)
	mem.heading = C.float(m.Heading)
	mem.hor_velocity = C.float(m.HorVelocity)
	mem.ver_velocity = C.float(m.VerVelocity)
	mem.altitude_type = C.uint8_t(m.AltitudeType)
	mem.emitter_type = C.uint8_t(m.EmitterType)
	builtin_interfaces_msg.DurationTypeSupport.AsCStruct(unsafe.Pointer(&mem.tslc), &m.Tslc)
	mem.flags = C.uint16_t(m.Flags)
	mem.squawk = C.uint16_t(m.Squawk)
}

func (t _ADSBVehicleTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ADSBVehicle)
	mem := (*C.mavros_msgs__msg__ADSBVehicle)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.IcaoAddress = uint32(mem.icao_address)
	primitives.StringAsGoStruct(&m.Callsign, unsafe.Pointer(&mem.callsign))
	m.Latitude = float64(mem.latitude)
	m.Longitude = float64(mem.longitude)
	m.Altitude = float32(mem.altitude)
	m.Heading = float32(mem.heading)
	m.HorVelocity = float32(mem.hor_velocity)
	m.VerVelocity = float32(mem.ver_velocity)
	m.AltitudeType = uint8(mem.altitude_type)
	m.EmitterType = uint8(mem.emitter_type)
	builtin_interfaces_msg.DurationTypeSupport.AsGoStruct(&m.Tslc, unsafe.Pointer(&mem.tslc))
	m.Flags = uint16(mem.flags)
	m.Squawk = uint16(mem.squawk)
}

func (t _ADSBVehicleTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__ADSBVehicle())
}

type CADSBVehicle = C.mavros_msgs__msg__ADSBVehicle
type CADSBVehicle__Sequence = C.mavros_msgs__msg__ADSBVehicle__Sequence

func ADSBVehicle__Sequence_to_Go(goSlice *[]ADSBVehicle, cSlice CADSBVehicle__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ADSBVehicle, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ADSBVehicleTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ADSBVehicle__Sequence_to_C(cSlice *CADSBVehicle__Sequence, goSlice []ADSBVehicle) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__ADSBVehicle)(C.malloc(C.sizeof_struct_mavros_msgs__msg__ADSBVehicle * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ADSBVehicleTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ADSBVehicle__Array_to_Go(goSlice []ADSBVehicle, cSlice []CADSBVehicle) {
	for i := 0; i < len(cSlice); i++ {
		ADSBVehicleTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ADSBVehicle__Array_to_C(cSlice []CADSBVehicle, goSlice []ADSBVehicle) {
	for i := 0; i < len(goSlice); i++ {
		ADSBVehicleTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
