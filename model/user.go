package model

import "time"

type User struct {
	Gender     int8      `db:"gender"`
	UserID     int64     `db:"user_id"`
	Username   string    `db:"username"`
	Password   string    `db:"password"`
	Email      string    `db:"email"`
	Avatar     string    `db:"avatar"`
	Intro      string    `db:"intro"`
	CreateTime time.Time `db:"create_time"`
	UpdateTime time.Time `db:"update_time"`
}
