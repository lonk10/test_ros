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

#include <visualization_msgs/msg/interactive_marker_init.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("visualization_msgs/InteractiveMarkerInit", InteractiveMarkerInitTypeSupport)
	typemap.RegisterMessage("visualization_msgs/msg/InteractiveMarkerInit", InteractiveMarkerInitTypeSupport)
}

type InteractiveMarkerInit struct {
	ServerId string `yaml:"server_id"`// Identifying string. Must be unique in the topic namespacethat this server works on.
	SeqNum uint64 `yaml:"seq_num"`// Sequence number.The client will use this to detect if it has missed a subsequentupdate.  Every update message will have the same sequence number asan init message.  Clients will likely want to unsubscribe from theinit topic after a successful initialization to avoid receivingduplicate data.
	Markers []InteractiveMarker `yaml:"markers"`// All markers.
}

// NewInteractiveMarkerInit creates a new InteractiveMarkerInit with default values.
func NewInteractiveMarkerInit() *InteractiveMarkerInit {
	self := InteractiveMarkerInit{}
	self.SetDefaults()
	return &self
}

func (t *InteractiveMarkerInit) Clone() *InteractiveMarkerInit {
	c := &InteractiveMarkerInit{}
	c.ServerId = t.ServerId
	c.SeqNum = t.SeqNum
	if t.Markers != nil {
		c.Markers = make([]InteractiveMarker, len(t.Markers))
		CloneInteractiveMarkerSlice(c.Markers, t.Markers)
	}
	return c
}

func (t *InteractiveMarkerInit) CloneMsg() types.Message {
	return t.Clone()
}

func (t *InteractiveMarkerInit) SetDefaults() {
	t.ServerId = ""
	t.SeqNum = 0
	t.Markers = nil
}

func (t *InteractiveMarkerInit) GetTypeSupport() types.MessageTypeSupport {
	return InteractiveMarkerInitTypeSupport
}

// InteractiveMarkerInitPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type InteractiveMarkerInitPublisher struct {
	*rclgo.Publisher
}

// NewInteractiveMarkerInitPublisher creates and returns a new publisher for the
// InteractiveMarkerInit
func NewInteractiveMarkerInitPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*InteractiveMarkerInitPublisher, error) {
	pub, err := node.NewPublisher(topic_name, InteractiveMarkerInitTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &InteractiveMarkerInitPublisher{pub}, nil
}

func (p *InteractiveMarkerInitPublisher) Publish(msg *InteractiveMarkerInit) error {
	return p.Publisher.Publish(msg)
}

// InteractiveMarkerInitSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type InteractiveMarkerInitSubscription struct {
	*rclgo.Subscription
}

// InteractiveMarkerInitSubscriptionCallback type is used to provide a subscription
// handler function for a InteractiveMarkerInitSubscription.
type InteractiveMarkerInitSubscriptionCallback func(msg *InteractiveMarkerInit, info *rclgo.MessageInfo, err error)

// NewInteractiveMarkerInitSubscription creates and returns a new subscription for the
// InteractiveMarkerInit
func NewInteractiveMarkerInitSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback InteractiveMarkerInitSubscriptionCallback) (*InteractiveMarkerInitSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg InteractiveMarkerInit
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, InteractiveMarkerInitTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &InteractiveMarkerInitSubscription{sub}, nil
}

func (s *InteractiveMarkerInitSubscription) TakeMessage(out *InteractiveMarkerInit) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneInteractiveMarkerInitSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneInteractiveMarkerInitSlice(dst, src []InteractiveMarkerInit) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var InteractiveMarkerInitTypeSupport types.MessageTypeSupport = _InteractiveMarkerInitTypeSupport{}

type _InteractiveMarkerInitTypeSupport struct{}

func (t _InteractiveMarkerInitTypeSupport) New() types.Message {
	return NewInteractiveMarkerInit()
}

func (t _InteractiveMarkerInitTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.visualization_msgs__msg__InteractiveMarkerInit
	return (unsafe.Pointer)(C.visualization_msgs__msg__InteractiveMarkerInit__create())
}

func (t _InteractiveMarkerInitTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.visualization_msgs__msg__InteractiveMarkerInit__destroy((*C.visualization_msgs__msg__InteractiveMarkerInit)(pointer_to_free))
}

func (t _InteractiveMarkerInitTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*InteractiveMarkerInit)
	mem := (*C.visualization_msgs__msg__InteractiveMarkerInit)(dst)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.server_id), m.ServerId)
	mem.seq_num = C.uint64_t(m.SeqNum)
	InteractiveMarker__Sequence_to_C(&mem.markers, m.Markers)
}

func (t _InteractiveMarkerInitTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*InteractiveMarkerInit)
	mem := (*C.visualization_msgs__msg__InteractiveMarkerInit)(ros2_message_buffer)
	primitives.StringAsGoStruct(&m.ServerId, unsafe.Pointer(&mem.server_id))
	m.SeqNum = uint64(mem.seq_num)
	InteractiveMarker__Sequence_to_Go(&m.Markers, mem.markers)
}

func (t _InteractiveMarkerInitTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__visualization_msgs__msg__InteractiveMarkerInit())
}

type CInteractiveMarkerInit = C.visualization_msgs__msg__InteractiveMarkerInit
type CInteractiveMarkerInit__Sequence = C.visualization_msgs__msg__InteractiveMarkerInit__Sequence

func InteractiveMarkerInit__Sequence_to_Go(goSlice *[]InteractiveMarkerInit, cSlice CInteractiveMarkerInit__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]InteractiveMarkerInit, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		InteractiveMarkerInitTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func InteractiveMarkerInit__Sequence_to_C(cSlice *CInteractiveMarkerInit__Sequence, goSlice []InteractiveMarkerInit) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.visualization_msgs__msg__InteractiveMarkerInit)(C.malloc(C.sizeof_struct_visualization_msgs__msg__InteractiveMarkerInit * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		InteractiveMarkerInitTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func InteractiveMarkerInit__Array_to_Go(goSlice []InteractiveMarkerInit, cSlice []CInteractiveMarkerInit) {
	for i := 0; i < len(cSlice); i++ {
		InteractiveMarkerInitTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func InteractiveMarkerInit__Array_to_C(cSlice []CInteractiveMarkerInit, goSlice []InteractiveMarkerInit) {
	for i := 0; i < len(goSlice); i++ {
		InteractiveMarkerInitTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
