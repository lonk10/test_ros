// generated from rosidl_generator_c/resource/idl__functions.c.em
// with input from aburos_msgs:msg/AbuRule.idl
// generated code does not contain a copyright notice
#include "aburos_msgs/msg/detail/abu_rule__functions.h"

#include <assert.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

#include "rcutils/allocator.h"


// Include directives for member types
// Member `condition`
// Member `actions`
// Member `remote_resources`
#include "rosidl_runtime_c/string_functions.h"
// Member `local_resources`
#include "rosidl_runtime_c/primitives_sequence_functions.h"

bool
aburos_msgs__msg__AbuRule__init(aburos_msgs__msg__AbuRule * msg)
{
  if (!msg) {
    return false;
  }
  // condition
  if (!rosidl_runtime_c__String__init(&msg->condition)) {
    aburos_msgs__msg__AbuRule__fini(msg);
    return false;
  }
  // actions
  if (!rosidl_runtime_c__String__Sequence__init(&msg->actions, 0)) {
    aburos_msgs__msg__AbuRule__fini(msg);
    return false;
  }
  // local_resources
  if (!rosidl_runtime_c__octet__Sequence__init(&msg->local_resources, 0)) {
    aburos_msgs__msg__AbuRule__fini(msg);
    return false;
  }
  // remote_resources
  if (!rosidl_runtime_c__String__Sequence__init(&msg->remote_resources, 0)) {
    aburos_msgs__msg__AbuRule__fini(msg);
    return false;
  }
  return true;
}

void
aburos_msgs__msg__AbuRule__fini(aburos_msgs__msg__AbuRule * msg)
{
  if (!msg) {
    return;
  }
  // condition
  rosidl_runtime_c__String__fini(&msg->condition);
  // actions
  rosidl_runtime_c__String__Sequence__fini(&msg->actions);
  // local_resources
  rosidl_runtime_c__octet__Sequence__fini(&msg->local_resources);
  // remote_resources
  rosidl_runtime_c__String__Sequence__fini(&msg->remote_resources);
}

bool
aburos_msgs__msg__AbuRule__are_equal(const aburos_msgs__msg__AbuRule * lhs, const aburos_msgs__msg__AbuRule * rhs)
{
  if (!lhs || !rhs) {
    return false;
  }
  // condition
  if (!rosidl_runtime_c__String__are_equal(
      &(lhs->condition), &(rhs->condition)))
  {
    return false;
  }
  // actions
  if (!rosidl_runtime_c__String__Sequence__are_equal(
      &(lhs->actions), &(rhs->actions)))
  {
    return false;
  }
  // local_resources
  if (!rosidl_runtime_c__octet__Sequence__are_equal(
      &(lhs->local_resources), &(rhs->local_resources)))
  {
    return false;
  }
  // remote_resources
  if (!rosidl_runtime_c__String__Sequence__are_equal(
      &(lhs->remote_resources), &(rhs->remote_resources)))
  {
    return false;
  }
  return true;
}

bool
aburos_msgs__msg__AbuRule__copy(
  const aburos_msgs__msg__AbuRule * input,
  aburos_msgs__msg__AbuRule * output)
{
  if (!input || !output) {
    return false;
  }
  // condition
  if (!rosidl_runtime_c__String__copy(
      &(input->condition), &(output->condition)))
  {
    return false;
  }
  // actions
  if (!rosidl_runtime_c__String__Sequence__copy(
      &(input->actions), &(output->actions)))
  {
    return false;
  }
  // local_resources
  if (!rosidl_runtime_c__octet__Sequence__copy(
      &(input->local_resources), &(output->local_resources)))
  {
    return false;
  }
  // remote_resources
  if (!rosidl_runtime_c__String__Sequence__copy(
      &(input->remote_resources), &(output->remote_resources)))
  {
    return false;
  }
  return true;
}

aburos_msgs__msg__AbuRule *
aburos_msgs__msg__AbuRule__create()
{
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  aburos_msgs__msg__AbuRule * msg = (aburos_msgs__msg__AbuRule *)allocator.allocate(sizeof(aburos_msgs__msg__AbuRule), allocator.state);
  if (!msg) {
    return NULL;
  }
  memset(msg, 0, sizeof(aburos_msgs__msg__AbuRule));
  bool success = aburos_msgs__msg__AbuRule__init(msg);
  if (!success) {
    allocator.deallocate(msg, allocator.state);
    return NULL;
  }
  return msg;
}

void
aburos_msgs__msg__AbuRule__destroy(aburos_msgs__msg__AbuRule * msg)
{
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  if (msg) {
    aburos_msgs__msg__AbuRule__fini(msg);
  }
  allocator.deallocate(msg, allocator.state);
}


