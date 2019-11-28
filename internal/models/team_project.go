package models

// TeamProject 团队项目
type TeamProject struct {
	Id         string  `json:"tpid"`        // 团队项目id
	Name       string  `json:"name"`        // 项目名称
	Team       Teams   `json:"team"`        // 团队
	Service    Service `json:"service"`     // 服务
	Status     int8    `json:"status"`      // 状态 0：正常 1：关闭
	CreateTime int     `json:"create_time"` // 创建时间
	ModifyTime int     `json:"modify_time"` // 更新时间
}

// getTeamProjectTable getTeamProjectTable
func getTeamProjectTable() string {
	return "team_project"
}
