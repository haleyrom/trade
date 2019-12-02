package params

// CreateTeamParam 创建团队参数
type CreateTeamParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Name   string     `json:"name" form:"name" binding:"required"`
}

// JoinTeamParam 加入团队参数
type JoinTeamParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
}

// ExitTeamParam 退出团队参数
type ExitTeamParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
}
