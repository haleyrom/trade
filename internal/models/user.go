package models

// Users 用户
type Users struct {
	Id         string `json:"id" bson:"_id"` // 用户主键
	Uid        string `json:"uid"`           // 用户id
	Name       string `json:"name"`          // 用户名称
	Email      string `json:"email"`         // 邮箱
	HeadImg    string `json:"headimg"`       // 头像
	Status     int8   `json:"status"`        // 状态 0：正常 1：暂停
	CreateTime int    `json:"create_time"`   // 创建时间
	ModifyTime int    `json:"modify_time"`   // 更新时间
}

// GetTable GetTable
func (u *Users) GetTable() string {
	return "users"
}
