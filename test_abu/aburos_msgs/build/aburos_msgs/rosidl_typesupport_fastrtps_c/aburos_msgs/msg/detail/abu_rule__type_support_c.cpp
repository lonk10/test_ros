// generated from rosidl_typesupport_fastrtps_c/resource/idl__type_support_c.cpp.em
// with input from aburos_msgs:msg/AbuRule.idl
// generated code does not contain a copyright notice
#include "aburos_msgs/msg/detail/abu_rule__rosidl_typesupport_fastrtps_c.h"


#include <cassert>
#include <limits>
#include <string>
#include "rosidl_typesupport_fastrtps_c/identifier.h"
#include "rosidl_typesupport_fastrtps_c/wstring_conversion.hpp"
#include "rosidl_typesupport_fastrtps_cpp/message_type_support.h"
#include "aburos_msgs/msg/rosidl_typesupport_fastrtps_c__visibility_control.h"
#include "aburos_msgs/msg/detail/abu_rule__struct.h"
#include "aburos_msgs/msg/detail/abu_rule__functions.h"
#include "fastcdr/Cdr.h"

#ifndef _WIN32
# pragma GCC diagnostic push
# pragma GCC diagnostic ignored "-Wunused-parameter"
# ifdef __clang__
#  pragma clang diagnostic ignored "-Wdeprecated-register"
#  pragma clang diagnostic ignored "-Wreturn-type-c-linkage"
# endif
#endif
#ifndef _WIN32
# pragma GCC diagnostic pop
#endif

// includes and forward declarations of message dependencies and their conversion functions

