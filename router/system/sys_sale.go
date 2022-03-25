package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type SaleRouter struct{}

func (s *SaleRouter) InitSaleRouter(Router *gin.RouterGroup) {
	casbinRouter := Router.Group("sale")
	casbinApi := v1.ApiGroupApp.SystemApiGroup.SaleApi
	{
		casbinRouter.POST("createSale", casbinApi.CreateSale)
	}
}
