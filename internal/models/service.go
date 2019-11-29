package models

// Service 服务
type Service struct {
	Id         string `json:"sid" bson:"_id"`                 // 服务id
	Name       string `json:"name" bson:"name"`               // 服务名称
	Status     int8   `json:"status" bson:"status"`           // 状态 0:开启 1:关闭
	CreateTime int    `json:"create_time" bson:"create_time"` // 创建时间
	ModifyTime int    `json:"modify_time" bson:"modify_time"` // 更新时间
}

// GetTable GetTable
func (s *Service) GetTable() string {
	return "service"
}
