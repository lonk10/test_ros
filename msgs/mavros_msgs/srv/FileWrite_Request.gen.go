// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/srv/file_write.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/FileWrite_Request", FileWrite_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/FileWrite_Request", FileWrite_RequestTypeSupport)
}

type FileWrite_Request struct {
	FilePath string `yaml:"file_path"`
	Offset uint64 `yaml:"offset"`
	Data []uint8 `yaml:"data"`
}

// NewFileWrite_Request creates a new FileWrite_Request with default values.
func NewFileWrite_Request() *FileWrite_Request {
	self := FileWrite_Request{}
	self.SetDefaults()
	return &self
}

func (t *FileWrite_Request) Clone() *FileWrite_Request {
	c := &FileWrite_Request{}
	c.FilePath = t.FilePath
	c.Offset = t.Offset
	if t.Data != nil {
		c.Data = make([]uint8, len(t.Data))
		copy(c.Data, t.Data)
	}
	return c
}

func (t *FileWrite_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *FileWrite_Request) SetDefaults() {
	t.FilePath = ""
	t.Offset = 0
	t.Data = nil
}

func (t *FileWrite_Request) GetTypeSupport() types.MessageTypeSupport {
	return FileWrite_RequestTypeSupport
}

// FileWrite_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type FileWrite_RequestPublisher struct {
	*rclgo.Publisher
}

// NewFileWrite_RequestPublisher creates and returns a new publisher for the
// FileWrite_Request
func NewFileWrite_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*FileWrite_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, FileWrite_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &FileWrite_RequestPublisher{pub}, nil
}

func (p *FileWrite_RequestPublisher) Publish(msg *FileWrite_Request) error {
	return p.Publisher.Publish(msg)
}

// FileWrite_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type FileWrite_RequestSubscription struct {
	*rclgo.Subscription
}

// FileWrite_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a FileWrite_RequestSubscription.
type FileWrite_RequestSubscriptionCallback func(msg *FileWrite_Request, info *rclgo.MessageInfo, err error)

// NewFileWrite_RequestSubscription creates and returns a new subscription for the
// FileWrite_Request
func NewFileWrite_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback FileWrite_RequestSubscriptionCallback) (*FileWrite_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg FileWrite_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, FileWrite_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &FileWrite_RequestSubscription{sub}, nil
}

func (s *FileWrite_RequestSubscription) TakeMessage(out *FileWrite_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFileWrite_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFileWrite_RequestSlice(dst, src []FileWrite_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var FileWrite_RequestTypeSupport types.MessageTypeSupport = _FileWrite_RequestTypeSupport{}

type _FileWrite_RequestTypeSupport struct{}

func (t _FileWrite_RequestTypeSupport) New() types.Message {
	return NewFileWrite_Request()
}

func (t _FileWrite_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__FileWrite_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__FileWrite_Request__create())
}

func (t _FileWrite_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__FileWrite_Request__destroy((*C.mavros_msgs__srv__FileWrite_Request)(pointer_to_free))
}

func (t _FileWrite_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*FileWrite_Request)
	mem := (*C.mavros_msgs__srv__FileWrite_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.file_path), m.FilePath)
	mem.offset = C.uint64_t(m.Offset)
	primitives.Uint8__Sequence_to_C((*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)), m.Data)
}

func (t _FileWrite_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*FileWrite_Request)
	mem := (*C.mavros_msgs__srv__FileWrite_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.FilePath, unsafe.Pointer(&mem.file_path))
	m.Offset = uint64(mem.offset)
	primitives.Uint8__Sequence_to_Go(&m.Data, *(*primitives.CUint8__Sequence)(unsafe.Pointer(&mem.data)))
}

func (t _FileWrite_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__FileWrite_Request())
}

type CFileWrite_Request = C.mavros_msgs__srv__FileWrite_Request
type CFileWrite_Request__Sequence = C.mavros_msgs__srv__FileWrite_Request__Sequence

func FileWrite_Request__Sequence_to_Go(goSlice *[]FileWrite_Request, cSlice CFileWrite_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FileWrite_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		FileWrite_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func FileWrite_Request__Sequence_to_C(cSlice *CFileWrite_Request__Sequence, goSlice []FileWrite_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__FileWrite_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__FileWrite_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		FileWrite_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func FileWrite_Request__Array_to_Go(goSlice []FileWrite_Request, cSlice []CFileWrite_Request) {
	for i := 0; i < len(cSlice); i++ {
		FileWrite_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func FileWrite_Request__Array_to_C(cSlice []CFileWrite_Request, goSlice []FileWrite_Request) {
	for i := 0; i < len(goSlice); i++ {
		FileWrite_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
