package system

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"

	systemReq "gandi.icu/demo/model/system/request"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ClientApi struct{}

func (clientApi *ClientApi) CreateClient(c *gin.Context) {
	var r system.SysClient
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.CreateClientVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientService.CreateClient(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (clientApi *ClientApi) GetClientList(c *gin.Context) {
	var r systemReq.SearchClientParams
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r.PageInfo, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := clientService.GetClientList(r); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("获取失败", err, c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     r.Page,
			PageSize: r.PageSize,
		}, "", c)
	}
}

func (clientApi *ClientApi) GetClientById(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if clientRes, err := clientService.GetClientById(r.ID); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("获取失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysClientResponse{Client: clientRes}, "获取成功", c)
	}
}

func (clientApi *ClientApi) UpdateClient(c *gin.Context) {
	var r system.SysClient
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.UpdateClientVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientService.UpdateClient(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("更新失败", err, c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (clientApi *ClientApi) DeleteClient(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := clientService.DeleteClient(r.ID); err != nil {
		global.AM_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("删除失败", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