#if defined(__cplusplus)
extern "C"
{
#endif

#include "rosidl_runtime_c/primitives_sequence.h"  // local_resources
#include "rosidl_runtime_c/primitives_sequence_functions.h"  // local_resources
#include "rosidl_runtime_c/string.h"  // actions, condition, remote_resources
#include "rosidl_runtime_c/string_functions.h"  // actions, condition, remote_resources

// forward declare type support functions


using _AbuRule__ros_msg_type = aburos_msgs__msg__AbuRule;

static bool _AbuRule__cdr_serialize(
  const void * untyped_ros_message,
  eprosima::fastcdr::Cdr & cdr)
{
  if (!untyped_ros_message) {
    fprintf(stderr, "ros message handle is null\n");
    return false;
  }
  const _AbuRule__ros_msg_type * ros_message = static_cast<const _AbuRule__ros_msg_type *>(untyped_ros_message);
  // Field name: condition
  {
    const rosidl_runtime_c__String * str = &ros_message->condition;
    if (str->capacity == 0 || str->capacity <= str->size) {
      fprintf(stderr, "string capacity not greater than size\n");
      return false;
    }
    if (str->data[str->size] != '\0') {
      fprintf(stderr, "string not null-terminated\n");
      return false;
    }
    cdr << str->data;
  }

  // Field name: actions
  {
    size_t size = ros_message->actions.size;
    auto array_ptr = ros_message->actions.data;
    cdr << static_cast<uint32_t>(size);
    for (size_t i = 0; i < size; ++i) {
      const rosidl_runtime_c__String * str = &array_ptr[i];
      if (str->capacity == 0 || str->capacity <= str->size) {
        fprintf(stderr, "string capacity not greater than size\n");
        return false;
      }
      if (str->data[str->size] != '\0') {
        fprintf(stderr, "string not null-terminated\n");
        return false;
      }
      cdr << str->data;
    }
  }

  // Field name: local_resources
  {
    size_t size = ros_message->local_resources.size;
    auto array_ptr = ros_message->local_resources.data;
    cdr << static_cast<uint32_t>(size);
    cdr.serializeArray(array_ptr, size);
  }

  // Field name: remote_resources
  {
    size_t size = ros_message->remote_resources.size;
    auto array_ptr = ros_message->remote_resources.data;
    cdr << static_cast<uint32_t>(size);
    for (size_t i = 0; i < size; ++i) {
      const rosidl_runtime_c__String * str = &array_ptr[i];
      if (str->capacity == 0 || str->capacity <= str->size) {
        fprintf(stderr, "string capacity not greater than size\n");
        return false;
      }
      if (str->data[str->size] != '\0') {
        fprintf(stderr, "string not null-terminated\n");
        return false;
      }
      cdr << str->data;
    }
  }

  return true;
}

static bool _AbuRule__cdr_deserialize(
  eprosima::fastcdr::Cdr & cdr,
  void * untyped_ros_message)
{
  if (!untyped_ros_message) {
    fprintf(stderr, "ros message handle is null\n");
    return false;
  }
  _AbuRule__ros_msg_type * ros_message = static_cast<_AbuRule__ros_msg_type *>(untyped_ros_message);
  // Field name: condition
  {
    std::string tmp;
    cdr >> tmp;
    if (!ros_message->condition.data) {
      rosidl_runtime_c__String__init(&ros_message->condition);
    }
    bool succeeded = rosidl_runtime_c__String__assign(
      &ros_message->condition,
      tmp.c_str());
    if (!succeeded) {
      fprintf(stderr, "failed to assign string into field 'condition'\n");
      return false;
    }
  }

  // Field name: actions
  {
    uint32_t cdrSize;
    cdr >> cdrSize;
    size_t size = static_cast<size_t>(cdrSize);
    if (ros_message->actions.data) {
      rosidl_runtime_c__String__Sequence__fini(&ros_message->actions);
    }
    if (!rosidl_runtime_c__String__Sequence__init(&ros_message->actions, size)) {
      fprintf(stderr, "failed to create array for field 'actions'");
      return false;
    }
    auto array_ptr = ros_message->actions.data;
    for (size_t i = 0; i < size; ++i) {
      std::string tmp;
      cdr >> tmp;
      auto & ros_i = array_ptr[i];
      if (!ros_i.data) {
        rosidl_runtime_c__String__init(&ros_i);
      }
      bool succeeded = rosidl_runtime_c__String__assign(
        &ros_i,
        tmp.c_str());
      if (!succeeded) {
        fprintf(stderr, "failed to assign string into field 'actions'\n");
        return false;
      }
    }
  }

  // Field name: local_resources
  {
    uint32_t cdrSize;
    cdr >> cdrSize;
    size_t size = static_cast<size_t>(cdrSize);
    if (ros_message->local_resources.data) {
      rosidl_runtime_c__octet__Sequence__fini(&ros_message->local_resources);
    }
    if (!rosidl_runtime_c__octet__Sequence__init(&ros_message->local_resources, size)) {
      fprintf(stderr, "failed to create array for field 'local_resources'");
      return false;
    }
    auto array_ptr = ros_message->local_resources.data;
    cdr.deserializeArray(array_ptr, size);
  }

  // Field name: remote_resources
  {
    uint32_t cdrSize;
    cdr >> cdrSize;
    size_t size = static_cast<size_t>(cdrSize);
    if (ros_message->remote_resources.data) {
      rosidl_runtime_c__String__Sequence__fini(&ros_message->remote_resources);
    }
    if (!rosidl_runtime_c__String__Sequence__init(&ros_message->remote_resources, size)) {
      fprintf(stderr, "failed to create array for field 'remote_resources'");
      return false;
    }
    auto array_ptr = ros_message->remote_resources.data;
    for (size_t i = 0; i < size; ++i) {
      std::string tmp;
      cdr >> tmp;
      auto & ros_i = array_ptr[i];
      if (!ros_i.data) {
        rosidl_runtime_c__String__init(&ros_i);
      }
      bool succeeded = rosidl_runtime_c__String__assign(
        &ros_i,
        tmp.c_str());
      if (!succeeded) {
        fprintf(stderr, "failed to assign string into field 'remote_resources'\n");
        return false;
      }
    }
  }

  return true;
}  // NOLINT(readability/fn_size)

ROSIDL_TYPESUPPORT_FASTRTPS_C_PUBLIC_aburos_msgs
size_t get_serialized_size_aburos_msgs__msg__AbuRule(
  const void * untyped_ros_message,
  size_t current_alignment)
{
  const _AbuRule__ros_msg_type * ros_message = static_cast<const _AbuRule__ros_msg_type *>(untyped_ros_message);
  (void)ros_message;
  size_t initial_alignment = current_alignment;

  const size_t padding = 4;
  const size_t wchar_size = 4;
  (void)padding;
  (void)wchar_size;

  // field.name condition
  current_alignment += padding +
    eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
    (ros_message->condition.size + 1);
  // field.name actions
  {
    size_t array_size = ros_message->actions.size;
    auto array_ptr = ros_message->actions.data;
    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);
    for (size_t index = 0; index < array_size; ++index) {
      current_alignment += padding +
        eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
        (array_ptr[index].size + 1);
    }
  }
  // field.name local_resources
  {
    size_t array_size = ros_message->local_resources.size;
    auto array_ptr = ros_message->local_resources.data;
    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);
    (void)array_ptr;
    size_t item_size = sizeof(array_ptr[0]);
    current_alignment += array_size * item_size +
      eprosima::fastcdr::Cdr::alignment(current_alignment, item_size);
  }
  // field.name remote_resources
  {
    size_t array_size = ros_message->remote_resources.size;
    auto array_ptr = ros_message->remote_resources.data;
    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);
    for (size_t index = 0; index < array_size; ++index) {
      current_alignment += padding +
        eprosima::fastcdr::Cdr::alignment(current_alignment, padding) +
        (array_ptr[index].size + 1);
    }
  }

  return current_alignment - initial_alignment;
}

