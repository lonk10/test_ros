// Code generated by rclgo-gen. DO NOT EDIT.

package std_srvs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <std_srvs/srv/set_bool.h>
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
	typemap.RegisterService("std_srvs/SetBool", SetBoolTypeSupport)
	typemap.RegisterService("std_srvs/srv/SetBool", SetBoolTypeSupport)
}

type _SetBoolTypeSupport struct {}

func (s _SetBoolTypeSupport) Request() types.MessageTypeSupport {
	return SetBool_RequestTypeSupport
}

func (s _SetBoolTypeSupport) Response() types.MessageTypeSupport {
	return SetBool_ResponseTypeSupport
}

func (s _SetBoolTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__std_srvs__srv__SetBool())
}

// Modifying this variable is undefined behavior.
var SetBoolTypeSupport types.ServiceTypeSupport = _SetBoolTypeSupport{}

// SetBoolClient wraps rclgo.Client to provide type safe helper
// functions
type SetBoolClient struct {
	*rclgo.Client
}

// NewSetBoolClient creates and returns a new client for the
// SetBool
func NewSetBoolClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*SetBoolClient, error) {
	client, err := node.NewClient(serviceName, SetBoolTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SetBoolClient{client}, nil
}

func (s *SetBoolClient) Send(ctx context.Context, req *SetBool_Request) (*SetBool_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*SetBool_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type SetBoolServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s SetBoolServiceResponseSender) SendResponse(resp *SetBool_Response) error {
	return s.sender.SendResponse(resp)
}

type SetBoolServiceRequestHandler func(*rclgo.ServiceInfo, *SetBool_Request, SetBoolServiceResponseSender)

// SetBoolService wraps rclgo.Service to provide type safe helper
// functions
type SetBoolService struct {
	*rclgo.Service
}

// NewSetBoolService creates and returns a new service for the
// SetBool
func NewSetBoolService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler SetBoolServiceRequestHandler) (*SetBoolService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*SetBool_Request)
		responseSender := SetBoolServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, SetBoolTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &SetBoolService{service}, nil
}