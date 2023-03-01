package user_run

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"py_gomall/v2/user_web/user_dial"
	"py_gomall/v2/user_web/user_dial/client_balancer"
	userpb "py_gomall/v2/user_web/user_proto/user_proto_gen"
	"syscall"
	"time"
)

var UserWebSrv *http.Server
var R *gin.Engine
var AppHost string

func Run(r *gin.Engine) {
	// 方式一: 手动注册服务并发现
	//err := user_dial.RegistrationServices()
	//if err != nil {
	//	log.Fatal(err)
	//}
	// 方式二: 使用consul api client 注册服务并使用 服务发现
	// 使用时代码再进一步组织 这里有点啰嗦 为了方便展示
	if err := user_dial.NewConsulClient(); err != nil {
		zap.L().Fatal("new consul client error", zap.Error(err))
	}
	if err := user_dial.Register(); err != nil {
		zap.L().Fatal("consul client register user_web error", zap.Error(err))
	}
	//if err := user_dial.ServiceList(); err != nil {
	//	zap.L().Fatal("consul client register user_web error", zap.Error(err))
	//}
	// 通过负载均衡库获取grpc User client
	// user_dial.UserBalancerConn = user_dial.InitSrvConn("gomall_user_srv")
	// grpc 拉取服务 负载均衡方式二
	userConn, err := client_balancer.NewUserBalancer("gomall_user_srv")
	if err != nil {
		log.Fatal(err)
	}
	user_dial.UserBalancerConn = userConn
	u := userpb.NewUserClient(user_dial.UserBalancerConn)
	list, err := u.UserList(context.Background(), &userpb.PageRequest{
		Page: 1,
		Size: 5,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(list)

	srv := &http.Server{
		Addr:    AppHost,
		Handler: r,
	}
	go func() {
		zap.L().Info("user_web listen", zap.String("on: ", AppHost))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("user_web run error", zap.Error(err))
		}
	}()
	UserWebSrv = srv
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Debug("Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// TODO:: close store's client or conn ...
		//for _, info := range user_dial.ServerInfoList {
		//	_ = info.Conn.Close()
		//}
		//_ = user_dial.Deregister()
		_ = user_dial.UserBalancerConn.Close()
		_ = zap.L().Sync()
		cancel()
	}()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("关闭服务错误", zap.Error(err))
	}
	zap.L().Debug("Shutdown server and OSS server success")
	return
}

func ReStart() {
	zap.L().Info("配置文件更新,自动重启服务...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if err := UserWebSrv.Shutdown(ctx); err != nil {
		zap.L().Fatal("关闭服务错误", zap.Error(err))
		cancel()
		return
	}
	for _, info := range user_dial.ServerInfoList {
		_ = info.Conn.Close()
	}
	_ = zap.L().Sync()
	cancel()
	for i := 5; i >= 0; i-- {
		fmt.Printf("重启倒计时: %d\n", i)
		time.Sleep(time.Second)
	}
	Run(R)
}
