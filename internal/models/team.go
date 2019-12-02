package models

import (
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/internal/params"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Team 团队
type Teams struct {
	Id         bson.ObjectId `json:"tid" bson:"_id"`                 // 团队id
	Name       string        `json:"name" bson:"name"`               // 团队名称
	Creator    Users         `json:"creator" bson:"creator"`         // 团队创建人
	Status     int8          `json:"status" bson:"status"`           // 状态 0：正常 1：解散
	CreateTime int           `json:"create_time" bson:"create_time"` // 创建时间
	ModifyTime int           `json:"modify_time" bson:"modify_time"` // 更新时间
}

const (
	// TeamStatusDismiss 团队解散
	TeamStatusDismiss int8 = 1
)

// GetTable GetTable
func (t *Teams) GetTable() string {
	return "teams"
}

// NewTeam 初始化团队
func NewTeam() *Teams {
	return &Teams{}
}

// CreateTeams 创建团队
func (t *Teams) CreateTeam(p *params.CreateTeamParam) error {
	var err error
	timer := int(time.Now().Unix())
	t.Id, t.Name = bson.NewObjectId(), p.Name
	t.CreateTime, t.ModifyTime = timer, timer
	t.Creator = AssignUsers(p.Claims)

	if err = core.Orm.InsertAll(t.GetTable(), []interface{}{*t}); err == nil {
		teamUser := TeamUser{
			Team:   *t,
			User:   t.Creator,
			Role:   Roles{},
			Type:   TeamUserTypeOwner,
			Status: 0,
		}
		err = teamUser.JoinTeamUser()
	}
	return err
}

// IsExistTeam 判断是否存在团队
func (t *Teams) IsExistTeam(tid string) error {
	query := bson.M{
		"_id": bson.ObjectIdHex(tid),
	}
	return core.Orm.One(t.GetTable(), query, t)
}

// DismissTeam 解散团队
func (t *Teams) DismissTeam(tid string) error {
	var err error
	update := bson.M{
		"$set": bson.M{
			"status":      TeamStatusDismiss,
			"modify_time": int(time.Now().Unix()),
		},
	}
	query := bson.M{
		"_id": bson.ObjectIdHex(tid),
	}
	if err = core.Orm.Update(t.GetTable(), query, update); err == nil {
		err = NewTeamUser().DismissTeam(tid)
	}
	return err
}
