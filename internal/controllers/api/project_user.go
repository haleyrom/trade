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

// ReadProjectUser 读取项目用户
// @Tags 12. ReadProjectUser
// @Summary 读取项目用户接口
// @Description 读取项目用户项目
// @Produce json
// @Param pid query string true "项目id"
// @Param uid query string true "用户id"
// @Success 200
// @Router /api/project/user [post]
func ReadProjectUser(c *gin.Context) {
	p := &params.ReadProjectUserParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO: 权限查看/测试
	projectUser := models.NewProjectUser()
	if err := projectUser.ReadProjectUser(p.Pid, p.Uid); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotUser))
		return
	}
	core.GResp.Success(projectUser.User)
	return
}

// ReadListProjectUser 读取项目用户列表
// @Tags 13. ReadListProjectUser
// @Summary 读取项目用户列表接口
// @Description 读取项目用户列表
// @Produce json
// @Param pid query string true "项目id"
// @Param page query string true "页数默认为1"
// @Param size query string true "页数数量默认为20"
// @Success 200
// @Router /api/project/user [post]
func ReadListProjectUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultPostForm("page", "1"))
	size, _ := strconv.Atoi(c.DefaultPostForm("size", "20"))
	p := &params.ReadListProjectUserParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
		Page:   page,
		Size:   size,
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO： 是否存在权限查看/测试

	// 判断是否存在该团队
	project := models.NewTeamProject()
	if err := project.IsExistProject(p.Pid); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotProject))
		return
	}

	data, err := models.NewProjectUser().PageProjectUser(p.Pid, (page-1)*size, size)
	if err != nil {
		core.GResp.Failure(err)
		return
	}
	core.GResp.Success(data)
	return
}

// ExitProject 退出项目
// @Tags 14. ExitProject
// @Summary 退出项目接口
// @Description 退出项目列表
// @Produce json
// @Param pid query string true "项目id"
// @Success 200
// @Router /api/project/exit [post]
func ExitProject(c *gin.Context) {
	p := &params.ExitProjectParam{
		Claims: core.UserInfoPool.Get().(*params.BaseParam),
	}

	// 绑定参数
	if err := c.ShouldBind(p); err != nil {
		core.GResp.Failure(err)
		return
	}

	// TODO: 权限/团队最后一人/测试

	// 判断是否存在该团队
	projectUser := models.NewProjectUser()
	if err := projectUser.IsExistJoinProject(p.Pid, p.Claims.ID); err != nil {
		core.GResp.Failure(fmt.Errorf("%d", resp.CodeNotProject))
		return
	}

	if err := projectUser.ExitProject(p.Pid, p.Claims.ID); err != nil {
		core.GResp.Failure(err)
		return
	}
	core.GResp.Success(resp.EmptyData())
	return
}
