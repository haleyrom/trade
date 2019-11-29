package models

// TeamProject 团队项目
type TeamProject struct {
	Id         string  `json:"tpid" bson:"_id"`                 // 团队项目id
	Name       string  `json:"name"  bson:"name"`               // 项目名称
	Team       Teams   `json:"team"  bson:"team"`               // 团队
	Service    Service `json:"service"  bson:"service"`         // 服务
	Status     int8    `json:"status"  bson:"status"`           // 状态 0：正常 1：关闭
	CreateTime int     `json:"create_time"  bson:"create_time"` // 创建时间
	ModifyTime int     `json:"modify_time"  bson:"modify_time"` // 更新时间
}

// GetTable GetTable
func (t *TeamProject) GetTable() string {
	return "team_project"
}
