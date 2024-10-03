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

#include <mavros_msgs/msg/gpsraw.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/GPSRAW", GPSRAWTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/GPSRAW", GPSRAWTypeSupport)
}
const (
	GPSRAW_GPS_FIX_TYPE_NO_GPS uint8 = 0// No GPS connected. # GPS_FIX_TYPE enum
	GPSRAW_GPS_FIX_TYPE_NO_FIX uint8 = 1// No position information, GPS is connected
	GPSRAW_GPS_FIX_TYPE_2D_FIX uint8 = 2// 2D position
	GPSRAW_GPS_FIX_TYPE_3D_FIX uint8 = 3// 3D position
	GPSRAW_GPS_FIX_TYPE_DGPS uint8 = 4// DGPS/SBAS aided 3D position
	GPSRAW_GPS_FIX_TYPE_RTK_FLOAT uint8 = 5// RTK float, 3D position
	GPSRAW_GPS_FIX_TYPE_RTK_FIXED uint8 = 6// RTK Fixed, 3D position
	GPSRAW_GPS_FIX_TYPE_STATIC uint8 = 7// Static fixed, typically used for base stations
	GPSRAW_GPS_FIX_TYPE_PPP uint8 = 8// PPP, 3D position
)

type GPSRAW struct {
	Header std_msgs_msg.Header `yaml:"header"`
	FixType uint8 `yaml:"fix_type"`// [GPS_FIX_TYPE] GPS fix type
	Lat int32 `yaml:"lat"`// [degE7] Latitude (WGS84, EGM96 ellipsoid)
	Lon int32 `yaml:"lon"`// [degE7] Longitude (WGS84, EGM96 ellipsoid)
	Alt int32 `yaml:"alt"`// [mm] Altitude (MSL). Positive for up. Note that virtually all GPS modules provide the MSL altitude in addition to the WGS84 altitude.
	Eph uint16 `yaml:"eph"`// GPS HDOP horizontal dilution of position (unitless). If unknown, set to: UINT16_MAX
	Epv uint16 `yaml:"epv"`// GPS VDOP vertical dilution of position (unitless). If unknown, set to: UINT16_MAX
	Vel uint16 `yaml:"vel"`// [cm/s] GPS ground speed. If unknown, set to: UINT16_MAX
	Cog uint16 `yaml:"cog"`// [cdeg] Course over ground (NOT heading, but direction of movement), 0.0..359.99 degrees. If unknown, set to: UINT16_MAX
	SatellitesVisible uint8 `yaml:"satellites_visible"`// Number of satellites visible. If unknown, set to 255
	AltEllipsoid int32 `yaml:"alt_ellipsoid"`// [mm] Altitude (above WGS84, EGM96 ellipsoid). Positive for up.. -*- only available with MAVLink v2.0 and GPS_RAW_INT messages -*-
	HAcc uint32 `yaml:"h_acc"`// [mm] Position uncertainty. Positive for up.
	VAcc uint32 `yaml:"v_acc"`// [mm] Altitude uncertainty. Positive for up.
	VelAcc uint32 `yaml:"vel_acc"`// [mm] Speed uncertainty. Positive for up.
	HdgAcc int32 `yaml:"hdg_acc"`// [degE5] Heading / track uncertainty
	Yaw uint16 `yaml:"yaw"`// [cdeg] Yaw in earth frame from north.
	DgpsNumch uint8 `yaml:"dgps_numch"`// Number of DGPS satellites. -*- only available with MAVLink v2.0 and GPS2_RAW messages -*-
	DgpsAge uint32 `yaml:"dgps_age"`// [ms] Age of DGPS info
}

// NewGPSRAW creates a new GPSRAW with default values.
func NewGPSRAW() *GPSRAW {
	self := GPSRAW{}
	self.SetDefaults()
	return &self
}

func (t *GPSRAW) Clone() *GPSRAW {
	c := &GPSRAW{}
	c.Header = *t.Header.Clone()
	c.FixType = t.FixType
	c.Lat = t.Lat
	c.Lon = t.Lon
	c.Alt = t.Alt
	c.Eph = t.Eph
	c.Epv = t.Epv
	c.Vel = t.Vel
	c.Cog = t.Cog
	c.SatellitesVisible = t.SatellitesVisible
	c.AltEllipsoid = t.AltEllipsoid
	c.HAcc = t.HAcc
	c.VAcc = t.VAcc
	c.VelAcc = t.VelAcc
	c.HdgAcc = t.HdgAcc
	c.Yaw = t.Yaw
	c.DgpsNumch = t.DgpsNumch
	c.DgpsAge = t.DgpsAge
	return c
}

func (t *GPSRAW) CloneMsg() types.Message {
	return t.Clone()
}

func (t *GPSRAW) SetDefaults() {
	t.Header.SetDefaults()
	t.FixType = 0
	t.Lat = 0
	t.Lon = 0
	t.Alt = 0
	t.Eph = 0
	t.Epv = 0
	t.Vel = 0
	t.Cog = 0
	t.SatellitesVisible = 0
	t.AltEllipsoid = 0
	t.HAcc = 0
	t.VAcc = 0
	t.VelAcc = 0
	t.HdgAcc = 0
	t.Yaw = 0
	t.DgpsNumch = 0
	t.DgpsAge = 0
}

