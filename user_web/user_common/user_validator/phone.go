package user_validator

import "github.com/go-playground/validator/v10"

func Phone(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	return VerifyMobileFormat(phone)
}
