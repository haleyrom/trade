package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// LogsConf 日志配置
type LogsConf struct {
	Path   string `yaml:"path"`
	Name   string `yaml:"name"`
	Suffix string `yaml:"suffix"`
}

// MongoConf mongo配置
type MongoConf struct {
	Addr      string `yaml:"addr"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	PoolLimit int    `yaml:"limit"`
	Database  string `yaml:"database"`
}

// LoadConfig 加载配置
type Configure struct {
	AppName  string    `yaml:"appname"`
	HttpPort string    `yaml:"httpport"`
	RunMode  string    `yaml:"runmode"`
	Logs     LogsConf  `yaml:"logs"`
	Mongo    MongoConf `yaml:"mongo"`
}

// Init 初始化
func (c *Configure) Init(path string) error {
	var err error
	if data, err := ioutil.ReadFile(path); err == nil {
		err = yaml.Unmarshal(data, &c)
	}
	return err
}