func (t *GPSRAW) GetTypeSupport() types.MessageTypeSupport {
	return GPSRAWTypeSupport
}

// GPSRAWPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type GPSRAWPublisher struct {
	*rclgo.Publisher
}

// NewGPSRAWPublisher creates and returns a new publisher for the
// GPSRAW
func NewGPSRAWPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*GPSRAWPublisher, error) {
	pub, err := node.NewPublisher(topic_name, GPSRAWTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GPSRAWPublisher{pub}, nil
}

func (p *GPSRAWPublisher) Publish(msg *GPSRAW) error {
	return p.Publisher.Publish(msg)
}

// GPSRAWSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type GPSRAWSubscription struct {
	*rclgo.Subscription
}

// GPSRAWSubscriptionCallback type is used to provide a subscription
// handler function for a GPSRAWSubscription.
type GPSRAWSubscriptionCallback func(msg *GPSRAW, info *rclgo.MessageInfo, err error)

// NewGPSRAWSubscription creates and returns a new subscription for the
// GPSRAW
func NewGPSRAWSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback GPSRAWSubscriptionCallback) (*GPSRAWSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg GPSRAW
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, GPSRAWTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &GPSRAWSubscription{sub}, nil
}

func (s *GPSRAWSubscription) TakeMessage(out *GPSRAW) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneGPSRAWSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneGPSRAWSlice(dst, src []GPSRAW) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var GPSRAWTypeSupport types.MessageTypeSupport = _GPSRAWTypeSupport{}

type _GPSRAWTypeSupport struct{}

func (t _GPSRAWTypeSupport) New() types.Message {
	return NewGPSRAW()
}

func (t _GPSRAWTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__GPSRAW
	return (unsafe.Pointer)(C.mavros_msgs__msg__GPSRAW__create())
}

func (t _GPSRAWTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__GPSRAW__destroy((*C.mavros_msgs__msg__GPSRAW)(pointer_to_free))
}

func (t _GPSRAWTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*GPSRAW)
	mem := (*C.mavros_msgs__msg__GPSRAW)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.fix_type = C.uint8_t(m.FixType)
	mem.lat = C.int32_t(m.Lat)
	mem.lon = C.int32_t(m.Lon)
	mem.alt = C.int32_t(m.Alt)
	mem.eph = C.uint16_t(m.Eph)
	mem.epv = C.uint16_t(m.Epv)
	mem.vel = C.uint16_t(m.Vel)
	mem.cog = C.uint16_t(m.Cog)
	mem.satellites_visible = C.uint8_t(m.SatellitesVisible)
	mem.alt_ellipsoid = C.int32_t(m.AltEllipsoid)
	mem.h_acc = C.uint32_t(m.HAcc)
	mem.v_acc = C.uint32_t(m.VAcc)
	mem.vel_acc = C.uint32_t(m.VelAcc)
	mem.hdg_acc = C.int32_t(m.HdgAcc)
	mem.yaw = C.uint16_t(m.Yaw)
	mem.dgps_numch = C.uint8_t(m.DgpsNumch)
	mem.dgps_age = C.uint32_t(m.DgpsAge)
}

func (t _GPSRAWTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*GPSRAW)
	mem := (*C.mavros_msgs__msg__GPSRAW)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.FixType = uint8(mem.fix_type)
	m.Lat = int32(mem.lat)
	m.Lon = int32(mem.lon)
	m.Alt = int32(mem.alt)
	m.Eph = uint16(mem.eph)
	m.Epv = uint16(mem.epv)
	m.Vel = uint16(mem.vel)
	m.Cog = uint16(mem.cog)
	m.SatellitesVisible = uint8(mem.satellites_visible)
	m.AltEllipsoid = int32(mem.alt_ellipsoid)
	m.HAcc = uint32(mem.h_acc)
	m.VAcc = uint32(mem.v_acc)
	m.VelAcc = uint32(mem.vel_acc)
	m.HdgAcc = int32(mem.hdg_acc)
	m.Yaw = uint16(mem.yaw)
	m.DgpsNumch = uint8(mem.dgps_numch)
	m.DgpsAge = uint32(mem.dgps_age)
}

func (t _GPSRAWTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__GPSRAW())
}

type CGPSRAW = C.mavros_msgs__msg__GPSRAW
type CGPSRAW__Sequence = C.mavros_msgs__msg__GPSRAW__Sequence

func GPSRAW__Sequence_to_Go(goSlice *[]GPSRAW, cSlice CGPSRAW__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]GPSRAW, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		GPSRAWTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func GPSRAW__Sequence_to_C(cSlice *CGPSRAW__Sequence, goSlice []GPSRAW) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__GPSRAW)(C.malloc(C.sizeof_struct_mavros_msgs__msg__GPSRAW * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		GPSRAWTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func GPSRAW__Array_to_Go(goSlice []GPSRAW, cSlice []CGPSRAW) {
	for i := 0; i < len(cSlice); i++ {
		GPSRAWTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func GPSRAW__Array_to_C(cSlice []CGPSRAW, goSlice []GPSRAW) {
	for i := 0; i < len(goSlice); i++ {
		GPSRAWTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
