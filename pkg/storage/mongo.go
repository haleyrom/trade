package storage

import (
	"github.com/haleyrom/trade/pkg/config"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

// MongoClient mongo数据库
type MongoClient struct {
	Client   *mgo.Session
	Database string `json:"database"`
}

// Init Init
func (m *MongoClient) Init(c config.Configure) error {
	var err error
	conf := &mgo.DialInfo{
		Addrs:     []string{c.Mongo.Addr},
		Username:  c.Mongo.Username,
		Password:  c.Mongo.Password,
		PoolLimit: c.Mongo.PoolLimit,
	}
	// 连接服务器
	if m.Client, err = mgo.DialWithInfo(conf); err != nil {
		logrus.Errorf("mongo client link failure : %s", err)
		return err
	}

	m.Client.SetMode(mgo.Monotonic, true)
	m.Database = c.Mongo.Database
	// default is 4096
	logrus.Infof("Connected MongoDB!")
	return nil
}

// table 表
func (m *MongoClient) table(table string) *mgo.Collection {
	return m.Client.DB(m.Database).C(table)
}

// InsertAll 插入
func (m *MongoClient) InsertAll(table string, data []interface{}) error {
	return m.table(table).Insert(data...)
}

// One 单条查询
func (m *MongoClient) One(table string, condition map[string]interface{}, obj interface{}) error {
	return m.table(table).Find(condition).One(obj)
}

// Update 更新
func (m *MongoClient) Update(table string, condition map[string]interface{}, obj interface{}) error {
	return m.table(table).Update(condition, obj)
}

// Clone
func (m *MongoClient) Clone() {
	m.Client.Clone()
	return
}
