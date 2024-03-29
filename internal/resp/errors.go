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

	// CodeAuth 暂无权限
	CodeAuth StatusCode = http.StatusUnauthorized

	// CodeInternalServerError 内部服务出错
	CodeInternalServerError StatusCode = http.StatusInternalServerError

	// CodeNoToken 请求参数必需要有token
	CodeNoToken StatusCode = 101101

	// CodeIllegalToken token不合法
	CodeIllegalToken StatusCode = 101102

	// CodeNotTeam 团队不存在
	CodeNotTeam StatusCode = 101103

	// CodeExistTeam 团队存在
	CodeExistTeam StatusCode = 101104

	// CodeNotUser 用户不存在
	CodeNotUser StatusCode = 101105

	// CodeExistUser 用户存在
	CodeExistUser StatusCode = 101106

	// CodeNotProject 项目不存在
	CodeNotProject StatusCode = 101107

	// CodeExistProject 项目存在
	CodeExistProject StatusCode = 101108
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
		CodeAuth:                "暂无权限",
		CodeNotTeam:             "该团队不存在/已解散",
		CodeExistTeam:           "已加入该团队,您可以直接进入",
		CodeNotUser:             "用户不存在",
		CodeExistUser:           "用户存在",
		CodeNotProject:          "该项目不存在/已解散",
		CodeExistProject:        "已加入该项目,您可以直接进入",
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
