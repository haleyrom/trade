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
// @Tags 9. JoinProject
// @Summary 加入项目接口
// @Description 加入项目
// @Produce json
// @Param pid query string true "项目id"
// @Param tid query string true "团队id"
// @Success 200
// @Router /api/project/join [post]
func JoinProject(c *gin.Context) {
	p := &params.JoinProjectParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}
	// 判断是否存在团队里面
	teamUser := models.NewTeamUser()
	if err := teamUser.IsExistJoinTeam(p.Tid, p.Claims.ID); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeExistTeam))
		return
	}

	// 判断是否存在该项目
	project := models.NewTeamProject()
	if err := project.IsExistTeam(p.Pid); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotProject))
		return
	}

	if err := models.NewProjectUser().IsExistJoinProject(p.Pid, p.Claims.ID); err == nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeExistProject))
		return
	}

	projectUser := models.NewProjectUser()
	projectUser.Project, projectUser.User = *project, models.AssignUsers(p.Claims)
	if err := projectUser.JoinTeamProject(); err != nil {
		core.GResp.Failure(err)
		return
	}

	core.GResp.Success(resp.EmptyData())
	return

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
