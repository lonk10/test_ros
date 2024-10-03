// generated from rosidl_generator_cpp/resource/idl__struct.hpp.em
// with input from aburos_msgs:msg/AbuRule.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_RULE__STRUCT_HPP_
#define ABUROS_MSGS__MSG__DETAIL__ABU_RULE__STRUCT_HPP_

#include <algorithm>
#include <array>
#include <memory>
#include <string>
#include <vector>

#include "rosidl_runtime_cpp/bounded_vector.hpp"
#include "rosidl_runtime_cpp/message_initialization.hpp"


#ifndef _WIN32
# define DEPRECATED__aburos_msgs__msg__AbuRule __attribute__((deprecated))
#else
# define DEPRECATED__aburos_msgs__msg__AbuRule __declspec(deprecated)
#endif

namespace aburos_msgs
{

namespace msg
{

// message struct
template<class ContainerAllocator>
struct AbuRule_
{
  using Type = AbuRule_<ContainerAllocator>;

  explicit AbuRule_(rosidl_runtime_cpp::MessageInitialization _init = rosidl_runtime_cpp::MessageInitialization::ALL)
  {
    if (rosidl_runtime_cpp::MessageInitialization::ALL == _init ||
      rosidl_runtime_cpp::MessageInitialization::ZERO == _init)
    {
      this->condition = "";
    }
  }

  explicit AbuRule_(const ContainerAllocator & _alloc, rosidl_runtime_cpp::MessageInitialization _init = rosidl_runtime_cpp::MessageInitialization::ALL)
  : condition(_alloc)
  {
    if (rosidl_runtime_cpp::MessageInitialization::ALL == _init ||
      rosidl_runtime_cpp::MessageInitialization::ZERO == _init)
    {
      this->condition = "";
    }
  }

  // field types and members
  using _condition_type =
    std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>;
  _condition_type condition;
  using _actions_type =
    std::vector<std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>>>;
  _actions_type actions;
  using _local_resources_type =
    std::vector<unsigned char, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<unsigned char>>;
  _local_resources_type local_resources;
  using _remote_resources_type =
    std::vector<std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>>>;
  _remote_resources_type remote_resources;

  // setters for named parameter idiom
  Type & set__condition(
    const std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>> & _arg)
  {
    this->condition = _arg;
    return *this;
  }
  Type & set__actions(
    const std::vector<std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>>> & _arg)
  {
    this->actions = _arg;
    return *this;
  }
  Type & set__local_resources(
    const std::vector<unsigned char, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<unsigned char>> & _arg)
  {
    this->local_resources = _arg;
    return *this;
  }
  Type & set__remote_resources(
    const std::vector<std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>>> & _arg)
  {
    this->remote_resources = _arg;
    return *this;
  }

  // constant declarations

  // pointer types
  using RawPtr =
    aburos_msgs::msg::AbuRule_<ContainerAllocator> *;
  using ConstRawPtr =
    const aburos_msgs::msg::AbuRule_<ContainerAllocator> *;
  using SharedPtr =
    std::shared_ptr<aburos_msgs::msg::AbuRule_<ContainerAllocator>>;
  using ConstSharedPtr =
    std::shared_ptr<aburos_msgs::msg::AbuRule_<ContainerAllocator> const>;

  template<typename Deleter = std::default_delete<
      aburos_msgs::msg::AbuRule_<ContainerAllocator>>>
  using UniquePtrWithDeleter =
    std::unique_ptr<aburos_msgs::msg::AbuRule_<ContainerAllocator>, Deleter>;

  using UniquePtr = UniquePtrWithDeleter<>;

  template<typename Deleter = std::default_delete<
      aburos_msgs::msg::AbuRule_<ContainerAllocator>>>
  using ConstUniquePtrWithDeleter =
    std::unique_ptr<aburos_msgs::msg::AbuRule_<ContainerAllocator> const, Deleter>;
  using ConstUniquePtr = ConstUniquePtrWithDeleter<>;

  using WeakPtr =
    std::weak_ptr<aburos_msgs::msg::AbuRule_<ContainerAllocator>>;
  using ConstWeakPtr =
    std::weak_ptr<aburos_msgs::msg::AbuRule_<ContainerAllocator> const>;

  // pointer types similar to ROS 1, use SharedPtr / ConstSharedPtr instead
  // NOTE: Can't use 'using' here because GNU C++ can't parse attributes properly
  typedef DEPRECATED__aburos_msgs__msg__AbuRule
    std::shared_ptr<aburos_msgs::msg::AbuRule_<ContainerAllocator>>
    Ptr;
  typedef DEPRECATED__aburos_msgs__msg__AbuRule
    std::shared_ptr<aburos_msgs::msg::AbuRule_<ContainerAllocator> const>
    ConstPtr;

  // comparison operators
  bool operator==(const AbuRule_ & other) const
  {
    if (this->condition != other.condition) {
      return false;
    }
    if (this->actions != other.actions) {
      return false;
    }
    if (this->local_resources != other.local_resources) {
      return false;
    }
    if (this->remote_resources != other.remote_resources) {
      return false;
    }
    return true;
  }
  bool operator!=(const AbuRule_ & other) const
  {
    return !this->operator==(other);
  }
};  // struct AbuRule_

// alias to use template instance with default allocator
using AbuRule =
  aburos_msgs::msg::AbuRule_<std::allocator<void>>;

// constant definitions

}  // namespace msg

}  // namespace aburos_msgs

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_RULE__STRUCT_HPP_
