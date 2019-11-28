package models

// Team 团队
type Teams struct {
	Id         string `json:"tid"`         // 团队id
	Name       string `json:"name"`        // 团队名称
	Creator    Users  `json:"creator"`     // 团队创建人
	Status     int8   `json:"status"`      // 状态 0：正常 1：解散
	CreateTime int    `json:"create_time"` // 创建时间
	ModifyTime int    `json:"modify_time"` // 更新时间
}

// getTeamsTable getTeamsTable
func getTeamsTable() string {
	return "teams"
}
