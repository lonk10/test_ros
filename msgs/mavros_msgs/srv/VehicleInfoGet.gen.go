// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/vehicle_info_get.h>
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
	typemap.RegisterService("mavros_msgs/VehicleInfoGet", VehicleInfoGetTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/VehicleInfoGet", VehicleInfoGetTypeSupport)
}

type _VehicleInfoGetTypeSupport struct {}

func (s _VehicleInfoGetTypeSupport) Request() types.MessageTypeSupport {
	return VehicleInfoGet_RequestTypeSupport
}

func (s _VehicleInfoGetTypeSupport) Response() types.MessageTypeSupport {
	return VehicleInfoGet_ResponseTypeSupport
}

func (s _VehicleInfoGetTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__VehicleInfoGet())
}

// Modifying this variable is undefined behavior.
var VehicleInfoGetTypeSupport types.ServiceTypeSupport = _VehicleInfoGetTypeSupport{}

// VehicleInfoGetClient wraps rclgo.Client to provide type safe helper
// functions
type VehicleInfoGetClient struct {
	*rclgo.Client
}

// NewVehicleInfoGetClient creates and returns a new client for the
// VehicleInfoGet
func NewVehicleInfoGetClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*VehicleInfoGetClient, error) {
	client, err := node.NewClient(serviceName, VehicleInfoGetTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &VehicleInfoGetClient{client}, nil
}

func (s *VehicleInfoGetClient) Send(ctx context.Context, req *VehicleInfoGet_Request) (*VehicleInfoGet_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*VehicleInfoGet_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type VehicleInfoGetServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s VehicleInfoGetServiceResponseSender) SendResponse(resp *VehicleInfoGet_Response) error {
	return s.sender.SendResponse(resp)
}

type VehicleInfoGetServiceRequestHandler func(*rclgo.ServiceInfo, *VehicleInfoGet_Request, VehicleInfoGetServiceResponseSender)

// VehicleInfoGetService wraps rclgo.Service to provide type safe helper
// functions
type VehicleInfoGetService struct {
	*rclgo.Service
}

// NewVehicleInfoGetService creates and returns a new service for the
// VehicleInfoGet
func NewVehicleInfoGetService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler VehicleInfoGetServiceRequestHandler) (*VehicleInfoGetService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*VehicleInfoGet_Request)
		responseSender := VehicleInfoGetServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, VehicleInfoGetTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &VehicleInfoGetService{service}, nil
}