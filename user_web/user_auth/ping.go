package user_auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"py_gomall/v2/user_web/user_common/user_msg"
)

func Ping(ctx *gin.Context) {
	ip := ctx.ClientIP()
	ctxMobile, exists := ctx.Get("mobile")
	if exists {
		if mobile, ok := ctxMobile.(string); ok {
			ip = fmt.Sprintf("mobile:%v\nip:%s", mobile, ip)
		}
	}
	user_msg.Success(ctx, codes.OK, "pong", ip)
	return
}
