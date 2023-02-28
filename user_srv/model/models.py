import datetime
from peewee import *
from user_srv.settings import conf


# from common.pwd import pwd_passlib


class BaseModel(Model):
    id = BigAutoField(primary_key=True)
    created_at = DateTimeField(default=datetime.datetime.now)
    updated_at = DateTimeField(null=True, default=None)
    deleted_at = DateTimeField(null=True, default=None)

    class Meta:
        database = conf.DB


class User(BaseModel):
    Gender_Choices = (("unknown", "保密"), ("female", "女"), ("male", "男"))
    Roler_Choices = (("user", 1), ("manager", 2))
    mobile = CharField(max_length=11, index=True, null=False, unique=True, verbose_name="电话")
    password = CharField(max_length=200, null=False, default="", verbose_name="密码")
    nickname = CharField(max_length=20, null=False, default="", verbose_name="昵称")
    icon = CharField(max_length=255, null=False, default="", verbose_name="头像")
    birthday = DateField(null=True, verbose_name="生日")
    addr = CharField(max_length=200, null=False, default="", verbose_name="地址")
    desc = TextField(null=False, default="", verbose_name="简介")
    gender = CharField(max_length=8, choices=Gender_Choices, null=False, default="保密", verbose_name="性别")
    role = SmallIntegerField(choices=Roler_Choices, null=False, default=1, verbose_name="用户角色,1:用户,2:管理员")


if __name__ == "__main__":
    if not conf.DB.table_exists("user"):
        conf.DB.create_tables([User])
    # conf.DB.create_tables([User])
    # infos = []
    # i = 0
    # while i < 10:
    #     info = {"mobile":f"1366666666{i}","password":pwd_passlib.parse_sha256("admin123"),"nickname":f"admin123{i}"}
    #     infos.append(info)
    #     i = i + 1
    # res = User.insert_many(infos).execute()

    # i = 10
    # while i < 20:
    #     user = User()
    #     user.mobile = f"136666666{i}"
    #     user.password = pwd_passlib.parse_sha256("admin123")
    #     user.nickname = f"admin123{i}"
    #     user.save()
    #     i = i + 1
