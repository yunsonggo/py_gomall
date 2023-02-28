import datetime

import grpc
from user_srv.user_proto.user_gen import user_pb2_grpc as user_grpc
from user_srv.user_proto.user_gen import user_pb2 as user_pb
from user_srv.settings import conf


class UserTest:
    def __init__(self):
        # 连接GRPC服务
        self.stub = user_grpc.UserStub(grpc.insecure_channel(f"{conf.SERVER_HOST}:{conf.SERVER_PORT}"))
        # self.stub = user_grpc.UserStub(grpc.insecure_channel("192.168.1.136:5051"))

    def user_list(self, page=1, size=5):
        request = user_pb.PageRequest()
        request.page = page
        request.size = size
        resp = self.stub.UserList(request)
        return resp

    def user_first(self, int_id=0, mobile=""):
        request = user_pb.IDRequest()
        request.intID = int_id
        request.strID = mobile
        return self.stub.UserFirst(request)

    def user_create(self, mobile, nickname, passwd, birthday):
        request = user_pb.UserRequest()
        request.mobile = mobile
        request.nickname = nickname
        request.password = passwd
        request.birthday = birthday
        return self.stub.UserCreate(request)

    def user_edit(self, id, nickname):
        request = user_pb.UserEditRequest()
        request.ids.intID = id
        request.info.nickname = nickname
        return self.stub.UserEdit(request)

    def user_remove(self, id):
        request = user_pb.IDRequest
        request.intID = id
        return self.stub.UserRemove(request)

    def user_edit_passwd(self, int_id: int = 0, mobile: str = "", old: str = "", new: str = ""):
        request = user_pb.PasswdEditRequest()
        request.ids.intID = int_id
        request.ids.strID = mobile
        request.oldPasswd = old
        request.newPasswd = new
        return self.stub.UserEditPasswd(request)

    def check_passwd(self, mobile, passwd):
        request = user_pb.UserRequest()
        request.mobile = mobile
        request.password = passwd
        return self.stub.UserCheckPasswd(request)


if __name__ == "__main__":
    user = UserTest()
    # result = user.user_list(1, 5)
    # print(result.total)
    # for u in result.data:
    #     print(u)
    # resp = user.user_first(mobile="13666666664")
    # resp = user.user_first(int_id=5)
    # print(resp)
    # bd = int(datetime.datetime.now().timestamp())
    # resp = user.user_create("13999999998", "bee8", "bee123", bd)
    # print(resp)
    # resp = user.user_edit(id=1, nickname="editBee55")
    # print(resp)
    # resp = user.user_remove(id=14)
    # resp = user.user_edit_passwd(1, "", "admin1234", "admin123")
    resp = user.check_passwd("13666666660", "admin1234")
    print(resp)
