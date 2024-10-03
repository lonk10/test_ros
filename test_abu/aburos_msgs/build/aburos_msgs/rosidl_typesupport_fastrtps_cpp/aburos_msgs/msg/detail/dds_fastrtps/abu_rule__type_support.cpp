// generated from rosidl_typesupport_fastrtps_cpp/resource/idl__type_support.cpp.em
// with input from aburos_msgs:msg/AbuRule.idl
// generated code does not contain a copyright notice
#include "aburos_msgs/msg/detail/abu_rule__rosidl_typesupport_fastrtps_cpp.hpp"
#include "aburos_msgs/msg/detail/abu_rule__struct.hpp"

#include <limits>
#include <stdexcept>
#include <string>
#include "rosidl_typesupport_cpp/message_type_support.hpp"
#include "rosidl_typesupport_fastrtps_cpp/identifier.hpp"
#include "rosidl_typesupport_fastrtps_cpp/message_type_support.h"
#include "rosidl_typesupport_fastrtps_cpp/message_type_support_decl.hpp"
#include "rosidl_typesupport_fastrtps_cpp/wstring_conversion.hpp"
#include "fastcdr/Cdr.h"


// forward declaration of message dependencies and their conversion functions

namespace aburos_msgs
{

namespace msg
{

namespace typesupport_fastrtps_cpp
{

bool
ROSIDL_TYPESUPPORT_FASTRTPS_CPP_PUBLIC_aburos_msgs
cdr_serialize(
  const aburos_msgs::msg::AbuRule & ros_message,
  eprosima::fastcdr::Cdr & cdr)
{
  // Member: condition
  cdr << ros_message.condition;
  // Member: actions
  {
    cdr << ros_message.actions;
  }
  // Member: local_resources
  {
    cdr << ros_message.local_resources;
  }
  // Member: remote_resources
  {
    cdr << ros_message.remote_resources;
  }
  return true;
}

bool
ROSIDL_TYPESUPPORT_FASTRTPS_CPP_PUBLIC_aburos_msgs
cdr_deserialize(
  eprosima::fastcdr::Cdr & cdr,
  aburos_msgs::msg::AbuRule & ros_message)
{
  // Member: condition
  cdr >> ros_message.condition;

  // Member: actions
  {
    cdr >> ros_message.actions;
  }

  // Member: local_resources
  {
    cdr >> ros_message.local_resources;
  }

  // Member: remote_resources
  {
    cdr >> ros_message.remote_resources;
  }

  return true;
}

size_t
ROSIDL_TYPESUPPORT_FASTRTPS_CPP_PUBLIC_aburos_msgs
get_serialized_size(
  const aburos_msgs::msg::AbuRule & ros_message,
  size_t current_alignment)
{
  size_t initial_alignment = current_alignment;

  const size_t padding = 4;
  const size_t wchar_size = 4;
  (void)padding;
  (void)wchar_size;

  // Member: condition
  current_alignment += padding +
    eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
    (ros_message.condition.size() + 1);
  // Member: actions
  {
    size_t array_size = ros_message.actions.size();

    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);
    for (size_t index = 0; index < array_size; ++index) {
      current_alignment += padding +
        eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
        (ros_message.actions[index].size() + 1);
    }
  }
  // Member: local_resources
  {
    size_t array_size = ros_message.local_resources.size();

    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);
    size_t item_size = sizeof(ros_message.local_resources[0]);
    current_alignment += array_size * item_size +
      eprosima::fastcdr::Cdr::alignment(current_alignment, item_size);
  }
  // Member: remote_resources
  {
    size_t array_size = ros_message.remote_resources.size();

    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);
    for (size_t index = 0; index < array_size; ++index) {
      current_alignment += padding +
        eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
        (ros_message.remote_resources[index].size() + 1);
    }
  }

  return current_alignment - initial_alignment;
}

