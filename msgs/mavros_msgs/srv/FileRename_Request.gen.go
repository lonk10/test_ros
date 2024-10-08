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

#include <mavros_msgs/srv/file_rename.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/FileRename_Request", FileRename_RequestTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/FileRename_Request", FileRename_RequestTypeSupport)
}

type FileRename_Request struct {
	OldPath string `yaml:"old_path"`
	NewPath string `yaml:"new_path"`
}

// NewFileRename_Request creates a new FileRename_Request with default values.
func NewFileRename_Request() *FileRename_Request {
	self := FileRename_Request{}
	self.SetDefaults()
	return &self
}

func (t *FileRename_Request) Clone() *FileRename_Request {
	c := &FileRename_Request{}
	c.OldPath = t.OldPath
	c.NewPath = t.NewPath
	return c
}

func (t *FileRename_Request) CloneMsg() types.Message {
	return t.Clone()
}

func (t *FileRename_Request) SetDefaults() {
	t.OldPath = ""
	t.NewPath = ""
}

func (t *FileRename_Request) GetTypeSupport() types.MessageTypeSupport {
	return FileRename_RequestTypeSupport
}

// FileRename_RequestPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type FileRename_RequestPublisher struct {
	*rclgo.Publisher
}

// NewFileRename_RequestPublisher creates and returns a new publisher for the
// FileRename_Request
func NewFileRename_RequestPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*FileRename_RequestPublisher, error) {
	pub, err := node.NewPublisher(topic_name, FileRename_RequestTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &FileRename_RequestPublisher{pub}, nil
}

func (p *FileRename_RequestPublisher) Publish(msg *FileRename_Request) error {
	return p.Publisher.Publish(msg)
}

// FileRename_RequestSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type FileRename_RequestSubscription struct {
	*rclgo.Subscription
}

// FileRename_RequestSubscriptionCallback type is used to provide a subscription
// handler function for a FileRename_RequestSubscription.
type FileRename_RequestSubscriptionCallback func(msg *FileRename_Request, info *rclgo.MessageInfo, err error)

// NewFileRename_RequestSubscription creates and returns a new subscription for the
// FileRename_Request
func NewFileRename_RequestSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback FileRename_RequestSubscriptionCallback) (*FileRename_RequestSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg FileRename_Request
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, FileRename_RequestTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &FileRename_RequestSubscription{sub}, nil
}

func (s *FileRename_RequestSubscription) TakeMessage(out *FileRename_Request) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFileRename_RequestSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFileRename_RequestSlice(dst, src []FileRename_Request) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var FileRename_RequestTypeSupport types.MessageTypeSupport = _FileRename_RequestTypeSupport{}

type _FileRename_RequestTypeSupport struct{}

func (t _FileRename_RequestTypeSupport) New() types.Message {
	return NewFileRename_Request()
}

func (t _FileRename_RequestTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__FileRename_Request
	return (unsafe.Pointer)(C.mavros_msgs__srv__FileRename_Request__create())
}

func (t _FileRename_RequestTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__FileRename_Request__destroy((*C.mavros_msgs__srv__FileRename_Request)(pointer_to_free))
}

func (t _FileRename_RequestTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*FileRename_Request)
	mem := (*C.mavros_msgs__srv__FileRename_Request)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.old_path), m.OldPath)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.new_path), m.NewPath)
}

func (t _FileRename_RequestTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*FileRename_Request)
	mem := (*C.mavros_msgs__srv__FileRename_Request)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.OldPath, unsafe.Pointer(&mem.old_path))
	primitives.StringAsGoStruct(&m.NewPath, unsafe.Pointer(&mem.new_path))
}

func (t _FileRename_RequestTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__FileRename_Request())
}

type CFileRename_Request = C.mavros_msgs__srv__FileRename_Request
type CFileRename_Request__Sequence = C.mavros_msgs__srv__FileRename_Request__Sequence

func FileRename_Request__Sequence_to_Go(goSlice *[]FileRename_Request, cSlice CFileRename_Request__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FileRename_Request, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		FileRename_RequestTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func FileRename_Request__Sequence_to_C(cSlice *CFileRename_Request__Sequence, goSlice []FileRename_Request) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__FileRename_Request)(C.malloc(C.sizeof_struct_mavros_msgs__srv__FileRename_Request * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		FileRename_RequestTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func FileRename_Request__Array_to_Go(goSlice []FileRename_Request, cSlice []CFileRename_Request) {
	for i := 0; i < len(cSlice); i++ {
		FileRename_RequestTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func FileRename_Request__Array_to_C(cSlice []CFileRename_Request, goSlice []FileRename_Request) {
	for i := 0; i < len(goSlice); i++ {
		FileRename_RequestTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
