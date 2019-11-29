package api

import (
	"github.com/gin-gonic/gin"
	"github.com/haleyrom/trade/core"
	"github.com/haleyrom/trade/internal/models"
	"github.com/haleyrom/trade/internal/params"
	"github.com/haleyrom/trade/internal/resp"
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
	param := &params.CreateTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(param); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO： 暂未验证权限
	team := &models.Teams{}
	if err := team.CreateTeam(param); err != nil {
		core.GResp.Failure(err)
		return
	}
	core.GResp.Success(resp.EmptyData())
	return
}

// JoinTeam 加入团队
func JoinTeam(c *gin.Context) {
	param := &params.JoinTeamParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(param); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO
	//team := &models.Teams{}

}

// ExitTeam 退出团队
func ExitTeam(c *gin.Context) {
	// TODO
}

// ReadListTeam 查看团队列表
func ReadListTeam(c *gin.Context) {
	// TODO
}

// ReadInfoTeam 查看团队信息
func ReadInfoTeam(c *gin.Context) {
	// TODO
}

// ReadUserListTeam 查看团队成员信息
func ReadUserListTeam(c *gin.Context) {
	// TODO
}
