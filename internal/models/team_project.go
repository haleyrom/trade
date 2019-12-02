package models

import (
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/internal/params"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// TeamProject 团队项目
type TeamProject struct {
	Id         bson.ObjectId `json:"tpid" bson:"_id"`                 // 团队项目id
	Name       string        `json:"name"  bson:"name"`               // 项目名称
	Team       Teams         `json:"team"  bson:"team"`               // 团队
	Service    Service       `json:"service"  bson:"service"`         // 服务
	Status     int8          `json:"status"  bson:"status"`           // 状态 0：正常 1：关闭 2 : 删除
	CreateTime int           `json:"create_time"  bson:"create_time"` // 创建时间
	ModifyTime int           `json:"modify_time"  bson:"modify_time"` // 更新时间
}

const (
	// TeamProjectStatusPublic 团队项目正常状态
	TeamProjectStatusPublic int8 = 0
	// TeamProjectStatusClose 团队项目关闭状态
	TeamProjectStatusClose int8 = 1
	// TeamProjectStatusRm 团队项目删除
	TeamProjectStatusRm int8 = 2
)

// GetTable GetTable
func (t *TeamProject) GetTable() string {
	return "team_project"
}

// NewTeamProject 初始化团队项目
func NewTeamProject() *TeamProject {
	return &TeamProject{}
}

// CreateTeams 创建团队
func (t *TeamProject) CreateProject(p *params.CreateProjectParam) error {
	var err error
	timer := int(time.Now().Unix())
	t.Id, t.Name = bson.NewObjectId(), p.Name
	t.CreateTime, t.ModifyTime = timer, timer

	if err = core.Orm.InsertAll(t.GetTable(), []interface{}{*t}); err == nil {
		projectUser := ProjectUser{
			Project: *t,
			User:    AssignUsers(p.Claims),
		}
		err = projectUser.JoinTeamProject()
	}
	return err
}

// IsExistTeam 判断是否存在团队
func (t *TeamProject) IsExistTeam(pid string) error {
	query := bson.M{
		"_id":    bson.ObjectIdHex(pid),
		"status": TeamProjectStatusPublic,
	}
	return core.Orm.One(t.GetTable(), query, t)
}
