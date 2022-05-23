package ERRTell

const (
	ErrorServerFile = 5001
)

var codeMsg = map[int]string{
	ErrorServerFile: "服务端返回了一个空的文件",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
