package model

type User struct {
	Gender   int8   `db:"gender"`
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Email    string `db:"email"`
	Avatar   string `db:"avatar"`
	Intro    string `db:"intro"`
	Token    string
}
