package user_middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"py_gomall/v2/user_web/user_common/user_msg"
	"py_gomall/v2/user_web/user_common/user_token"
	"py_gomall/v2/user_web/user_dial"
	userpb "py_gomall/v2/user_web/user_proto/user_proto_gen"
	"strings"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//var err error
		var token string
		var mobile string
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			user_msg.Failed(ctx, codes.Unauthenticated, "Please retry or login again", errors.New("header error"))
			ctx.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			user_msg.Failed(ctx, codes.Unauthenticated, "Please retry or login again", errors.New("token error"))
			ctx.Abort()
			return
		}
		token = parts[1]
		_, claims, tokenErr := user_token.Parse(token)
		if tokenErr != nil {
			user_msg.Failed(ctx, codes.Unauthenticated, "token expired", tokenErr)
			ctx.Abort()
			return
		} else {
			mobile = claims.Mobile
		}
		// TODO::调用用户服务
		//user, err := dial.SnakeServers[0].SC.FirstUser(ctx, &snakepb.IDRequest{IdString: userID})
		user, err := user_dial.NewUserClient().UserFirst(ctx, &userpb.IDRequest{StrID: mobile})

		if user.Id == 0 {
			user_msg.Failed(ctx, codes.Unauthenticated, "token error", err)
			ctx.Abort()
			return
		}
		ctx.Set("mobile", mobile)
		ctx.Set("token", token)
		ctx.Next()
	}
}
