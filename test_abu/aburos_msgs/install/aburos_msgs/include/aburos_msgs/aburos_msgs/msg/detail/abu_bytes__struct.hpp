// generated from rosidl_generator_cpp/resource/idl__struct.hpp.em
// with input from aburos_msgs:msg/AbuBytes.idl
// generated code does not contain a copyright notice

#ifndef ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__STRUCT_HPP_
#define ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__STRUCT_HPP_

#include <algorithm>
#include <array>
#include <memory>
#include <string>
#include <vector>

#include "rosidl_runtime_cpp/bounded_vector.hpp"
#include "rosidl_runtime_cpp/message_initialization.hpp"


#ifndef _WIN32
# define DEPRECATED__aburos_msgs__msg__AbuBytes __attribute__((deprecated))
#else
# define DEPRECATED__aburos_msgs__msg__AbuBytes __declspec(deprecated)
#endif

namespace aburos_msgs
{

namespace msg
{

// message struct
template<class ContainerAllocator>
struct AbuBytes_
{
  using Type = AbuBytes_<ContainerAllocator>;

  explicit AbuBytes_(rosidl_runtime_cpp::MessageInitialization _init = rosidl_runtime_cpp::MessageInitialization::ALL)
  {
    if (rosidl_runtime_cpp::MessageInitialization::ALL == _init ||
      rosidl_runtime_cpp::MessageInitialization::ZERO == _init)
    {
      this->origin = "";
    }
  }

  explicit AbuBytes_(const ContainerAllocator & _alloc, rosidl_runtime_cpp::MessageInitialization _init = rosidl_runtime_cpp::MessageInitialization::ALL)
  : origin(_alloc)
  {
    if (rosidl_runtime_cpp::MessageInitialization::ALL == _init ||
      rosidl_runtime_cpp::MessageInitialization::ZERO == _init)
    {
      this->origin = "";
    }
  }

  // field types and members
  using _data_type =
    std::vector<unsigned char, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<unsigned char>>;
  _data_type data;
  using _origin_type =
    std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>>;
  _origin_type origin;

  // setters for named parameter idiom
  Type & set__data(
    const std::vector<unsigned char, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<unsigned char>> & _arg)
  {
    this->data = _arg;
    return *this;
  }
  Type & set__origin(
    const std::basic_string<char, std::char_traits<char>, typename std::allocator_traits<ContainerAllocator>::template rebind_alloc<char>> & _arg)
  {
    this->origin = _arg;
    return *this;
  }

  // constant declarations

  // pointer types
  using RawPtr =
    aburos_msgs::msg::AbuBytes_<ContainerAllocator> *;
  using ConstRawPtr =
    const aburos_msgs::msg::AbuBytes_<ContainerAllocator> *;
  using SharedPtr =
    std::shared_ptr<aburos_msgs::msg::AbuBytes_<ContainerAllocator>>;
  using ConstSharedPtr =
    std::shared_ptr<aburos_msgs::msg::AbuBytes_<ContainerAllocator> const>;

  template<typename Deleter = std::default_delete<
      aburos_msgs::msg::AbuBytes_<ContainerAllocator>>>
  using UniquePtrWithDeleter =
    std::unique_ptr<aburos_msgs::msg::AbuBytes_<ContainerAllocator>, Deleter>;

  using UniquePtr = UniquePtrWithDeleter<>;

  template<typename Deleter = std::default_delete<
      aburos_msgs::msg::AbuBytes_<ContainerAllocator>>>
  using ConstUniquePtrWithDeleter =
    std::unique_ptr<aburos_msgs::msg::AbuBytes_<ContainerAllocator> const, Deleter>;
  using ConstUniquePtr = ConstUniquePtrWithDeleter<>;

  using WeakPtr =
    std::weak_ptr<aburos_msgs::msg::AbuBytes_<ContainerAllocator>>;
  using ConstWeakPtr =
    std::weak_ptr<aburos_msgs::msg::AbuBytes_<ContainerAllocator> const>;

  // pointer types similar to ROS 1, use SharedPtr / ConstSharedPtr instead
  // NOTE: Can't use 'using' here because GNU C++ can't parse attributes properly
  typedef DEPRECATED__aburos_msgs__msg__AbuBytes
    std::shared_ptr<aburos_msgs::msg::AbuBytes_<ContainerAllocator>>
    Ptr;
  typedef DEPRECATED__aburos_msgs__msg__AbuBytes
    std::shared_ptr<aburos_msgs::msg::AbuBytes_<ContainerAllocator> const>
    ConstPtr;

  // comparison operators
  bool operator==(const AbuBytes_ & other) const
  {
    if (this->data != other.data) {
      return false;
    }
    if (this->origin != other.origin) {
      return false;
    }
    return true;
  }
  bool operator!=(const AbuBytes_ & other) const
  {
    return !this->operator==(other);
  }
};  // struct AbuBytes_

// alias to use template instance with default allocator
using AbuBytes =
  aburos_msgs::msg::AbuBytes_<std::allocator<void>>;

// constant definitions

}  // namespace msg

}  // namespace aburos_msgs

#endif  // ABUROS_MSGS__MSG__DETAIL__ABU_BYTES__STRUCT_HPP_
