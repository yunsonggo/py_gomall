import abc


class PyServerCenter(metaclass=abc.ABCMeta):
    @abc.abstractmethod
    def register(self,  name, id, address, port, timeout, interval, remove, tags):
        pass

    @abc.abstractmethod
    def deregister(self, id):
        pass

    @abc.abstractmethod
    def services(self):
        pass

    @abc.abstractmethod
    def get_server(self, id):
        pass
