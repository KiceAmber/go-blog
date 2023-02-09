package controller

type RespCode int32

const (
	CodeSuccess RespCode = 1000 + iota
	CodeInvalidParams
	CodeUserExist
)

var CodeMapMsg = map[RespCode]string{
	CodeSuccess:       "成功",
	CodeInvalidParams: "参数不足或错误",
	CodeUserExist:     "用户存在",
}

func (code RespCode) getMapping() string {
	return CodeMapMsg[code]
}
