package params

// teamBaseParam 团队通用参数
type teamBaseParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
}

// CreateTeamParam 创建团队参数
type CreateTeamParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Name   string     `json:"name" form:"name" binding:"required"`
}

// JoinTeamParam 加入团队参数
type JoinTeamParam teamBaseParam

// ExitTeamParam 退出团队参数
type ExitTeamParam teamBaseParam

// DismissTeam 解散团队
type DismissTeam teamBaseParam
