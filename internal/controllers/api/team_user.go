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
