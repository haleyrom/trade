package core

import (
	"flag"
	"fmt"
	"github.com/haleyrom/trade/core"
	"github.com/sirupsen/logrus"
)

// Init init
func Init() {
	InitConf()
	InitStorage()
}

// InitConf  初始化配置
func InitConf() {
	// 获取配置
	configFilePath := flag.String("C", "assets/config/conf.yaml", "config file path")
	if err := core.Conf.Init(*configFilePath); err != nil {
		logrus.Error("err parsing  config file:", err)
		panic(fmt.Errorf("err parsing  config file:", err))
	}
}

// InitStorage 初始化数据库
func InitStorage() {
	if err := core.Orm.Init(core.Conf); err != nil {
		logrus.Error("init orm client fail:", err)
		panic(fmt.Errorf("init orm client fail:", err))
	}
}
