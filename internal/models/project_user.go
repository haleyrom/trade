package models

import (
	"github.com/haleyrom/trade/core"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// ProjectUser 项目用户
type ProjectUser struct {
	Id         bson.ObjectId `json:"puid" bson:"_id"`                // 项目用户id
	Project    TeamProject   `json:"project" bson:"project"`         // 项目
	User       Users         `json:"user" bson:"user"`               // 用户
	Role       Roles         `json:"role" bson:"role"`               // 权限
	Status     int8          `json:"status" bson:"status"`           // 状态 0：正常 1：退出 2：解散
	CreateTime int           `json:"create_time" bson:"create_time"` // 创建时间
	ModifyTime int           `json:"modify_time" bson:"modify_time"` // 更新时间
}

const (
	// ProjectUserOnline 项目用户正常
	ProjectUserOnline int8 = 0
	// ProjectUserExit 退出
	ProjectUserExit int8 = 1
	// ProjectUserStatusOnline  项目状态正常
	ProjectUserStatusOnline int8 = 0
	// ProjectUserStatusExit  退出项目
	ProjectUserStatusExit int8 = 1
	// ProjectUserStatusDismiss 解散
	ProjectUserStatusDismiss int8 = 2
)

// GetTable GetTable
func (p *ProjectUser) GetTable() string {
	return "project_user"
}

// NewProjectUser 初始化项目用户
func NewProjectUser() *ProjectUser {
	return &ProjectUser{}
}

// IsExistJoinTeam 判断是否存在团队
func (p *ProjectUser) IsExistJoinProject(pid, uid string) error {
	query := bson.M{
		"project._id": bson.ObjectIdHex(pid),
		"user._id":    bson.ObjectIdHex(uid),
		"status":      ProjectUserStatusOnline,
	}
	return core.Orm.One(p.GetTable(), query, p)
}

// JoinTeamProject 加入项目成员
func (p *ProjectUser) JoinTeamProject() error {
	var err error
	timer := int(time.Now().Unix())
	p.Status, p.ModifyTime = ProjectUserOnline, timer

	if len(p.Id) == core.DefaultNilNum {
		p.Id, p.CreateTime = bson.NewObjectId(), timer
		err = core.Orm.InsertAll(p.GetTable(), []interface{}{*p})
	} else {
		query := bson.M{"_id": p.Id}
		err = core.Orm.Update(p.GetTable(), query, p)
	}
	return err
}
