package user_api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"py_gomall/v2/user_web/user_common/user_msg"
)

func Ping(ctx *gin.Context) {
	ip := ctx.ClientIP()
	user_msg.Success(ctx, codes.OK, "pong", ip)
}
