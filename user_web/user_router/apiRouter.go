package user_router

import (
	"github.com/gin-gonic/gin"
	"py_gomall/v2/user_web/user_api"
)

func ApiRouter(r *gin.Engine) {
	ag := r.Group("/api/user")
	{
		ag.GET("/ping", user_api.Ping)
		// 点击验证码
		ag.GET("/captcha", user_api.Captcha)
		// 验证点击验证码
		ag.POST("/captcha/check", user_api.CaptchaCheck)
		// 注册
		ag.POST("/signup", user_api.Signup)
		// 登录
		ag.POST("/login", user_api.Login)
	}
}
