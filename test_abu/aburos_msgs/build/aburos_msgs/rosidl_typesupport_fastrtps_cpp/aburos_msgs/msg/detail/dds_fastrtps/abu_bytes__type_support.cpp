// generated from rosidl_typesupport_fastrtps_cpp/resource/idl__type_support.cpp.em
// with input from aburos_msgs:msg/AbuBytes.idl
// generated code does not contain a copyright notice
#include "aburos_msgs/msg/detail/abu_bytes__rosidl_typesupport_fastrtps_cpp.hpp"
#include "aburos_msgs/msg/detail/abu_bytes__struct.hpp"

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
  const aburos_msgs::msg::AbuBytes & ros_message,
  eprosima::fastcdr::Cdr & cdr)
{
  // Member: data
  {
    cdr << ros_message.data;
  }
  // Member: origin
  cdr << ros_message.origin;
  return true;
}

bool
ROSIDL_TYPESUPPORT_FASTRTPS_CPP_PUBLIC_aburos_msgs
cdr_deserialize(
  eprosima::fastcdr::Cdr & cdr,
  aburos_msgs::msg::AbuBytes & ros_message)
{
  // Member: data
  {
    cdr >> ros_message.data;
  }

  // Member: origin
  cdr >> ros_message.origin;

  return true;
}

size_t
ROSIDL_TYPESUPPORT_FASTRTPS_CPP_PUBLIC_aburos_msgs
get_serialized_size(
  const aburos_msgs::msg::AbuBytes & ros_message,
  size_t current_alignment)
{
  size_t initial_alignment = current_alignment;

  const size_t padding = 4;
  const size_t wchar_size = 4;
  (void)padding;
  (void)wchar_size;

  // Member: data
  {
    size_t array_size = ros_message.data.size();

    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);
    size_t item_size = sizeof(ros_message.data[0]);
    current_alignment += array_size * item_size +
      eprosima::fastcdr::Cdr::alignment(current_alignment, item_size);
  }
  // Member: origin
  current_alignment += padding +
    eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
    (ros_message.origin.size() + 1);

  return current_alignment - initial_alignment;
}

size_t
ROSIDL_TYPESUPPORT_FASTRTPS_CPP_PUBLIC_aburos_msgs
max_serialized_size_AbuBytes(
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


  // Member: data
  {
    size_t array_size = 0;
    full_bounded = false;
    is_plain = false;
    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);

    last_member_size = array_size * sizeof(uint8_t);
    current_alignment += array_size * sizeof(uint8_t);
  }

  // Member: origin
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

  size_t ret_val = current_alignment - initial_alignment;
  if (is_plain) {
    // All members are plain, and type is not empty.
    // We still need to check that the in-memory alignment
    // is the same as the CDR mandated alignment.
    using DataType = aburos_msgs::msg::AbuBytes;
    is_plain =
      (
      offsetof(DataType, origin) +
      last_member_size
      ) == ret_val;
  }

  return ret_val;
}

static bool _AbuBytes__cdr_serialize(
  const void * untyped_ros_message,
  eprosima::fastcdr::Cdr & cdr)
{
  auto typed_message =
    static_cast<const aburos_msgs::msg::AbuBytes *>(
    untyped_ros_message);
  return cdr_serialize(*typed_message, cdr);
}

static bool _AbuBytes__cdr_deserialize(
  eprosima::fastcdr::Cdr & cdr,
  void * untyped_ros_message)
{
  auto typed_message =
    static_cast<aburos_msgs::msg::AbuBytes *>(
    untyped_ros_message);
  return cdr_deserialize(cdr, *typed_message);
}

static uint32_t _AbuBytes__get_serialized_size(
  const void * untyped_ros_message)
{
  auto typed_message =
    static_cast<const aburos_msgs::msg::AbuBytes *>(
    untyped_ros_message);
  return static_cast<uint32_t>(get_serialized_size(*typed_message, 0));
}

static size_t _AbuBytes__max_serialized_size(char & bounds_info)
{
  bool full_bounded;
  bool is_plain;
  size_t ret_val;

  ret_val = max_serialized_size_AbuBytes(full_bounded, is_plain, 0);

  bounds_info =
    is_plain ? ROSIDL_TYPESUPPORT_FASTRTPS_PLAIN_TYPE :
    full_bounded ? ROSIDL_TYPESUPPORT_FASTRTPS_BOUNDED_TYPE : ROSIDL_TYPESUPPORT_FASTRTPS_UNBOUNDED_TYPE;
  return ret_val;
}

static message_type_support_callbacks_t _AbuBytes__callbacks = {
  "aburos_msgs::msg",
  "AbuBytes",
  _AbuBytes__cdr_serialize,
  _AbuBytes__cdr_deserialize,
  _AbuBytes__get_serialized_size,
  _AbuBytes__max_serialized_size
};

static rosidl_message_type_support_t _AbuBytes__handle = {
  rosidl_typesupport_fastrtps_cpp::typesupport_identifier,
  &_AbuBytes__callbacks,
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
get_message_type_support_handle<aburos_msgs::msg::AbuBytes>()
{
  return &aburos_msgs::msg::typesupport_fastrtps_cpp::_AbuBytes__handle;
}

}  // namespace rosidl_typesupport_fastrtps_cpp

#ifdef __cplusplus
extern "C"
{
#endif

const rosidl_message_type_support_t *
ROSIDL_TYPESUPPORT_INTERFACE__MESSAGE_SYMBOL_NAME(rosidl_typesupport_fastrtps_cpp, aburos_msgs, msg, AbuBytes)() {
  return &aburos_msgs::msg::typesupport_fastrtps_cpp::_AbuBytes__handle;
}

#ifdef __cplusplus
}
#endif
