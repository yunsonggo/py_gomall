package by_http

import (
	"github.com/hashicorp/consul/api"
	"log"
)

func ConsulHttpRegister(consulHost string, addr string, port int, name string, tags []string, id string) {
	cfg := api.DefaultConfig()
	cfg.Address = consulHost

	client, err := api.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Agent().ServiceRegister(&api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Tags:    tags,
		Port:    port,
		Address: addr,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("注册成功")
}

func ConsulHttpDeregister(consulHost string, name string) {
	cfg := api.DefaultConfig()
	cfg.Address = consulHost

	client, err := api.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Agent().ServiceDeregister(name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ok")
}
