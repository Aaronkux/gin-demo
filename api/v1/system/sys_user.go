package system

import (
	"errors"
	"fmt"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
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
	userRes, err := userService.Register(r)
	if err != nil {
		global.AM_LOG.Error("注册失败!", zap.Error(err))
		errMsg := "注册失败"
		fmt.Println(errors.Is(err, &response.CusError{}))
		if cusError, ok := err.(*response.CusError); ok {
			errMsg = cusError.Error()
		}
		response.FailWithMessage(errMsg, c)
	} else {
		response.OkWithDetailed(systemRes.SysUserResponse{User: userRes}, "注册成功", c)
	}
}
