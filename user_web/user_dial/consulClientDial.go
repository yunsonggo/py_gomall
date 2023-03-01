package user_dial

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"py_gomall/v2/common/go_consul"
)

var ConsulClient *go_consul.ConsulClient

func NewConsulClient() error {
	client, err := go_consul.NewConsulClient(ConsulHost, ConsulPort)
	ConsulClient = client
	return err
}

func Register() error {
	healthURL := fmt.Sprintf("http://%s:%d/api/consul/health", Host, Port)
	serverUUID := fmt.Sprintf("%s", uuid.NewV4())
	fmt.Println("serverUUID:", serverUUID)
	item := go_consul.ServerItem{
		Address:            Host,
		Port:               Port,
		Name:               Name,
		Tage:               Tages,
		ID:                 serverUUID,
		HealthURL:          healthURL,
		TimeoutSec:         TimeoutSec,
		IntervalSec:        IntervalSec,
		DeregisterAfterSec: RemoveAfterSec,
	}
	fmt.Println(item)
	return ConsulClient.Register(item)
}

func Deregister() error {
	return ConsulClient.Deregister(Name)
}

func ServiceList() error {
	services, err := ConsulClient.Services(Services)
	if err != nil {
		return err
	}
	ServerInfoList = make(map[string]*ServerInfo)
	for name, info := range services {
		sInfo := ServerInfo{
			Name: info.Name,
			Host: info.Host,
			Port: info.Port,
			Conn: info.Conn,
		}
		ServerInfoList[name] = &sInfo
		fmt.Printf("%+v\n", sInfo)
	}
	return nil
}

func GetServer() (*go_consul.ServerInfo, error) {
	return ConsulClient.GetServer(Name)
}
