// Code generated by rclgo-gen. DO NOT EDIT.

package rosbag2_interfaces_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <rosbag2_interfaces/srv/resume.h>
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
	typemap.RegisterService("rosbag2_interfaces/Resume", ResumeTypeSupport)
	typemap.RegisterService("rosbag2_interfaces/srv/Resume", ResumeTypeSupport)
}

type _ResumeTypeSupport struct {}

func (s _ResumeTypeSupport) Request() types.MessageTypeSupport {
	return Resume_RequestTypeSupport
}

func (s _ResumeTypeSupport) Response() types.MessageTypeSupport {
	return Resume_ResponseTypeSupport
}

func (s _ResumeTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__rosbag2_interfaces__srv__Resume())
}

// Modifying this variable is undefined behavior.
var ResumeTypeSupport types.ServiceTypeSupport = _ResumeTypeSupport{}

// ResumeClient wraps rclgo.Client to provide type safe helper
// functions
type ResumeClient struct {
	*rclgo.Client
}

// NewResumeClient creates and returns a new client for the
// Resume
func NewResumeClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*ResumeClient, error) {
	client, err := node.NewClient(serviceName, ResumeTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &ResumeClient{client}, nil
}

func (s *ResumeClient) Send(ctx context.Context, req *Resume_Request) (*Resume_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*Resume_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type ResumeServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s ResumeServiceResponseSender) SendResponse(resp *Resume_Response) error {
	return s.sender.SendResponse(resp)
}

type ResumeServiceRequestHandler func(*rclgo.ServiceInfo, *Resume_Request, ResumeServiceResponseSender)

// ResumeService wraps rclgo.Service to provide type safe helper
// functions
type ResumeService struct {
	*rclgo.Service
}

// NewResumeService creates and returns a new service for the
// Resume
func NewResumeService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler ResumeServiceRequestHandler) (*ResumeService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*Resume_Request)
		responseSender := ResumeServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, ResumeTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &ResumeService{service}, nil
}