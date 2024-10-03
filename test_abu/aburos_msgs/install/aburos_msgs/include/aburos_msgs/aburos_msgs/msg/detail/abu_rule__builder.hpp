// generated from rosidl_generator_cpp/resource/idl__builder.hpp.em
// with input from aburos_msgs:msg/AbuRule.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_RULE__BUILDER_HPP_
#define ABUROS_MSGS__MSG__DETAIL__ABU_RULE__BUILDER_HPP_

#include <algorithm>
#include <utility>

#include "aburos_msgs/msg/detail/abu_rule__struct.hpp"
#include "rosidl_runtime_cpp/message_initialization.hpp"


namespace aburos_msgs
{

namespace msg
{

namespace builder
{

class Init_AbuRule_remote_resources
{
public:
  explicit Init_AbuRule_remote_resources(::aburos_msgs::msg::AbuRule & msg)
  : msg_(msg)
  {}
  ::aburos_msgs::msg::AbuRule remote_resources(::aburos_msgs::msg::AbuRule::_remote_resources_type arg)
  {
    msg_.remote_resources = std::move(arg);
    return std::move(msg_);
  }

private:
  ::aburos_msgs::msg::AbuRule msg_;
};

class Init_AbuRule_local_resources
{
public:
  explicit Init_AbuRule_local_resources(::aburos_msgs::msg::AbuRule & msg)
  : msg_(msg)
  {}
  Init_AbuRule_remote_resources local_resources(::aburos_msgs::msg::AbuRule::_local_resources_type arg)
  {
    msg_.local_resources = std::move(arg);
    return Init_AbuRule_remote_resources(msg_);
  }

private:
  ::aburos_msgs::msg::AbuRule msg_;
};

class Init_AbuRule_actions
{
public:
  explicit Init_AbuRule_actions(::aburos_msgs::msg::AbuRule & msg)
  : msg_(msg)
  {}
  Init_AbuRule_local_resources actions(::aburos_msgs::msg::AbuRule::_actions_type arg)
  {
    msg_.actions = std::move(arg);
    return Init_AbuRule_local_resources(msg_);
  }

private:
  ::aburos_msgs::msg::AbuRule msg_;
};

class Init_AbuRule_condition
{
public:
  Init_AbuRule_condition()
  : msg_(::rosidl_runtime_cpp::MessageInitialization::SKIP)
  {}
  Init_AbuRule_actions condition(::aburos_msgs::msg::AbuRule::_condition_type arg)
  {
    msg_.condition = std::move(arg);
    return Init_AbuRule_actions(msg_);
  }

private:
  ::aburos_msgs::msg::AbuRule msg_;
};

}  // namespace builder

}  // namespace msg

template<typename MessageType>
auto build();

template<>
inline
auto build<::aburos_msgs::msg::AbuRule>()
{
  return aburos_msgs::msg::builder::Init_AbuRule_condition();
}

}  // namespace aburos_msgs

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_RULE__BUILDER_HPP_
