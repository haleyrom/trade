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

// ReadUserListTeam 查看团队成员列表
// @Tags 7. ReadUserListTeam
// @Summary 查看团队成员列表接口
// @Description 查看团队成员列表
// @Produce json
// @Param page query string true "页数默认为1"
// @Param size query string true "页数数量默认为20"
// @Success 200
// @Router /api/team/user_list [post]
func ReadUserListTeam(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "20"))
	p := &params.ReadUserListTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
		Page:   page,
		Size:   size,
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO： 是否存在权限查看

	// 判断是否存在该团队
	team := models.NewTeam()
	if err := team.IsExistTeam(p.Tid); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotTeam))
		return
	}

	data, err := models.NewTeamUser().PageTeamUser(p.Tid, (page-1)*size, size)
	if err != nil {
		core.GResp.Failure(err)
		return
	}
	core.GResp.Success(data)
	return
}

// ReadTeamUser 读取团队用户
// @Tags 8. ReadTeamUser
// @Summary 读取团队用户接口
// @Description 读取团队用户
// @Produce json
// @Param uid query string true "用户id"
// @Param tid query string true "团队id"
// @Success 200
// @Router /api/team/user [post]
func ReadTeamUser(c *gin.Context) {
	p := &params.ReadTeamUserParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	teamUser := models.NewTeamUser()
	if err := teamUser.ReadTeamUser(p.Tid, p.Claims.ID); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotUser))
		return
	}
	core.GResp.Success(teamUser.User)
	return
}
