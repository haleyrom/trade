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

// JoinProjectParam 创建项目参数
type JoinProjectParam projectBaseParam
