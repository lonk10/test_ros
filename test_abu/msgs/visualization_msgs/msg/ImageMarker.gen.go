// Code generated by rclgo-gen. DO NOT EDIT.

package visualization_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	builtin_interfaces_msg "test/msgs/builtin_interfaces/msg"
	geometry_msgs_msg "test/msgs/geometry_msgs/msg"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/image_marker.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/ImageMarker", ImageMarkerTypeSupport)
	typemap.RegisterMessage("visualization_msgs/msg/ImageMarker", ImageMarkerTypeSupport)
}
const (
	ImageMarker_CIRCLE int32 = 0
	ImageMarker_LINE_STRIP int32 = 1
	ImageMarker_LINE_LIST int32 = 2
	ImageMarker_POLYGON int32 = 3
	ImageMarker_POINTS int32 = 4
	ImageMarker_ADD int32 = 0
	ImageMarker_REMOVE int32 = 1
)

type ImageMarker struct {
	Header std_msgs_msg.Header `yaml:"header"`
	Ns string `yaml:"ns"`// Namespace which is used with the id to form a unique id.
	Id int32 `yaml:"id"`// Unique id within the namespace.
	Type int32 `yaml:"type"`// One of the above types, e.g. CIRCLE, LINE_STRIP, etc.
	Action int32 `yaml:"action"`// Either ADD or REMOVE.
	Position geometry_msgs_msg.Point `yaml:"position"`// Two-dimensional coordinate position, in pixel-coordinates.
	Scale float32 `yaml:"scale"`// The scale of the object, e.g. the diameter for a CIRCLE.
	OutlineColor std_msgs_msg.ColorRGBA `yaml:"outline_color"`// The outline color of the marker.
	Filled uint8 `yaml:"filled"`// Whether or not to fill in the shape with color.
	FillColor std_msgs_msg.ColorRGBA `yaml:"fill_color"`// Fill color; in the range: [0.0-1.0]
	Lifetime builtin_interfaces_msg.Duration `yaml:"lifetime"`// How long the object should last before being automatically deleted.0 indicates forever.
	Points []geometry_msgs_msg.Point `yaml:"points"`// Coordinates in 2D in pixel coords. Used for LINE_STRIP, LINE_LIST, POINTS, etc.
	OutlineColors []std_msgs_msg.ColorRGBA `yaml:"outline_colors"`// The color for each line, point, etc. in the points field.
}

// NewImageMarker creates a new ImageMarker with default values.
func NewImageMarker() *ImageMarker {
	self := ImageMarker{}
	self.SetDefaults()
	return &self
}

func (t *ImageMarker) Clone() *ImageMarker {
	c := &ImageMarker{}
	c.Header = *t.Header.Clone()
	c.Ns = t.Ns
	c.Id = t.Id
	c.Type = t.Type
	c.Action = t.Action
	c.Position = *t.Position.Clone()
	c.Scale = t.Scale
	c.OutlineColor = *t.OutlineColor.Clone()
	c.Filled = t.Filled
	c.FillColor = *t.FillColor.Clone()
	c.Lifetime = *t.Lifetime.Clone()
	if t.Points != nil {
		c.Points = make([]geometry_msgs_msg.Point, len(t.Points))
		geometry_msgs_msg.ClonePointSlice(c.Points, t.Points)
	}
	if t.OutlineColors != nil {
		c.OutlineColors = make([]std_msgs_msg.ColorRGBA, len(t.OutlineColors))
		std_msgs_msg.CloneColorRGBASlice(c.OutlineColors, t.OutlineColors)
	}
	return c
}

func (t *ImageMarker) CloneMsg() types.Message {
	return t.Clone()
}

func (t *ImageMarker) SetDefaults() {
	t.Header.SetDefaults()
	t.Ns = ""
	t.Id = 0
	t.Type = 0
	t.Action = 0
	t.Position.SetDefaults()
	t.Scale = 0
	t.OutlineColor.SetDefaults()
	t.Filled = 0
	t.FillColor.SetDefaults()
	t.Lifetime.SetDefaults()
	t.Points = nil
	t.OutlineColors = nil
}

func (t *ImageMarker) GetTypeSupport() types.MessageTypeSupport {
	return ImageMarkerTypeSupport
}

// ImageMarkerPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type ImageMarkerPublisher struct {
	*rclgo.Publisher
}

// NewImageMarkerPublisher creates and returns a new publisher for the
// ImageMarker
func NewImageMarkerPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*ImageMarkerPublisher, error) {
	pub, err := node.NewPublisher(topic_name, ImageMarkerTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ImageMarkerPublisher{pub}, nil
}

func (p *ImageMarkerPublisher) Publish(msg *ImageMarker) error {
	return p.Publisher.Publish(msg)
}

// ImageMarkerSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type ImageMarkerSubscription struct {
	*rclgo.Subscription
}

