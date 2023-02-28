package user_translator

import (
	"errors"
	"fmt"
	"py_gomall/v2/user_web/user_common/user_validator"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	vali "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var TS ut.Translator

func NewTranslator(locale string) (err error) {
	var trans ut.Translator
	v, ok := binding.Validator.Engine().(*vali.Validate)
	if !ok {
		err = errors.New("校验器类型错误")
		return
	}
	// 注册一个获取JSON tag的方法
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	zhT := zh.New()
	enT := en.New()
	// 第一个参数是备用（fallback）的语言环境
	// 后面的参数是应该支持的语言环境（支持多个）
	// uni := ut.New(zhT, zhT) 也是可以的
	uni := ut.New(enT, zhT, enT)
	// 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找
	trans, ok = uni.GetTranslator(locale)
	if !ok {
		return fmt.Errorf("获取翻译器(%s)错误", locale)
	}
	// 注册翻译器
	switch locale {
	case "en":
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	case "zh":
		err = zhTranslations.RegisterDefaultTranslations(v, trans)
	default:
		err = enTranslations.RegisterDefaultTranslations(v, trans)
	}
	TS = trans
	// 注册自定义验证器验证手机号
	if err = v.RegisterValidation("phone", user_validator.Phone); err != nil {
		return
	}
	err = v.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "{0} 手机号码格式错误!", true)
	}, func(ut ut.Translator, fe vali.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})
	return
}
