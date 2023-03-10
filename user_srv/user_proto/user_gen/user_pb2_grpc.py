# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from user_srv.user_proto.user_gen import user_pb2 as user__pb2


class UserStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.UserList = channel.unary_unary(
                '/user_proto.User/UserList',
                request_serializer=user__pb2.PageRequest.SerializeToString,
                response_deserializer=user__pb2.UsersResponse.FromString,
                )
        self.UserFirst = channel.unary_unary(
                '/user_proto.User/UserFirst',
                request_serializer=user__pb2.IDRequest.SerializeToString,
                response_deserializer=user__pb2.UserResponse.FromString,
                )
        self.UserCreate = channel.unary_unary(
                '/user_proto.User/UserCreate',
                request_serializer=user__pb2.UserRequest.SerializeToString,
                response_deserializer=user__pb2.UserResponse.FromString,
                )
        self.UserEdit = channel.unary_unary(
                '/user_proto.User/UserEdit',
                request_serializer=user__pb2.UserEditRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.UserRemove = channel.unary_unary(
                '/user_proto.User/UserRemove',
                request_serializer=user__pb2.IDRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )
        self.UserCheckPasswd = channel.unary_unary(
                '/user_proto.User/UserCheckPasswd',
                request_serializer=user__pb2.UserRequest.SerializeToString,
                response_deserializer=user__pb2.CheckPasswdResponse.FromString,
                )
        self.UserEditPasswd = channel.unary_unary(
                '/user_proto.User/UserEditPasswd',
                request_serializer=user__pb2.PasswdEditRequest.SerializeToString,
                response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
                )


class UserServicer(object):
    """Missing associated documentation comment in .proto file."""

    def UserList(self, request, context):
        """user ????????????
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UserFirst(self, request, context):
        """ID???????????? ????????????
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UserCreate(self, request, context):
        """????????????
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UserEdit(self, request, context):
        """????????????
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UserRemove(self, request, context):
        """????????????
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UserCheckPasswd(self, request, context):
        """????????????
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UserEditPasswd(self, request, context):
        """????????????
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_UserServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'UserList': grpc.unary_unary_rpc_method_handler(
                    servicer.UserList,
                    request_deserializer=user__pb2.PageRequest.FromString,
                    response_serializer=user__pb2.UsersResponse.SerializeToString,
            ),
            'UserFirst': grpc.unary_unary_rpc_method_handler(
                    servicer.UserFirst,
                    request_deserializer=user__pb2.IDRequest.FromString,
                    response_serializer=user__pb2.UserResponse.SerializeToString,
            ),
            'UserCreate': grpc.unary_unary_rpc_method_handler(
                    servicer.UserCreate,
                    request_deserializer=user__pb2.UserRequest.FromString,
                    response_serializer=user__pb2.UserResponse.SerializeToString,
            ),
            'UserEdit': grpc.unary_unary_rpc_method_handler(
                    servicer.UserEdit,
                    request_deserializer=user__pb2.UserEditRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'UserRemove': grpc.unary_unary_rpc_method_handler(
                    servicer.UserRemove,
                    request_deserializer=user__pb2.IDRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
            'UserCheckPasswd': grpc.unary_unary_rpc_method_handler(
                    servicer.UserCheckPasswd,
                    request_deserializer=user__pb2.UserRequest.FromString,
                    response_serializer=user__pb2.CheckPasswdResponse.SerializeToString,
            ),
            'UserEditPasswd': grpc.unary_unary_rpc_method_handler(
                    servicer.UserEditPasswd,
                    request_deserializer=user__pb2.PasswdEditRequest.FromString,
                    response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'user_proto.User', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class User(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def UserList(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/user_proto.User/UserList',
            user__pb2.PageRequest.SerializeToString,
            user__pb2.UsersResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UserFirst(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/user_proto.User/UserFirst',
            user__pb2.IDRequest.SerializeToString,
            user__pb2.UserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UserCreate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/user_proto.User/UserCreate',
            user__pb2.UserRequest.SerializeToString,
            user__pb2.UserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UserEdit(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/user_proto.User/UserEdit',
            user__pb2.UserEditRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UserRemove(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/user_proto.User/UserRemove',
            user__pb2.IDRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UserCheckPasswd(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/user_proto.User/UserCheckPasswd',
            user__pb2.UserRequest.SerializeToString,
            user__pb2.CheckPasswdResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UserEditPasswd(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/user_proto.User/UserEditPasswd',
            user__pb2.PasswdEditRequest.SerializeToString,
            google_dot_protobuf_dot_empty__pb2.Empty.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
