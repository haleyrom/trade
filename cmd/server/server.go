package main

import (
	"context"
	"fmt"
	cmd "github.com/haleyrom/trade/cmd/core"
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/pkg/middleware"
	"github.com/haleyrom/trade/pkg/version"
	"github.com/haleyrom/trade/router"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var srv *http.Server

// init init
func init() {
	cmd.Init()
}

// main main
// @title 测试
// @version 0.0.1
// @description  测试
// @BasePath /api/v1
func main() {
	// 注册路由
	r := router.InitRouter()
	// 日志中间件
	r.Use(middleware.LoggerToFile(core.Conf))
	_, out := middleware.OpenLoggerFile(core.Conf)
	// 日志文件落地
	logrus.SetOutput(out)
	// 版本信息
	version.LogAppInfo()

	srv = &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	if len(core.Conf.HttpPort) > core.DefaultNilNum {
		srv.Addr = core.Conf.HttpPort
	}

	defer func() {
		clone()
	}()

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(fmt.Sprintf("listen: %s\n", err))
		}
	}()
	fmt.Printf("Listening and serving HTTP on %s\n", srv.Addr)

}

// clone 退出
func clone() {
	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
		core.Orm.Clone()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server Shutdown:", err)
	}
	logrus.Println("Server exiting")
}
