package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	menuRouter := Router.Group("menu")
	menuApi := v1.ApiGroupApp.SystemApiGroup.MenuApi
	{
		menuRouter.POST("createMenu", menuApi.CreateMenu)
		menuRouter.POST("getMenuList", menuApi.GetMenuList)
		menuRouter.POST("getMenuListAll", menuApi.GetMenuListAll)
		menuRouter.POST("getMenuById", menuApi.GetMenuById)
		menuRouter.POST("updateMenu", menuApi.UpdateMenu)
		menuRouter.POST("deleteMenu", menuApi.DeleteMenu)
		menuRouter.POST("getMenuKeysByUserAuthority", menuApi.GetMenuKeysByUserAuthority)
	}
}
