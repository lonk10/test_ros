// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/param_get.h>
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
	typemap.RegisterService("mavros_msgs/ParamGet", ParamGetTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/ParamGet", ParamGetTypeSupport)
}

type _ParamGetTypeSupport struct {}

func (s _ParamGetTypeSupport) Request() types.MessageTypeSupport {
	return ParamGet_RequestTypeSupport
}

func (s _ParamGetTypeSupport) Response() types.MessageTypeSupport {
	return ParamGet_ResponseTypeSupport
}

func (s _ParamGetTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__ParamGet())
}

// Modifying this variable is undefined behavior.
var ParamGetTypeSupport types.ServiceTypeSupport = _ParamGetTypeSupport{}

// ParamGetClient wraps rclgo.Client to provide type safe helper
// functions
type ParamGetClient struct {
	*rclgo.Client
}

// NewParamGetClient creates and returns a new client for the
// ParamGet
func NewParamGetClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*ParamGetClient, error) {
	client, err := node.NewClient(serviceName, ParamGetTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ParamGetClient{client}, nil
}

func (s *ParamGetClient) Send(ctx context.Context, req *ParamGet_Request) (*ParamGet_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*ParamGet_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type ParamGetServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s ParamGetServiceResponseSender) SendResponse(resp *ParamGet_Response) error {
	return s.sender.SendResponse(resp)
}

type ParamGetServiceRequestHandler func(*rclgo.ServiceInfo, *ParamGet_Request, ParamGetServiceResponseSender)

// ParamGetService wraps rclgo.Service to provide type safe helper
// functions
type ParamGetService struct {
	*rclgo.Service
}

// NewParamGetService creates and returns a new service for the
// ParamGet
func NewParamGetService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler ParamGetServiceRequestHandler) (*ParamGetService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*ParamGet_Request)
		responseSender := ParamGetServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, ParamGetTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &ParamGetService{service}, nil
}