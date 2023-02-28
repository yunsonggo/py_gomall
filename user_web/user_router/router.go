package user_router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"py_gomall/v2/user_web/user_common/user_logger"
	"py_gomall/v2/user_web/user_common/user_translator"
	"py_gomall/v2/user_web/user_config"
	"py_gomall/v2/user_web/user_middleware"
)

func NewRouter() *gin.Engine {
	// 加载配置
	user_config.ParseNacosConf()
	// 初始化日志
	user_logger.NewLogger()
	// 初始化校验器

	// 初始化校验翻译器
	if err := user_translator.NewTranslator("en"); err != nil {
		log.Fatal(err)
	}
	r := gin.New()
	if user_config.AppConf.Mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r.Use(user_middleware.CorsMiddleware())
	r.StaticFS("/api/user/static", http.Dir("./user_static/public"))
	ConsulRouter(r)
	ApiRouter(r)
	AuthRouter(r)
	r.NoRoute(func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{
			"msg": "not found",
		})
	})
	return r
}
