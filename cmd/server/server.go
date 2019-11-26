package main

import (
	"github.com/haleyrom/trade/cmd"
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/pkg/middleware"
	"github.com/haleyrom/trade/router"
	"github.com/haleyrom/trade/version"
	"github.com/sirupsen/logrus"
)

func init() {
	cmd.Init()
}

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
