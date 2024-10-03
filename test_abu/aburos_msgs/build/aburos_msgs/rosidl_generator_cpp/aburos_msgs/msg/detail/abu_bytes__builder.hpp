// generated from rosidl_generator_cpp/resource/idl__builder.hpp.em
// with input from aburos_msgs:msg/AbuBytes.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__BUILDER_HPP_
#define ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__BUILDER_HPP_

#include <algorithm>
#include <utility>

#include "aburos_msgs/msg/detail/abu_bytes__struct.hpp"
#include "rosidl_runtime_cpp/message_initialization.hpp"


namespace aburos_msgs
{

namespace msg
{

namespace builder
{

class Init_AbuBytes_data
{
public:
  Init_AbuBytes_data()
  : msg_(::rosidl_runtime_cpp::MessageInitialization::SKIP)
  {}
  ::aburos_msgs::msg::AbuBytes data(::aburos_msgs::msg::AbuBytes::_data_type arg)
  {
    msg_.data = std::move(arg);
    return std::move(msg_);
  }

private:
  ::aburos_msgs::msg::AbuBytes msg_;
};

}  // namespace builder

}  // namespace msg

template<typename MessageType>
auto build();

template<>
inline
auto build<::aburos_msgs::msg::AbuBytes>()
{
  return aburos_msgs::msg::builder::Init_AbuBytes_data();
}

}  // namespace aburos_msgs

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__BUILDER_HPP_
