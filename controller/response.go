package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespData struct {
	Code    RespCode
	Message string
	Data    interface{}
}

func RespError(c *gin.Context, code RespCode, data interface{}) {
	c.JSON(http.StatusOK, RespData{
		Code:    code,
		Message: code.getMapping(),
		Data:    data,
	})
}

func RespSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, RespData{
		Code:    CodeSuccess,
		Message: CodeSuccess.getMapping(),
		Data:    data,
	})
}
