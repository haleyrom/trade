package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/internal/models"
	"github.com/haleyrom/trade/internal/params"
	"github.com/haleyrom/trade/internal/resp"
	"strconv"
)

// CreateTeam 创建团队
// @Tags 1. CreateTeam
// @Summary 创建团队接口
// @Description 创建团队
// @Produce json
// @Param name query string true "团队名称"
// @Success 200
// @Router /api/team/create [post]
func CreateTeam(c *gin.Context) {
	p := &params.CreateTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO： 暂未验证权限
	if err := models.NewTeam().CreateTeam(p); err != nil {
		core.GResp.Failure(err)
		return
	}
	core.GResp.Success(resp.EmptyData())
	return
}

// JoinTeam 加入团队
// @Tags 2. JoinTeam
// @Summary 加入团队接口
// @Description 加入团队
// @Produce json
// @Param tid query string true "团队id"
// @Success 200
// @Router /api/team/join [post]
func JoinTeam(c *gin.Context) {
	p := &params.JoinTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO: 判断邀请码/权限判断

	// 判断是否存在该团队
	team := models.NewTeam()
	if err := team.IsExistTeam(p.Tid); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotTeam))
		return
	}

	// 判断是否存在团队里面
	teamUser := models.NewTeamUser()
	if err := teamUser.IsExistJoinTeam(p.Tid, p.Claims.ID); err == nil {
		// 是否退出
		if teamUser.Status == models.TeamUserStatusOnline {
			core.GResp.Failure(fmt.Errorf("%d", resp.CodeExistTeam))
			return
		}
	}

	teamUser.Team, teamUser.User = *team, models.AssignUsers(p.Claims)
	if err := teamUser.JoinTeamUser(); err != nil {
		core.GResp.Failure(err)
		return
	}

	core.GResp.Success(resp.EmptyData())
	return

}

// ExitTeam 退出团队
// @Tags 3. ExitTeam
// @Summary 退出团队接口
// @Description 退出团队
// @Produce json
// @Param tid query string true "团队id"
// @Success 200
// @Router /api/team/exit [post]
func ExitTeam(c *gin.Context) {
	p := &params.ExitTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO: 权限

	// 判断是否存在该团队
	teamUser := models.NewTeamUser()
	if err := teamUser.IsExistJoinTeam(p.Tid, p.Claims.ID); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotTeam))
		return
	}

	// 团队队长不能退团
	if teamUser.Type == models.TeamUserTypeOwner {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeAuth))
		return
	}

	if teamUser.Status == models.TeamUserStatusOnline {
		if err := teamUser.ExitTeam(p.Tid, p.Claims.ID); err != nil {
			core.GResp.Failure(err)
			return
		}
	}
	core.GResp.Success(resp.EmptyData())
	return
}

// DismissTeam 解散团队
// @Tags 4. DismissTeam
// @Summary 解散团队接口
// @Description 解散团队
// @Produce json
// @Param tid query string true "团队id"
// @Success 200
// @Router /api/team/dismiss [post]
func DismissTeam(c *gin.Context) {
	p := &params.DismissTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	teamUser := models.NewTeamUser()
	if err := teamUser.IsExistJoinTeam(p.Tid, p.Claims.ID); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotTeam))
		return
	}

	// 团队队长
	if teamUser.Type == models.TeamUserTypeOwner {
		if err := models.NewTeam().DismissTeam(p.Tid); err != nil {
			core.GResp.Failure(err)
			return
		}
		core.GResp.Success(resp.EmptyData())
		return
	}
	core.GResp.Failure(fmt.Errorf("%d", resp.CodeAuth))
	return
}

// ReadListTeam 查看团队列表
// @Tags 5. ReadListTeam
// @Summary 查看团队列表接口
// @Description 查看团队列表
// @Produce json
// @Param page query string true "页数默认为1"
// @Param size query string true "页数数量默认为20"
// @Success 200
// @Router /api/team/list [post]
func ReadListTeam(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "20"))
	p := &params.ReadListTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
		Page:   page,
		Size:   size,
	}

	data, err := models.NewTeamUser().PageTeam(p.Claims.ID, (page-1)*size, size)
	if err != nil {
		core.GResp.Failure(err)
		return
	}
	core.GResp.Success(data)
	return
}

// ReadInfoTeam 查看团队信息
// @Tags 6. ReadInfoTeam
// @Summary 查看团队信息接口
// @Description 查看团队信息
// @Produce json
// @Param tid query string true "团队id"
// @Success 200
// @Router /api/team/info [post]
func ReadInfoTeam(c *gin.Context) {
	p := &params.ReadInfoTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO: 是否存在权限查看

	// 判断是否存在该团队
	team := models.NewTeam()
	if err := team.ReadTeamInfo(p.Tid); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotTeam))
		return
	}

	core.GResp.Success(team)
	return
}
