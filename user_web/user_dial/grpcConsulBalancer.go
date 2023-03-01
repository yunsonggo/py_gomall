package user_dial

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	_ "py_gomall/v2/common/go_consul_resolver"
)

var UserBalancerConn *grpc.ClientConn

// InitSrvConn grpc-consul-resolver 引入grpc-consul负载均衡库 注册 拉取服务
func InitSrvConn(dialSrvName string) *grpc.ClientConn {
	conn, err := grpc.Dial(
		"consul://192.168.1.136:8500/gomall_user_srv?wait=14s",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// 前提: 启动多个服务实例,注册入consul时,服务ID必须唯一才能有负载均衡的效果
		// 第一步 引入 resolver
		// 第二步 配置 loadBalancingPolicy
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal("InitSrvConn 链接服务失败")
	}
	return conn
}
