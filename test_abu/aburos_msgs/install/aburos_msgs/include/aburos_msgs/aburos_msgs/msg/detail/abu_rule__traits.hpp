// generated from rosidl_generator_cpp/resource/idl__traits.hpp.em
// with input from aburos_msgs:msg/AbuRule.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_RULE__TRAITS_HPP_
#define ABUROS_MSGS__MSG__DETAIL__ABU_RULE__TRAITS_HPP_

#include <stdint.h>

#include <sstream>
#include <string>
#include <type_traits>

#include "aburos_msgs/msg/detail/abu_rule__struct.hpp"
#include "rosidl_runtime_cpp/traits.hpp"

namespace aburos_msgs
{

namespace msg
{

inline void to_flow_style_yaml(
  const AbuRule & msg,
  std::ostream & out)
{
  out << "{";
  // member: condition
  {
    out << "condition: ";
    rosidl_generator_traits::value_to_yaml(msg.condition, out);
    out << ", ";
  }

  // member: actions
  {
    if (msg.actions.size() == 0) {
      out << "actions: []";
    } else {
      out << "actions: [";
      size_t pending_items = msg.actions.size();
      for (auto item : msg.actions) {
        rosidl_generator_traits::value_to_yaml(item, out);
        if (--pending_items > 0) {
          out << ", ";
        }
      }
      out << "]";
    }
    out << ", ";
  }

  // member: local_resources
  {
    if (msg.local_resources.size() == 0) {
      out << "local_resources: []";
    } else {
      out << "local_resources: [";
      size_t pending_items = msg.local_resources.size();
      for (auto item : msg.local_resources) {
        rosidl_generator_traits::character_value_to_yaml(item, out);
        if (--pending_items > 0) {
          out << ", ";
        }
      }
      out << "]";
    }
    out << ", ";
  }

  // member: remote_resources
  {
    if (msg.remote_resources.size() == 0) {
      out << "remote_resources: []";
    } else {
      out << "remote_resources: [";
      size_t pending_items = msg.remote_resources.size();
      for (auto item : msg.remote_resources) {
        rosidl_generator_traits::value_to_yaml(item, out);
        if (--pending_items > 0) {
          out << ", ";
        }
      }
      out << "]";
    }
  }
  out << "}";
}  // NOLINT(readability/fn_size)

inline void to_block_style_yaml(
  const AbuRule & msg,
  std::ostream & out, size_t indentation = 0)
{
  // member: condition
  {
    if (indentation > 0) {
      out << std::string(indentation, ' ');
    }
    out << "condition: ";
    rosidl_generator_traits::value_to_yaml(msg.condition, out);
    out << "\n";
  }

  // member: actions
  {
    if (indentation > 0) {
      out << std::string(indentation, ' ');
    }
    if (msg.actions.size() == 0) {
      out << "actions: []\n";
    } else {
      out << "actions:\n";
      for (auto item : msg.actions) {
        if (indentation > 0) {
          out << std::string(indentation, ' ');
        }
        out << "- ";
        rosidl_generator_traits::value_to_yaml(item, out);
        out << "\n";
      }
    }
  }

  // member: local_resources
  {
    if (indentation > 0) {
      out << std::string(indentation, ' ');
    }
    if (msg.local_resources.size() == 0) {
      out << "local_resources: []\n";
    } else {
      out << "local_resources:\n";
      for (auto item : msg.local_resources) {
        if (indentation > 0) {
          out << std::string(indentation, ' ');
        }
        out << "- ";
        rosidl_generator_traits::character_value_to_yaml(item, out);
        out << "\n";
      }
    }
  }

  // member: remote_resources
  {
    if (indentation > 0) {
      out << std::string(indentation, ' ');
    }
    if (msg.remote_resources.size() == 0) {
      out << "remote_resources: []\n";
    } else {
      out << "remote_resources:\n";
      for (auto item : msg.remote_resources) {
        if (indentation > 0) {
          out << std::string(indentation, ' ');
        }
        out << "- ";
        rosidl_generator_traits::value_to_yaml(item, out);
        out << "\n";
      }
    }
  }
}  // NOLINT(readability/fn_size)

inline std::string to_yaml(const AbuRule & msg, bool use_flow_style = false)
{
  std::ostringstream out;
  if (use_flow_style) {
    to_flow_style_yaml(msg, out);
  } else {
    to_block_style_yaml(msg, out);
  }
  return out.str();
}

}  // namespace msg

}  // namespace aburos_msgs

namespace rosidl_generator_traits
{

[[deprecated("use aburos_msgs::msg::to_block_style_yaml() instead")]]
inline void to_yaml(
  const aburos_msgs::msg::AbuRule & msg,
  std::ostream & out, size_t indentation = 0)
{
  aburos_msgs::msg::to_block_style_yaml(msg, out, indentation);
}

[[deprecated("use aburos_msgs::msg::to_yaml() instead")]]
inline std::string to_yaml(const aburos_msgs::msg::AbuRule & msg)
{
  return aburos_msgs::msg::to_yaml(msg);
}

template<>
inline const char * data_type<aburos_msgs::msg::AbuRule>()
{
  return "aburos_msgs::msg::AbuRule";
}

template<>
inline const char * name<aburos_msgs::msg::AbuRule>()
{
  return "aburos_msgs/msg/AbuRule";
}

template<>
struct has_fixed_size<aburos_msgs::msg::AbuRule>
  : std::integral_constant<bool, false> {};

template<>
struct has_bounded_size<aburos_msgs::msg::AbuRule>
  : std::integral_constant<bool, false> {};

template<>
struct is_message<aburos_msgs::msg::AbuRule>
  : std::true_type {};

}  // namespace rosidl_generator_traits

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_RULE__TRAITS_HPP_
