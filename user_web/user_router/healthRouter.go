package user_router

import (
	"github.com/gin-gonic/gin"
	"py_gomall/v2/user_web/consul_api"
)

func ConsulRouter(r *gin.Engine) {
	hr := r.Group("/api")
	{
		// consul 健康检查
		hr.GET("/consul/health", consul_api.Health)
	}
}
