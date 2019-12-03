package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/internal/params"
	"github.com/haleyrom/trade/internal/resp"
	"github.com/haleyrom/trade/pkg/jwt"
	"github.com/pkg/errors"
)

var (
	// HttpHeadToken http请求包头的token数据
	HttpHeadToken string = "token"
)

// HttpInterceptor 拦截器
func HttpInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		if token := c.Request.Header.Get(HttpHeadToken); len(token) > core.DefaultNilNum {
			j := jwt.NewJWT()
			claims := &jwt.CustomClaims{}
			// parseToken 解析token包含的信息
			if claims, err = j.ParseToken(token); err == nil {
				info := core.UserInfoPool.Get().(*params.BaseParam)
				info.ID = claims.ID
				info.Name = claims.Name
				info.Mobile = claims.Mobile
				info.StandardClaims = claims.StandardClaims
				core.UserInfoPool.Put(info)
			}
		} else {
			err = errors.Errorf("%d", resp.CodeNoToken)
		}

		switch err {
		case nil:
			c.Next()
		case jwt.TokenExpired:
			fallthrough
		case jwt.TokenNotValidYet:
			fallthrough
		case jwt.TokenMalformed:
			fallthrough
		case jwt.TokenInvalid:
			err = errors.Errorf("%d", resp.CodeIllegalToken)
			fallthrough
		default:
			core.GResp.Failure(err)
			c.Abort()
			return
		}
	}
}

// HttpBindGResp HttpBindGResp
func HttpBindGResp() gin.HandlerFunc {
	return func(c *gin.Context) {
		core.GResp = &resp.Resp{
			c,
		}
	}
}
