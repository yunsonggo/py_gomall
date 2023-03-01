package user_dial

import (
	"context"
	"fmt"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"py_gomall/v2/common/go_grpc_pool"
	userpb "py_gomall/v2/user_web/user_proto/user_proto_gen"
	"strconv"
	"strings"
	"time"
)

var Services, Tages []string
var Host, Name, ConsulHost, TimeoutSec, IntervalSec, RemoveAfterSec string
var Port, ConsulPort int
var ServerInfoList map[string]*ServerInfo
var UserClientPool *go_grpc_pool.Pool

type ServerInfo struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
	Conn *grpc.ClientConn
}

// RegistrationServices 手动注册服务
func RegistrationServices() error {
	ServerInfoList = make(map[string]*ServerInfo)
	for _, info := range Services {
		info = strings.TrimSpace(info)
		infos := strings.Split(info, ",")
		var s ServerInfo
		for i := 0; i < len(infos); i++ {
			switch i {
			case 0:
				s.Name = infos[i]
			case 1:
				s.Host = infos[i]
			case 2:
				port, err := strconv.Atoi(infos[i])
				if err != nil {
					zap.L().Debug("RegistrationServices error", zap.Error(err))
					return err
				}
				s.Port = port
			}
		}
		addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
		if err != nil {
			zap.L().Debug("dial: 连接grpc服务错误", zap.Error(err))
			break
		}
		s.Conn = conn
		ServerInfoList[s.Name] = &s
	}
	return nil
}

//func NewUserClient() userpb.UserClient {
//	fmt.Println(ServerInfoList)
//	return userpb.NewUserClient(ServerInfoList["gomall_user_srv"].Conn)
//}

func NewUserClient() userpb.UserClient {
	return userpb.NewUserClient(UserBalancerConn)
}

func NewUserClientPool(host string, port int, length, cap int, timeout time.Duration) (*go_grpc_pool.Pool, error) {
	addr := fmt.Sprintf("%s:%d", host, port)
	f := func(addr string) (*grpc.ClientConn, error) {
		return grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
	}
	return go_grpc_pool.New(addr, f, length, cap, timeout)
}

func UserClient(ctx context.Context) userpb.UserClient {
	conn, err := UserClientPool.Get(ctx)
	if err != nil {
		return userpb.NewUserClient(conn)
	}
	return nil
}
