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

#include <mavros_msgs/srv/file_truncate.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("mavros_msgs/FileTruncate_Response", FileTruncate_ResponseTypeSupport)
	typemap.RegisterMessage("mavros_msgs/srv/FileTruncate_Response", FileTruncate_ResponseTypeSupport)
}

type FileTruncate_Response struct {
	Success bool `yaml:"success"`
	RErrno int32 `yaml:"r_errno"`
}

// NewFileTruncate_Response creates a new FileTruncate_Response with default values.
func NewFileTruncate_Response() *FileTruncate_Response {
	self := FileTruncate_Response{}
	self.SetDefaults()
	return &self
}

func (t *FileTruncate_Response) Clone() *FileTruncate_Response {
	c := &FileTruncate_Response{}
	c.Success = t.Success
	c.RErrno = t.RErrno
	return c
}

func (t *FileTruncate_Response) CloneMsg() types.Message {
	return t.Clone()
}

func (t *FileTruncate_Response) SetDefaults() {
	t.Success = false
	t.RErrno = 0
}

func (t *FileTruncate_Response) GetTypeSupport() types.MessageTypeSupport {
	return FileTruncate_ResponseTypeSupport
}

// FileTruncate_ResponsePublisher wraps rclgo.Publisher to provide type safe helper
// functions
type FileTruncate_ResponsePublisher struct {
	*rclgo.Publisher
}

// NewFileTruncate_ResponsePublisher creates and returns a new publisher for the
// FileTruncate_Response
func NewFileTruncate_ResponsePublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*FileTruncate_ResponsePublisher, error) {
	pub, err := node.NewPublisher(topic_name, FileTruncate_ResponseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &FileTruncate_ResponsePublisher{pub}, nil
}

func (p *FileTruncate_ResponsePublisher) Publish(msg *FileTruncate_Response) error {
	return p.Publisher.Publish(msg)
}

// FileTruncate_ResponseSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type FileTruncate_ResponseSubscription struct {
	*rclgo.Subscription
}

// FileTruncate_ResponseSubscriptionCallback type is used to provide a subscription
// handler function for a FileTruncate_ResponseSubscription.
type FileTruncate_ResponseSubscriptionCallback func(msg *FileTruncate_Response, info *rclgo.MessageInfo, err error)

// NewFileTruncate_ResponseSubscription creates and returns a new subscription for the
// FileTruncate_Response
func NewFileTruncate_ResponseSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback FileTruncate_ResponseSubscriptionCallback) (*FileTruncate_ResponseSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg FileTruncate_Response
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, FileTruncate_ResponseTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &FileTruncate_ResponseSubscription{sub}, nil
}

func (s *FileTruncate_ResponseSubscription) TakeMessage(out *FileTruncate_Response) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneFileTruncate_ResponseSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneFileTruncate_ResponseSlice(dst, src []FileTruncate_Response) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var FileTruncate_ResponseTypeSupport types.MessageTypeSupport = _FileTruncate_ResponseTypeSupport{}

type _FileTruncate_ResponseTypeSupport struct{}

func (t _FileTruncate_ResponseTypeSupport) New() types.Message {
	return NewFileTruncate_Response()
}

func (t _FileTruncate_ResponseTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.mavros_msgs__srv__FileTruncate_Response
	return (unsafe.Pointer)(C.mavros_msgs__srv__FileTruncate_Response__create())
}

func (t _FileTruncate_ResponseTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.mavros_msgs__srv__FileTruncate_Response__destroy((*C.mavros_msgs__srv__FileTruncate_Response)(pointer_to_free))
}

func (t _FileTruncate_ResponseTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*FileTruncate_Response)
	mem := (*C.mavros_msgs__srv__FileTruncate_Response)(dst)
	mem.success = C.bool(m.Success)
	mem.r_errno = C.int32_t(m.RErrno)
}

func (t _FileTruncate_ResponseTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*FileTruncate_Response)
	mem := (*C.mavros_msgs__srv__FileTruncate_Response)(ros2_message_buffer)
	m.Success = bool(mem.success)
	m.RErrno = int32(mem.r_errno)
}

func (t _FileTruncate_ResponseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__mavros_msgs__srv__FileTruncate_Response())
}

type CFileTruncate_Response = C.mavros_msgs__srv__FileTruncate_Response
type CFileTruncate_Response__Sequence = C.mavros_msgs__srv__FileTruncate_Response__Sequence

func FileTruncate_Response__Sequence_to_Go(goSlice *[]FileTruncate_Response, cSlice CFileTruncate_Response__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]FileTruncate_Response, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		FileTruncate_ResponseTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func FileTruncate_Response__Sequence_to_C(cSlice *CFileTruncate_Response__Sequence, goSlice []FileTruncate_Response) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.mavros_msgs__srv__FileTruncate_Response)(C.malloc(C.sizeof_struct_mavros_msgs__srv__FileTruncate_Response * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		FileTruncate_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func FileTruncate_Response__Array_to_Go(goSlice []FileTruncate_Response, cSlice []CFileTruncate_Response) {
	for i := 0; i < len(cSlice); i++ {
		FileTruncate_ResponseTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func FileTruncate_Response__Array_to_C(cSlice []CFileTruncate_Response, goSlice []FileTruncate_Response) {
	for i := 0; i < len(goSlice); i++ {
		FileTruncate_ResponseTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
