package user_api

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"py_gomall/v2/user_web/user_common/user_captcha"
	"py_gomall/v2/user_web/user_common/user_msg"
)

func Captcha(ctx *gin.Context) {
	data, err := user_captcha.GenBigCaptcha()
	if err != nil {
		user_msg.Failed(ctx, codes.Internal, "get captcha error", err)
		return
	}
	resp := user_captcha.GoCaptchaResponse{
		Base64:      data.Base64,
		ThumbBase64: data.ThumbBase64,
		Key:         data.Key,
	}
	user_msg.Success(ctx, codes.OK, "ok", resp)
}

func CaptchaCheck(ctx *gin.Context) {
	var request user_captcha.CheckCaptchaRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		user_msg.Failed(ctx, codes.InvalidArgument, "param error", err)
		return
	}
	ok, err := user_captcha.VerifyCaptcha(&request)
	if !ok || err != nil {
		user_msg.Failed(ctx, codes.InvalidArgument, "checked failed", err)
		return
	}
	user_msg.Success(ctx, codes.OK, "ok", nil)
	return
}
