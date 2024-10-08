// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/file_close.h>
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
	typemap.RegisterService("mavros_msgs/FileClose", FileCloseTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/FileClose", FileCloseTypeSupport)
}

type _FileCloseTypeSupport struct {}

func (s _FileCloseTypeSupport) Request() types.MessageTypeSupport {
	return FileClose_RequestTypeSupport
}

func (s _FileCloseTypeSupport) Response() types.MessageTypeSupport {
	return FileClose_ResponseTypeSupport
}

func (s _FileCloseTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__FileClose())
}

// Modifying this variable is undefined behavior.
var FileCloseTypeSupport types.ServiceTypeSupport = _FileCloseTypeSupport{}

// FileCloseClient wraps rclgo.Client to provide type safe helper
// functions
type FileCloseClient struct {
	*rclgo.Client
}

// NewFileCloseClient creates and returns a new client for the
// FileClose
func NewFileCloseClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*FileCloseClient, error) {
	client, err := node.NewClient(serviceName, FileCloseTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &FileCloseClient{client}, nil
}

func (s *FileCloseClient) Send(ctx context.Context, req *FileClose_Request) (*FileClose_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*FileClose_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type FileCloseServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s FileCloseServiceResponseSender) SendResponse(resp *FileClose_Response) error {
	return s.sender.SendResponse(resp)
}

type FileCloseServiceRequestHandler func(*rclgo.ServiceInfo, *FileClose_Request, FileCloseServiceResponseSender)

// FileCloseService wraps rclgo.Service to provide type safe helper
// functions
type FileCloseService struct {
	*rclgo.Service
}

// NewFileCloseService creates and returns a new service for the
// FileClose
func NewFileCloseService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler FileCloseServiceRequestHandler) (*FileCloseService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*FileClose_Request)
		responseSender := FileCloseServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, FileCloseTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &FileCloseService{service}, nil
}