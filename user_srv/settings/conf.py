# mysql 连接池
import nacos
import os
import yaml
import json
from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin

# 获取当前脚本所在文件夹路径
curPath = os.path.dirname(os.path.realpath(__file__))
# 获取yaml文件路径
yamlPath = os.path.join(curPath, "nacos_conf/user_nacos_conf.yaml")
f = open(yamlPath, 'r', encoding='utf-8')
data = yaml.load(f.read(), Loader=yaml.FullLoader)
client = nacos.NacosClient(server_addresses=data['nacos_addr'], namespace=data['namespace_id'])
str_data = client.get_config(data_id=data['data_id'], group=data['group'], timeout=data['timeout_ms'])
json_data = json.loads(str_data)
APP_CONF = json_data


def parse_conf(args):
    print(args)


class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase):
    pass


DB = ReconnectMysqlDatabase(
    APP_CONF["mysql"]["db"],
    host=APP_CONF["mysql"]["host"],
    port=APP_CONF["mysql"]["port"],
    user=APP_CONF["mysql"]["user"],
    password=APP_CONF["mysql"]["pass"],
)
