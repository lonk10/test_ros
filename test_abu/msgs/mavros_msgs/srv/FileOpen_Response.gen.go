// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <mavros_msgs/srv/file_open.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/FileOpen_Response", FileOpen_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/FileOpen_Response", FileOpen_ResponseTypeSupport)
}

type FileOpen_Response struct {
	Size uint32 `yaml:"size"`
	Success bool `yaml:"success"`
	RErrno int32 `yaml:"r_errno"`
}

// NewFileOpen_Response creates a new FileOpen_Response with default values.
func NewFileOpen_Response() *FileOpen_Response {
	self := FileOpen_Response{}
	self.SetDefaults()
	return &self
}

func (t *FileOpen_Response) Clone() *FileOpen_Response {
	c := &FileOpen_Response{}
	c.Size = t.Size
	c.Success = t.Success
	c.RErrno = t.RErrno
	return c
}

func (t *FileOpen_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *FileOpen_Response) SetDefaults() {
	t.Size = 0
	t.Success = false
	t.RErrno = 0
}

func (t *FileOpen_Response) GetTypeSupport() types.MessageTypeSupport {
	return FileOpen_ResponseTypeSupport
}

// FileOpen_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type FileOpen_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewFileOpen_ResponsePublisher creates and returns a new publisher for the
// FileOpen_Response
func NewFileOpen_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*FileOpen_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, FileOpen_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &FileOpen_ResponsePublisher{pub}, nil
}

func (p *FileOpen_ResponsePublisher) Publish(msg *FileOpen_Response) error {
	return p.Publisher.Publish(msg)
}

// FileOpen_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type FileOpen_ResponseSubscription struct {
	*rclgo.Subscription
}

// FileOpen_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a FileOpen_ResponseSubscription.
type FileOpen_ResponseSubscriptionCallback func(msg *FileOpen_Response, info *rclgo.MessageInfo, err error)

// NewFileOpen_ResponseSubscription creates and returns a new subscription for the
// FileOpen_Response
func NewFileOpen_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback FileOpen_ResponseSubscriptionCallback) (*FileOpen_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg FileOpen_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, FileOpen_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &FileOpen_ResponseSubscription{sub}, nil
}

func (s *FileOpen_ResponseSubscription) TakeMessage(out *FileOpen_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFileOpen_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFileOpen_ResponseSlice(dst, src []FileOpen_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var FileOpen_ResponseTypeSupport types.MessageTypeSupport = _FileOpen_ResponseTypeSupport{}

type _FileOpen_ResponseTypeSupport struct{}

func (t _FileOpen_ResponseTypeSupport) New() types.Message {
	return NewFileOpen_Response()
}

func (t _FileOpen_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__FileOpen_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__FileOpen_Response__create())
}

func (t _FileOpen_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__FileOpen_Response__destroy((*C.mavros_msgs__srv__FileOpen_Response)(pointer_to_free))
}

func (t _FileOpen_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*FileOpen_Response)
	mem := (*C.mavros_msgs__srv__FileOpen_Response)(dst)
	mem.size = C.uint32_t(m.Size)
	mem.success = C.bool(m.Success)
	mem.r_errno = C.int32_t(m.RErrno)
}

func (t _FileOpen_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*FileOpen_Response)
	mem := (*C.mavros_msgs__srv__FileOpen_Response)(ros2_message_buffer)
	m.Size = uint32(mem.size)
	m.Success = bool(mem.success)
	m.RErrno = int32(mem.r_errno)
}

func (t _FileOpen_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__FileOpen_Response())
}

type CFileOpen_Response = C.mavros_msgs__srv__FileOpen_Response
type CFileOpen_Response__Sequence = C.mavros_msgs__srv__FileOpen_Response__Sequence

func FileOpen_Response__Sequence_to_Go(goSlice *[]FileOpen_Response, cSlice CFileOpen_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FileOpen_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		FileOpen_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func FileOpen_Response__Sequence_to_C(cSlice *CFileOpen_Response__Sequence, goSlice []FileOpen_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__FileOpen_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__FileOpen_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		FileOpen_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func FileOpen_Response__Array_to_Go(goSlice []FileOpen_Response, cSlice []CFileOpen_Response) {
	for i := 0; i < len(cSlice); i++ {
		FileOpen_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func FileOpen_Response__Array_to_C(cSlice []CFileOpen_Response, goSlice []FileOpen_Response) {
	for i := 0; i < len(goSlice); i++ {
		FileOpen_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
