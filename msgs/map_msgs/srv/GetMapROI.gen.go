// Code generated by rclgo-gen. DO NOT EDIT.

package map_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <map_msgs/srv/get_map_roi.h>
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
	typemap.RegisterService("map_msgs/GetMapROI", GetMapROITypeSupport)
	typemap.RegisterService("map_msgs/srv/GetMapROI", GetMapROITypeSupport)
}

type _GetMapROITypeSupport struct {}

func (s _GetMapROITypeSupport) Request() types.MessageTypeSupport {
	return GetMapROI_RequestTypeSupport
}

func (s _GetMapROITypeSupport) Response() types.MessageTypeSupport {
	return GetMapROI_ResponseTypeSupport
}

func (s _GetMapROITypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__map_msgs__srv__GetMapROI())
}

// Modifying this variable is undefined behavior.
var GetMapROITypeSupport types.ServiceTypeSupport = _GetMapROITypeSupport{}

// GetMapROIClient wraps rclgo.Client to provide type safe helper
// functions
type GetMapROIClient struct {
	*rclgo.Client
}

// NewGetMapROIClient creates and returns a new client for the
// GetMapROI
func NewGetMapROIClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*GetMapROIClient, error) {
	client, err := node.NewClient(serviceName, GetMapROITypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GetMapROIClient{client}, nil
}

func (s *GetMapROIClient) Send(ctx context.Context, req *GetMapROI_Request) (*GetMapROI_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*GetMapROI_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type GetMapROIServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s GetMapROIServiceResponseSender) SendResponse(resp *GetMapROI_Response) error {
	return s.sender.SendResponse(resp)
}

type GetMapROIServiceRequestHandler func(*rclgo.ServiceInfo, *GetMapROI_Request, GetMapROIServiceResponseSender)

// GetMapROIService wraps rclgo.Service to provide type safe helper
// functions
type GetMapROIService struct {
	*rclgo.Service
}

// NewGetMapROIService creates and returns a new service for the
// GetMapROI
func NewGetMapROIService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler GetMapROIServiceRequestHandler) (*GetMapROIService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*GetMapROI_Request)
		responseSender := GetMapROIServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, GetMapROITypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &GetMapROIService{service}, nil
}