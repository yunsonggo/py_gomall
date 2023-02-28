package user_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"py_gomall/v2/user_web/user_common/user_ctx"
	"py_gomall/v2/user_web/user_common/user_msg"
	"py_gomall/v2/user_web/user_common/user_token"
	"py_gomall/v2/user_web/user_dial"
	"py_gomall/v2/user_web/user_params"
	userpb "py_gomall/v2/user_web/user_proto/user_proto_gen"
)

func Login(ctx *gin.Context) {
	var param user_params.LoginParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		user_msg.Failed(ctx, codes.InvalidArgument, "request param error,Phone number and password are required.", err)
		return
	}
	request := userpb.UserRequest{
		Mobile:   param.Mobile,
		Password: param.Password,
	}
	uc := user_dial.NewUserClient()
	resp, err := uc.UserCheckPasswd(ctx, &request)
	fmt.Printf("resp:%+v\n", resp)
	fmt.Printf("err:%+v\n", err)
	code, detail := user_ctx.PyCode(ctx)
	fmt.Printf("code:%+v\n", code)
	fmt.Printf("detail:%+v\n", detail)
	if err != nil || resp.IsChecked != userpb.CheckedMessage_CHECKED_YES {
		err = errors.New("password verification failed")
		user_msg.Failed(ctx, code, detail, err)
		return
	}
	token, err := user_token.Release(request.Mobile)
	if err != nil {
		user_msg.Failed(ctx, codes.Internal, "release token code error", err)
		return
	}
	user_msg.Success(ctx, codes.OK, "Bearer ok", token)
	return
}

func Signup(ctx *gin.Context) {
	var param user_params.SignupParam
	if err := ctx.ShouldBindJSON(&param); err != nil {
		user_msg.Failed(ctx, codes.InvalidArgument, "param error", err)
		return
	}
	request := user_params.SignupParamToReq(param)
	uc := user_dial.NewUserClient()
	resp, err := uc.UserCreate(ctx, request)
	if err != nil {
		code := codes.Internal
		if co, ok := status.FromError(err); ok {
			code = co.Code()
		}
		user_msg.Failed(ctx, code, "create user error", err)
		return
	}
	info := user_params.UserRespToParam(resp)
	user_msg.Success(ctx, codes.OK, "ok", info)
	return
}
