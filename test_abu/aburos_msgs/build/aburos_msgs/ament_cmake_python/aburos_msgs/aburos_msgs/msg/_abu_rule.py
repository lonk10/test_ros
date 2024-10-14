# generated from rosidl_generator_py/resource/_idl.py.em
# with input from aburos_msgs:msg/AbuRule.idl
# generated code does not contain a copyright notice


# Import statements for member types

import builtins  # noqa: E402, I100

import rosidl_parser.definition  # noqa: E402, I100


class Metaclass_AbuRule(type):
    """Metaclass of message 'AbuRule'."""

    _CREATE_ROS_MESSAGE = None
    _CONVERT_FROM_PY = None
    _CONVERT_TO_PY = None
    _DESTROY_ROS_MESSAGE = None
    _TYPE_SUPPORT = None

    __constants = {
    }

    @classmethod
    def __import_type_support__(cls):
        try:
            from rosidl_generator_py import import_type_support
            module = import_type_support('aburos_msgs')
        except ImportError:
            import logging
            import traceback
            logger = logging.getLogger(
                'aburos_msgs.msg.AbuRule')
            logger.debug(
                'Failed to import needed modules for type support:\n' +
                traceback.format_exc())
        else:
            cls._CREATE_ROS_MESSAGE = module.create_ros_message_msg__msg__abu_rule
            cls._CONVERT_FROM_PY = module.convert_from_py_msg__msg__abu_rule
            cls._CONVERT_TO_PY = module.convert_to_py_msg__msg__abu_rule
            cls._TYPE_SUPPORT = module.type_support_msg__msg__abu_rule
            cls._DESTROY_ROS_MESSAGE = module.destroy_ros_message_msg__msg__abu_rule

    @classmethod
    def __prepare__(cls, name, bases, **kwargs):
        # list constant names here so that they appear in the help text of
        # the message class under "Data and other attributes defined here:"
        # as well as populate each message instance
        return {
        }


class AbuRule(metaclass=Metaclass_AbuRule):
    """Message class 'AbuRule'."""

    __slots__ = [
        '_condition',
        '_actions',
        '_local_resources',
        '_remote_resources',
    ]

    _fields_and_field_types = {
        'condition': 'string',
        'actions': 'sequence<string>',
        'local_resources': 'sequence<octet>',
        'remote_resources': 'sequence<string>',
    }

    SLOT_TYPES = (
        rosidl_parser.definition.UnboundedString(),  # noqa: E501
        rosidl_parser.definition.UnboundedSequence(rosidl_parser.definition.UnboundedString()),  # noqa: E501
        rosidl_parser.definition.UnboundedSequence(rosidl_parser.definition.BasicType('octet')),  # noqa: E501
        rosidl_parser.definition.UnboundedSequence(rosidl_parser.definition.UnboundedString()),  # noqa: E501
    )

    def __init__(self, **kwargs):
        assert all('_' + key in self.__slots__ for key in kwargs.keys()), \
            'Invalid arguments passed to constructor: %s' % \
            ', '.join(sorted(k for k in kwargs.keys() if '_' + k not in self.__slots__))
        self.condition = kwargs.get('condition', str())
        self.actions = kwargs.get('actions', [])
        self.local_resources = kwargs.get('local_resources', [])
        self.remote_resources = kwargs.get('remote_resources', [])

    def __repr__(self):
        typename = self.__class__.__module__.split('.')
        typename.pop()
        typename.append(self.__class__.__name__)
        args = []
        for s, t in zip(self.__slots__, self.SLOT_TYPES):
            field = getattr(self, s)
            fieldstr = repr(field)
            # We use Python array type for fields that can be directly stored
            # in them, and "normal" sequences for everything else.  If it is
            # a type that we store in an array, strip off the 'array' portion.
            if (
                isinstance(t, rosidl_parser.definition.AbstractSequence) and
                isinstance(t.value_type, rosidl_parser.definition.BasicType) and
                t.value_type.typename in ['float', 'double', 'int8', 'uint8', 'int16', 'uint16', 'int32', 'uint32', 'int64', 'uint64']
            ):
                if len(field) == 0:
                    fieldstr = '[]'
                else:
                    assert fieldstr.startswith('array(')
                    prefix = "array('X', "
                    suffix = ')'
                    fieldstr = fieldstr[len(prefix):-len(suffix)]
            args.append(s[1:] + '=' + fieldstr)
        return '%s(%s)' % ('.'.join(typename), ', '.join(args))

    def __eq__(self, other):
        if not isinstance(other, self.__class__):
            return False
        if self.condition != other.condition:
            return False
        if self.actions != other.actions:
            return False
        if self.local_resources != other.local_resources:
            return False
        if self.remote_resources != other.remote_resources:
            return False
        return True

    @classmethod
    def get_fields_and_field_types(cls):
        from copy import copy
        return copy(cls._fields_and_field_types)

    @builtins.property
    def condition(self):
        """Message field 'condition'."""
        return self._condition

    @condition.setter
    def condition(self, value):
        if __debug__:
            assert \
                isinstance(value, str), \
                "The 'condition' field must be of type 'str'"
        self._condition = value

    @builtins.property
    def actions(self):
        """Message field 'actions'."""
        return self._actions

    @actions.setter
    def actions(self, value):
        if __debug__:
            from collections.abc import Sequence
            from collections.abc import Set
            from collections import UserList
            from collections import UserString
            assert \
                ((isinstance(value, Sequence) or
                  isinstance(value, Set) or
                  isinstance(value, UserList)) and
                 not isinstance(value, str) and
                 not isinstance(value, UserString) and
                 all(isinstance(v, str) for v in value) and
                 True), \
                "The 'actions' field must be a set or sequence and each value of type 'str'"
        self._actions = value

    @builtins.property
    def local_resources(self):
        """Message field 'local_resources'."""
        return self._local_resources

    @local_resources.setter
    def local_resources(self, value):
        if __debug__:
            from collections.abc import Sequence
            from collections.abc import Set
            from collections import UserList
            from collections import UserString
            assert \
                ((isinstance(value, Sequence) or
                  isinstance(value, Set) or
                  isinstance(value, UserList)) and
                 not isinstance(value, str) and
                 not isinstance(value, UserString) and
                 all(isinstance(v, bytes) for v in value) and
                 True), \
                "The 'local_resources' field must be a set or sequence and each value of type 'bytes'"
        self._local_resources = value

    @builtins.property
    def remote_resources(self):
        """Message field 'remote_resources'."""
        return self._remote_resources

    @remote_resources.setter
    def remote_resources(self, value):
        if __debug__:
            from collections.abc import Sequence
            from collections.abc import Set
            from collections import UserList
            from collections import UserString
            assert \
                ((isinstance(value, Sequence) or
                  isinstance(value, Set) or
                  isinstance(value, UserList)) and
                 not isinstance(value, str) and
                 not isinstance(value, UserString) and
                 all(isinstance(v, str) for v in value) and
                 True), \
                "The 'remote_resources' field must be a set or sequence and each value of type 'str'"
        self._remote_resources = value
