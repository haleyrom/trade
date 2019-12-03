package models

import (
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/internal/resp"
	"gopkg.in/mgo.v2"
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

// PageProject 读取项目信息
func (p *ProjectUser) PageProject(uid string, start, end int) (resp.PageResp, error) {
	var (
		items []TeamProject
		err   error
	)
	data := resp.PageResp{
		Items: make([]Teams, 0),
		Page: resp.PageInfoResp{
			PageSize: end,
		},
	}

	query := []bson.M{
		{"$match": bson.M{
			"user._id": bson.ObjectIdHex(uid),
			"status":   ProjectUserOnline,
		}},
		{"$skip": start},
		{"$limit": end},
	}
	if err := core.Orm.All(p.GetTable(), query, &items); err != nil && err != mgo.ErrNotFound {
		return data, err
	}

	for _, val := range items {
		data.Items = append(data.Items.([]Teams), val.Team)
	}

	if data.Page.Count, err = p.CountProject(uid); err != nil {
		return data, err
	}

	return data, nil
}

// CountProject 统计项目
func (p *ProjectUser) CountProject(uid string) (int, error) {
	query := bson.M{
		"user._id": bson.ObjectIdHex(uid),
		"status":   ProjectUserOnline,
	}
	return core.Orm.Count(p.GetTable(), query)
}

// PageProjectUser 读取项目信息
func (p *ProjectUser) PageProjectUser(pid string, start, end int) (resp.PageResp, error) {
	var (
		items []ProjectUser
		err   error
	)
	data := resp.PageResp{
		Items: make([]Users, 0),
		Page: resp.PageInfoResp{
			PageSize: end,
		},
	}

	query := []bson.M{
		{"$match": bson.M{
			"project._id": bson.ObjectIdHex(pid),
			"status":      ProjectUserOnline,
		}},
		{"$skip": start},
		{"$limit": end},
	}
	if err := core.Orm.All(p.GetTable(), query, &items); err != nil && err != mgo.ErrNotFound {
		return data, err
	}

	for _, val := range items {
		data.Items = append(data.Items.([]Users), val.User)
	}

	if data.Page.Count, err = p.CountProjectUser(pid); err != nil {
		return data, err
	}

	return data, nil
}

// CountProjectUser 统计项目用户
func (p *ProjectUser) CountProjectUser(tid string) (int, error) {
	query := bson.M{
		"project._id": bson.ObjectIdHex(tid),
		"status":      ProjectUserOnline,
	}
	return core.Orm.Count(p.GetTable(), query)
}

// DismissProject 解散项目
func (p *ProjectUser) DismissProject(pid string) error {
	update := bson.M{
		"$set": bson.M{
			"status":      ProjectUserStatusDismiss,
			"modify_time": int(time.Now().Unix()),
		},
	}
	query := bson.M{
		"project._id": bson.ObjectIdHex(pid),
		"status":      ProjectUserStatusOnline,
	}

	return core.Orm.Update(p.GetTable(), query, update)
}

// ReadProjectUser 读取项目用户信息
func (p *ProjectUser) ReadProjectUser(pid, uid string) error {
	query := bson.M{
		"project._id": bson.ObjectIdHex(pid),
		"user._id":    bson.ObjectIdHex(uid),
		"status":      TeamUserStatusOnline,
	}
	return core.Orm.One(p.GetTable(), query, p)
}

// ExitProject 退出项目
func (p *ProjectUser) ExitProject(pid, uid string) error {
	update := bson.M{
		"$set": bson.M{
			"status":      ProjectUserStatusDismiss,
			"modify_time": int(time.Now().Unix()),
		},
	}

	query := bson.M{
		"project._id": bson.ObjectIdHex(pid),
		"user._id":    bson.ObjectIdHex(uid),
		"status":      ProjectUserOnline,
	}
	return core.Orm.Update(p.GetTable(), query, update)
}
