package vo

import (
	"github.com/gin-gonic/gin"
	"h-ui/model/constant"
	"net/http"
)

type result struct {
	Code    int         `json:"code"`
	Type    string      `json:"type"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	TypeSuccess = "ok"
	TypeError   = "no"
)

func Success(data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, result{
		Code: constant.CodeSuccess,
		Type: TypeSuccess,
		Data: data,
	})
}

func Fail(message string, c *gin.Context) {
	var code int
	if constant.UnauthorizedError == message {
		code = constant.CodeUnauthorizedError
	} else if constant.ForbiddenError == message {
		code = constant.CodeForbiddenError
	} else if constant.InvalidError == message {
		code = constant.CodeInvalidError
	} else {
		code = constant.CodeSysError
	}
	c.JSON(http.StatusOK, result{
		Code:    code,
		Type:    TypeError,
		Message: message,
		Data:    nil,
	})
}
