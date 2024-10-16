// Code generated by rclgo-gen. DO NOT EDIT.

package map_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <map_msgs/srv/get_point_map.h>
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
	typemap.RegisterService("map_msgs/GetPointMap", GetPointMapTypeSupport)
	typemap.RegisterService("map_msgs/srv/GetPointMap", GetPointMapTypeSupport)
}

type _GetPointMapTypeSupport struct {}

func (s _GetPointMapTypeSupport) Request() types.MessageTypeSupport {
	return GetPointMap_RequestTypeSupport
}

func (s _GetPointMapTypeSupport) Response() types.MessageTypeSupport {
	return GetPointMap_ResponseTypeSupport
}

func (s _GetPointMapTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__map_msgs__srv__GetPointMap())
}

// Modifying this variable is undefined behavior.
var GetPointMapTypeSupport types.ServiceTypeSupport = _GetPointMapTypeSupport{}

// GetPointMapClient wraps rclgo.Client to provide type safe helper
// functions
type GetPointMapClient struct {
	*rclgo.Client
}

// NewGetPointMapClient creates and returns a new client for the
// GetPointMap
func NewGetPointMapClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*GetPointMapClient, error) {
	client, err := node.NewClient(serviceName, GetPointMapTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GetPointMapClient{client}, nil
}

func (s *GetPointMapClient) Send(ctx context.Context, req *GetPointMap_Request) (*GetPointMap_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*GetPointMap_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type GetPointMapServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s GetPointMapServiceResponseSender) SendResponse(resp *GetPointMap_Response) error {
	return s.sender.SendResponse(resp)
}

type GetPointMapServiceRequestHandler func(*rclgo.ServiceInfo, *GetPointMap_Request, GetPointMapServiceResponseSender)

// GetPointMapService wraps rclgo.Service to provide type safe helper
// functions
type GetPointMapService struct {
	*rclgo.Service
}

// NewGetPointMapService creates and returns a new service for the
// GetPointMap
func NewGetPointMapService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler GetPointMapServiceRequestHandler) (*GetPointMapService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*GetPointMap_Request)
		responseSender := GetPointMapServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, GetPointMapTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &GetPointMapService{service}, nil
}