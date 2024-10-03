// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/command_vtol_transition.h>
*/
import "C"

import (
	"context"
	"errors"
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"
)

func init() {
	typemap.RegisterService("mavros_msgs/CommandVtolTransition", CommandVtolTransitionTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/CommandVtolTransition", CommandVtolTransitionTypeSupport)
}

type _CommandVtolTransitionTypeSupport struct {}

func (s _CommandVtolTransitionTypeSupport) Request() types.MessageTypeSupport {
	return CommandVtolTransition_RequestTypeSupport
}

func (s _CommandVtolTransitionTypeSupport) Response() types.MessageTypeSupport {
	return CommandVtolTransition_ResponseTypeSupport
}

func (s _CommandVtolTransitionTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__CommandVtolTransition())
}

// Modifying this variable is undefined behavior.
var CommandVtolTransitionTypeSupport types.ServiceTypeSupport = _CommandVtolTransitionTypeSupport{}

// CommandVtolTransitionClient wraps rclgo.Client to provide type safe helper
// functions
type CommandVtolTransitionClient struct {
	*rclgo.Client
}

// NewCommandVtolTransitionClient creates and returns a new client for the
// CommandVtolTransition
func NewCommandVtolTransitionClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*CommandVtolTransitionClient, error) {
	client, err := node.NewClient(serviceName, CommandVtolTransitionTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CommandVtolTransitionClient{client}, nil
}

func (s *CommandVtolTransitionClient) Send(ctx context.Context, req *CommandVtolTransition_Request) (*CommandVtolTransition_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*CommandVtolTransition_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type CommandVtolTransitionServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s CommandVtolTransitionServiceResponseSender) SendResponse(resp *CommandVtolTransition_Response) error {
	return s.sender.SendResponse(resp)
}

type CommandVtolTransitionServiceRequestHandler func(*rclgo.ServiceInfo, *CommandVtolTransition_Request, CommandVtolTransitionServiceResponseSender)

// CommandVtolTransitionService wraps rclgo.Service to provide type safe helper
// functions
type CommandVtolTransitionService struct {
	*rclgo.Service
}

// NewCommandVtolTransitionService creates and returns a new service for the
// CommandVtolTransition
func NewCommandVtolTransitionService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler CommandVtolTransitionServiceRequestHandler) (*CommandVtolTransitionService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*CommandVtolTransition_Request)
		responseSender := CommandVtolTransitionServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, CommandVtolTransitionTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &CommandVtolTransitionService{service}, nil
}