package core

import (
	"github.com/haleyrom/trade/internal/resp"
	"github.com/haleyrom/trade/pkg/config"
	"github.com/haleyrom/trade/pkg/storage"
)

var (
	// Conf 配置
	Conf config.Configure
	// Orm 数据
	Orm storage.MongoClient
	// GRespPool
	GResp *resp.Resp

	// DefaultNilString DefString
	DefaultNilString string = ""

	// DefaultNilNum DefaultNilNum
	DefaultNilNum int = 0
)

func init() {
	GResp = new(resp.Resp)
}