// ImageMarkerSubscriptionCallback type is used to provide a subscription
// handler function for a ImageMarkerSubscription.
type ImageMarkerSubscriptionCallback func(msg *ImageMarker, info *rclgo.MessageInfo, err error)

// NewImageMarkerSubscription creates and returns a new subscription for the
// ImageMarker
func NewImageMarkerSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback ImageMarkerSubscriptionCallback) (*ImageMarkerSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg ImageMarker
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, ImageMarkerTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &ImageMarkerSubscription{sub}, nil
}

func (s *ImageMarkerSubscription) TakeMessage(out *ImageMarker) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneImageMarkerSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneImageMarkerSlice(dst, src []ImageMarker) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var ImageMarkerTypeSupport types.MessageTypeSupport = _ImageMarkerTypeSupport{}

type _ImageMarkerTypeSupport struct{}

func (t _ImageMarkerTypeSupport) New() types.Message {
	return NewImageMarker()
}

func (t _ImageMarkerTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__ImageMarker
	return (unsafe.Pointer)(C.visualization_msgs__msg__ImageMarker__create())
}

func (t _ImageMarkerTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__ImageMarker__destroy((*C.visualization_msgs__msg__ImageMarker)(pointer_to_free))
}

func (t _ImageMarkerTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*ImageMarker)
	mem := (*C.visualization_msgs__msg__ImageMarker)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.ns), m.Ns)
	mem.id = C.int32_t(m.Id)
	mem._type = C.int32_t(m.Type)
	mem.action = C.int32_t(m.Action)
	geometry_msgs_msg.PointTypeSupport.AsCStruct(unsafe.Pointer(&mem.position), &m.Position)
	mem.scale = C.float(m.Scale)
	std_msgs_msg.ColorRGBATypeSupport.AsCStruct(unsafe.Pointer(&mem.outline_color), &m.OutlineColor)
	mem.filled = C.uint8_t(m.Filled)
	std_msgs_msg.ColorRGBATypeSupport.AsCStruct(unsafe.Pointer(&mem.fill_color), &m.FillColor)
	builtin_interfaces_msg.DurationTypeSupport.AsCStruct(unsafe.Pointer(&mem.lifetime), &m.Lifetime)
	geometry_msgs_msg.Point__Sequence_to_C((*geometry_msgs_msg.CPoint__Sequence)(unsafe.Pointer(&mem.points)), m.Points)
	std_msgs_msg.ColorRGBA__Sequence_to_C((*std_msgs_msg.CColorRGBA__Sequence)(unsafe.Pointer(&mem.outline_colors)), m.OutlineColors)
}

func (t _ImageMarkerTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*ImageMarker)
	mem := (*C.visualization_msgs__msg__ImageMarker)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	primitives.StringAsGoStruct(&m.Ns, unsafe.Pointer(&mem.ns))
	m.Id = int32(mem.id)
	m.Type = int32(mem._type)
	m.Action = int32(mem.action)
	geometry_msgs_msg.PointTypeSupport.AsGoStruct(&m.Position, unsafe.Pointer(&mem.position))
	m.Scale = float32(mem.scale)
	std_msgs_msg.ColorRGBATypeSupport.AsGoStruct(&m.OutlineColor, unsafe.Pointer(&mem.outline_color))
	m.Filled = uint8(mem.filled)
	std_msgs_msg.ColorRGBATypeSupport.AsGoStruct(&m.FillColor, unsafe.Pointer(&mem.fill_color))
	builtin_interfaces_msg.DurationTypeSupport.AsGoStruct(&m.Lifetime, unsafe.Pointer(&mem.lifetime))
	geometry_msgs_msg.Point__Sequence_to_Go(&m.Points, *(*geometry_msgs_msg.CPoint__Sequence)(unsafe.Pointer(&mem.points)))
	std_msgs_msg.ColorRGBA__Sequence_to_Go(&m.OutlineColors, *(*std_msgs_msg.CColorRGBA__Sequence)(unsafe.Pointer(&mem.outline_colors)))
}

func (t _ImageMarkerTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__ImageMarker())
}

type CImageMarker = C.visualization_msgs__msg__ImageMarker
type CImageMarker__Sequence = C.visualization_msgs__msg__ImageMarker__Sequence

func ImageMarker__Sequence_to_Go(goSlice *[]ImageMarker, cSlice CImageMarker__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]ImageMarker, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		ImageMarkerTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func ImageMarker__Sequence_to_C(cSlice *CImageMarker__Sequence, goSlice []ImageMarker) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__ImageMarker)(C.malloc(C.sizeof_struct_visualization_msgs__msg__ImageMarker * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		ImageMarkerTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func ImageMarker__Array_to_Go(goSlice []ImageMarker, cSlice []CImageMarker) {
	for i := 0; i < len(cSlice); i++ {
		ImageMarkerTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func ImageMarker__Array_to_C(cSlice []CImageMarker, goSlice []ImageMarker) {
	for i := 0; i < len(goSlice); i++ {
		ImageMarkerTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
