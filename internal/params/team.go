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

// ReadListTeamParam 团队列表参数
type ReadListTeamParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Page   int        `json:"page" form:"page"  binding:"required"`
	Size   int        `json:"size" form:"size"  binding:"required"`
}

// ReadUserListTeamParam 团队成员列表参数
type ReadUserListTeamParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
	Page   int        `json:"page" form:"page"  binding:"required"`
	Size   int        `json:"size" form:"size"  binding:"required"`
}

// ReadTeamUserParam 读取团队用户信息参数
type ReadTeamUserParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
	Uid    string     `json:"uid" form:"uid" binding:"required"`
}

// JoinTeamParam 加入团队参数
type JoinTeamParam teamBaseParam

// ExitTeamParam 退出团队参数
type ExitTeamParam teamBaseParam

// DismissTeamParam 解散团队参数
type DismissTeamParam teamBaseParam

// ReadInfoTeamParam 读取团队信息参数
type ReadInfoTeamParam teamBaseParam
