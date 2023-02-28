package go_consul

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"strings"
)

type ServerInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
	Conn *grpc.ClientConn
}

type ServerItem struct {
	Address            string
	Port               int
	Name               string
	Tage               []string
	ID                 string
	HealthURL          string
	TimeoutSec         string
	IntervalSec        string
	DeregisterAfterSec string
}

type ConsulClient struct {
	Client *api.Client
}

func NewConsulClient(host string, port int) (*ConsulClient, error) {
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", host, port)
	client, err := api.NewClient(cfg)
	if err != nil {
		zap.L().Error("连接consul服务错误", zap.Error(err))
		return nil, err
	}
	c := &ConsulClient{
		Client: client,
	}
	return c, nil
}

func (cc *ConsulClient) Register(item ServerItem) error {
	fmt.Println(item.HealthURL)
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		Name:                           item.Name,
		HTTP:                           item.HealthURL,
		Timeout:                        item.TimeoutSec,
		Interval:                       item.IntervalSec,
		DeregisterCriticalServiceAfter: item.DeregisterAfterSec,
	}
	// 生成注册对象
	registration := &api.AgentServiceRegistration{
		ID:                item.ID,
		Name:              item.Name,
		Tags:              item.Tage,
		Port:              item.Port,
		Address:           item.Address,
		EnableTagOverride: false,
		Check:             check,
	}
	err := cc.Client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.L().Error("consul注册服务成员错误", zap.Error(err))
	}
	return err
}

func (cc *ConsulClient) Deregister(itemID string) error {
	err := cc.Client.Agent().ServiceDeregister(itemID)
	if err != nil {
		zap.L().Error("consul注销服务成员错误", zap.Error(err))
	}
	return err
}

// 使用随机端口
func (cc *ConsulClient) Services(configServices []string) (map[string]*ServerInfo, error) {
	services := make(map[string]*ServerInfo)
	for _, info := range configServices {
		info = strings.TrimSpace(info)
		infos := strings.Split(info, ",")
		s, err := cc.GetServer(infos[0])
		if err != nil {
			return nil, err
		}
		services[s.Name] = s
	}
	return services, nil
}

// 固定配置服务端口
//func (cc *ConsulClient) Services(configServices []string) (map[string]*ServerInfo, error) {
//	services := make(map[string]*ServerInfo)
//	for _, info := range configServices {
//		info = strings.TrimSpace(info)
//		infos := strings.Split(info, ",")
//		var s ServerInfo
//		for i := 0; i < len(infos); i++ {
//			switch i {
//			case 0:
//				s.Name = infos[i]
//			case 1:
//				s.Host = infos[i]
//			case 2:
//				port, err := strconv.Atoi(infos[i])
//				if err != nil {
//					zap.L().Debug("RegistrationServices error", zap.Error(err))
//					return nil, err
//				}
//				s.Port = port
//			}
//		}
//		addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
//		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
//		if err != nil {
//			zap.L().Debug("dial: 连接grpc服务错误", zap.Error(err))
//			break
//		}
//		s.Conn = conn
//		services[s.Name] = &s
//	}
//	return services, nil
//}

func (cc *ConsulClient) GetServer(name string) (*ServerInfo, error) {
	data, err := cc.Client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, name))
	if err != nil {
		return nil, err
	}
	var s ServerInfo
	for _, v := range data {
		s = ServerInfo{
			ID:   v.ID,
			Name: v.Service,
			Host: v.Address,
			Port: v.Port,
		}
		addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()))
		if err != nil {
			zap.L().Debug("dial: 连接grpc服务错误", zap.Error(err))
			return nil, err
		}
		s.Conn = conn
	}
	return &s, nil
}
