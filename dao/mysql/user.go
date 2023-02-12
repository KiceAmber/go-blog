package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"go-blog/model"
)

const secret = "Amber"

// encryptPwd 密码加密
func encryptPwd(password string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(password)))
}

func CheckUserExist(username string) (err error) {
	sql := "select username, password from user where username = ?"
	rows, err := db.Query(sql, username)
	if err != nil {
		return
	}
	defer rows.Close()
	return
}

func CreateUser(user *model.User) (err error) {
	// 密码加密
	user.Password = encryptPwd(user.Password)
	sql := "insert into user(user_id, username, password) values(?, ?, ?)"
	_, err = db.Exec(sql, user.UserID, user.Username, user.Password)
	return
}

func QueryUser(username string, password string) (user *model.User, err error) {
	user = new(model.User)
	password = encryptPwd(password)
	sql := "select user_id, username, password from user where username = ? and password = ?"
	err = db.Get(user, sql, username, password)
	return
}
