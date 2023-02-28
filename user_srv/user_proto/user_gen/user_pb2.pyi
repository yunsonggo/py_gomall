from google.protobuf import empty_pb2 as _empty_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

CHECKED_NO: CheckedMessage
CHECKED_UNKNOWN: CheckedMessage
CHECKED_YES: CheckedMessage
DESCRIPTOR: _descriptor.FileDescriptor

class CheckPasswdResponse(_message.Message):
    __slots__ = ["isChecked"]
    ISCHECKED_FIELD_NUMBER: _ClassVar[int]
    isChecked: CheckedMessage
    def __init__(self, isChecked: _Optional[_Union[CheckedMessage, str]] = ...) -> None: ...

class HasCheckedResponse(_message.Message):
    __slots__ = ["checked"]
    CHECKED_FIELD_NUMBER: _ClassVar[int]
    checked: bool
    def __init__(self, checked: bool = ...) -> None: ...

class IDRequest(_message.Message):
    __slots__ = ["intID", "strID"]
    INTID_FIELD_NUMBER: _ClassVar[int]
    STRID_FIELD_NUMBER: _ClassVar[int]
    intID: int
    strID: str
    def __init__(self, intID: _Optional[int] = ..., strID: _Optional[str] = ...) -> None: ...

class PageRequest(_message.Message):
    __slots__ = ["page", "size"]
    PAGE_FIELD_NUMBER: _ClassVar[int]
    SIZE_FIELD_NUMBER: _ClassVar[int]
    page: int
    size: int
    def __init__(self, page: _Optional[int] = ..., size: _Optional[int] = ...) -> None: ...

class PasswdEditRequest(_message.Message):
    __slots__ = ["ids", "newPasswd", "oldPasswd"]
    IDS_FIELD_NUMBER: _ClassVar[int]
    NEWPASSWD_FIELD_NUMBER: _ClassVar[int]
    OLDPASSWD_FIELD_NUMBER: _ClassVar[int]
    ids: IDRequest
    newPasswd: str
    oldPasswd: str
    def __init__(self, ids: _Optional[_Union[IDRequest, _Mapping]] = ..., oldPasswd: _Optional[str] = ..., newPasswd: _Optional[str] = ...) -> None: ...

class UserEditRequest(_message.Message):
    __slots__ = ["ids", "info"]
    IDS_FIELD_NUMBER: _ClassVar[int]
    INFO_FIELD_NUMBER: _ClassVar[int]
    ids: IDRequest
    info: UserRequest
    def __init__(self, ids: _Optional[_Union[IDRequest, _Mapping]] = ..., info: _Optional[_Union[UserRequest, _Mapping]] = ...) -> None: ...

class UserRequest(_message.Message):
    __slots__ = ["addr", "birthday", "desc", "gender", "icon", "mobile", "nickname", "password", "role"]
    ADDR_FIELD_NUMBER: _ClassVar[int]
    BIRTHDAY_FIELD_NUMBER: _ClassVar[int]
    DESC_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    ICON_FIELD_NUMBER: _ClassVar[int]
    MOBILE_FIELD_NUMBER: _ClassVar[int]
    NICKNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    addr: str
    birthday: int
    desc: str
    gender: str
    icon: str
    mobile: str
    nickname: str
    password: str
    role: int
    def __init__(self, mobile: _Optional[str] = ..., password: _Optional[str] = ..., nickname: _Optional[str] = ..., icon: _Optional[str] = ..., birthday: _Optional[int] = ..., addr: _Optional[str] = ..., desc: _Optional[str] = ..., gender: _Optional[str] = ..., role: _Optional[int] = ...) -> None: ...

class UserResponse(_message.Message):
    __slots__ = ["addr", "birthday", "created_at", "deleted_at", "desc", "gender", "icon", "id", "mobile", "nickname", "password", "role", "updated_at"]
    ADDR_FIELD_NUMBER: _ClassVar[int]
    BIRTHDAY_FIELD_NUMBER: _ClassVar[int]
    CREATED_AT_FIELD_NUMBER: _ClassVar[int]
    DELETED_AT_FIELD_NUMBER: _ClassVar[int]
    DESC_FIELD_NUMBER: _ClassVar[int]
    GENDER_FIELD_NUMBER: _ClassVar[int]
    ICON_FIELD_NUMBER: _ClassVar[int]
    ID_FIELD_NUMBER: _ClassVar[int]
    MOBILE_FIELD_NUMBER: _ClassVar[int]
    NICKNAME_FIELD_NUMBER: _ClassVar[int]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    ROLE_FIELD_NUMBER: _ClassVar[int]
    UPDATED_AT_FIELD_NUMBER: _ClassVar[int]
    addr: str
    birthday: int
    created_at: int
    deleted_at: int
    desc: str
    gender: str
    icon: str
    id: int
    mobile: str
    nickname: str
    password: str
    role: int
    updated_at: int
    def __init__(self, id: _Optional[int] = ..., created_at: _Optional[int] = ..., updated_at: _Optional[int] = ..., deleted_at: _Optional[int] = ..., mobile: _Optional[str] = ..., password: _Optional[str] = ..., nickname: _Optional[str] = ..., icon: _Optional[str] = ..., birthday: _Optional[int] = ..., addr: _Optional[str] = ..., desc: _Optional[str] = ..., gender: _Optional[str] = ..., role: _Optional[int] = ...) -> None: ...

class UsersResponse(_message.Message):
    __slots__ = ["data", "total"]
    DATA_FIELD_NUMBER: _ClassVar[int]
    TOTAL_FIELD_NUMBER: _ClassVar[int]
    data: _containers.RepeatedCompositeFieldContainer[UserResponse]
    total: int
    def __init__(self, total: _Optional[int] = ..., data: _Optional[_Iterable[_Union[UserResponse, _Mapping]]] = ...) -> None: ...

class CheckedMessage(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
