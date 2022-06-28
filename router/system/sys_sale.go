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
		saleRouter.POST("getSaleById", saleApi.GetSaleById)
		saleRouter.POST("getSaleList", saleApi.GetSaleList)
		saleRouter.POST("getAllSales", saleApi.GetAllSales)
		saleRouter.POST("updateSale", saleApi.UpdateSale)
		saleRouter.POST("updateSaleAvatar", saleApi.UpdateSaleAvatar)
		saleRouter.POST("deleteSale", saleApi.DeleteSale)
	}
}
