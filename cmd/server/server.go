package main

import (
	cmd "github.com/haleyrom/trade/cmd/core"
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/pkg/middleware"
	"github.com/haleyrom/trade/pkg/version"
	"github.com/haleyrom/trade/router"
	"github.com/sirupsen/logrus"
)

// init init
func init() {
	cmd.Init()
}

// main main
func main() {
	// 注册路由
	r := router.InitRouter()
	// 日志中间件
	r.Use(middleware.LoggerToFile(core.Conf))
	_, out := middleware.OpenLoggerFile(core.Conf)
	// 日志文件落地
	logrus.SetOutput(out)
	version.LogAppInfo()

	port := ":8080"

	if len(core.Conf.HttpPort) > 0 {
		port = core.Conf.HttpPort
	}

	// Listen and Server in 0.0.0.0:8080
	r.Run(port)
}
