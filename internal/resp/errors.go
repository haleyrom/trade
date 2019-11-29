package resp

import (
	"net/http"
	"sync"
)

var (
	// okDataPool http响应成功数据池化
	okDataPool *sync.Pool

	// statusCodeMsgs 错误代码消息
	statusCodeMsgs map[StatusCode]string

	// emptyData 空数据
	emptyData = &struct{}{}
)

const (
	// CodeUnknow 未知
	CodeUnknow StatusCode = -1

	// CodeOk 请求响应
	CodeOk StatusCode = http.StatusOK

	// CodeInternalServerError 内部服务出错
	CodeInternalServerError StatusCode = http.StatusInternalServerError

	// CodeNoToken 请求参数必需要有token
	CodeNoToken StatusCode = 1001

	// CodeIllegalToken token不合法
	CodeIllegalToken StatusCode = 1002
)

// StatusCode 状态码
type StatusCode int

// Error 实现error接口
func (c StatusCode) Error() string {
	if msg, ok := statusCodeMsgs[c]; ok {
		return msg
	}
	return statusCodeMsgs[CodeUnknow]
}

func init() {
	okDataPool = &sync.Pool{
		New: func() interface{} {
			return &ResData{
				Code: int(CodeOk),
				Msg:  "ok",
			}
		},
	}

	statusCodeMsgs = map[StatusCode]string{
		CodeUnknow:              "unknow status",
		CodeOk:                  "请求成功",
		CodeInternalServerError: "服务繁忙,请稍后！",
		CodeNoToken:             "请求参数必需要有token",
		CodeIllegalToken:        "token不合法",
	}
}

// ResData http响应数据封包
type ResData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// EmptyData 空数据
func EmptyData() *struct{} {
	return emptyData
}

// Ok 返回成功响应数据封包
func Ok(data interface{}) *ResData {
	resData := okDataPool.Get().(*ResData)
	resData.Data = data
	return resData
}

// RecycleOk 回收Ok响应数据封包
func RecycleOk(data *ResData) {
	data.Data = nil // Notice:一定要赋nil避免泄内存
	okDataPool.Put(data)
}

// ErrCodeMsg 返回错误消息响应数据
func ErrCodeMsg(code StatusCode, msg ...string) *ResData {
	var errMsg string
	var exist bool
	if len(msg) > 0 {
		errMsg = msg[0]
	} else if errMsg, exist = statusCodeMsgs[code]; !exist {
		errMsg = statusCodeMsgs[CodeInternalServerError]
	}
	return &ResData{
		Code: int(code),
		Msg:  errMsg,
		Data: emptyData,
	}
}
