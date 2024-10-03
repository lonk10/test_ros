// Code generated by rclgo-gen. DO NOT EDIT.

package mavros_msgs_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <mavros_msgs/srv/stream_rate.h>
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
	typemap.RegisterService("mavros_msgs/StreamRate", StreamRateTypeSupport)
	typemap.RegisterService("mavros_msgs/srv/StreamRate", StreamRateTypeSupport)
}

type _StreamRateTypeSupport struct {}

func (s _StreamRateTypeSupport) Request() types.MessageTypeSupport {
	return StreamRate_RequestTypeSupport
}

func (s _StreamRateTypeSupport) Response() types.MessageTypeSupport {
	return StreamRate_ResponseTypeSupport
}

func (s _StreamRateTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__mavros_msgs__srv__StreamRate())
}

// Modifying this variable is undefined behavior.
var StreamRateTypeSupport types.ServiceTypeSupport = _StreamRateTypeSupport{}

// StreamRateClient wraps rclgo.Client to provide type safe helper
// functions
type StreamRateClient struct {
	*rclgo.Client
}

// NewStreamRateClient creates and returns a new client for the
// StreamRate
func NewStreamRateClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*StreamRateClient, error) {
	client, err := node.NewClient(serviceName, StreamRateTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &StreamRateClient{client}, nil
}

func (s *StreamRateClient) Send(ctx context.Context, req *StreamRate_Request) (*StreamRate_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*StreamRate_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type StreamRateServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s StreamRateServiceResponseSender) SendResponse(resp *StreamRate_Response) error {
	return s.sender.SendResponse(resp)
}

type StreamRateServiceRequestHandler func(*rclgo.ServiceInfo, *StreamRate_Request, StreamRateServiceResponseSender)

// StreamRateService wraps rclgo.Service to provide type safe helper
// functions
type StreamRateService struct {
	*rclgo.Service
}

// NewStreamRateService creates and returns a new service for the
// StreamRate
func NewStreamRateService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler StreamRateServiceRequestHandler) (*StreamRateService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*StreamRate_Request)
		responseSender := StreamRateServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, StreamRateTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &StreamRateService{service}, nil
}