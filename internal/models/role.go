package models

// Roles 角色
type Roles struct {
	Id         string `json:"rid" bson:"_id"`                 // 角色id
	Name       string `json:"name" bson:"name"`               // 角色名称
	Status     int8   `json:"status" bson:"status"`           // 状态 0：正常 1：启用
	Type       int8   `json:"type" bson:"type"`               // 类型 0：团队 1：角色
	Auth       string `json:"auth" bson:"auth"`               // 权限
	CreateTime int    `json:"create_time" bson:"create_time"` // 创建时间
	ModifyTime int    `json:"modify_time" bson:"modify_time"` // 更新时间
}

// GetTable GetTable
func (r *Roles) GetTable() string {
	return "roles"
}
