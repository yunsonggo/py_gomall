package user_router

import (
	"github.com/gin-gonic/gin"
	"py_gomall/v2/user_web/user_auth"
	"py_gomall/v2/user_web/user_middleware"
)

func AuthRouter(r *gin.Engine) {
	aag := r.Group("/api/user/auth")
	aag.Use(user_middleware.Auth())
	{
		aag.GET("/ping", user_auth.Ping)
		// 用户列表
		aag.GET("/list", user_auth.List)
	}
}
