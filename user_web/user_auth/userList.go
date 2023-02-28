package user_auth

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"py_gomall/v2/user_web/user_common/user_msg"
	"py_gomall/v2/user_web/user_dial"
	"py_gomall/v2/user_web/user_params"
	userpb "py_gomall/v2/user_web/user_proto/user_proto_gen"
	"strconv"
)

func List(ctx *gin.Context) {
	pageStr := ctx.DefaultQuery("page", "1")
	sizeStr := ctx.DefaultQuery("size", "5")
	page, err := strconv.Atoi(pageStr)
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		user_msg.Failed(ctx, codes.InvalidArgument, "param error", err)
		return
	}
	uc := user_dial.NewUserClient()
	resp, err := uc.UserList(ctx, &userpb.PageRequest{
		Page: int32(page),
		Size: int32(size),
	})
	if err != nil {
		user_msg.Failed(ctx, status.Code(err), "error", err)
		return
	}
	var users user_params.Users
	users.Total = resp.Total
	var data []user_params.User
	for _, info := range resp.Data {
		u := user_params.UserRespToParam(info)
		data = append(data, u)
	}
	users.Data = data
	user_msg.Success(ctx, codes.OK, "ok", users)
	return
}
