// Code generated by rclgo-gen. DO NOT EDIT.

package tf2_msgs_action

/*
#include <rosidl_runtime_c/message_type_support_struct.h>
#include <tf2_msgs/action/lookup_transform.h>
*/
import "C"

import (
	"context"
	"time"
	"unsafe"

	"github.com/tiiuae/rclgo/pkg/rclgo"
	"github.com/tiiuae/rclgo/pkg/rclgo/typemap"
	"github.com/tiiuae/rclgo/pkg/rclgo/types"

	action_msgs_msg "test/msgs/action_msgs/msg"
	action_msgs_srv "test/msgs/action_msgs/srv"
)

func init() {
	typemap.RegisterAction("tf2_msgs/LookupTransform", LookupTransformTypeSupport)
	typemap.RegisterAction("tf2_msgs/action/LookupTransform", LookupTransformTypeSupport)
}

type _LookupTransformTypeSupport struct {}

func (s _LookupTransformTypeSupport) Goal() types.MessageTypeSupport {
	return LookupTransform_GoalTypeSupport
}

func (s _LookupTransformTypeSupport) SendGoal() types.ServiceTypeSupport {
	return LookupTransform_SendGoalTypeSupport
}

func (s _LookupTransformTypeSupport) NewSendGoalResponse(accepted bool, stamp time.Duration) types.Message {
	msg := NewLookupTransform_SendGoal_Response()
	msg.Accepted = accepted
	secs := stamp.Truncate(time.Second)
	msg.Stamp.Sec = int32(secs)
	msg.Stamp.Nanosec = uint32(stamp - secs)
	return msg
}

func (s _LookupTransformTypeSupport) Result() types.MessageTypeSupport {
	return LookupTransform_ResultTypeSupport
}

func (s _LookupTransformTypeSupport) GetResult() types.ServiceTypeSupport {
	return LookupTransform_GetResultTypeSupport
}

func (s _LookupTransformTypeSupport) NewGetResultResponse(status int8, result types.Message) types.Message {
	msg := NewLookupTransform_GetResult_Response()
	msg.Status = status
	if result == nil {
		msg.Result = *NewLookupTransform_Result()
	} else {
		msg.Result = *result.(*LookupTransform_Result)
	}
	return msg
}

func (s _LookupTransformTypeSupport) CancelGoal() types.ServiceTypeSupport {
	return action_msgs_srv.CancelGoalTypeSupport
}

func (s _LookupTransformTypeSupport) Feedback() types.MessageTypeSupport {
	return LookupTransform_FeedbackTypeSupport
}

func (s _LookupTransformTypeSupport) FeedbackMessage() types.MessageTypeSupport {
	return LookupTransform_FeedbackMessageTypeSupport
}

func (s _LookupTransformTypeSupport) NewFeedbackMessage(goalID *types.GoalID, feedback types.Message) types.Message {
	msg := NewLookupTransform_FeedbackMessage()
	msg.GoalID.Uuid = *goalID
	msg.Feedback = *feedback.(*LookupTransform_Feedback)
	return msg
}

func (s _LookupTransformTypeSupport) GoalStatusArray() types.MessageTypeSupport {
	return action_msgs_msg.GoalStatusArrayTypeSupport
}

func (s _LookupTransformTypeSupport) TypeSupport() unsafe.Pointer {
	return unsafe.Pointer(C.rosidl_typesupport_c__get_action_type_support_handle__tf2_msgs__action__LookupTransform())
}

// Modifying this variable is undefined behavior.
var LookupTransformTypeSupport types.ActionTypeSupport = _LookupTransformTypeSupport{}

type LookupTransformFeedbackSender struct {
	sender rclgo.FeedbackSender
}

func (s *LookupTransformFeedbackSender) Send(msg *LookupTransform_Feedback) error {
	return s.sender.Send(msg)
}

type LookupTransformGoalHandle struct{
	*rclgo.GoalHandle

	Description *LookupTransform_Goal
}

func (g *LookupTransformGoalHandle) Accept() (*LookupTransformFeedbackSender, error) {
	s, err := g.GoalHandle.Accept()
	if err != nil {
		return nil, err
	}
	return &LookupTransformFeedbackSender{*s}, nil
}

type LookupTransformAction interface {
	ExecuteGoal(context.Context, *LookupTransformGoalHandle) (*LookupTransform_Result, error)
}

func NewLookupTransformAction(
	executeGoal func(context.Context, *LookupTransformGoalHandle) (*LookupTransform_Result, error),
) LookupTransformAction {
	return _LookupTransformFuncAction(executeGoal)
}

type _LookupTransformFuncAction func(context.Context, *LookupTransformGoalHandle) (*LookupTransform_Result, error)

