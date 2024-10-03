// Code generated by rclgo-gen. DO NOT EDIT.

package rosbag2_interfaces_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <rosbag2_interfaces/srv/pause.h>
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
	typemap.RegisterService("rosbag2_interfaces/Pause", PauseTypeSupport)
	typemap.RegisterService("rosbag2_interfaces/srv/Pause", PauseTypeSupport)
}

type _PauseTypeSupport struct {}

func (s _PauseTypeSupport) Request() types.MessageTypeSupport {
	return Pause_RequestTypeSupport
}

func (s _PauseTypeSupport) Response() types.MessageTypeSupport {
	return Pause_ResponseTypeSupport
}

func (s _PauseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__rosbag2_interfaces__srv__Pause())
}

// Modifying this variable is undefined behavior.
var PauseTypeSupport types.ServiceTypeSupport = _PauseTypeSupport{}

// PauseClient wraps rclgo.Client to provide type safe helper
// functions
type PauseClient struct {
	*rclgo.Client
}

// NewPauseClient creates and returns a new client for the
// Pause
func NewPauseClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*PauseClient, error) {
	client, err := node.NewClient(serviceName, PauseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &PauseClient{client}, nil
}

func (s *PauseClient) Send(ctx context.Context, req *Pause_Request) (*Pause_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*Pause_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type PauseServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s PauseServiceResponseSender) SendResponse(resp *Pause_Response) error {
	return s.sender.SendResponse(resp)
}

type PauseServiceRequestHandler func(*rclgo.ServiceInfo, *Pause_Request, PauseServiceResponseSender)

// PauseService wraps rclgo.Service to provide type safe helper
// functions
type PauseService struct {
	*rclgo.Service
}

// NewPauseService creates and returns a new service for the
// Pause
func NewPauseService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler PauseServiceRequestHandler) (*PauseService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*Pause_Request)
		responseSender := PauseServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, PauseTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &PauseService{service}, nil
}