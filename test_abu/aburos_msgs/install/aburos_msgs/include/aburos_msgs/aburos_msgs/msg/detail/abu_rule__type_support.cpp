// generated from rosidl_typesupport_introspection_cpp/resource/idl__type_support.cpp.em
// with input from aburos_msgs:msg/AbuRule.idl
// generated code does not contain a copyright notice

#include "array"
#include "cstddef"
#include "string"
#include "vector"
#include "rosidl_runtime_c/message_type_support_struct.h"
#include "rosidl_typesupport_cpp/message_type_support.hpp"
#include "rosidl_typesupport_interface/macros.h"
#include "aburos_msgs/msg/detail/abu_rule__struct.hpp"
#include "rosidl_typesupport_introspection_cpp/field_types.hpp"
#include "rosidl_typesupport_introspection_cpp/identifier.hpp"
#include "rosidl_typesupport_introspection_cpp/message_introspection.hpp"
#include "rosidl_typesupport_introspection_cpp/message_type_support_decl.hpp"
#include "rosidl_typesupport_introspection_cpp/visibility_control.h"

namespace aburos_msgs
{

namespace msg
{

namespace rosidl_typesupport_introspection_cpp
{

void AbuRule_init_function(
  void * message_memory, rosidl_runtime_cpp::MessageInitialization _init)
{
  new (message_memory) aburos_msgs::msg::AbuRule(_init);
}

void AbuRule_fini_function(void * message_memory)
{
  auto typed_message = static_cast<aburos_msgs::msg::AbuRule *>(message_memory);
  typed_message->~AbuRule();
}

size_t size_function__AbuRule__actions(const void * untyped_member)
{
  const auto * member = reinterpret_cast<const std::vector<std::string> *>(untyped_member);
  return member->size();
}

const void * get_const_function__AbuRule__actions(const void * untyped_member, size_t index)
{
  const auto & member =
    *reinterpret_cast<const std::vector<std::string> *>(untyped_member);
  return &member[index];
}

void * get_function__AbuRule__actions(void * untyped_member, size_t index)
{
  auto & member =
    *reinterpret_cast<std::vector<std::string> *>(untyped_member);
  return &member[index];
}

void fetch_function__AbuRule__actions(
  const void * untyped_member, size_t index, void * untyped_value)
{
  const auto & item = *reinterpret_cast<const std::string *>(
    get_const_function__AbuRule__actions(untyped_member, index));
  auto & value = *reinterpret_cast<std::string *>(untyped_value);
  value = item;
}

void assign_function__AbuRule__actions(
  void * untyped_member, size_t index, const void * untyped_value)
{
  auto & item = *reinterpret_cast<std::string *>(
    get_function__AbuRule__actions(untyped_member, index));
  const auto & value = *reinterpret_cast<const std::string *>(untyped_value);
  item = value;
}

void resize_function__AbuRule__actions(void * untyped_member, size_t size)
{
  auto * member =
    reinterpret_cast<std::vector<std::string> *>(untyped_member);
  member->resize(size);
}

size_t size_function__AbuRule__local_resources(const void * untyped_member)
{
  const auto * member = reinterpret_cast<const std::vector<unsigned char> *>(untyped_member);
  return member->size();
}

const void * get_const_function__AbuRule__local_resources(const void * untyped_member, size_t index)
{
  const auto & member =
    *reinterpret_cast<const std::vector<unsigned char> *>(untyped_member);
  return &member[index];
}

void * get_function__AbuRule__local_resources(void * untyped_member, size_t index)
{
  auto & member =
    *reinterpret_cast<std::vector<unsigned char> *>(untyped_member);
  return &member[index];
}

void fetch_function__AbuRule__local_resources(
  const void * untyped_member, size_t index, void * untyped_value)
{
  const auto & item = *reinterpret_cast<const unsigned char *>(
    get_const_function__AbuRule__local_resources(untyped_member, index));
  auto & value = *reinterpret_cast<unsigned char *>(untyped_value);
  value = item;
}

void assign_function__AbuRule__local_resources(
  void * untyped_member, size_t index, const void * untyped_value)
{
  auto & item = *reinterpret_cast<unsigned char *>(
    get_function__AbuRule__local_resources(untyped_member, index));
  const auto & value = *reinterpret_cast<const unsigned char *>(untyped_value);
  item = value;
}

void resize_function__AbuRule__local_resources(void * untyped_member, size_t size)
{
  auto * member =
    reinterpret_cast<std::vector<unsigned char> *>(untyped_member);
  member->resize(size);
}

size_t size_function__AbuRule__remote_resources(const void * untyped_member)
{
  const auto * member = reinterpret_cast<const std::vector<std::string> *>(untyped_member);
  return member->size();
}

const void * get_const_function__AbuRule__remote_resources(const void * untyped_member, size_t index)
{
  const auto & member =
    *reinterpret_cast<const std::vector<std::string> *>(untyped_member);
  return &member[index];
}

void * get_function__AbuRule__remote_resources(void * untyped_member, size_t index)
{
  auto & member =
    *reinterpret_cast<std::vector<std::string> *>(untyped_member);
  return &member[index];
}

void fetch_function__AbuRule__remote_resources(
  const void * untyped_member, size_t index, void * untyped_value)
{
  const auto & item = *reinterpret_cast<const std::string *>(
    get_const_function__AbuRule__remote_resources(untyped_member, index));
  auto & value = *reinterpret_cast<std::string *>(untyped_value);
  value = item;
}

void assign_function__AbuRule__remote_resources(
  void * untyped_member, size_t index, const void * untyped_value)
{
  auto & item = *reinterpret_cast<std::string *>(
    get_function__AbuRule__remote_resources(untyped_member, index));
  const auto & value = *reinterpret_cast<const std::string *>(untyped_value);
  item = value;
}

void resize_function__AbuRule__remote_resources(void * untyped_member, size_t size)
{
  auto * member =
    reinterpret_cast<std::vector<std::string> *>(untyped_member);
  member->resize(size);
}

static const ::rosidl_typesupport_introspection_cpp::MessageMember AbuRule_message_member_array[4] = {
  {
    "condition",  // name
    ::rosidl_typesupport_introspection_cpp::ROS_TYPE_STRING,  // type
    0,  // upper bound of string
    nullptr,  // members of sub message
    false,  // is array
    0,  // array size
    false,  // is upper bound
    offsetof(aburos_msgs::msg::AbuRule, condition),  // bytes offset in struct
    nullptr,  // default value
    nullptr,  // size() function pointer
    nullptr,  // get_const(index) function pointer
    nullptr,  // get(index) function pointer
    nullptr,  // fetch(index, &value) function pointer
    nullptr,  // assign(index, value) function pointer
    nullptr  // resize(index) function pointer
  },
  {
    "actions",  // name
    ::rosidl_typesupport_introspection_cpp::ROS_TYPE_STRING,  // type
    0,  // upper bound of string
    nullptr,  // members of sub message
    true,  // is array
    0,  // array size
    false,  // is upper bound
    offsetof(aburos_msgs::msg::AbuRule, actions),  // bytes offset in struct
    nullptr,  // default value
    size_function__AbuRule__actions,  // size() function pointer
    get_const_function__AbuRule__actions,  // get_const(index) function pointer
    get_function__AbuRule__actions,  // get(index) function pointer
    fetch_function__AbuRule__actions,  // fetch(index, &value) function pointer
    assign_function__AbuRule__actions,  // assign(index, value) function pointer
    resize_function__AbuRule__actions  // resize(index) function pointer
  },
  {
    "local_resources",  // name
    ::rosidl_typesupport_introspection_cpp::ROS_TYPE_OCTET,  // type
    0,  // upper bound of string
    nullptr,  // members of sub message
    true,  // is array
    0,  // array size
    false,  // is upper bound
    offsetof(aburos_msgs::msg::AbuRule, local_resources),  // bytes offset in struct
    nullptr,  // default value
    size_function__AbuRule__local_resources,  // size() function pointer
    get_const_function__AbuRule__local_resources,  // get_const(index) function pointer
    get_function__AbuRule__local_resources,  // get(index) function pointer
    fetch_function__AbuRule__local_resources,  // fetch(index, &value) function pointer
    assign_function__AbuRule__local_resources,  // assign(index, value) function pointer
    resize_function__AbuRule__local_resources  // resize(index) function pointer
  },
  {
    "remote_resources",  // name
    ::rosidl_typesupport_introspection_cpp::ROS_TYPE_STRING,  // type
    0,  // upper bound of string
    nullptr,  // members of sub message
    true,  // is array
    0,  // array size
    false,  // is upper bound
    offsetof(aburos_msgs::msg::AbuRule, remote_resources),  // bytes offset in struct
    nullptr,  // default value
    size_function__AbuRule__remote_resources,  // size() function pointer
    get_const_function__AbuRule__remote_resources,  // get_const(index) function pointer
    get_function__AbuRule__remote_resources,  // get(index) function pointer
    fetch_function__AbuRule__remote_resources,  // fetch(index, &value) function pointer
    assign_function__AbuRule__remote_resources,  // assign(index, value) function pointer
    resize_function__AbuRule__remote_resources  // resize(index) function pointer
  }
};

static const ::rosidl_typesupport_introspection_cpp::MessageMembers AbuRule_message_members = {
  "aburos_msgs::msg",  // message namespace
  "AbuRule",  // message name
  4,  // number of fields
  sizeof(aburos_msgs::msg::AbuRule),
  AbuRule_message_member_array,  // message members
  AbuRule_init_function,  // function to initialize message memory (memory has to be allocated)
  AbuRule_fini_function  // function to terminate message instance (will not free memory)
};

static const rosidl_message_type_support_t AbuRule_message_type_support_handle = {
  ::rosidl_typesupport_introspection_cpp::typesupport_identifier,
  &AbuRule_message_members,
  get_message_typesupport_handle_function,
};

}  // namespace rosidl_typesupport_introspection_cpp

}  // namespace msg

}  // namespace aburos_msgs


namespace rosidl_typesupport_introspection_cpp
{

template<>
ROSIDL_TYPESUPPORT_INTROSPECTION_CPP_PUBLIC
const rosidl_message_type_support_t *
get_message_type_support_handle<aburos_msgs::msg::AbuRule>()
{
  return &::aburos_msgs::msg::rosidl_typesupport_introspection_cpp::AbuRule_message_type_support_handle;
}

}  // namespace rosidl_typesupport_introspection_cpp

#ifdef __cplusplus
extern "C"
{
#endif

ROSIDL_TYPESUPPORT_INTROSPECTION_CPP_PUBLIC
const rosidl_message_type_support_t *
ROSIDL_TYPESUPPORT_INTERFACE__MESSAGE_SYMBOL_NAME(rosidl_typesupport_introspection_cpp, aburos_msgs, msg, AbuRule)() {
  return &::aburos_msgs::msg::rosidl_typesupport_introspection_cpp::AbuRule_message_type_support_handle;
}

#ifdef __cplusplus
}
#endif
