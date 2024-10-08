// Code generated by rclgo-gen. DO NOT EDIT.

package visualization_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <visualization_msgs/msg/mesh_file.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/MeshFile", MeshFileTypeSupport)
	typemap.RegisterMessage("visualization_msgs/msg/MeshFile", MeshFileTypeSupport)
}

type MeshFile struct {
	Filename string `yaml:"filename"`// The filename is used for both debug purposes and to provide a file extensionfor whatever parser is used.
	Data []uint8 `yaml:"data"`// This stores the raw text of the mesh file.
}

// NewMeshFile creates a new MeshFile with default values.
func NewMeshFile() *MeshFile {
	self := MeshFile{}
	self.SetDefaults()
	return &self
}

func (t *MeshFile) Clone() *MeshFile {
	c := &MeshFile{}
	c.Filename = t.Filename
	if t.Data != nil {
		c.Data = make([]uint8, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *MeshFile) CloneMsg() types.Message {
	return t.Clone()
}

func (t *MeshFile) SetDefaults() {
	t.Filename = ""
	t.Data = nil
}

func (t *MeshFile) GetTypeSupport() types.MessageTypeSupport {
	return MeshFileTypeSupport
}

// MeshFilePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type MeshFilePublisher struct {
	*rclgo.Publisher
}

// NewMeshFilePublisher creates and returns a new publisher for the
// MeshFile
func NewMeshFilePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*MeshFilePublisher, error) {
	pub, err := node.NewPublisher(topic_name, MeshFileTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &MeshFilePublisher{pub}, nil
}

func (p *MeshFilePublisher) Publish(msg *MeshFile) error {
	return p.Publisher.Publish(msg)
}

// MeshFileSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type MeshFileSubscription struct {
	*rclgo.Subscription
}

// MeshFileSubscriptionCallback type is used to provide a subscription
// handler function for a MeshFileSubscription.
type MeshFileSubscriptionCallback func(msg *MeshFile, info *rclgo.MessageInfo, err error)

// NewMeshFileSubscription creates and returns a new subscription for the
// MeshFile
func NewMeshFileSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback MeshFileSubscriptionCallback) (*MeshFileSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg MeshFile
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, MeshFileTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &MeshFileSubscription{sub}, nil
}

func (s *MeshFileSubscription) TakeMessage(out *MeshFile) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneMeshFileSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneMeshFileSlice(dst, src []MeshFile) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var MeshFileTypeSupport types.MessageTypeSupport = _MeshFileTypeSupport{}

type _MeshFileTypeSupport struct{}

func (t _MeshFileTypeSupport) New() types.Message {
	return NewMeshFile()
}

func (t _MeshFileTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__MeshFile
	return (unsafe.Pointer)(C.visualization_msgs__msg__MeshFile__create())
}

func (t _MeshFileTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__MeshFile__destroy((*C.visualization_msgs__msg__MeshFile)(pointer_to_free))
}

func (t _MeshFileTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*MeshFile)
	mem := (*C.visualization_msgs__msg__MeshFile)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.filename), m.Filename)
	primitives.Uint8__Sequence_to_C((*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _MeshFileTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*MeshFile)
	mem := (*C.visualization_msgs__msg__MeshFile)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.Filename, unsafe.Pointer(&mem.filename))
	primitives.Uint8__Sequence_to_Go(&m.Data, *(*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _MeshFileTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__MeshFile())
}

type CMeshFile = C.visualization_msgs__msg__MeshFile
type CMeshFile__Sequence = C.visualization_msgs__msg__MeshFile__Sequence

func MeshFile__Sequence_to_Go(goSlice *[]MeshFile, cSlice CMeshFile__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]MeshFile, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		MeshFileTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func MeshFile__Sequence_to_C(cSlice *CMeshFile__Sequence, goSlice []MeshFile) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__MeshFile)(C.malloc(C.sizeof_struct_visualization_msgs__msg__MeshFile * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		MeshFileTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func MeshFile__Array_to_Go(goSlice []MeshFile, cSlice []CMeshFile) {
	for i := 0; i < len(cSlice); i++ {
		MeshFileTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func MeshFile__Array_to_C(cSlice []CMeshFile, goSlice []MeshFile) {
	for i := 0; i < len(goSlice); i++ {
		MeshFileTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
