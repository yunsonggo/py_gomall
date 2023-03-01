package client_balancer

import (
	"fmt"
	"github.com/simplesurance/grpcconsulresolver/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"py_gomall/v2/user_web/user_dial"
)

func init() {
	resolver.Register(consul.NewBuilder())
}

func NewUserBalancer(name string) (*grpc.ClientConn, error) {
	//resolver.Register(consul.NewBuilder())
	//consul://%s:%d/%s?scheme=https&tags=primary,eu&health=fallbackToUnhealthy
	url := fmt.Sprintf("consul://%s:%d/%s?scheme=http", user_dial.ConsulHost, user_dial.ConsulPort, name)
	userClientConn, err := grpc.Dial(
		url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 前提: 启动多个服务实例,注册入consul时,服务ID必须唯一才能有负载均衡的效果
		// 第一步 引入 resolver
		// 第二步 配置 loadBalancingPolicy
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		fmt.Println("dial error", err)
		return nil, err
	}
	fmt.Println("dial ok")

	return userClientConn, nil
}
