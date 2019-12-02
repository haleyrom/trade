package models

import (
	"github.com/haleyrom/trade/internal/params"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Users 用户
type Users struct {
	Uid        bson.ObjectId `json:"uid" bson:"_id"`                 // 用户id
	Name       string        `json:"name" bson:"name"`               // 用户名称
	Email      string        `json:"email" bson:"email"`             // 邮箱
	HeadImg    string        `json:"headimg" bson:"headimg"`         // 头像
	Status     int8          `json:"status" bson:"status"`           // 状态 0：正常 1：暂停
	CreateTime int           `json:"create_time" bson:"create_time"` // 创建时间
	ModifyTime int           `json:"modify_time" bson:"modify_time"` // 更新时间
}

// GetTable GetTable
func (u *Users) GetTable() string {
	return "users"
}

// AssignUsers 赋值用户
func AssignUsers(p *params.BaseParam) Users {
	timer := int(time.Now().Unix())
	return Users{
		Uid:        bson.ObjectIdHex(p.ID),
		Name:       p.Name,
		Email:      "",
		HeadImg:    "",
		Status:     0,
		CreateTime: timer,
		ModifyTime: timer,
	}
}
