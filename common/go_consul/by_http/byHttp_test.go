package by_http

import "testing"

func TestByHttp(t *testing.T) {
	serverId := "gomall_user_web"
	url := "192.168.1.136:8500"
	//deUrl := "http://192.168.1.136:8500/v1/agent/service/deregister/" + serverId
	//ConsulHttpRegister(url, "192.168.1.136", 8000, serverId, []string{"gomall", "gomall_web"}, serverId)
	ConsulHttpDeregister(url, serverId)
}
