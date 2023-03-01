import argparse
import socket
import sys
import grpc
import signal
import os
import datetime
import uuid
from concurrent import futures
from types import FrameType
from functools import partial

# 获取项目根路径
BASE_DIR = os.path.dirname(os.path.abspath(os.path.dirname(__file__)))
# 强行引入包地址 非虚拟环境或 虚拟环境正常不用引入
sys.path.insert(0, "/Volumes/pypro/environment/py_web_grpc22.0/lib/python3.11/site-packages")
sys.path.insert(1, BASE_DIR)

from loguru import logger
from user_srv.user_proto.user_gen import user_pb2_grpc as user_grpc
from user_srv.handler.user import UserServicer
from common.py_grpc_health.v1 import grpc_health
from common.py_grpc_health.v1 import grpc_health_pb2_grpc as health_grpc
from common.py_consul import http_request_consul as request_consul
from common.py_consul.consul_server_center import ConsulClient
from user_srv.settings import conf


def user_server():
    today = f"logs/{datetime.date.today()}/"
    log_path = today + "user_srv_{time}.log"
    logger.add(
        log_path,
        rotation=conf.APP_CONF["log"]["maxsize_mb"],
        retention=conf.APP_CONF["log"]["maxbackups_day"],
        compression=conf.APP_CONF["log"]["compression"],
    )
    # 解析命令行参数
    parser = argparse.ArgumentParser()
    parser.add_argument('--host', nargs="?", type=str, default=conf.APP_CONF["listen"], help="user server host")
    parser.add_argument('--port', nargs="?", type=str, default=conf.APP_CONF["port"], help="user server port")
    parser.add_argument('--workers', nargs="?", type=int, default=conf.APP_CONF["workers"], help="max_workers")
    args = parser.parse_args()
    if args.port == 0 and conf.APP_CONF["port"] == 0:
        conf.APP_CONF["port"] = free_port()
    if args.host != conf.APP_CONF["listen"]:
        conf.APP_CONF["listen"] = args.host
    if args.workers >= conf.APP_CONF["workers"]:
        conf.APP_CONF["workers"] = args.workers

    s = grpc.server(futures.ThreadPoolExecutor(max_workers=conf.APP_CONF["workers"]))
    # 注册用户GRPC服务
    user_grpc.add_UserServicer_to_server(UserServicer(), s)

    # 注册到consul

    # 方法一: 直接使用http请求注册
    # request_consul.register_grpc(conf.SERVER_NAME, conf.SERVER_NAME, conf.SERVER_HOST, int(conf.SERVER_PORT))
    # 方法二: 基于consul库 注册
    consul_client = ConsulClient(host="192.168.1.136", port=8500)
    server_uuid = str(uuid.uuid1())
    is_registered = consul_client.register(
        name=conf.APP_CONF["name"],
        id=server_uuid,
        address=conf.APP_CONF["listen"],
        port=conf.APP_CONF["port"],
        timeout=conf.APP_CONF["consul"]["timeout_sec"],
        interval=conf.APP_CONF["consul"]["interval_sec"],
        remove=conf.APP_CONF["consul"]["remove_after_sec"],
        tags=conf.APP_CONF["consul"]["tages"],
    )
    # is_registered = consul_client.register(
    #     name="gomall_user_srv",
    #     id=server_uuid,
    #     address="192.168.1.136",
    #     port=8500,
    #     timeout="5s",
    #     interval="5s",
    #     remove="15s",
    #     tags=["gomall", "user_srv"],
    # )
    if not is_registered:
        logger.info("User Server consul register error")
        sys.exit(0)
    else:
        logger.info("User Server consul register success")
    # 注册consul健康检查服务
    health_server = grpc_health.HealthServicer()
    health_grpc.add_HealthServicer_to_server(health_server, s)

    # 监听: 主进程退出信号 win/linux: ctrl+c:SIGINT, kill:SIGTERM
    signal.signal(signal.SIGINT, partial(on_exit, server_uuid=server_uuid))
    signal.signal(signal.SIGTERM, partial(on_exit, server_uuid=server_uuid))
    # s.add_insecure_port(f'[::]:{port}')
    s.add_insecure_port(f'{conf.APP_CONF["listen"]}:{conf.APP_CONF["port"]}')
    s.start()
    logger.info(
        f"User Server started, listening on {conf.APP_CONF['listen']} : {conf.APP_CONF['port']} ,workers:{conf.APP_CONF['workers']}")
    s.wait_for_termination()


# 获取随机端口号
def free_port() -> int:
    tcp = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp.bind(("", 0))
    _, port = tcp.getsockname()
    tcp.close()
    return port


def on_exit(signa: int, frame: FrameType, server_uuid):
    logger.info(f"注销 {server_uuid} 服务")
    # 方式一: 什么都不做,consul会自动注销不健康服务
    # 方式二:http请求主动注销consul
    # request_consul.deregister(conf.SERVER_NAME)
    # 方式二: consul库注销
    consul_client = ConsulClient(conf.APP_CONF["consul"]["consul_host"], conf.APP_CONF["consul"]["consul_port"])
    consul_client.deregister(id=server_uuid)
    sys.exit(0)


if __name__ == "__main__":
    # conf.client.add_config_watcher(conf.data["data_id"], conf.data["group"], conf.parse_conf)
    user_server()
    # print("conf port:", conf.SERVER_PORT)
    # print(type(free_port()), free_port())
