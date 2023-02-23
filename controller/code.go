package controller

type RespCode int32

const (
	CodeSuccess RespCode = 1000 + iota
	CodeServiceBusy
	CodeInvalidParams
	CodeUserExist
	CodePasswordErr
	CodeNeedLogin
	CodeInvalidToken
)

var CodeMapMsg = map[RespCode]string{
	CodeSuccess:       "成功",
	CodeServiceBusy:   "服务繁忙",
	CodeInvalidParams: "参数不足或错误",
	CodeUserExist:     "用户存在",
	CodePasswordErr:   "用户名或密码错误",
	CodeNeedLogin:     "需要登录",
	CodeInvalidToken:  "无效的token",
}

func (code RespCode) getMapping() string {
	return CodeMapMsg[code]
}
