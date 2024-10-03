// Code generated by rclgo-gen. DO NOT EDIT.

package turtlesim_action

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <turtlesim/action/rotate_absolute.h>
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
	typemap.RegisterService("turtlesim/RotateAbsolute_SendGoal", RotateAbsolute_SendGoalTypeSupport)
	typemap.RegisterService("turtlesim/action/RotateAbsolute_SendGoal", RotateAbsolute_SendGoalTypeSupport)
}

type _RotateAbsolute_SendGoalTypeSupport struct {}

func (s _RotateAbsolute_SendGoalTypeSupport) Request() types.MessageTypeSupport {
	return RotateAbsolute_SendGoal_RequestTypeSupport
}

func (s _RotateAbsolute_SendGoalTypeSupport) Response() types.MessageTypeSupport {
	return RotateAbsolute_SendGoal_ResponseTypeSupport
}

func (s _RotateAbsolute_SendGoalTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__turtlesim__action__RotateAbsolute_SendGoal())
}

// Modifying this variable is undefined behavior.
var RotateAbsolute_SendGoalTypeSupport types.ServiceTypeSupport = _RotateAbsolute_SendGoalTypeSupport{}

// RotateAbsolute_SendGoalClient wraps rclgo.Client to provide type safe helper
// functions
type RotateAbsolute_SendGoalClient struct {
	*rclgo.Client
}

// NewRotateAbsolute_SendGoalClient creates and returns a new client for the
// RotateAbsolute_SendGoal
func NewRotateAbsolute_SendGoalClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*RotateAbsolute_SendGoalClient, error) {
	client, err := node.NewClient(serviceName, RotateAbsolute_SendGoalTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &RotateAbsolute_SendGoalClient{client}, nil
}

func (s *RotateAbsolute_SendGoalClient) Send(ctx context.Context, req *RotateAbsolute_SendGoal_Request) (*RotateAbsolute_SendGoal_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*RotateAbsolute_SendGoal_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type RotateAbsolute_SendGoalServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s RotateAbsolute_SendGoalServiceResponseSender) SendResponse(resp *RotateAbsolute_SendGoal_Response) error {
	return s.sender.SendResponse(resp)
}

type RotateAbsolute_SendGoalServiceRequestHandler func(*rclgo.ServiceInfo, *RotateAbsolute_SendGoal_Request, RotateAbsolute_SendGoalServiceResponseSender)

// RotateAbsolute_SendGoalService wraps rclgo.Service to provide type safe helper
// functions
type RotateAbsolute_SendGoalService struct {
	*rclgo.Service
}

// NewRotateAbsolute_SendGoalService creates and returns a new service for the
// RotateAbsolute_SendGoal
func NewRotateAbsolute_SendGoalService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler RotateAbsolute_SendGoalServiceRequestHandler) (*RotateAbsolute_SendGoalService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*RotateAbsolute_SendGoal_Request)
		responseSender := RotateAbsolute_SendGoalServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, RotateAbsolute_SendGoalTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &RotateAbsolute_SendGoalService{service}, nil
}