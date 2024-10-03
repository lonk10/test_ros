// generated from rosidl_generator_c/resource/idl__struct.h.em
// with input from aburos_msgs:msg/AbuRos.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_ROS__STRUCT_H_
#define ABUROS_MSGS__MSG__DETAIL__ABU_ROS__STRUCT_H_

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdbool.h>
#include <stddef.h>
#include <stdint.h>


// Constants defined in the message

// Include directives for member types
// Member 'data'
#include "rosidl_runtime_c/primitives_sequence.h"

/// Struct defined in msg/AbuRos in the package aburos_msgs.
typedef struct aburos_msgs__msg__AbuRos
{
  rosidl_runtime_c__octet__Sequence data;
} aburos_msgs__msg__AbuRos;

// Struct for a sequence of aburos_msgs__msg__AbuRos.
typedef struct aburos_msgs__msg__AbuRos__Sequence
{
  aburos_msgs__msg__AbuRos * data;
  /// The number of valid items in data
  size_t size;
  /// The number of allocated items in data
  size_t capacity;
} aburos_msgs__msg__AbuRos__Sequence;

#ifdef __cplusplus
}
#endif

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_ROS__STRUCT_H_
