package core

import (
	"github.com/haleyrom/trade/pkg/config"
	"github.com/haleyrom/trade/pkg/storage"
)

var (
	// Conf 配置
	Conf config.Configure
	// Orm 数据
	Orm storage.MongoClient

	// DefaultNilString DefString
	DefaultNilString string = ""
	// DefaultNilNum DefaultNilNum
	DefaultNilNum int = 0
)
