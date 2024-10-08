// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/command_long.h>
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
	typemap.RegisterService("mavros_msgs/CommandLong", CommandLongTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/CommandLong", CommandLongTypeSupport)
}

type _CommandLongTypeSupport struct {}

func (s _CommandLongTypeSupport) Request() types.MessageTypeSupport {
	return CommandLong_RequestTypeSupport
}

func (s _CommandLongTypeSupport) Response() types.MessageTypeSupport {
	return CommandLong_ResponseTypeSupport
}

func (s _CommandLongTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__CommandLong())
}

// Modifying this variable is undefined behavior.
var CommandLongTypeSupport types.ServiceTypeSupport = _CommandLongTypeSupport{}

// CommandLongClient wraps rclgo.Client to provide type safe helper
// functions
type CommandLongClient struct {
	*rclgo.Client
}

// NewCommandLongClient creates and returns a new client for the
// CommandLong
func NewCommandLongClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*CommandLongClient, error) {
	client, err := node.NewClient(serviceName, CommandLongTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CommandLongClient{client}, nil
}

func (s *CommandLongClient) Send(ctx context.Context, req *CommandLong_Request) (*CommandLong_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*CommandLong_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type CommandLongServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s CommandLongServiceResponseSender) SendResponse(resp *CommandLong_Response) error {
	return s.sender.SendResponse(resp)
}

type CommandLongServiceRequestHandler func(*rclgo.ServiceInfo, *CommandLong_Request, CommandLongServiceResponseSender)

// CommandLongService wraps rclgo.Service to provide type safe helper
// functions
type CommandLongService struct {
	*rclgo.Service
}

// NewCommandLongService creates and returns a new service for the
// CommandLong
func NewCommandLongService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler CommandLongServiceRequestHandler) (*CommandLongService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*CommandLong_Request)
		responseSender := CommandLongServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, CommandLongTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &CommandLongService{service}, nil
}