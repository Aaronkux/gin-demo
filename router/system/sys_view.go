package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type ViewRouter struct{}

func (s *ViewRouter) InitViewRouter(Router *gin.RouterGroup) {
	viewRouter := Router.Group("view")
	viewApi := v1.ApiGroupApp.SystemApiGroup.ViewApi
	{
		viewRouter.POST("createView", viewApi.CreateView)
		viewRouter.POST("getViewByUserIdAndType", viewApi.GetViewByUserIdAndType)
		viewRouter.POST("getViewById", viewApi.GetViewById)
		viewRouter.POST("updateView", viewApi.UpdateView)
		viewRouter.POST("deleteView", viewApi.DeleteView)
	}
}
