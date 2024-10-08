// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/file_truncate.h>
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
	typemap.RegisterService("mavros_msgs/FileTruncate", FileTruncateTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/FileTruncate", FileTruncateTypeSupport)
}

type _FileTruncateTypeSupport struct {}

func (s _FileTruncateTypeSupport) Request() types.MessageTypeSupport {
	return FileTruncate_RequestTypeSupport
}

func (s _FileTruncateTypeSupport) Response() types.MessageTypeSupport {
	return FileTruncate_ResponseTypeSupport
}

func (s _FileTruncateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__FileTruncate())
}

// Modifying this variable is undefined behavior.
var FileTruncateTypeSupport types.ServiceTypeSupport = _FileTruncateTypeSupport{}

// FileTruncateClient wraps rclgo.Client to provide type safe helper
// functions
type FileTruncateClient struct {
	*rclgo.Client
}

// NewFileTruncateClient creates and returns a new client for the
// FileTruncate
func NewFileTruncateClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*FileTruncateClient, error) {
	client, err := node.NewClient(serviceName, FileTruncateTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &FileTruncateClient{client}, nil
}

func (s *FileTruncateClient) Send(ctx context.Context, req *FileTruncate_Request) (*FileTruncate_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*FileTruncate_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type FileTruncateServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s FileTruncateServiceResponseSender) SendResponse(resp *FileTruncate_Response) error {
	return s.sender.SendResponse(resp)
}

type FileTruncateServiceRequestHandler func(*rclgo.ServiceInfo, *FileTruncate_Request, FileTruncateServiceResponseSender)

// FileTruncateService wraps rclgo.Service to provide type safe helper
// functions
type FileTruncateService struct {
	*rclgo.Service
}

// NewFileTruncateService creates and returns a new service for the
// FileTruncate
func NewFileTruncateService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler FileTruncateServiceRequestHandler) (*FileTruncateService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*FileTruncate_Request)
		responseSender := FileTruncateServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, FileTruncateTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &FileTruncateService{service}, nil
}