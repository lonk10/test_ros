// generated from rosidl_typesupport_introspection_c/resource/idl__type_support.c.em
// with input from aburos_msgs:msg/AbuBytes.idl
// generated code does not contain a copyright notice

#include <stddef.h>
#include "aburos_msgs/msg/detail/abu_bytes__rosidl_typesupport_introspection_c.h"
#include "aburos_msgs/msg/rosidl_typesupport_introspection_c__visibility_control.h"
#include "rosidl_typesupport_introspection_c/field_types.h"
#include "rosidl_typesupport_introspection_c/identifier.h"
#include "rosidl_typesupport_introspection_c/message_introspection.h"
#include "aburos_msgs/msg/detail/abu_bytes__functions.h"
#include "aburos_msgs/msg/detail/abu_bytes__struct.h"


// Include directives for member types
// Member `data`
#include "rosidl_runtime_c/primitives_sequence_functions.h"

#ifdef __cplusplus
extern "C"
{
#endif

void aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_init_function(
  void * message_memory, enum rosidl_runtime_c__message_initialization _init)
{
  // TODO(karsten1987): initializers are not yet implemented for typesupport c
  // see https://github.com/ros2/ros2/issues/397
  (void) _init;
  aburos_msgs__msg__AbuBytes__init(message_memory);
}

void aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_fini_function(void * message_memory)
{
  aburos_msgs__msg__AbuBytes__fini(message_memory);
}

size_t aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__size_function__AbuBytes__data(
  const void * untyped_member)
{
  const rosidl_runtime_c__octet__Sequence * member =
    (const rosidl_runtime_c__octet__Sequence *)(untyped_member);
  return member->size;
}

const void * aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__get_const_function__AbuBytes__data(
  const void * untyped_member, size_t index)
{
  const rosidl_runtime_c__octet__Sequence * member =
    (const rosidl_runtime_c__octet__Sequence *)(untyped_member);
  return &member->data[index];
}

void * aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__get_function__AbuBytes__data(
  void * untyped_member, size_t index)
{
  rosidl_runtime_c__octet__Sequence * member =
    (rosidl_runtime_c__octet__Sequence *)(untyped_member);
  return &member->data[index];
}

void aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__fetch_function__AbuBytes__data(
  const void * untyped_member, size_t index, void * untyped_value)
{
  const uint8_t * item =
    ((const uint8_t *)
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__get_const_function__AbuBytes__data(untyped_member, index));
  uint8_t * value =
    (uint8_t *)(untyped_value);
  *value = *item;
}

void aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__assign_function__AbuBytes__data(
  void * untyped_member, size_t index, const void * untyped_value)
{
  uint8_t * item =
    ((uint8_t *)
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__get_function__AbuBytes__data(untyped_member, index));
  const uint8_t * value =
    (const uint8_t *)(untyped_value);
  *item = *value;
}

bool aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__resize_function__AbuBytes__data(
  void * untyped_member, size_t size)
{
  rosidl_runtime_c__octet__Sequence * member =
    (rosidl_runtime_c__octet__Sequence *)(untyped_member);
  rosidl_runtime_c__octet__Sequence__fini(member);
  return rosidl_runtime_c__octet__Sequence__init(member, size);
}

static rosidl_typesupport_introspection_c__MessageMember aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_message_member_array[1] = {
  {
    "data",  // name
    rosidl_typesupport_introspection_c__ROS_TYPE_OCTET,  // type
    0,  // upper bound of string
    NULL,  // members of sub message
    true,  // is array
    0,  // array size
    false,  // is upper bound
    offsetof(aburos_msgs__msg__AbuBytes, data),  // bytes offset in struct
    NULL,  // default value
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__size_function__AbuBytes__data,  // size() function pointer
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__get_const_function__AbuBytes__data,  // get_const(index) function pointer
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__get_function__AbuBytes__data,  // get(index) function pointer
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__fetch_function__AbuBytes__data,  // fetch(index, &value) function pointer
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__assign_function__AbuBytes__data,  // assign(index, value) function pointer
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__resize_function__AbuBytes__data  // resize(index) function pointer
  }
};

static const rosidl_typesupport_introspection_c__MessageMembers aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_message_members = {
  "aburos_msgs__msg",  // message namespace
  "AbuBytes",  // message name
  1,  // number of fields
  sizeof(aburos_msgs__msg__AbuBytes),
  aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_message_member_array,  // message members
  aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_init_function,  // function to initialize message memory (memory has to be allocated)
  aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_fini_function  // function to terminate message instance (will not free memory)
};

// this is not const since it must be initialized on first access
// since C does not allow non-integral compile-time constants
static rosidl_message_type_support_t aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_message_type_support_handle = {
  0,
  &aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_message_members,
  get_message_typesupport_handle_function,
};

ROSIDL_TYPESUPPORT_INTROSPECTION_C_EXPORT_aburos_msgs
const rosidl_message_type_support_t *
ROSIDL_TYPESUPPORT_INTERFACE__MESSAGE_SYMBOL_NAME(rosidl_typesupport_introspection_c, aburos_msgs, msg, AbuBytes)() {
  if (!aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_message_type_support_handle.typesupport_identifier) {
    aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_message_type_support_handle.typesupport_identifier =
      rosidl_typesupport_introspection_c__identifier;
  }
  return &aburos_msgs__msg__AbuBytes__rosidl_typesupport_introspection_c__AbuBytes_message_type_support_handle;
}
#ifdef __cplusplus
}
#endif
