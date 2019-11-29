package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/pkg/jwt"
	"net/http"
)

var (
	// HttpHeadToken http请求包头的token数据
	HttpHeadToken string = "token"
)

// HTTPInterceptor 拦截器
func HTTPInterceptor() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		if token := c.Request.Header.Get(HttpHeadToken); len(token) > core.DefaultNilNum {
			j := jwt.NewJWT()
			// parseToken 解析token包含的信息
			if claims, err := j.ParseToken(token); err == nil {
				c.Set("claims", claims)
			}
		} else {
			err = fmt.Errorf("请求未携带token，无权限访问")
		}

		switch err {
		case nil:
			c.Next()
		case jwt.TokenExpired:
			err = fmt.Errorf("授权已过期")
			fallthrough
		default:
			c.JSON(http.StatusOK, gin.H{
				"status":  -1,
				"message": err,
			})
			c.Abort()
			return
		}
	}

}
