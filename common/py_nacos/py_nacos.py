import time

import nacos
import yaml
import os
import json


def nacos_at(path: str):
    # 获取当前脚本所在文件夹路径
    curPath = os.path.dirname(os.path.realpath(__file__))
    # 获取yaml文件路径
    yamlPath = os.path.join(curPath, path)
    f = open(yamlPath, 'r', encoding='utf-8')
    # str_data = f.read()
    data = yaml.load(f.read(), Loader=yaml.FullLoader)
    client = nacos.NacosClient(server_addresses=data['nacos_addr'], namespace=data['namespace_id'])
    str_data = client.get_config(data_id=data['data_id'], group=data['group'], timeout=data['timeout_ms'])
    json_data = json.loads(str_data)
    return client, data, json_data


# def test_cb(args):
#     print('配置更新')
#     print(args)


if __name__ == "__main__":
    nacos_at("nacos_conf/nacos_conf.yaml")
