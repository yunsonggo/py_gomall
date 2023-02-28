import requests

print('直接通过http请求 向consul 注册 服务')

headers = {"contentType": "application/json"}


# 注册HTTP服务
def register(name, id, address, port):
    url = "http://192.168.1.136:8500/v1/agent/service/register"
    res = requests.put(url, headers=headers, json={
        "Name": name,
        "ID": id,
        "Tags": ["gomall", "user_srv"],
        "Address": address,
        "Port": port,
        "Check": {
            "HTTP": f"http://{address}:{port}/consul/health",
            "Timeout": "5s",
            "Interval": "5s",
            "DeregisterCriticalServiceAfter": "1m",
        }
    })
    if res.status_code == 200:
        print("注册成功")
    else:
        print(f'注册失败:{res.status_code}')


# 注册GRPC服务
def register_grpc(name, id, address, port):
    url = "http://192.168.1.136:8500/v1/agent/service/register"
    res = requests.put(url, headers=headers, json={
        "Name": name,
        "ID": id,
        "Tags": ["gomall", "user_srv"],
        "Address": address,
        "Port": port,
        "Check": {
            "GRPC": f"{address}:{port}",
            "GRPCUseTLS": False,
            "Timeout": "5s",
            "Interval": "5s",
            "DeregisterCriticalServiceAfter": "1m",
        }
    })
    if res.status_code == 200:
        print("注册成功")
    else:
        print(f'注册失败:{res.status_code}')


# 注销服务
def deregister(id):
    url = f"http://192.168.1.136:8500/v1/agent/service/deregister/{id}"
    res = requests.put(url)
    if res.status_code == 200:
        print("注销成功")
    else:
        print(f'注销失败:{res.status_code}')


if __name__ == "__main__":
    # register("gomall_user_web", "gomall_user_web", "192.168.1.136", 8000)
    deregister("gomall_user_srv")
    # register_grpc("gomall_user_srv", "gomall_user_srv", "192.168.1.136", 50051)
