package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"h-ui/model/constant"
	"h-ui/model/vo"
	"net/http"
	"regexp"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
	_ = validate.RegisterValidation("validateStr", validateStr)
}

func validateStr(f validator.FieldLevel) bool {
	field := f.Field().String()
	// 字符串必须是字母和数字的组合
	reg := "^[A-Za-z0-9]+$"
	compile := regexp.MustCompile(reg)
	return field == "" || compile.MatchString(field)
}

func validateField[T interface{}](c *gin.Context, field T) (T, error) {
	if c.Request.Method == http.MethodGet {
		_ = c.ShouldBindQuery(&field)
	} else if c.Request.Method == http.MethodPost ||
		c.Request.Method == http.MethodPut ||
		c.Request.Method == http.MethodDelete {
		_ = c.ShouldBindJSON(&field)
	}
	if err := validate.Struct(&field); err != nil {
		vo.Fail(constant.InvalidError, c)
		return field, fmt.Errorf(constant.InvalidError)
	}
	return field, nil
}
