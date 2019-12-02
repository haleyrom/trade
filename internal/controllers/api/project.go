package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/internal/models"
	"github.com/haleyrom/trade/internal/params"
	"github.com/haleyrom/trade/internal/resp"
)

// CreateProject 创建项目
// @Tags 8. CreateProject
// @Summary 创建项目接口
// @Description 创建项目列表
// @Produce json
// @Param name query string true "项目名称"
// @Param tid query string true "团队id"
// @Success 200
// @Router /api/project/create [post]
func CreateProject(c *gin.Context) {
	p := &params.CreateProjectParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO: 权限创建项目

	// 判断是否存在该团队
	team := models.NewTeam()
	if err := team.IsExistTeam(p.Tid); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotTeam))
		return
	}

	// 判断是否为团队内部成员
	if err := models.NewTeamUser().IsExistJoinTeam(p.Tid, p.Claims.ID); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeAuth))
		return
	}

	// TODO: 服务是否存在
	project := &models.TeamProject{
		Team: *team,
	}
	if err := project.CreateProject(p); err != nil {
		core.GResp.Failure(err)
		return
	}
	core.GResp.Success(resp.EmptyData())
	return
}

// JoinProject 加入项目
func JoinProject(c *gin.Context) {
	// TODO
}

// ReadListProject 项目列表
func ReadListProject(c *gin.Context) {
	// TODO
}

// ReadListTeamProject 读取团队项目列表
func ReadListTeamProject(c *gin.Context) {
	// TODO
}

// DismissProject 解散项目
func DismissProject(c *gin.Context) {
	// TODO
}
