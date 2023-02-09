package service

import (
	"go-blog/api"
	"go-blog/dao/mysql"
	"go-blog/model"
	"go-blog/pkg/snowflake"
)

func Signup(param *api.UserSignup) (err error) {

	// 验证用户是否存在
	if err = mysql.CheckUserExist(param.Username); err != nil {
		return
	}
	// 创建用户，雪花算法生成ID
	userID := snowflake.GenID()
	user := &model.User{
		UserID:   userID,
		Username: param.Username,
		Password: param.Password,
	}
	err = mysql.CreateUser(user)
	return
}
