// Code generated by rclgo-gen. DO NOT EDIT.

package lifecycle_msgs_msg
import (
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	primitives "github.com/tiiuae/rclgo/pkg/rclgo/primitives"
	
)
/*
#include <rosidl_runtime_c/message_type_support_struct.h>

#include <lifecycle_msgs/msg/transition.h>

*/
import "C"

func init() {
	typemap.RegisterMessage("lifecycle_msgs/Transition", TransitionTypeSupport)
	typemap.RegisterMessage("lifecycle_msgs/msg/Transition", TransitionTypeSupport)
}
const (
	Transition_TRANSITION_CREATE uint8 = 0// This transition will instantiate the node, but will not run any code beyondthe constructor.
	Transition_TRANSITION_CONFIGURE uint8 = 1// The node's onConfigure callback will be called to allow the node to load itsconfiguration and conduct any required setup.
	Transition_TRANSITION_CLEANUP uint8 = 2// The node's callback onCleanup will be called in this transition to allow thenode to load its configuration and conduct any required setup.
	Transition_TRANSITION_ACTIVATE uint8 = 3// The node's callback onActivate will be executed to do any final preparationsto start executing.
	Transition_TRANSITION_DEACTIVATE uint8 = 4// The node's callback onDeactivate will be executed to do any cleanup to startexecuting, and reverse the onActivate changes.
	Transition_TRANSITION_UNCONFIGURED_SHUTDOWN uint8 = 5// This signals shutdown during an unconfigured state, the node's callbackonShutdown will be executed to do any cleanup necessary before destruction.
	Transition_TRANSITION_INACTIVE_SHUTDOWN uint8 = 6// This signals shutdown during an inactive state, the node's callback onShutdownwill be executed to do any cleanup necessary before destruction.
	Transition_TRANSITION_ACTIVE_SHUTDOWN uint8 = 7// This signals shutdown during an active state, the node's callback onShutdownwill be executed to do any cleanup necessary before destruction.
	Transition_TRANSITION_DESTROY uint8 = 8// This transition will simply cause the deallocation of the node.
	Transition_TRANSITION_ON_CONFIGURE_SUCCESS uint8 = 10// Reserved [10-69], private transitionsThese transitions are not publicly available and cannot be invoked by a user.The following transitions are implicitly invoked based on the callbackfeedback of the intermediate transition states.
	Transition_TRANSITION_ON_CONFIGURE_FAILURE uint8 = 11
	Transition_TRANSITION_ON_CONFIGURE_ERROR uint8 = 12
	Transition_TRANSITION_ON_CLEANUP_SUCCESS uint8 = 20
	Transition_TRANSITION_ON_CLEANUP_FAILURE uint8 = 21
	Transition_TRANSITION_ON_CLEANUP_ERROR uint8 = 22
	Transition_TRANSITION_ON_ACTIVATE_SUCCESS uint8 = 30
	Transition_TRANSITION_ON_ACTIVATE_FAILURE uint8 = 31
	Transition_TRANSITION_ON_ACTIVATE_ERROR uint8 = 32
	Transition_TRANSITION_ON_DEACTIVATE_SUCCESS uint8 = 40
	Transition_TRANSITION_ON_DEACTIVATE_FAILURE uint8 = 41
	Transition_TRANSITION_ON_DEACTIVATE_ERROR uint8 = 42
	Transition_TRANSITION_ON_SHUTDOWN_SUCCESS uint8 = 50
	Transition_TRANSITION_ON_SHUTDOWN_FAILURE uint8 = 51
	Transition_TRANSITION_ON_SHUTDOWN_ERROR uint8 = 52
	Transition_TRANSITION_ON_ERROR_SUCCESS uint8 = 60
	Transition_TRANSITION_ON_ERROR_FAILURE uint8 = 61
	Transition_TRANSITION_ON_ERROR_ERROR uint8 = 62
	Transition_TRANSITION_CALLBACK_SUCCESS uint8 = 97// The transition callback successfully performed its required functionality.
	Transition_TRANSITION_CALLBACK_FAILURE uint8 = 98// The transition callback failed to perform its required functionality.
	Transition_TRANSITION_CALLBACK_ERROR uint8 = 99// The transition callback encountered an error that requires special cleanup, ifpossible.
)

type Transition struct {
	Id uint8 `yaml:"id"`// The transition id from above definitions.
	Label string `yaml:"label"`// A text label of the transition.
}

// NewTransition creates a new Transition with default values.
func NewTransition() *Transition {
	self := Transition{}
	self.SetDefaults()
	return &self
}

func (t *Transition) Clone() *Transition {
	c := &Transition{}
	c.Id = t.Id
	c.Label = t.Label
	return c
}

func (t *Transition) CloneMsg() types.Message {
	return t.Clone()
}

func (t *Transition) SetDefaults() {
	t.Id = 0
	t.Label = ""
}

func (t *Transition) GetTypeSupport() types.MessageTypeSupport {
	return TransitionTypeSupport
}

// TransitionPublisher wraps rclgo.Publisher to provide type safe helper
// functions
type TransitionPublisher struct {
	*rclgo.Publisher
}

