package params

// CreateProjectParam 创建项目参数
type CreateProjectParam struct {
	Claims *BaseParam `json:"claims" form:"claims"`
	Name   string     `json:"name" form:"name" binding:"required"`
	Tid    string     `json:"tid" form:"tid" binding:"required"`
	Sid    string     `json:"sid" form:"sid" `
}
