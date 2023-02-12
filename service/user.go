package service

import (
	"go-blog/api"
	"go-blog/dao/mysql"
	"go-blog/model"
	"go-blog/pkg/jwt"
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

func Login(param *api.UserLogin) (user *model.User, err error) {
	// 查询是否用户名与密码匹配
	user, err = mysql.QueryUser(param.Username, param.Password)
	if err != nil {
		return
	}

	// 如果查询出了用户，生成token
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token

	return
}
