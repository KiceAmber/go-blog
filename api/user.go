package api

// UserSignup 用户注册接口模型
type UserSignup struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
}
