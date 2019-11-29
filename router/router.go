package router

import (
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/internal/controllers/api"
	"github.com/haleyrom/trade/pkg/middleware"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.HTTPInterceptor())

	registerRouter(r)
	return r
}

// registerRouter 注册路由
func registerRouter(r *gin.Engine) {
	v1 := r.Group("/api")
	{
		v1.GET("/index", api.CreateTeam)
	}
}
