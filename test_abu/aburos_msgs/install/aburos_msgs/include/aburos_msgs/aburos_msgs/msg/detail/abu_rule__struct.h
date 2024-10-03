// generated from rosidl_generator_c/resource/idl__struct.h.em
// with input from aburos_msgs:msg/AbuRule.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_RULE__STRUCT_H_
#define ABUROS_MSGS__MSG__DETAIL__ABU_RULE__STRUCT_H_

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>


// Constants defined in the message

// Include directives for member types
// Member 'condition'
// Member 'actions'
// Member 'remote_resources'
#include "rosidl_runtime_c/string.h"
// Member 'local_resources'
#include "rosidl_runtime_c/primitives_sequence.h"

/// Struct defined in msg/AbuRule in the package aburos_msgs.
typedef struct aburos_msgs__msg__AbuRule
{
  rosidl_runtime_c__String condition;
  rosidl_runtime_c__String__Sequence actions;
  rosidl_runtime_c__octet__Sequence local_resources;
  rosidl_runtime_c__String__Sequence remote_resources;
} aburos_msgs__msg__AbuRule;

// Struct for a sequence of aburos_msgs__msg__AbuRule.
typedef struct aburos_msgs__msg__AbuRule__Sequence
{
  aburos_msgs__msg__AbuRule * data;
  /// The number of valid items in data
  size_t size;
  /// The number of allocated items in data
  size_t capacity;
} aburos_msgs__msg__AbuRule__Sequence;

#ifdef __cplusplus
}
#endif

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_RULE__STRUCT_H_
