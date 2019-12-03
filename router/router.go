package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/core"
	_ "github.com/haleyrom/trade/docs"
	"github.com/haleyrom/trade/internal/controllers/api"
	"github.com/haleyrom/trade/pkg/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	registerSwagger(r)
	r.Use(middleware.HttpBindGResp(), middleware.HttpInterceptor())
	registerRouter(r)
	return r
}

// registerSwagger 注册swagger
func registerSwagger(r *gin.Engine) {
	url := ginSwagger.URL(fmt.Sprintf("http://localhost%s/swagger/doc.json", core.Conf.HttpPort))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

// registerRouter 注册路由
func registerRouter(r *gin.Engine) {
	v1 := r.Group("/api")
	{
		// team api
		v1.POST("/team/create", api.CreateTeam)
		v1.POST("/team/join", api.JoinTeam)
		v1.POST("/team/exit", api.ExitTeam)
		v1.POST("/team/dismiss", api.DismissTeam)
		v1.POST("/team/list", api.ReadListTeam)
		v1.POST("/team/info", api.ReadInfoTeam)
		v1.POST("/team/user_list", api.ReadUserListTeam)
		v1.POST("/team/user", api.ReadTeamUser)
		// project api
		v1.POST("/project/create", api.CreateProject)
		v1.POST("/project/join", api.JoinProject)
		v1.POST("/project/list", api.ReadListProject)
		v1.POST("/project/dismiss", api.DismissProject)
	}

}
