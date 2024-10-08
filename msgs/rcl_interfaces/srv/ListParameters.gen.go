// Code generated by rclgo-gen. DO NOT EDIT.

package rcl_interfaces_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <rcl_interfaces/srv/list_parameters.h>
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
	typemap.RegisterService("rcl_interfaces/ListParameters", ListParametersTypeSupport)
	typemap.RegisterService("rcl_interfaces/srv/ListParameters", ListParametersTypeSupport)
}

type _ListParametersTypeSupport struct {}

func (s _ListParametersTypeSupport) Request() types.MessageTypeSupport {
	return ListParameters_RequestTypeSupport
}

func (s _ListParametersTypeSupport) Response() types.MessageTypeSupport {
	return ListParameters_ResponseTypeSupport
}

func (s _ListParametersTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__rcl_interfaces__srv__ListParameters())
}

// Modifying this variable is undefined behavior.
var ListParametersTypeSupport types.ServiceTypeSupport = _ListParametersTypeSupport{}

// ListParametersClient wraps rclgo.Client to provide type safe helper
// functions
type ListParametersClient struct {
	*rclgo.Client
}

// NewListParametersClient creates and returns a new client for the
// ListParameters
func NewListParametersClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*ListParametersClient, error) {
	client, err := node.NewClient(serviceName, ListParametersTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ListParametersClient{client}, nil
}

func (s *ListParametersClient) Send(ctx context.Context, req *ListParameters_Request) (*ListParameters_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*ListParameters_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type ListParametersServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s ListParametersServiceResponseSender) SendResponse(resp *ListParameters_Response) error {
	return s.sender.SendResponse(resp)
}

type ListParametersServiceRequestHandler func(*rclgo.ServiceInfo, *ListParameters_Request, ListParametersServiceResponseSender)

// ListParametersService wraps rclgo.Service to provide type safe helper
// functions
type ListParametersService struct {
	*rclgo.Service
}

// NewListParametersService creates and returns a new service for the
// ListParameters
func NewListParametersService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler ListParametersServiceRequestHandler) (*ListParametersService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*ListParameters_Request)
		responseSender := ListParametersServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, ListParametersTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &ListParametersService{service}, nil
}