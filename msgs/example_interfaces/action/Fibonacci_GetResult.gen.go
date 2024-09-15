// Code generated by rclgo-gen. DO NOT EDIT.

package example_interfaces_action

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <example_interfaces/action/fibonacci.h>
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
	typemap.RegisterService("example_interfaces/Fibonacci_GetResult", Fibonacci_GetResultTypeSupport)
	typemap.RegisterService("example_interfaces/action/Fibonacci_GetResult", Fibonacci_GetResultTypeSupport)
}

type _Fibonacci_GetResultTypeSupport struct {}

func (s _Fibonacci_GetResultTypeSupport) Request() types.MessageTypeSupport {
	return Fibonacci_GetResult_RequestTypeSupport
}

func (s _Fibonacci_GetResultTypeSupport) Response() types.MessageTypeSupport {
	return Fibonacci_GetResult_ResponseTypeSupport
}

func (s _Fibonacci_GetResultTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__example_interfaces__action__Fibonacci_GetResult())
}

// Modifying this variable is undefined behavior.
var Fibonacci_GetResultTypeSupport types.ServiceTypeSupport = _Fibonacci_GetResultTypeSupport{}

// Fibonacci_GetResultClient wraps rclgo.Client to provide type safe helper
// functions
type Fibonacci_GetResultClient struct {
	*rclgo.Client
}

// NewFibonacci_GetResultClient creates and returns a new client for the
// Fibonacci_GetResult
func NewFibonacci_GetResultClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*Fibonacci_GetResultClient, error) {
	client, err := node.NewClient(serviceName, Fibonacci_GetResultTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_GetResultClient{client}, nil
}

func (s *Fibonacci_GetResultClient) Send(ctx context.Context, req *Fibonacci_GetResult_Request) (*Fibonacci_GetResult_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*Fibonacci_GetResult_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type Fibonacci_GetResultServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s Fibonacci_GetResultServiceResponseSender) SendResponse(resp *Fibonacci_GetResult_Response) error {
	return s.sender.SendResponse(resp)
}

type Fibonacci_GetResultServiceRequestHandler func(*rclgo.ServiceInfo, *Fibonacci_GetResult_Request, Fibonacci_GetResultServiceResponseSender)

// Fibonacci_GetResultService wraps rclgo.Service to provide type safe helper
// functions
type Fibonacci_GetResultService struct {
	*rclgo.Service
}

// NewFibonacci_GetResultService creates and returns a new service for the
// Fibonacci_GetResult
func NewFibonacci_GetResultService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler Fibonacci_GetResultServiceRequestHandler) (*Fibonacci_GetResultService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*Fibonacci_GetResult_Request)
		responseSender := Fibonacci_GetResultServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, Fibonacci_GetResultTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &Fibonacci_GetResultService{service}, nil
}