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

#include <mavros_msgs/msg/terrain_report.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/TerrainReport", TerrainReportTypeSupport)
	typemap.RegisterMessage("mavros_msgs/msg/TerrainReport", TerrainReportTypeSupport)
}

type TerrainReport struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Latitude float64 `yaml:"latitude"`
	Longitude float64 `yaml:"longitude"`
	Spacing uint16 `yaml:"spacing"`
	TerrainHeight float32 `yaml:"terrain_height"`// in meters, terrain height
	CurrentHeight float32 `yaml:"current_height"`// in meters, vehicle height above terrain
	Pending uint16 `yaml:"pending"`
	Loaded uint16 `yaml:"loaded"`
}

// NewTerrainReport creates a new TerrainReport with default values.
func NewTerrainReport() *TerrainReport {
	self := TerrainReport{}
	self.SetDefaults()
	return &self
}

func (t *TerrainReport) Clone() *TerrainReport {
	c := &TerrainReport{}
	c.Header = *t.Header.Clone()
	c.Latitude = t.Latitude
	c.Longitude = t.Longitude
	c.Spacing = t.Spacing
	c.TerrainHeight = t.TerrainHeight
	c.CurrentHeight = t.CurrentHeight
	c.Pending = t.Pending
	c.Loaded = t.Loaded
	return c
}

func (t *TerrainReport) CloneMsg() types.Message {
	return t.Clone()
}

func (t *TerrainReport) SetDefaults() {
	t.Header.SetDefaults()
	t.Latitude = 0
	t.Longitude = 0
	t.Spacing = 0
	t.TerrainHeight = 0
	t.CurrentHeight = 0
	t.Pending = 0
	t.Loaded = 0
}

func (t *TerrainReport) GetTypeSupport() types.MessageTypeSupport {
	return TerrainReportTypeSupport
}

// TerrainReportPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TerrainReportPublisher struct {
	*rclgo.Publisher
}

// NewTerrainReportPublisher creates and returns a new publisher for the
// TerrainReport
func NewTerrainReportPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TerrainReportPublisher, error) {
	pub, err := node.NewPublisher(topic_name, TerrainReportTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TerrainReportPublisher{pub}, nil
}

func (p *TerrainReportPublisher) Publish(msg *TerrainReport) error {
	return p.Publisher.Publish(msg)
}

// TerrainReportSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TerrainReportSubscription struct {
	*rclgo.Subscription
}

// TerrainReportSubscriptionCallback type is used to provide a subscription
// handler function for a TerrainReportSubscription.
type TerrainReportSubscriptionCallback func(msg *TerrainReport, info *rclgo.MessageInfo, err error)

// NewTerrainReportSubscription creates and returns a new subscription for the
// TerrainReport
func NewTerrainReportSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback TerrainReportSubscriptionCallback) (*TerrainReportSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg TerrainReport
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TerrainReportTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &TerrainReportSubscription{sub}, nil
}

func (s *TerrainReportSubscription) TakeMessage(out *TerrainReport) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTerrainReportSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTerrainReportSlice(dst, src []TerrainReport) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TerrainReportTypeSupport types.MessageTypeSupport = _TerrainReportTypeSupport{}

type _TerrainReportTypeSupport struct{}

func (t _TerrainReportTypeSupport) New() types.Message {
	return NewTerrainReport()
}

func (t _TerrainReportTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__msg__TerrainReport
	return (unsafe.Pointer)(C.mavros_msgs__msg__TerrainReport__create())
}

func (t _TerrainReportTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__msg__TerrainReport__destroy((*C.mavros_msgs__msg__TerrainReport)(pointer_to_free))
}

func (t _TerrainReportTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*TerrainReport)
	mem := (*C.mavros_msgs__msg__TerrainReport)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	mem.latitude = C.double(m.Latitude)
	mem.longitude = C.double(m.Longitude)
	mem.spacing = C.uint16_t(m.Spacing)
	mem.terrain_height = C.float(m.TerrainHeight)
	mem.current_height = C.float(m.CurrentHeight)
	mem.pending = C.uint16_t(m.Pending)
	mem.loaded = C.uint16_t(m.Loaded)
}

func (t _TerrainReportTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*TerrainReport)
	mem := (*C.mavros_msgs__msg__TerrainReport)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	m.Latitude = float64(mem.latitude)
	m.Longitude = float64(mem.longitude)
	m.Spacing = uint16(mem.spacing)
	m.TerrainHeight = float32(mem.terrain_height)
	m.CurrentHeight = float32(mem.current_height)
	m.Pending = uint16(mem.pending)
	m.Loaded = uint16(mem.loaded)
}

func (t _TerrainReportTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__msg__TerrainReport())
}

type CTerrainReport = C.mavros_msgs__msg__TerrainReport
type CTerrainReport__Sequence = C.mavros_msgs__msg__TerrainReport__Sequence

func TerrainReport__Sequence_to_Go(goSlice *[]TerrainReport, cSlice CTerrainReport__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]TerrainReport, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		TerrainReportTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func TerrainReport__Sequence_to_C(cSlice *CTerrainReport__Sequence, goSlice []TerrainReport) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__msg__TerrainReport)(C.malloc(C.sizeof_struct_mavros_msgs__msg__TerrainReport * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		TerrainReportTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func TerrainReport__Array_to_Go(goSlice []TerrainReport, cSlice []CTerrainReport) {
	for i := 0; i < len(cSlice); i++ {
		TerrainReportTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func TerrainReport__Array_to_C(cSlice []CTerrainReport, goSlice []TerrainReport) {
	for i := 0; i < len(goSlice); i++ {
		TerrainReportTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
