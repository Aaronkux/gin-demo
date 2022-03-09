package system

import (
	"fmt"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) Register(c *gin.Context) {
	var r systemReq.Register
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.RegisterVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 检查权限ID是否存在
	var authorities []system.SysAuthority
	if err := global.AM_DB.Where("id In ?", r.AuthorityIds).Find(&authorities).Error; err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(r.AuthorityIds)
	if len(authorities) != len(r.AuthorityIds) {
		response.FailWithMessage("权限ID不存在", c)
		return
	}
	user := &system.SysUser{Email: r.Email, Nickname: r.Nickname, Password: r.Password, Avatar: r.Avatar, Authorities: authorities}
	user.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.AM_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
	}
}