size_t
ROSIDL_TYPESUPPORT_FASTRTPS_CPP_PUBLIC_aburos_msgs
max_serialized_size_AbuRule(
  bool & full_bounded,
  bool & is_plain,
  size_t current_alignment)
{
  size_t initial_alignment = current_alignment;

  const size_t padding = 4;
  const size_t wchar_size = 4;
  size_t last_member_size = 0;
  (void)last_member_size;
  (void)padding;
  (void)wchar_size;

  full_bounded = true;
  is_plain = true;


  // Member: condition
  {
    size_t array_size = 1;

    full_bounded = false;
    is_plain = false;
    for (size_t index = 0; index < array_size; ++index) {
      current_alignment += padding +
        eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
        1;
    }
  }

  // Member: actions
  {
    size_t array_size = 0;
    full_bounded = false;
    is_plain = false;
    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);

    full_bounded = false;
    is_plain = false;
    for (size_t index = 0; index < array_size; ++index) {
      current_alignment += padding +
        eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
        1;
    }
  }

  // Member: local_resources
  {
    size_t array_size = 0;
    full_bounded = false;
    is_plain = false;
    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);

    last_member_size = array_size * sizeof(uint8_t);
    current_alignment += array_size * sizeof(uint8_t);
  }

  // Member: remote_resources
  {
    size_t array_size = 0;
    full_bounded = false;
    is_plain = false;
    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);

    full_bounded = false;
    is_plain = false;
    for (size_t index = 0; index < array_size; ++index) {
      current_alignment += padding +
        eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
        1;
    }
  }

  size_t ret_val = current_alignment - initial_alignment;
  if (is_plain) {
    // All members are plain, and type is not empty.
    // We still need to check that the in-memory alignment
    // is the same as the CDR mandated alignment.
    using DataType = aburos_msgs::msg::AbuRule;
    is_plain =
      (
      offsetof(DataType, remote_resources) +
      last_member_size
      ) == ret_val;
  }

  return ret_val;
}

static bool _AbuRule__cdr_serialize(
  const void * untyped_ros_message,
  eprosima::fastcdr::Cdr & cdr)
{
  auto typed_message =
    static_cast<const aburos_msgs::msg::AbuRule *>(
    untyped_ros_message);
  return cdr_serialize(*typed_message, cdr);
}

static bool _AbuRule__cdr_deserialize(
  eprosima::fastcdr::Cdr & cdr,
  void * untyped_ros_message)
{
  auto typed_message =
    static_cast<aburos_msgs::msg::AbuRule *>(
    untyped_ros_message);
  return cdr_deserialize(cdr, *typed_message);
}

static uint32_t _AbuRule__get_serialized_size(
  const void * untyped_ros_message)
{
  auto typed_message =
    static_cast<const aburos_msgs::msg::AbuRule *>(
    untyped_ros_message);
  return static_cast<uint32_t>(get_serialized_size(*typed_message, 0));
}

static size_t _AbuRule__max_serialized_size(char & bounds_info)
{
  bool full_bounded;
  bool is_plain;
  size_t ret_val;

  ret_val = max_serialized_size_AbuRule(full_bounded, is_plain, 0);

  bounds_info =
    is_plain ? ROSIDL_TYPESUPPORT_FASTRTPS_PLAIN_TYPE :
    full_bounded ? ROSIDL_TYPESUPPORT_FASTRTPS_BOUNDED_TYPE : ROSIDL_TYPESUPPORT_FASTRTPS_UNBOUNDED_TYPE;
  return ret_val;
}

static message_type_support_callbacks_t _AbuRule__callbacks = {
  "aburos_msgs::msg",
  "AbuRule",
  _AbuRule__cdr_serialize,
  _AbuRule__cdr_deserialize,
  _AbuRule__get_serialized_size,
  _AbuRule__max_serialized_size
};

static rosidl_message_type_support_t _AbuRule__handle = {
  rosidl_typesupport_fastrtps_cpp::typesupport_identifier,
  &_AbuRule__callbacks,
  get_message_typesupport_handle_function,
};

}  // namespace typesupport_fastrtps_cpp

}  // namespace msg

}  // namespace aburos_msgs

namespace rosidl_typesupport_fastrtps_cpp
{

template<>
ROSIDL_TYPESUPPORT_FASTRTPS_CPP_EXPORT_aburos_msgs
const rosidl_message_type_support_t *
get_message_type_support_handle<aburos_msgs::msg::AbuRule>()
{
  return &aburos_msgs::msg::typesupport_fastrtps_cpp::_AbuRule__handle;
}

}  // namespace rosidl_typesupport_fastrtps_cpp

#ifdef __cplusplus
extern "C"
{
#endif

const rosidl_message_type_support_t *
ROSIDL_TYPESUPPORT_INTERFACE__MESSAGE_SYMBOL_NAME(rosidl_typesupport_fastrtps_cpp, aburos_msgs, msg, AbuRule)() {
  return &aburos_msgs::msg::typesupport_fastrtps_cpp::_AbuRule__handle;
}

#ifdef __cplusplus
}
#endif
