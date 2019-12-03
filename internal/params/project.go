package params

// projectBaseParam 项目通用结构
type projectBaseParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
	Pid    string     `json:"pid" form:"pid" binding:"required"`
}

// CreateProjectParam 创建项目参数
type CreateProjectParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Name   string     `json:"name" form:"name" binding:"required"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
	Sid    string     `json:"sid" form:"sid"`
}

// ReadListProjectParam 项目列表参数
type ReadListProjectParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Page   int        `json:"page" form:"page"  binding:"required"`
	Size   int        `json:"size" form:"size"  binding:"required"`
}

// ReadListTeamProjectParam 团队项目列表参数
type ReadListTeamProjectParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Page   int        `json:"page" form:"page"  binding:"required"`
	Size   int        `json:"size" form:"size"  binding:"required"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
}

// ReadProjectUserParam  读取项目用户参数
type ReadProjectUserParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Uid    string     `json:"uid" form:"uid" binding:"required"`
	Pid    string     `json:"pid" form:"pid" binding:"required"`
}

// ReadListProjectUserParam 读取项目用户列表参数
type ReadListProjectUserParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Page   int        `json:"page" form:"page"  binding:"required"`
	Size   int        `json:"size" form:"size"  binding:"required"`
	Pid    string     `json:"pid" form:"pid" binding:"required"`
}

// ExitProjectParam 退出项目
type ExitProjectParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Pid    string     `json:"pid" form:"pid" binding:"required"`
}

// JoinProjectParam 创建项目参数
type JoinProjectParam projectBaseParam

// ReadInfoProjectParam 解散项目参数
type DismissProjectParam projectBaseParam

// ReadInfoProjectParam 读取项目参数
type ReadInfoProjectParam projectBaseParam
