package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/pkg/config"
	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

// OpenLoggerFile 打开资源文件
func OpenLoggerFile(c config.Configure) (string, io.Writer) {
	//日志文件
	fileName := path.Join(c.Logs.Path, fmt.Sprintf("%s%s", c.Logs.Name, ".log"))
	//写入文件logs
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	if err != nil {
		fmt.Println("写入日志文件错误：", err)
	}

	return path.Join(c.Logs.Path, c.Logs.Suffix), src
}

// 日志记录到文件
func LoggerToFile(c config.Configure) gin.HandlerFunc {
	//实例化
	logger := logrus.New()
	var apiLogPath string
	//设置输出
	apiLogPath, logger.Out = OpenLoggerFile(c)

	//设置日志级别
	if c.RunMode == "dev" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}

	logWriter, _ := rotatelogs.New(
		apiLogPath+".%Y-%m-%d-%H-%M.log",
		rotatelogs.WithLinkName(apiLogPath),       // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.TextFormatter{})
	logger.AddHook(lfHook)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}
