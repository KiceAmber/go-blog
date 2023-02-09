package model

import "time"

type User struct {
	Gender     int8
	UserID     int64
	Username   string
	Password   string
	Email      string
	Avatar     string
	Intro      string
	CreateTime time.Time
	UpdateTime time.Time
}
