// Code generated by rclgo-gen. DO NOT EDIT.

package geographic_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <geographic_msgs/srv/get_route_plan.h>
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
	typemap.RegisterService("geographic_msgs/GetRoutePlan", GetRoutePlanTypeSupport)
	typemap.RegisterService("geographic_msgs/srv/GetRoutePlan", GetRoutePlanTypeSupport)
}

type _GetRoutePlanTypeSupport struct {}

func (s _GetRoutePlanTypeSupport) Request() types.MessageTypeSupport {
	return GetRoutePlan_RequestTypeSupport
}

func (s _GetRoutePlanTypeSupport) Response() types.MessageTypeSupport {
	return GetRoutePlan_ResponseTypeSupport
}

func (s _GetRoutePlanTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__geographic_msgs__srv__GetRoutePlan())
}

// Modifying this variable is undefined behavior.
var GetRoutePlanTypeSupport types.ServiceTypeSupport = _GetRoutePlanTypeSupport{}

// GetRoutePlanClient wraps rclgo.Client to provide type safe helper
// functions
type GetRoutePlanClient struct {
	*rclgo.Client
}

// NewGetRoutePlanClient creates and returns a new client for the
// GetRoutePlan
func NewGetRoutePlanClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*GetRoutePlanClient, error) {
	client, err := node.NewClient(serviceName, GetRoutePlanTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &GetRoutePlanClient{client}, nil
}

func (s *GetRoutePlanClient) Send(ctx context.Context, req *GetRoutePlan_Request) (*GetRoutePlan_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*GetRoutePlan_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type GetRoutePlanServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s GetRoutePlanServiceResponseSender) SendResponse(resp *GetRoutePlan_Response) error {
	return s.sender.SendResponse(resp)
}

type GetRoutePlanServiceRequestHandler func(*rclgo.ServiceInfo, *GetRoutePlan_Request, GetRoutePlanServiceResponseSender)

// GetRoutePlanService wraps rclgo.Service to provide type safe helper
// functions
type GetRoutePlanService struct {
	*rclgo.Service
}

// NewGetRoutePlanService creates and returns a new service for the
// GetRoutePlan
func NewGetRoutePlanService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler GetRoutePlanServiceRequestHandler) (*GetRoutePlanService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*GetRoutePlan_Request)
		responseSender := GetRoutePlanServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, GetRoutePlanTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &GetRoutePlanService{service}, nil
}