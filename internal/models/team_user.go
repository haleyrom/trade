package models

import (
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/internal/params"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// TeamUser 团队成员
type TeamUser struct {
	Id         bson.ObjectId `json:"tuid" bson:"_id"`                // 团队成员id
	Team       Teams         `json:"team" bson:"team"`               // 团队
	User       Users         `json:"user" bson:"user"`               // 用户
	Role       Roles         `json:"role" bson:"role"`               // 身份
	Type       int8          `json:"type" bson:"type"`               // 类型 0：正常 1：队长
	Status     int8          `json:"status" bson:"status"`           // 状态 0：正常 1：退出
	CreateTime int           `json:"create_time" bson:"create_time"` // 创建时间
	ModifyTime int           `json:"modify_time" bson:"modify_time"` // 更新时间
}

const (
	// teamUserTypePublic 普通用户
	TeamUserPublic int8 = 0 + iota
	// TeamUserTypeOwner 创建用户
	TeamUserTypeOwner
)

// GetTable GetTable
func (t *TeamUser) GetTable() string {
	return "team_user"
}

// NewTeamUser 初始化团队成员
func NewTeamUser() *TeamUser {
	return &TeamUser{}
}

// JoinTeamUser 加入团队成员
func (t *TeamUser) JoinTeamUser() error {
	t.Id = bson.NewObjectId()
	timer := int(time.Now().Unix())
	t.CreateTime, t.ModifyTime = timer, timer
	err := core.Orm.InsertAll(t.GetTable(), []interface{}{*t})
	return err
}

// IsExistJoinTeam 判断是否存在团队
func (t *TeamUser) IsExistJoinTeam(p *params.JoinTeamParam) error {
	query := bson.M{
		"team._id": bson.ObjectIdHex(p.Tid),
		"user._id": bson.ObjectIdHex(p.Claims.ID),
	}
	return core.Orm.One(t.GetTable(), query, t)
}
