package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-blog/api"
	"go-blog/service"
	"go.uber.org/zap"
)

func Signup(c *gin.Context) {
	var param = new(api.UserSignup)
	if err := c.ShouldBindJSON(&param); err != nil {
		// 表单数据规则校验
		zap.L().Error("Signup with invalid params...", zap.Error(err))
		RespError(c, CodeInvalidParams, nil)
		return
	}
	// 校验没问题，逻辑业务判断
	if err := service.Signup(param); err != nil {
		zap.L().Error("service.Signup() failed", zap.Error(err))
		RespError(c, CodeUserExist, nil)
		return
	}
	// 返回 JSON
	RespSuccess(c, nil)
}

func Login(c *gin.Context) {
	var param = new(api.UserLogin)
	if err := c.ShouldBindJSON(param); err != nil {
		zap.L().Error("Signup with invalid params...", zap.Error(err))
		RespError(c, CodeInvalidParams, nil)
		return
	}
	// 业务判断
	user, err := service.Login(param)
	if err != nil {
		zap.L().Error("service.Login() failed", zap.Error(err))
		RespError(c, CodePasswordErr, nil)
		return
	}
	// 返回JSON
	RespSuccess(c, gin.H{
		"user_id":  fmt.Sprintf("%d", user.UserID),
		"username": user.Username,
	})
}
