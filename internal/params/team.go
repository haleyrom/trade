package params

// CreateTeamParam 创建团队参数
type CreateTeamParam struct {
	Claims *BaseParam `json:"claims" form:"claims" `
	Name   string     `json:"name" form:"name" binding:"required"`
}
