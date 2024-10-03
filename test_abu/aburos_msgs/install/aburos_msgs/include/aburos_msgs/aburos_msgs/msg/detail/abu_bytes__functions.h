// generated from rosidl_generator_c/resource/idl__functions.h.em
// with input from aburos_msgs:msg/AbuBytes.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__FUNCTIONS_H_
#define ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__FUNCTIONS_H_

#ifdef __cplusplus
extern "C"
{
#endif

#include <stdbool.h>
#include <stdlib.h>

#include "rosidl_runtime_c/visibility_control.h"
#include "aburos_msgs/msg/rosidl_generator_c__visibility_control.h"

#include "aburos_msgs/msg/detail/abu_bytes__struct.h"

/// Initialize msg/AbuBytes message.
/**
 * If the init function is called twice for the same message without
 * calling fini inbetween previously allocated memory will be leaked.
 * \param[in,out] msg The previously allocated message pointer.
 * Fields without a default value will not be initialized by this function.
 * You might want to call memset(msg, 0, sizeof(
 * aburos_msgs__msg__AbuBytes
 * )) before or use
 * aburos_msgs__msg__AbuBytes__create()
 * to allocate and initialize the message.
 * \return true if initialization was successful, otherwise false
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
bool
aburos_msgs__msg__AbuBytes__init(aburos_msgs__msg__AbuBytes * msg);

/// Finalize msg/AbuBytes message.
/**
 * \param[in,out] msg The allocated message pointer.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
void
aburos_msgs__msg__AbuBytes__fini(aburos_msgs__msg__AbuBytes * msg);

/// Create msg/AbuBytes message.
/**
 * It allocates the memory for the message, sets the memory to zero, and
 * calls
 * aburos_msgs__msg__AbuBytes__init().
 * \return The pointer to the initialized message if successful,
 * otherwise NULL
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
aburos_msgs__msg__AbuBytes *
aburos_msgs__msg__AbuBytes__create();

/// Destroy msg/AbuBytes message.
/**
 * It calls
 * aburos_msgs__msg__AbuBytes__fini()
 * and frees the memory of the message.
 * \param[in,out] msg The allocated message pointer.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
void
aburos_msgs__msg__AbuBytes__destroy(aburos_msgs__msg__AbuBytes * msg);

/// Check for msg/AbuBytes message equality.
/**
 * \param[in] lhs The message on the left hand size of the equality operator.
 * \param[in] rhs The message on the right hand size of the equality operator.
 * \return true if messages are equal, otherwise false.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
bool
aburos_msgs__msg__AbuBytes__are_equal(const aburos_msgs__msg__AbuBytes * lhs, const aburos_msgs__msg__AbuBytes * rhs);

/// Copy a msg/AbuBytes message.
/**
 * This functions performs a deep copy, as opposed to the shallow copy that
 * plain assignment yields.
 *
 * \param[in] input The source message pointer.
 * \param[out] output The target message pointer, which must
 *   have been initialized before calling this function.
 * \return true if successful, or false if either pointer is null
 *   or memory allocation fails.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
bool
aburos_msgs__msg__AbuBytes__copy(
  const aburos_msgs__msg__AbuBytes * input,
  aburos_msgs__msg__AbuBytes * output);

/// Initialize array of msg/AbuBytes messages.
/**
 * It allocates the memory for the number of elements and calls
 * aburos_msgs__msg__AbuBytes__init()
 * for each element of the array.
 * \param[in,out] array The allocated array pointer.
 * \param[in] size The size / capacity of the array.
 * \return true if initialization was successful, otherwise false
 * If the array pointer is valid and the size is zero it is guaranteed
 # to return true.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
bool
aburos_msgs__msg__AbuBytes__Sequence__init(aburos_msgs__msg__AbuBytes__Sequence * array, size_t size);

/// Finalize array of msg/AbuBytes messages.
/**
 * It calls
 * aburos_msgs__msg__AbuBytes__fini()
 * for each element of the array and frees the memory for the number of
 * elements.
 * \param[in,out] array The initialized array pointer.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
void
aburos_msgs__msg__AbuBytes__Sequence__fini(aburos_msgs__msg__AbuBytes__Sequence * array);

/// Create array of msg/AbuBytes messages.
/**
 * It allocates the memory for the array and calls
 * aburos_msgs__msg__AbuBytes__Sequence__init().
 * \param[in] size The size / capacity of the array.
 * \return The pointer to the initialized array if successful, otherwise NULL
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
aburos_msgs__msg__AbuBytes__Sequence *
aburos_msgs__msg__AbuBytes__Sequence__create(size_t size);

/// Destroy array of msg/AbuBytes messages.
/**
 * It calls
 * aburos_msgs__msg__AbuBytes__Sequence__fini()
 * on the array,
 * and frees the memory of the array.
 * \param[in,out] array The initialized array pointer.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
void
aburos_msgs__msg__AbuBytes__Sequence__destroy(aburos_msgs__msg__AbuBytes__Sequence * array);

/// Check for msg/AbuBytes message array equality.
/**
 * \param[in] lhs The message array on the left hand size of the equality operator.
 * \param[in] rhs The message array on the right hand size of the equality operator.
 * \return true if message arrays are equal in size and content, otherwise false.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
bool
aburos_msgs__msg__AbuBytes__Sequence__are_equal(const aburos_msgs__msg__AbuBytes__Sequence * lhs, const aburos_msgs__msg__AbuBytes__Sequence * rhs);

/// Copy an array of msg/AbuBytes messages.
/**
 * This functions performs a deep copy, as opposed to the shallow copy that
 * plain assignment yields.
 *
 * \param[in] input The source array pointer.
 * \param[out] output The target array pointer, which must
 *   have been initialized before calling this function.
 * \return true if successful, or false if either pointer
 *   is null or memory allocation fails.
 */
ROSIDL_GENERATOR_C_PUBLIC_aburos_msgs
bool
aburos_msgs__msg__AbuBytes__Sequence__copy(
  const aburos_msgs__msg__AbuBytes__Sequence * input,
  aburos_msgs__msg__AbuBytes__Sequence * output);

#ifdef __cplusplus
}
#endif

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__FUNCTIONS_H_
