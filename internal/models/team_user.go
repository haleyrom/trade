package models

// TeamUser 团队成员
type TeamUser struct {
	Id         string `json:"tuid" bson:"_id"` // 团队成员id
	Team       Teams  `json:"team"`            // 团队
	User       Users  `json:"user"`            // 用户
	Role       Roles  `json:"role"`            // 身份
	Status     int8   `json:"status"`          // 状态 0：正常 1：退出
	CreateTime int    `json:"create_time"`     // 创建时间
	ModifyTime int    `json:"modify_time"`     // 更新时间
}

// GetTable GetTable
func (t *TeamUser) GetTable() string {
	return "team_user"
}
