// generated from rosidl_generator_cpp/resource/idl__traits.hpp.em
// with input from aburos_msgs:msg/AbuBytes.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__TRAITS_HPP_
#define ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__TRAITS_HPP_

#include <stdint.h>

#include <sstream>
#include <string>
#include <type_traits>

#include "aburos_msgs/msg/detail/abu_bytes__struct.hpp"
#include "rosidl_runtime_cpp/traits.hpp"

namespace aburos_msgs
{

namespace msg
{

inline void to_flow_style_yaml(
  const AbuBytes & msg,
  std::ostream & out)
{
  out << "{";
  // member: data
  {
    if (msg.data.size() == 0) {
      out << "data: []";
    } else {
      out << "data: [";
      size_t pending_items = msg.data.size();
      for (auto item : msg.data) {
        rosidl_generator_traits::character_value_to_yaml(item, out);
        if (--pending_items > 0) {
          out << ", ";
        }
      }
      out << "]";
    }
    out << ", ";
  }

  // member: origin
  {
    out << "origin: ";
    rosidl_generator_traits::value_to_yaml(msg.origin, out);
  }
  out << "}";
}  // NOLINT(readability/fn_size)

inline void to_block_style_yaml(
  const AbuBytes & msg,
  std::ostream & out, size_t indentation = 0)
{
  // member: data
  {
    if (indentation > 0) {
      out << std::string(indentation, ' ');
    }
    if (msg.data.size() == 0) {
      out << "data: []\n";
    } else {
      out << "data:\n";
      for (auto item : msg.data) {
        if (indentation > 0) {
          out << std::string(indentation, ' ');
        }
        out << "- ";
        rosidl_generator_traits::character_value_to_yaml(item, out);
        out << "\n";
      }
    }
  }

  // member: origin
  {
    if (indentation > 0) {
      out << std::string(indentation, ' ');
    }
    out << "origin: ";
    rosidl_generator_traits::value_to_yaml(msg.origin, out);
    out << "\n";
  }
}  // NOLINT(readability/fn_size)

inline std::string to_yaml(const AbuBytes & msg, bool use_flow_style = false)
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
  const aburos_msgs::msg::AbuBytes & msg,
  std::ostream & out, size_t indentation = 0)
{
  aburos_msgs::msg::to_block_style_yaml(msg, out, indentation);
}

[[deprecated("use aburos_msgs::msg::to_yaml() instead")]]
inline std::string to_yaml(const aburos_msgs::msg::AbuBytes & msg)
{
  return aburos_msgs::msg::to_yaml(msg);
}

template<>
inline const char * data_type<aburos_msgs::msg::AbuBytes>()
{
  return "aburos_msgs::msg::AbuBytes";
}

template<>
inline const char * name<aburos_msgs::msg::AbuBytes>()
{
  return "aburos_msgs/msg/AbuBytes";
}

template<>
struct has_fixed_size<aburos_msgs::msg::AbuBytes>
  : std::integral_constant<bool, false> {};

template<>
struct has_bounded_size<aburos_msgs::msg::AbuBytes>
  : std::integral_constant<bool, false> {};

template<>
struct is_message<aburos_msgs::msg::AbuBytes>
  : std::true_type {};

}  // namespace rosidl_generator_traits

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__TRAITS_HPP_
