package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespData struct {
	Code    RespCode    `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func RespError(c *gin.Context, code RespCode, data interface{}) {
	c.JSON(http.StatusOK, &RespData{
		Code:    code,
		Message: code.getMapping(),
		Data:    data,
	})
}

func RespSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, &RespData{
		Code:    CodeSuccess,
		Message: CodeSuccess.getMapping(),
		Data:    data,
	})
}
