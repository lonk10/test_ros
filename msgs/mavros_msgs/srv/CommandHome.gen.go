// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/command_home.h>
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
	typemap.RegisterService("mavros_msgs/CommandHome", CommandHomeTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/CommandHome", CommandHomeTypeSupport)
}

type _CommandHomeTypeSupport struct {}

func (s _CommandHomeTypeSupport) Request() types.MessageTypeSupport {
	return CommandHome_RequestTypeSupport
}

func (s _CommandHomeTypeSupport) Response() types.MessageTypeSupport {
	return CommandHome_ResponseTypeSupport
}

func (s _CommandHomeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__CommandHome())
}

// Modifying this variable is undefined behavior.
var CommandHomeTypeSupport types.ServiceTypeSupport = _CommandHomeTypeSupport{}

// CommandHomeClient wraps rclgo.Client to provide type safe helper
// functions
type CommandHomeClient struct {
	*rclgo.Client
}

// NewCommandHomeClient creates and returns a new client for the
// CommandHome
func NewCommandHomeClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*CommandHomeClient, error) {
	client, err := node.NewClient(serviceName, CommandHomeTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &CommandHomeClient{client}, nil
}

func (s *CommandHomeClient) Send(ctx context.Context, req *CommandHome_Request) (*CommandHome_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*CommandHome_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type CommandHomeServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s CommandHomeServiceResponseSender) SendResponse(resp *CommandHome_Response) error {
	return s.sender.SendResponse(resp)
}

type CommandHomeServiceRequestHandler func(*rclgo.ServiceInfo, *CommandHome_Request, CommandHomeServiceResponseSender)

// CommandHomeService wraps rclgo.Service to provide type safe helper
// functions
type CommandHomeService struct {
	*rclgo.Service
}

// NewCommandHomeService creates and returns a new service for the
// CommandHome
func NewCommandHomeService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler CommandHomeServiceRequestHandler) (*CommandHomeService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*CommandHome_Request)
		responseSender := CommandHomeServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, CommandHomeTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &CommandHomeService{service}, nil
}