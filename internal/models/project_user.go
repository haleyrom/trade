package models

// ProjectUser 项目用户
type ProjectUser struct {
	Id         string      `json:"puid" bson:"_id"` // 项目用户id
	Project    TeamProject `json:"project"`         // 项目
	User       Users       `json:"user"`            // 用户
	Role       Roles       `json:"role"`            // 权限
	Status     int8        `json:"status"`          // 状态 0：正常 1：退出
	CreateTime int         `json:"create_time"`     // 创建时间
	ModifyTime int         `json:"modify_time"`     // 更新时间
}

// GetTable GetTable
func (p *ProjectUser) GetTable() string {
	return "project_user"
}
