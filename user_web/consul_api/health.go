package consul_api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"py_gomall/v2/user_web/user_common/user_msg"
)

func Health(ctx *gin.Context) {
	user_msg.Success(ctx, codes.OK, "checked", "ok")
}
