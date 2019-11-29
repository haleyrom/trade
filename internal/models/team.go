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

// GetTable GetTable
func (t *Teams) GetTable() string {
	return "teams"
}

// CreateTeams 创建团队
func (t *Teams) CreateTeam(param *params.CreateTeamParam) error {
	var err error
	timer := int(time.Now().Unix())
	t.Id, t.Name = bson.NewObjectId(), param.Name
	t.CreateTime, t.ModifyTime = timer, timer

	t.Creator = Users{
		Uid:        bson.ObjectIdHex(param.Claims.ID),
		Name:       param.Claims.Name,
		CreateTime: timer,
		ModifyTime: timer,
	}

	if err = core.Orm.InsertAll(t.GetTable(), []interface{}{*t}); err == nil {
		teamUser := TeamUser{
			Id:     bson.NewObjectId(),
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

// ExistTeam 判断是否存在团队
func (t *Teams) IsExistTeam(param *params.JoinTeamParam) error {

	return nil
}
