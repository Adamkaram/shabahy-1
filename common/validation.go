package common

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"strings"
)

type GetError func(err validator.ValidationErrors) string

func ValidateErrors(requestError error) string {
	return validate(requestError.(validator.ValidationErrors))
}

func validate(errors validator.ValidationErrors) string {
	resultErrors := ""
	for _, err := range errors {
		switch err.Tag() {
		case "required":
			resultErrors += err.Field() + " is required\n "
		case "email":
			resultErrors += err.Field() + " must me valid email\n "
		case "min":
			resultErrors += err.Field() + " must be " + err.Param() + " length at least\n"
		case "Enum":
			replacer := *strings.NewReplacer("_", ",")
			resultErrors += err.Field() + " must be one of " + replacer.Replace(err.Param())


		default:
			resultErrors += "error in filed " + err.Tag()
		}
	}
	return resultErrors
}

func ValidateNotFound(err error, errorString string) string {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errorString
	}
	return err.Error()
}

func Enum(
	fl validator.FieldLevel,
) bool {
	enumString := fl.Param()
	value := fl.Field().String()
	enumSlice := strings.Split(enumString, "_")
	for _, v := range enumSlice {
		if value == v {
			return true
		}
	}
	return false
}