static uint32_t _AbuRule__get_serialized_size(const void * untyped_ros_message)
{
  return static_cast<uint32_t>(
    get_serialized_size_aburos_msgs__msg__AbuRule(
      untyped_ros_message, 0));
}

ROSIDL_TYPESUPPORT_FASTRTPS_C_PUBLIC_aburos_msgs
size_t max_serialized_size_aburos_msgs__msg__AbuRule(
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

  // member: condition
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
  // member: actions
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
  // member: local_resources
  {
    size_t array_size = 0;
    full_bounded = false;
    is_plain = false;
    current_alignment += padding +
      eprosima::fastcdr::Cdr::alignment(current_alignment, padding);

    last_member_size = array_size * sizeof(uint8_t);
    current_alignment += array_size * sizeof(uint8_t);
  }
  // member: remote_resources
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
    using DataType = aburos_msgs__msg__AbuRule;
    is_plain =
      (
      offsetof(DataType, remote_resources) +
      last_member_size
      ) == ret_val;
  }

  return ret_val;
}

static size_t _AbuRule__max_serialized_size(char & bounds_info)
{
  bool full_bounded;
  bool is_plain;
  size_t ret_val;

  ret_val = max_serialized_size_aburos_msgs__msg__AbuRule(
    full_bounded, is_plain, 0);

  bounds_info =
    is_plain ? ROSIDL_TYPESUPPORT_FASTRTPS_PLAIN_TYPE :
    full_bounded ? ROSIDL_TYPESUPPORT_FASTRTPS_BOUNDED_TYPE : ROSIDL_TYPESUPPORT_FASTRTPS_UNBOUNDED_TYPE;
  return ret_val;
}


static message_type_support_callbacks_t __callbacks_AbuRule = {
  "aburos_msgs::msg",
  "AbuRule",
  _AbuRule__cdr_serialize,
  _AbuRule__cdr_deserialize,
  _AbuRule__get_serialized_size,
  _AbuRule__max_serialized_size
};

static rosidl_message_type_support_t _AbuRule__type_support = {
  rosidl_typesupport_fastrtps_c__identifier,
  &__callbacks_AbuRule,
  get_message_typesupport_handle_function,
};

const rosidl_message_type_support_t *
ROSIDL_TYPESUPPORT_INTERFACE__MESSAGE_SYMBOL_NAME(rosidl_typesupport_fastrtps_c, aburos_msgs, msg, AbuRule)() {
  return &_AbuRule__type_support;
}

#if defined(__cplusplus)
}
#endif
