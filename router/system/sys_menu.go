package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("menu")
	menuApi := v1.ApiGroupApp.SystemApiGroup.MenuApi
	{
		userRouter.POST("createMenu", menuApi.CreateMenu)
		userRouter.POST("getMenuList", menuApi.GetMenuList)
		userRouter.POST("getMenuListAll", menuApi.GetMenuListAll)
		userRouter.POST("getMenuById", menuApi.GetMenuById)
		userRouter.POST("updateMenu", menuApi.UpdateMenu)
		userRouter.POST("deleteMenu", menuApi.DeleteMenu)
		userRouter.POST("getMenuKeysByUserAuthority", menuApi.GetMenuKeysByUserAuthority)
	}
}
