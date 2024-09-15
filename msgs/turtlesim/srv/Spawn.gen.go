// Code generated by rclgo-gen. DO NOT EDIT.

package turtlesim_srv

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <turtlesim/srv/spawn.h>
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
	typemap.RegisterService("turtlesim/Spawn", SpawnTypeSupport)
	typemap.RegisterService("turtlesim/srv/Spawn", SpawnTypeSupport)
}

type _SpawnTypeSupport struct {}

func (s _SpawnTypeSupport) Request() types.MessageTypeSupport {
	return Spawn_RequestTypeSupport
}

func (s _SpawnTypeSupport) Response() types.MessageTypeSupport {
	return Spawn_ResponseTypeSupport
}

func (s _SpawnTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_service_type_support_handle__turtlesim__srv__Spawn())
}

// Modifying this variable is undefined behavior.
var SpawnTypeSupport types.ServiceTypeSupport = _SpawnTypeSupport{}

// SpawnClient wraps rclgo.Client to provide type safe helper
// functions
type SpawnClient struct {
	*rclgo.Client
}

// NewSpawnClient creates and returns a new client for the
// Spawn
func NewSpawnClient(node *rclgo.Node, serviceName string, options *rclgo.ClientOptions) (*SpawnClient, error) {
	client, err := node.NewClient(serviceName, SpawnTypeSupport, options)
	if err != nil {
		return nil, err
	}
	return &SpawnClient{client}, nil
}

func (s *SpawnClient) Send(ctx context.Context, req *Spawn_Request) (*Spawn_Response, *rclgo.ServiceInfo, error) {
	msg, rmw, err := s.Client.Send(ctx, req)
	if err != nil {
		return nil, rmw, err
	}
	typedMessage, ok := msg.(*Spawn_Response)
	if !ok {
		return nil, rmw, errors.New("invalid message type returned")
	}
	return typedMessage, rmw, err
}

type SpawnServiceResponseSender struct {
	sender rclgo.ServiceResponseSender
}

func (s SpawnServiceResponseSender) SendResponse(resp *Spawn_Response) error {
	return s.sender.SendResponse(resp)
}

type SpawnServiceRequestHandler func(*rclgo.ServiceInfo, *Spawn_Request, SpawnServiceResponseSender)

// SpawnService wraps rclgo.Service to provide type safe helper
// functions
type SpawnService struct {
	*rclgo.Service
}

// NewSpawnService creates and returns a new service for the
// Spawn
func NewSpawnService(node *rclgo.Node, name string, options *rclgo.ServiceOptions, handler SpawnServiceRequestHandler) (*SpawnService, error) {
	h := func(rmw *rclgo.ServiceInfo, msg types.Message, rs rclgo.ServiceResponseSender) {
		m := msg.(*Spawn_Request)
		responseSender := SpawnServiceResponseSender{sender: rs} 
		handler(rmw, m, responseSender)
	}
	service, err := node.NewService(name, SpawnTypeSupport, options, h)
	if err != nil {
		return nil, err
	}
	return &SpawnService{service}, nil
}