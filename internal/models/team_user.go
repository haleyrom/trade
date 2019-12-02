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
	TeamUserPublic int8 = 0
	// TeamUserTypeOwner 创建用户
	TeamUserTypeOwner int8 = 1

	// TeamUserStatusOnline 团队用户正常
	TeamUserStatusOnline int8 = 0
	// TeamUserStatusExit 退出
	TeamUserStatusExit int8 = 1
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
	var err error
	timer := int(time.Now().Unix())
	t.Status, t.ModifyTime = TeamUserStatusOnline, timer

	if len(t.Id) == core.DefaultNilNum {
		t.Id, t.CreateTime = bson.NewObjectId(), timer
		err = core.Orm.InsertAll(t.GetTable(), []interface{}{*t})
	} else {
		query := bson.M{"_id": t.Id}
		err = core.Orm.Update(t.GetTable(), query, t)
	}
	return err
}

// IsExistJoinTeam 判断是否存在团队
func (t *TeamUser) IsExistJoinTeam(tid, uid string) error {
	query := bson.M{
		"team._id": bson.ObjectIdHex(tid),
		"user._id": bson.ObjectIdHex(uid),
	}
	return core.Orm.One(t.GetTable(), query, t)
}

// ExitTeam 退出团队
func (t *TeamUser) ExitTeam(p *params.ExitTeamParam) error {
	update := bson.M{
		"$set": bson.M{
			"status":      TeamUserStatusExit,
			"modify_time": int(time.Now().Unix()),
		},
	}

	query := bson.M{
		"team._id": bson.ObjectIdHex(p.Tid),
		"user._id": bson.ObjectIdHex(p.Claims.ID),
	}
	return core.Orm.Update(t.GetTable(), query, update)
}
