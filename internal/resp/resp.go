package resp

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// Resp Resp
type Resp struct {
	*gin.Context
}

// Success 获取数据成功后的json数据响应
func (r *Resp) Success(data interface{}) {
	okData := Ok(data)
	defer RecycleOk(okData)
	r.JSON(http.StatusOK, okData)
}

// Failure 获取数据失败后的json数据响应
func (r *Resp) Failure(err error) {
	var (
		data *ResData
		code StatusCode
	)

	if ok := errors.As(err, &code); ok {
		data = ErrCodeMsg(code)
	} else {
		data = ErrCodeMsg(CodeInternalServerError)
		if code, _ := strconv.Atoi(err.Error()); code > 0 {
			data = ErrCodeMsg(StatusCode(code))
		}
	}

	logrus.Error(err)
	r.JSON(http.StatusInternalServerError, data)
}