// NewTransitionPublisher creates and returns a new publisher for the
// Transition
func NewTransitionPublisher(node *rclgo.Node, topic_name string, options *rclgo.PublisherOptions) (*TransitionPublisher, error) {
	pub, err := node.NewPublisher(topic_name, TransitionTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &TransitionPublisher{pub}, nil
}

func (p *TransitionPublisher) Publish(msg *Transition) error {
	return p.Publisher.Publish(msg)
}

// TransitionSubscription wraps rclgo.Subscription to provide type safe helper
// functions
type TransitionSubscription struct {
	*rclgo.Subscription
}

// TransitionSubscriptionCallback type is used to provide a subscription
// handler function for a TransitionSubscription.
type TransitionSubscriptionCallback func(msg *Transition, info *rclgo.MessageInfo, err error)

// NewTransitionSubscription creates and returns a new subscription for the
// Transition
func NewTransitionSubscription(node *rclgo.Node, topic_name string, opts *rclgo.SubscriptionOptions, subscriptionCallback TransitionSubscriptionCallback) (*TransitionSubscription, error) {
	callback := func(s *rclgo.Subscription) {
		var msg Transition
		info, err := s.TakeMessage(&msg)
		subscriptionCallback(&msg, info, err)
	}
	sub, err := node.NewSubscription(topic_name, TransitionTypeSupport, opts, callback)
	if err != nil {
		return nil, err
	}
	return &TransitionSubscription{sub}, nil
}

func (s *TransitionSubscription) TakeMessage(out *Transition) (*rclgo.MessageInfo, error) {
	return s.Subscription.TakeMessage(out)
}

// CloneTransitionSlice clones src to dst by calling Clone for each element in
// src. Panics if len(dst) < len(src).
func CloneTransitionSlice(dst, src []Transition) {
	for i := range src {
		dst[i] = *src[i].Clone()
	}
}

// Modifying this variable is undefined behavior.
var TransitionTypeSupport types.MessageTypeSupport = _TransitionTypeSupport{}

type _TransitionTypeSupport struct{}

func (t _TransitionTypeSupport) New() types.Message {
	return NewTransition()
}

func (t _TransitionTypeSupport) PrepareMemory() unsafe.Pointer { //returns *C.lifecycle_msgs__msg__Transition
	return (unsafe.Pointer)(C.lifecycle_msgs__msg__Transition__create())
}

func (t _TransitionTypeSupport) ReleaseMemory(pointer_to_free unsafe.Pointer) {
	C.lifecycle_msgs__msg__Transition__destroy((*C.lifecycle_msgs__msg__Transition)(pointer_to_free))
}

func (t _TransitionTypeSupport) AsCStruct(dst unsafe.Pointer, msg types.Message) {
	m := msg.(*Transition)
	mem := (*C.lifecycle_msgs__msg__Transition)(dst)
	mem.id = C.uint8_t(m.Id)
	primitives.StringAsCStruct(unsafe.Pointer(&mem.label), m.Label)
}

func (t _TransitionTypeSupport) AsGoStruct(msg types.Message, ros2_message_buffer unsafe.Pointer) {
	m := msg.(*Transition)
	mem := (*C.lifecycle_msgs__msg__Transition)(ros2_message_buffer)
	m.Id = uint8(mem.id)
	primitives.StringAsGoStruct(&m.Label, unsafe.Pointer(&mem.label))
}

func (t _TransitionTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_message_type_support_handle__lifecycle_msgs__msg__Transition())
}

type CTransition = C.lifecycle_msgs__msg__Transition
type CTransition__Sequence = C.lifecycle_msgs__msg__Transition__Sequence

func Transition__Sequence_to_Go(goSlice *[]Transition, cSlice CTransition__Sequence) {
	if cSlice.size == 0 {
		return
	}
	*goSlice = make([]Transition, cSlice.size)
	src := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range src {
		TransitionTypeSupport.AsGoStruct(&(*goSlice)[i], unsafe.Pointer(&src[i]))
	}
}
func Transition__Sequence_to_C(cSlice *CTransition__Sequence, goSlice []Transition) {
	if len(goSlice) == 0 {
		cSlice.data = nil
		cSlice.capacity = 0
		cSlice.size = 0
		return
	}
	cSlice.data = (*C.lifecycle_msgs__msg__Transition)(C.malloc(C.sizeof_struct_lifecycle_msgs__msg__Transition * C.size_t(len(goSlice))))
	cSlice.capacity = C.size_t(len(goSlice))
	cSlice.size = cSlice.capacity
	dst := unsafe.Slice(cSlice.data, cSlice.size)
	for i := range goSlice {
		TransitionTypeSupport.AsCStruct(unsafe.Pointer(&dst[i]), &goSlice[i])
	}
}
func Transition__Array_to_Go(goSlice []Transition, cSlice []CTransition) {
	for i := 0; i < len(cSlice); i++ {
		TransitionTypeSupport.AsGoStruct(&goSlice[i], unsafe.Pointer(&cSlice[i]))
	}
}
func Transition__Array_to_C(cSlice []CTransition, goSlice []Transition) {
	for i := 0; i < len(goSlice); i++ {
		TransitionTypeSupport.AsCStruct(unsafe.Pointer(&cSlice[i]), &goSlice[i])
	}
}
