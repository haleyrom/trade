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

// InsertAll 插入
func (m *MongoClient) InsertAll(table string, data []interface{}) error {
	collection := m.Client.DB(m.Database).C(table)
	return collection.Insert(data...)
}

// Clone
func (m *MongoClient) Clone() {
	m.Client.Clone()
	return
}
