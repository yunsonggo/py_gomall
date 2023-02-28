import datetime
import time

import grpc
from google.protobuf.empty_pb2 import Empty
from user_srv.user_proto.user_gen import user_pb2 as user_pb
from user_srv.user_proto.user_gen import user_pb2_grpc as user_grpc
from user_srv.model.models import User

from common.paginate import paginate
from loguru import logger
from peewee import DoesNotExist
from common.pwd import pwd_passlib


class UserServicer(user_grpc.UserServicer):
    @logger.catch
    def UserList(self, request: user_pb.PageRequest, context):
        # 分页数据
        resp = user_pb.UsersResponse()
        users = User.select().where(User.deleted_at.is_null(True))
        resp.total = users.count()
        if not request.page:
            request.page = 1
        if not request.size:
            request.size = 5
        size, offset = paginate.paginate(request.page, request.size)
        users = users.limit(size).offset(offset)
        for user in users:
            resp.data.append(self.user_model_resp(user))
        return resp

    @logger.catch
    def UserFirst(self, request: user_pb.IDRequest, context):
        print(request)
        # ID或手机号 查询用户
        context.set_code(grpc.StatusCode.OK)
        context.set_details('OK')
        response = user_pb.UserResponse
        user = User()
        if request.intID > 0:
            try:
                user = User.select().where((User.id == request.intID) & (User.deleted_at.is_null(True))).get()
            except DoesNotExist as e:
                context.set_code(grpc.StatusCode.NOT_FOUND)
                context.set_details(f'not found user where id:{request.intID}')
            else:
                response = self.user_model_resp(user)
        elif request.strID != "":
            try:
                user = User.select().where((User.mobile == request.strID) & (User.deleted_at.is_null(True))).get()
            except DoesNotExist as e:
                context.set_code(grpc.StatusCode.NOT_FOUND)
                context.set_details(f'not found user where mobile:{request.strID}')
            else:
                response = self.user_model_resp(user)
        else:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details('param error')
        print(response)
        # print(result)
        return response

    @logger.catch
    def UserCreate(self, request: user_pb.UserRequest, context):
        # 添加用户
        context.set_code(grpc.StatusCode.OK)
        context.set_details('ok')
        response = user_pb.UserResponse()
        if not request.mobile or request.mobile == "":
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details('param error,mobile required')
            return response
        else:
            try:
                user = User.get(User.mobile == request.mobile)
            except DoesNotExist as e:
                create_user = self.user_req_model(request)
                row = create_user.save()
                if row == 0:
                    context.set_code(grpc.StatusCode.INTERNAL)
                    context.set_details(f'create user error where mobile:{request.mobile}')
                else:
                    response = self.user_model_resp(create_user)
            else:
                context.set_code(grpc.StatusCode.ALREADY_EXISTS)
                context.set_details(f'user already exists where mobile:{request.mobile}')
            finally:
                return response

    @logger.catch
    def UserEdit(self, request: user_pb.UserEditRequest, context):
        # 修改用户
        context.set_code(grpc.StatusCode.OK)
        context.set_details('ok')
        id_request = request.ids
        info_request = request.info
        info_request.password = ""
        try:
            user = User()
            if id_request.intID:
                user = User.select().where((User.id == id_request.intID) & (User.deleted_at.is_null(True))).get()
            elif id_request.strID:
                user = User.select().where((User.id == id_request.intID) & (User.deleted_at.is_null(True))).get()
            else:
                context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
                context.set_details('param error')
                return user_pb.EmptyMessage().empty()
        except DoesNotExist as e:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details(f'not found user where id:{id_request.intID},mobile:{id_request.strID},err:{e}')
        else:
            if info_request.nickname:
                user.nickname = info_request.nickname
            if info_request.icon:
                user.icon = info_request.icon
            if info_request.birthday > 0:
                user.birthday = datetime.datetime.fromtimestamp(info_request.birthday)
            if info_request.addr:
                user.addr = info_request.addr
            if info_request.desc:
                user.desc = info_request.desc
            if info_request.gender != user.gender:
                user.gender = info_request.gender
            user.updated_at = datetime.datetime.now()
            row = user.save()
            if row == 0:
                context.set_code(grpc.StatusCode.INTERNAL)
                context.set_details(f'edit user error where id:{id_request.intID}')
        finally:
            return Empty()

    @logger.catch
    def UserRemove(self, request, context):
        # 删除用户
        context.set_code(grpc.StatusCode.OK)
        context.set_details('ok')
        if not request.intID:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details('param error')
            return Empty()
        now_time = datetime.datetime.now()
        res = User.update({User.deleted_at: now_time}).where(User.id == request.intID).execute()
        print(res)
        if res == 0:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f'delete user where id {request.intID} error')
        return Empty()

    @logger.catch
    def UserCheckPasswd(self, request: user_pb.UserRequest, context):
        context.set_code(grpc.StatusCode.OK)
        context.set_details('ok')
        response = user_pb.CheckPasswdResponse()
        response.isChecked = user_pb.CHECKED_UNKNOWN
        try:
            user = User.select().where((User.mobile == request.mobile) & (User.deleted_at.is_null(True))).get()
            print(user)
        except DoesNotExist as e:
            context.set_details(f'not found user where mobile:{request.mobile},error:{e}')
            response.isChecked = user_pb.CHECKED_NO
        else:
            if not pwd_passlib.verify_sha256(request.password, user.password):
                context.set_details(f'old password error')
                response.isChecked = user_pb.CHECKED_NO
            else:
                response.isChecked = user_pb.CHECKED_YES
        finally:
            print(response)
            return response

    @logger.catch
    def UserEditPasswd(self, request: user_pb.PasswdEditRequest, context):
        # 修改密码
        context.set_code(grpc.StatusCode.OK)
        context.set_details('ok')
        user = User()
        if request.ids.intID > 0:
            try:
                user = User.select().where((User.id == request.ids.intID) & (User.deleted_at.is_null(True))).get()
            except DoesNotExist as e:
                context.set_code(grpc.StatusCode.NOT_FOUND)
                context.set_details(f'not found user where id:{request.ids.intID},error:{e}')
                return Empty()
        elif request.ids.strID != "":
            try:
                user = User.select().where((User.mobile == request.ids.strID) & (User.deleted_at.is_null(True))).get()
            except DoesNotExist as e:
                context.set_code(grpc.StatusCode.NOT_FOUND)
                context.set_details(f'not found user where mobile:{request.ids.strID},error:{e}')
                return Empty()
        else:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details('param error')
        # 校验密码
        if not pwd_passlib.verify_sha256(request.oldPasswd, user.password):
            context.set_code(grpc.StatusCode.PERMISSION_DENIED)
            context.set_details(f'old password error')
            return Empty()
        new_passwd = pwd_passlib.parse_sha256(request.newPasswd)
        user.updated_at = datetime.datetime.now()
        user.password = new_passwd
        row = user.save()
        if row == 0:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details(f'edit user password error where id:{request.ids.intID} or mobile:{request.ids.strID}')
        return Empty()

    @logger.catch
    def user_model_resp(self, user: User):
        info = user_pb.UserResponse()
        info.id = user.id
        info.mobile = user.mobile
        if user.nickname:
            info.nickname = user.nickname
        if user.desc:
            info.desc = user.desc
        info.password = user.password
        info.gender = user.gender
        if user.addr:
            info.addr = user.addr
        if user.icon:
            info.icon = user.icon
        if user.birthday:
            info.birthday = int(time.mktime(user.birthday.timetuple()))
        if user.created_at:
            info.created_at = int(time.mktime(user.created_at.timetuple()))
        if user.updated_at:
            info.updated_at = int(time.mktime(user.updated_at.timetuple()))
        if user.deleted_at:
            info.deleted_at = int(time.mktime(user.deleted_at.timetuple()))
        return info

    @logger.catch
    def user_req_model(self, req: user_pb.UserRequest):
        create_info = User()
        create_info.mobile = req.mobile
        create_info.password = pwd_passlib.parse_sha256(req.password)
        if req.nickname:
            create_info.nickname = req.nickname
        if req.icon:
            create_info.icon = req.icon
        if req.birthday > 0:
            create_info.birthday = datetime.datetime.fromtimestamp(req.birthday)
        if req.addr:
            create_info.addr = req.addr
        if req.desc:
            create_info.desc = req.desc
        create_info.gender = req.gender
        create_info.role = req.role
        return create_info
