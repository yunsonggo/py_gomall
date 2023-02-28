package user_ctx

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
)

func PyCode(ctx *gin.Context) (code codes.Code, detail string) {
	if c, isExist := ctx.Get("code"); isExist {
		if v, ok := c.(codes.Code); ok {
			code = v
		}
	}
	if d, isExist := ctx.Get("detail"); isExist {
		if v, ok := d.(string); ok {
			detail = v
		}
	}
	return
}
