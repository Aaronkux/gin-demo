package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type SaleRouter struct{}

func (s *SaleRouter) InitSaleRouter(Router *gin.RouterGroup) {
	saleRouter := Router.Group("sale")
	saleApi := v1.ApiGroupApp.SystemApiGroup.SaleApi
	{
		saleRouter.POST("createSale", saleApi.CreateSale)
		saleRouter.POST("getSaleList", saleApi.GetSaleList)
		saleRouter.POST("updateSale", saleApi.UpdateSale)
		saleRouter.POST("deleteSale", saleApi.DeleteSale)
	}
}