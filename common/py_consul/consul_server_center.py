import consul
import requests

from common.py_server_center.py_server_center_interface import PyServerCenter as sc


class ConsulClient(sc):
    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.c = consul.Consul(host=host, port=port)

    def register(self, name, id, address, port, timeout, interval, remove, tags) -> bool:
        check = {
            "GRPC": f"{address}:{port}",
            "GRPCUseTLS": False,
            "Timeout": timeout,
            "Interval": interval,
            "DeregisterCriticalServiceAfter": remove,
        }
        result = self.c.agent.service.register(name=name, service_id=id, address=address, port=port, check=check,
                                               tags=tags)
        return result

    def deregister(self, id):
        return self.c.agent.service.deregister(id)

    def services(self):
        return self.c.agent.services()

    def get_server(self, id):
        url = f"http://{self.host}:{self.port}/v1/agent/services"
        params = {
            "filter": f'Service == "{id}"'
        }

        res = requests.get(url, params=params).json()
        for k, v in res.items():
            print(k, v)
        return res