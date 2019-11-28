package storage

import "github.com/haleyrom/trade/pkg/config"

// Storage 存储相关接口
type Storage interface {
	Init(conf config.Configure) error
	InsertAll(table string, data []interface{}) error
	Clone()
}
