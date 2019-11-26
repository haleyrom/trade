package router

import "github.com/gin-gonic/gin"

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()

	registerRouter(r)
	return r
}

// registerRouter 注册路由
func registerRouter(r *gin.Engine) {
	// TODO: 待定
}
