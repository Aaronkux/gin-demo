package system

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/common/response"
	systemReq "gandi.icu/demo/model/system/request"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SaleApi struct{}

func (s *SaleApi) CreateSale(c *gin.Context) {
	var r systemReq.CreateSale
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.SaleCreateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if saleRes, err := saleService.CreateSale(r); err != nil {
		global.AM_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithCustomErrorOrDefault("创建失败", err, c)
	} else {
		response.OkWithDetailed(systemRes.SysSaleResponse{Sale: saleRes}, "创建成功", c)
	}
}

func (s *SaleApi) GetSaleList(c *gin.Context) {
	var r systemReq.SearchSaleParams
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r.PageInfo, utils.GetSaleListVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if list, total, err := saleService.GetSaleList(r); err != nil {
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

func (s *SaleApi) GetAllSales(c *gin.Context) {
	if list, total, err := saleService.GetAllSales(); err != nil {
		global.AM_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:  list,
			Total: total,
		}, "", c)
	}
}

func (s *SaleApi) UpdateSale(c *gin.Context) {
	var r systemReq.UpdateSale
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.SaleUpdateVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := saleService.UpdateSale(r); err != nil {
		global.AM_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

func (s *SaleApi) DeleteSale(c *gin.Context) {
	var r request.GetById
	_ = c.ShouldBindJSON(&r)

	if err := utils.Verify(r, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := saleService.DeleteSale(r.ID); err != nil {
		global.AM_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
