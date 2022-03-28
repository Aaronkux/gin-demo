package system

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MenuApi struct{}

func (m *MenuApi) CreateMenu(c *gin.Context) {
	var r systemReq.CreateMenu
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.MenuCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if menuRes, err := menuService.CreateMenu(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysMenuResponse{Menu: menuRes}, "创建成功", c)
	}
}

func (m *MenuApi) GetMenuList(c *gin.Context) {
	var r systemReq.SearchMenuParams
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.PageInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := menuService.GetMenuList(r); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     r.Page,
			PageSize: r.PageSize,
		}, "", c)
	}
}

func (m *MenuApi) GetMenuById(c *gin.Context) {
	var r system.SysMenu
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if menuRes, err := menuService.GetMenuById(r); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysMenuResponse{Menu: menuRes}, "", c)
	}
}

func (m *MenuApi) UpdateMenu(c *gin.Context) {
	var r systemReq.UpdateMenu
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.MenuUpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if menuRes, err := menuService.UpdateMenu(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithDetailed(systemRes.SysMenuResponse{Menu: menuRes}, "更新成功", c)
	}
}

func (m *MenuApi) DeleteMenu(c *gin.Context) {
	var r system.SysMenu
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := menuService.DeleteMenu(r); err != nil {
		global.AM_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("删除失败", err, c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
