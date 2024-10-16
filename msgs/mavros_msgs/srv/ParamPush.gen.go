// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/param_push.h>
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
	typemap.RegisterService("mavros_msgs/ParamPush", ParamPushTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/ParamPush", ParamPushTypeSupport)
}

type _ParamPushTypeSupport struct {}

func (s _ParamPushTypeSupport) Request() types.MessageTypeSupport {
	return ParamPush_RequestTypeSupport
}

func (s _ParamPushTypeSupport) Response() types.MessageTypeSupport {
	return ParamPush_ResponseTypeSupport
}

func (s _ParamPushTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__ParamPush())
}

// Modifying this variable is undefined behavior.
var ParamPushTypeSupport types.ServiceTypeSupport = _ParamPushTypeSupport{}

// ParamPushClient wraps rclgo.Client to provide type safe helper
// functions
type ParamPushClient struct {
	*rclgo.Client
}

// NewParamPushClient creates and returns a new client for the
// ParamPush
func NewParamPushClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*ParamPushClient, error) {
	client, err := node.NewClient(serviceName, ParamPushTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ParamPushClient{client}, nil
}

func (s *ParamPushClient) Send(ctx context.Context, req *ParamPush_Request) (*ParamPush_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*ParamPush_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type ParamPushServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s ParamPushServiceResponseSender) SendResponse(resp *ParamPush_Response) error {
	return s.sender.SendResponse(resp)
}

type ParamPushServiceRequestHandler func(*rclgo.ServiceInfo, *ParamPush_Request, ParamPushServiceResponseSender)

// ParamPushService wraps rclgo.Service to provide type safe helper
// functions
type ParamPushService struct {
	*rclgo.Service
}

// NewParamPushService creates and returns a new service for the
// ParamPush
func NewParamPushService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler ParamPushServiceRequestHandler) (*ParamPushService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*ParamPush_Request)
		responseSender := ParamPushServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, ParamPushTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &ParamPushService{service}, nil
}