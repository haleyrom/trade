package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/internal/controllers/api"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	registerRouter(r)
	return r
}

// registerRouter 注册路由
func registerRouter(r *gin.Engine) {
	v1 := r.Group("/api")
	{
		v1.GET("/index", api.Index)
	}
}
