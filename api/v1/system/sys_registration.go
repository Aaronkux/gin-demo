package system

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RegistrationApi struct{}

func (clientApi *ClientApi) CreateIndividualClient(c *gin.Context) {
	var r system.SysClient
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.CreateClientVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if clientRes, err := clientService.CreateClient(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysClientResponse{Client: clientRes}, "创建成功", c)
	}
}

func (clientApi *ClientApi) CreateCompanyClient(c *gin.Context) {
	var r system.SysClient
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.CreateClientVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if clientRes, err := clientService.CreateClient(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysClientResponse{Client: clientRes}, "创建成功", c)
	}
}
