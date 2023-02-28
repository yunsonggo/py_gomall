# mysql 连接池
from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin


class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase):
    pass


SQL_DB = "py_gomall"
SQL_HOST = "192.168.1.136"
SQL_PORT = 13308
SQL_USER = "root"
SQL_PASS = "123456"

MAX_WORKERS = 10
SERVER_NAME = "gomall_user_srv"
SERVER_HOST = "192.168.1.136"
SERVER_PORT = 0

LOG_MAX_MB = "500 MB"
LOG_OLD_WEEK = "1 week"
LOG_CLEANUP_DAYS = "10 days"
LOG_COMP = "zip"


CONSUL_HOST = "192.168.1.136"
CONSUL_PORT = 8500
CONSUL_TAGS = ["gomall", "gomall_user_srv"]
CONSUL_TIMEOUT = "5s"
CONSUL_INTERVAL = "5s"
CONSUL_REMOVE = "90m"


DB = ReconnectMysqlDatabase(SQL_DB, host=SQL_HOST, port=SQL_PORT, user=SQL_USER, password=SQL_PASS)