func (a _LookupTransformFuncAction) ExecuteGoal(
	ctx context.Context, goal *LookupTransformGoalHandle,
) (*LookupTransform_Result, error) {
	return a(ctx, goal)
}

type _LookupTransformAction struct {
	action LookupTransformAction
}

func (a _LookupTransformAction) ExecuteGoal(ctx context.Context, handle *rclgo.GoalHandle) (types.Message, error) {
	return a.action.ExecuteGoal(ctx, &LookupTransformGoalHandle{
		GoalHandle:  handle,
		Description: handle.Description.(*LookupTransform_Goal),
	})
}

func (a _LookupTransformAction) TypeSupport() types.ActionTypeSupport {
	return LookupTransformTypeSupport
}

type LookupTransformServer struct{
	*rclgo.ActionServer
}

func NewLookupTransformServer(node *rclgo.Node, name string, action LookupTransformAction, opts *rclgo.ActionServerOptions) (*LookupTransformServer, error) {
	server, err := node.NewActionServer(name, _LookupTransformAction{action}, opts)
	if err != nil {
		return nil, err
	}
	return &LookupTransformServer{server}, nil
}

type LookupTransformFeedbackHandler func(context.Context, *LookupTransform_FeedbackMessage)

type LookupTransformStatusHandler func(context.Context, *action_msgs_msg.GoalStatus)

type LookupTransformClient struct{
	*rclgo.ActionClient
}

func NewLookupTransformClient(node *rclgo.Node, name string, opts *rclgo.ActionClientOptions) (*LookupTransformClient, error) {
	client, err := node.NewActionClient(name, LookupTransformTypeSupport, opts)
	if err != nil {
		return nil, err
	}
	return &LookupTransformClient{client}, nil
}

func (c *LookupTransformClient) WatchGoal(ctx context.Context, goal *LookupTransform_Goal, onFeedback LookupTransformFeedbackHandler) (*LookupTransform_GetResult_Response, *types.GoalID, error) {
	var resp types.Message
	var goalID *types.GoalID
	var err error
	if onFeedback == nil {
		resp, goalID, err = c.ActionClient.WatchGoal(ctx, goal, nil)
	} else {
		resp, goalID, err = c.ActionClient.WatchGoal(ctx, goal, func(ctx context.Context, msg types.Message) {
			onFeedback(ctx, msg.(*LookupTransform_FeedbackMessage))
		})
	}
	if r, ok := resp.(*LookupTransform_GetResult_Response); ok {
		return r, goalID, err
	}
	return nil, goalID, err
}

func (c *LookupTransformClient) SendGoal(ctx context.Context, goal *LookupTransform_Goal) (*LookupTransform_SendGoal_Response, *types.GoalID, error) {
	resp, id, err := c.ActionClient.SendGoal(ctx, goal)
	if r, ok := resp.(*LookupTransform_SendGoal_Response); ok {
		return r, id, err
	}
	return nil, id, err
}

func (c *LookupTransformClient) SendGoalRequest(ctx context.Context, request *LookupTransform_SendGoal_Request) (*LookupTransform_SendGoal_Response, error) {
	resp, err := c.ActionClient.SendGoalRequest(ctx, request)
	if r, ok := resp.(*LookupTransform_SendGoal_Response); ok {
		return r, err
	}
	return nil, err
}

func (c *LookupTransformClient) GetResult(ctx context.Context, goalID *types.GoalID) (*LookupTransform_GetResult_Response, error) {
	resp, err := c.ActionClient.GetResult(ctx, goalID)
	if r, ok := resp.(*LookupTransform_GetResult_Response); ok {
		return r, err
	}
	return nil, err
}

func (c *LookupTransformClient) CancelGoal(ctx context.Context, request *action_msgs_srv.CancelGoal_Request) (*action_msgs_srv.CancelGoal_Response, error) {
	resp, err := c.ActionClient.CancelGoal(ctx, request)
	if r, ok := resp.(*action_msgs_srv.CancelGoal_Response); ok {
		return r, err
	}
	return nil, err
}

func (c *LookupTransformClient) WatchFeedback(ctx context.Context, goalID *types.GoalID, handler LookupTransformFeedbackHandler) <-chan error {
	return c.ActionClient.WatchFeedback(ctx, goalID, func(ctx context.Context, msg types.Message) {
		handler(ctx, msg.(*LookupTransform_FeedbackMessage))
	})
}

func (c *LookupTransformClient) WatchStatus(ctx context.Context, goalID *types.GoalID, handler LookupTransformStatusHandler) <-chan error {
	return c.ActionClient.WatchStatus(ctx, goalID, func(ctx context.Context, msg types.Message) {
		handler(ctx, msg.(*action_msgs_msg.GoalStatus))
	})
}
