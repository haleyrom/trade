package core

import (
	"github.com/haleyrom/trade/internal/params"
	"github.com/haleyrom/trade/internal/resp"
	"github.com/haleyrom/trade/pkg/config"
	"github.com/haleyrom/trade/pkg/storage"
	"sync"
)

var (
	// Conf 配置
	Conf config.Configure
	// Orm 数据
	Orm storage.MongoClient
	// GResp 返回
	GResp *resp.Resp
	// UserInfo 用户信息
	UserInfoPool *sync.Pool

	// DefaultNilString DefString
	DefaultNilString string = ""

	// DefaultNilNum DefaultNilNum
	DefaultNilNum int = 0
)

func init() {
	GResp = new(resp.Resp)
	// 用户信息磁化
	UserInfoPool = &sync.Pool{
		New: func() interface{} {
			return &params.BaseParam{}
		},
	}
}
