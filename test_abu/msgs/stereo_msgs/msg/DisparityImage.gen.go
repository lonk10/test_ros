// Code generated by rclgo-gen. DO NOT EDIT.

package stereo_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	sensor_msgs_msg "test/msgs/sensor_msgs/msg"
	std_msgs_msg "test/msgs/std_msgs/msg"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <stereo_msgs/msg/disparity_image.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("stereo_msgs/DisparityImage", DisparityImageTypeSupport)
	typemap.RegisterMessage("stereo_msgs/msg/DisparityImage", DisparityImageTypeSupport)
}

type DisparityImage struct {
	Header std_msgs_msg.Header `yaml:"header"`// Separate header for compatibility with current TimeSynchronizer.Likely to be removed in a later release, use image.header instead.
	Image sensor_msgs_msg.Image `yaml:"image"`// Floating point disparity image. The disparities are pre-adjusted for anyx-offset between the principal points of the two cameras (in the casethat they are verged). That is: d = x_l - x_r - (cx_l - cx_r)
	F float32 `yaml:"f"`// Focal length, pixels. Stereo geometry. For disparity d, the depth from the camera is Z = fT/d.
	T float32 `yaml:"t"`// Baseline, world units
	ValidWindow sensor_msgs_msg.RegionOfInterest `yaml:"valid_window"`// Subwindow of (potentially) valid disparity values.
	MinDisparity float32 `yaml:"min_disparity"`// The range of disparities searched.In the disparity image, any disparity less than min_disparity is invalid.The disparity search range defines the horopter, or 3D volume that thestereo algorithm can "see". Points with Z outside of:Z_min = fT / max_disparityZ_max = fT / min_disparitycould not be found.
	MaxDisparity float32 `yaml:"max_disparity"`
	DeltaD float32 `yaml:"delta_d"`// Smallest allowed disparity increment. The smallest achievable depth rangeresolution is delta_Z = (Z^2/fT)*delta_d.
}

// NewDisparityImage creates a new DisparityImage with default values.
func NewDisparityImage() *DisparityImage {
	self := DisparityImage{}
	self.SetDefaults()
	return &self
}

func (t *DisparityImage) Clone() *DisparityImage {
	c := &DisparityImage{}
	c.Header = *t.Header.Clone()
	c.Image = *t.Image.Clone()
	c.F = t.F
	c.T = t.T
	c.ValidWindow = *t.ValidWindow.Clone()
	c.MinDisparity = t.MinDisparity
	c.MaxDisparity = t.MaxDisparity
	c.DeltaD = t.DeltaD
	return c
}

func (t *DisparityImage) CloneMsg() types.Message {
	return t.Clone()
}

func (t *DisparityImage) SetDefaults() {
	t.Header.SetDefaults()
	t.Image.SetDefaults()
	t.F = 0
	t.T = 0
	t.ValidWindow.SetDefaults()
	t.MinDisparity = 0
	t.MaxDisparity = 0
	t.DeltaD = 0
}

func (t *DisparityImage) GetTypeSupport() types.MessageTypeSupport {
	return DisparityImageTypeSupport
}

// DisparityImagePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type DisparityImagePublisher struct {
	*rclgo.Publisher
}

// NewDisparityImagePublisher creates and returns a new publisher for the
// DisparityImage
func NewDisparityImagePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*DisparityImagePublisher, error) {
	pub, err := node.NewPublisher(topic_name, DisparityImageTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &DisparityImagePublisher{pub}, nil
}

func (p *DisparityImagePublisher) Publish(msg *DisparityImage) error {
	return p.Publisher.Publish(msg)
}

// DisparityImageSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type DisparityImageSubscription struct {
	*rclgo.Subscription
}

// DisparityImageSubscriptionCallback type is used to provide a subscription
// handler function for a DisparityImageSubscription.
type DisparityImageSubscriptionCallback func(msg *DisparityImage, info *rclgo.MessageInfo, err error)

// NewDisparityImageSubscription creates and returns a new subscription for the
// DisparityImage
func NewDisparityImageSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback DisparityImageSubscriptionCallback) (*DisparityImageSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg DisparityImage
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, DisparityImageTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &DisparityImageSubscription{sub}, nil
}

func (s *DisparityImageSubscription) TakeMessage(out *DisparityImage) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneDisparityImageSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneDisparityImageSlice(dst, src []DisparityImage) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var DisparityImageTypeSupport types.MessageTypeSupport = _DisparityImageTypeSupport{}

type _DisparityImageTypeSupport struct{}

func (t _DisparityImageTypeSupport) New() types.Message {
	return NewDisparityImage()
}

func (t _DisparityImageTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.stereo_msgs__msg__DisparityImage
	return (unsafe.Pointer)(C.stereo_msgs__msg__DisparityImage__create())
}

func (t _DisparityImageTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.stereo_msgs__msg__DisparityImage__destroy((*C.stereo_msgs__msg__DisparityImage)(pointer_to_free))
}

func (t _DisparityImageTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*DisparityImage)
	mem := (*C.stereo_msgs__msg__DisparityImage)(dst)
	std_msgs_msg.HeaderTypeSupport.AsCStruct(unsafe.Pointer(&mem.header), &m.Header)
	sensor_msgs_msg.ImageTypeSupport.AsCStruct(unsafe.Pointer(&mem.image), &m.Image)
	mem.f = C.float(m.F)
	mem.t = C.float(m.T)
	sensor_msgs_msg.RegionOfInterestTypeSupport.AsCStruct(unsafe.Pointer(&mem.valid_window), &m.ValidWindow)
	mem.min_disparity = C.float(m.MinDisparity)
	mem.max_disparity = C.float(m.MaxDisparity)
	mem.delta_d = C.float(m.DeltaD)
}

func (t _DisparityImageTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*DisparityImage)
	mem := (*C.stereo_msgs__msg__DisparityImage)(ros2_message_buffer)
	std_msgs_msg.HeaderTypeSupport.AsGoStruct(&m.Header, unsafe.Pointer(&mem.header))
	sensor_msgs_msg.ImageTypeSupport.AsGoStruct(&m.Image, unsafe.Pointer(&mem.image))
	m.F = float32(mem.f)
	m.T = float32(mem.t)
	sensor_msgs_msg.RegionOfInterestTypeSupport.AsGoStruct(&m.ValidWindow, unsafe.Pointer(&mem.valid_window))
	m.MinDisparity = float32(mem.min_disparity)
	m.MaxDisparity = float32(mem.max_disparity)
	m.DeltaD = float32(mem.delta_d)
}

func (t _DisparityImageTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__stereo_msgs__msg__DisparityImage())
}

type CDisparityImage = C.stereo_msgs__msg__DisparityImage
type CDisparityImage__Sequence = C.stereo_msgs__msg__DisparityImage__Sequence

func DisparityImage__Sequence_to_Go(goSlice *[]DisparityImage, cSlice CDisparityImage__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]DisparityImage, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		DisparityImageTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func DisparityImage__Sequence_to_C(cSlice *CDisparityImage__Sequence, goSlice []DisparityImage) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.stereo_msgs__msg__DisparityImage)(C.malloc(C.sizeof_struct_stereo_msgs__msg__DisparityImage * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		DisparityImageTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func DisparityImage__Array_to_Go(goSlice []DisparityImage, cSlice []CDisparityImage) {
	for i := 0; i < len(cSlice); i++ {
		DisparityImageTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func DisparityImage__Array_to_C(cSlice []CDisparityImage, goSlice []DisparityImage) {
	for i := 0; i < len(goSlice); i++ {
		DisparityImageTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