bool
aburos_msgs__msg__AbuRule__Sequence__init(aburos_msgs__msg__AbuRule__Sequence * array, size_t size)
{
  if (!array) {
    return false;
  }
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  aburos_msgs__msg__AbuRule * data = NULL;

  if (size) {
    data = (aburos_msgs__msg__AbuRule *)allocator.zero_allocate(size, sizeof(aburos_msgs__msg__AbuRule), allocator.state);
    if (!data) {
      return false;
    }
    // initialize all array elements
    size_t i;
    for (i = 0; i < size; ++i) {
      bool success = aburos_msgs__msg__AbuRule__init(&data[i]);
      if (!success) {
        break;
      }
    }
    if (i < size) {
      // if initialization failed finalize the already initialized array elements
      for (; i > 0; --i) {
        aburos_msgs__msg__AbuRule__fini(&data[i - 1]);
      }
      allocator.deallocate(data, allocator.state);
      return false;
    }
  }
  array->data = data;
  array->size = size;
  array->capacity = size;
  return true;
}

void
aburos_msgs__msg__AbuRule__Sequence__fini(aburos_msgs__msg__AbuRule__Sequence * array)
{
  if (!array) {
    return;
  }
  rcutils_allocator_t allocator = rcutils_get_default_allocator();

  if (array->data) {
    // ensure that data and capacity values are consistent
    assert(array->capacity > 0);
    // finalize all array elements
    for (size_t i = 0; i < array->capacity; ++i) {
      aburos_msgs__msg__AbuRule__fini(&array->data[i]);
    }
    allocator.deallocate(array->data, allocator.state);
    array->data = NULL;
    array->size = 0;
    array->capacity = 0;
  } else {
    // ensure that data, size, and capacity values are consistent
    assert(0 == array->size);
    assert(0 == array->capacity);
  }
}

aburos_msgs__msg__AbuRule__Sequence *
aburos_msgs__msg__AbuRule__Sequence__create(size_t size)
{
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  aburos_msgs__msg__AbuRule__Sequence * array = (aburos_msgs__msg__AbuRule__Sequence *)allocator.allocate(sizeof(aburos_msgs__msg__AbuRule__Sequence), allocator.state);
  if (!array) {
    return NULL;
  }
  bool success = aburos_msgs__msg__AbuRule__Sequence__init(array, size);
  if (!success) {
    allocator.deallocate(array, allocator.state);
    return NULL;
  }
  return array;
}

void
aburos_msgs__msg__AbuRule__Sequence__destroy(aburos_msgs__msg__AbuRule__Sequence * array)
{
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  if (array) {
    aburos_msgs__msg__AbuRule__Sequence__fini(array);
  }
  allocator.deallocate(array, allocator.state);
}

bool
aburos_msgs__msg__AbuRule__Sequence__are_equal(const aburos_msgs__msg__AbuRule__Sequence * lhs, const aburos_msgs__msg__AbuRule__Sequence * rhs)
{
  if (!lhs || !rhs) {
    return false;
  }
  if (lhs->size != rhs->size) {
    return false;
  }
  for (size_t i = 0; i < lhs->size; ++i) {
    if (!aburos_msgs__msg__AbuRule__are_equal(&(lhs->data[i]), &(rhs->data[i]))) {
      return false;
    }
  }
  return true;
}

bool
aburos_msgs__msg__AbuRule__Sequence__copy(
  const aburos_msgs__msg__AbuRule__Sequence * input,
  aburos_msgs__msg__AbuRule__Sequence * output)
{
  if (!input || !output) {
    return false;
  }
  if (output->capacity < input->size) {
    const size_t allocation_size =
      input->size * sizeof(aburos_msgs__msg__AbuRule);
    rcutils_allocator_t allocator = rcutils_get_default_allocator();
    aburos_msgs__msg__AbuRule * data =
      (aburos_msgs__msg__AbuRule *)allocator.reallocate(
      output->data, allocation_size, allocator.state);
    if (!data) {
      return false;
    }
    // If reallocation succeeded, memory may or may not have been moved
    // to fulfill the allocation request, invalidating output->data.
    output->data = data;
    for (size_t i = output->capacity; i < input->size; ++i) {
      if (!aburos_msgs__msg__AbuRule__init(&output->data[i])) {
        // If initialization of any new item fails, roll back
        // all previously initialized items. Existing items
        // in output are to be left unmodified.
        for (; i-- > output->capacity; ) {
          aburos_msgs__msg__AbuRule__fini(&output->data[i]);
        }
        return false;
      }
    }
    output->capacity = input->size;
  }
  output->size = input->size;
  for (size_t i = 0; i < input->size; ++i) {
    if (!aburos_msgs__msg__AbuRule__copy(
        &(input->data[i]), &(output->data[i])))
    {
      return false;
    }
  }
  return true;
}
