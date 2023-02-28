package user_msg

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"net/http"
	"py_gomall/v2/user_web/user_common/user_translator"
	"strings"
)

func Success(ctx *gin.Context, code codes.Code, message string, data interface{}) {
	resMsg := ""
	if message == "" {
		resMsg = code.String()
	} else {
		resMsg = message
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  resMsg,
		"data": data,
	})
	return
}

func Failed(ctx *gin.Context, code codes.Code, message string, err error) {
	resMsg := ""
	if message == "" {
		resMsg = code.String()
	} else {
		resMsg = message
	}
	if err == nil {
		err = errors.New("")
	}
	// 是否验证错误类型
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   resMsg,
			"error": err.Error(),
		})
	} else {
		resultErr := make(map[string]string)
		if user_translator.TS != nil {
			resultErr = RemoveStructFields(errs.Translate(user_translator.TS))
		} else {
			for i := 0; i < len(errs); i++ {
				fe := errs[i]
				resultErr[fe.Namespace()] = fe.Error()
			}
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code":  code,
			"msg":   resMsg,
			"error": resultErr,
		})
	}
	return
}

// RemoveStructFields 剔除验证字段结构体名字
func RemoveStructFields(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}
