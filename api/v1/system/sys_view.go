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

type ViewApi struct{}

func (viewApi *ViewApi) CreateView(c *gin.Context) {
	var r system.SysView
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.ViewCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if viewRes, err := viewService.CreateView(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysViewResponse{View: viewRes}, "创建成功", c)
	}
}

func (viewApi *ViewApi) GetViewByUserIdAndType(c *gin.Context) {
	var r systemReq.SearchViewParams
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.GetViewByUserIdAndTypeVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if viewRes, err := viewService.GetViewByUserIdAndType(r); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("获取失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysViewResponse{View: viewRes}, "", c)
	}
}

func (viewApi *ViewApi) GetViewById(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if viewRes, err := viewService.GetViewById(r.ID); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("获取失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysViewResponse{View: viewRes}, "", c)
	}
}

func (viewApi *ViewApi) UpdateView(c *gin.Context) {
	var r system.SysView
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.ViewUpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := viewService.UpdateView(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("更新失败", err, c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (viewApi *ViewApi) DeleteView(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBind(&r)
	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := viewService.DeleteView(r.ID); err != nil {
		global.AM_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("删除失败", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
