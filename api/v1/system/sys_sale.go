package system

import (
	"gandi.icu/demo/global"
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
