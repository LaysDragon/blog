package web

import (
	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ValidateEnum(fl validator.FieldLevel) bool {
	enum := fl.Field().Interface().(domain.Enum)
	return enum.IsValid()
}

func GetValidator() *validator.Validate {
	return binding.Validator.Engine().(*validator.Validate)
}

func SetupValidation(validator *validator.Validate) {
	validator.RegisterValidation("enum", ValidateEnum)
}
