package models

// Service 服务
type Service struct {
	Id         string `json:"sid"`         // 服务id
	Name       string `json:"name"`        // 服务名称
	Status     int8   `json:"status"`      // 状态 0:开启 1:关闭
	CreateTime int    `json:"create_time"` // 创建时间
	ModifyTime int    `json:"modify_time"` // 更新时间
}

// getServiceTable getServiceTable
func getServiceTable() string {
	return "service"
}
