package storage

import "github.com/haleyrom/trade/pkg/config"

// Storage 存储相关接口
type Storage interface {
	Init(conf config.Configure) error
	One(table string, condition map[string]interface{}, obj interface{}) error
	InsertAll(table string, data []interface{}) error
	Update(table string, condition map[string]interface{}, obj interface{}) error
	Count(table string, condition map[string]interface{}) (int, error)
	All(table string, condition map[string]interface{})
	Clone()
}
