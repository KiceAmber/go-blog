package controller

type RespCode int32

const (
	CodeSuccess RespCode = 1000 + iota
)

var CodeMapMsg = map[RespCode]string{
	CodeSuccess: "成功",
}

func (code RespCode) GetMapping() string {
	return CodeMapMsg[code]
}
