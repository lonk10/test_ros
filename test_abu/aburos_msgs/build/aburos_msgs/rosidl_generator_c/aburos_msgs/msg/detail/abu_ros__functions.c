// generated from rosidl_generator_c/resource/idl__functions.c.em
// with input from aburos_msgs:msg/AbuRos.idl
// generated code does not contain a copyright notice
#include "aburos_msgs/msg/detail/abu_ros__functions.h"

#include <assert.h>
#include <stdbool.h>
#include <stdlib.h>
#include <string.h>

#include "rcutils/allocator.h"


// Include directives for member types
// Member `data`
#include "rosidl_runtime_c/primitives_sequence_functions.h"

bool
aburos_msgs__msg__AbuRos__init(aburos_msgs__msg__AbuRos * msg)
{
  if (!msg) {
    return false;
  }
  // data
  if (!rosidl_runtime_c__octet__Sequence__init(&msg->data, 0)) {
    aburos_msgs__msg__AbuRos__fini(msg);
    return false;
  }
  return true;
}

void
aburos_msgs__msg__AbuRos__fini(aburos_msgs__msg__AbuRos * msg)
{
  if (!msg) {
    return;
  }
  // data
  rosidl_runtime_c__octet__Sequence__fini(&msg->data);
}

bool
aburos_msgs__msg__AbuRos__are_equal(const aburos_msgs__msg__AbuRos * lhs, const aburos_msgs__msg__AbuRos * rhs)
{
  if (!lhs || !rhs) {
    return false;
  }
  // data
  if (!rosidl_runtime_c__octet__Sequence__are_equal(
      &(lhs->data), &(rhs->data)))
  {
    return false;
  }
  return true;
}

bool
aburos_msgs__msg__AbuRos__copy(
  const aburos_msgs__msg__AbuRos * input,
  aburos_msgs__msg__AbuRos * output)
{
  if (!input || !output) {
    return false;
  }
  // data
  if (!rosidl_runtime_c__octet__Sequence__copy(
      &(input->data), &(output->data)))
  {
    return false;
  }
  return true;
}

aburos_msgs__msg__AbuRos *
aburos_msgs__msg__AbuRos__create()
{
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  aburos_msgs__msg__AbuRos * msg = (aburos_msgs__msg__AbuRos *)allocator.allocate(sizeof(aburos_msgs__msg__AbuRos), allocator.state);
  if (!msg) {
    return NULL;
  }
  memset(msg, 0, sizeof(aburos_msgs__msg__AbuRos));
  bool success = aburos_msgs__msg__AbuRos__init(msg);
  if (!success) {
    allocator.deallocate(msg, allocator.state);
    return NULL;
  }
  return msg;
}

void
aburos_msgs__msg__AbuRos__destroy(aburos_msgs__msg__AbuRos * msg)
{
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  if (msg) {
    aburos_msgs__msg__AbuRos__fini(msg);
  }
  allocator.deallocate(msg, allocator.state);
}


bool
aburos_msgs__msg__AbuRos__Sequence__init(aburos_msgs__msg__AbuRos__Sequence * array, size_t size)
{
  if (!array) {
    return false;
  }
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  aburos_msgs__msg__AbuRos * data = NULL;

  if (size) {
    data = (aburos_msgs__msg__AbuRos *)allocator.zero_allocate(size, sizeof(aburos_msgs__msg__AbuRos), allocator.state);
    if (!data) {
      return false;
    }
    // initialize all array elements
    size_t i;
    for (i = 0; i < size; ++i) {
      bool success = aburos_msgs__msg__AbuRos__init(&data[i]);
      if (!success) {
        break;
      }
    }
    if (i < size) {
      // if initialization failed finalize the already initialized array elements
      for (; i > 0; --i) {
        aburos_msgs__msg__AbuRos__fini(&data[i - 1]);
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
aburos_msgs__msg__AbuRos__Sequence__fini(aburos_msgs__msg__AbuRos__Sequence * array)
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
      aburos_msgs__msg__AbuRos__fini(&array->data[i]);
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

aburos_msgs__msg__AbuRos__Sequence *
aburos_msgs__msg__AbuRos__Sequence__create(size_t size)
{
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  aburos_msgs__msg__AbuRos__Sequence * array = (aburos_msgs__msg__AbuRos__Sequence *)allocator.allocate(sizeof(aburos_msgs__msg__AbuRos__Sequence), allocator.state);
  if (!array) {
    return NULL;
  }
  bool success = aburos_msgs__msg__AbuRos__Sequence__init(array, size);
  if (!success) {
    allocator.deallocate(array, allocator.state);
    return NULL;
  }
  return array;
}

void
aburos_msgs__msg__AbuRos__Sequence__destroy(aburos_msgs__msg__AbuRos__Sequence * array)
{
  rcutils_allocator_t allocator = rcutils_get_default_allocator();
  if (array) {
    aburos_msgs__msg__AbuRos__Sequence__fini(array);
  }
  allocator.deallocate(array, allocator.state);
}

bool
aburos_msgs__msg__AbuRos__Sequence__are_equal(const aburos_msgs__msg__AbuRos__Sequence * lhs, const aburos_msgs__msg__AbuRos__Sequence * rhs)
{
  if (!lhs || !rhs) {
    return false;
  }
  if (lhs->size != rhs->size) {
    return false;
  }
  for (size_t i = 0; i < lhs->size; ++i) {
    if (!aburos_msgs__msg__AbuRos__are_equal(&(lhs->data[i]), &(rhs->data[i]))) {
      return false;
    }
  }
  return true;
}

bool
aburos_msgs__msg__AbuRos__Sequence__copy(
  const aburos_msgs__msg__AbuRos__Sequence * input,
  aburos_msgs__msg__AbuRos__Sequence * output)
{
  if (!input || !output) {
    return false;
  }
  if (output->capacity < input->size) {
    const size_t allocation_size =
      input->size * sizeof(aburos_msgs__msg__AbuRos);
    rcutils_allocator_t allocator = rcutils_get_default_allocator();
    aburos_msgs__msg__AbuRos * data =
      (aburos_msgs__msg__AbuRos *)allocator.reallocate(
      output->data, allocation_size, allocator.state);
    if (!data) {
      return false;
    }
    // If reallocation succeeded, memory may or may not have been moved
    // to fulfill the allocation request, invalidating output->data.
    output->data = data;
    for (size_t i = output->capacity; i < input->size; ++i) {
      if (!aburos_msgs__msg__AbuRos__init(&output->data[i])) {
        // If initialization of any new item fails, roll back
        // all previously initialized items. Existing items
        // in output are to be left unmodified.
        for (; i-- > output->capacity; ) {
          aburos_msgs__msg__AbuRos__fini(&output->data[i]);
        }
        return false;
      }
    }
    output->capacity = input->size;
  }
  output->size = input->size;
  for (size_t i = 0; i < input->size; ++i) {
    if (!aburos_msgs__msg__AbuRos__copy(
        &(input->data[i]), &(output->data[i])))
    {
      return false;
    }
  }
  return true;
}
