package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)   // 创建Api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)   // 删除Api
		apiRouter.POST("getApiById", apiRouterApi.GetApiById) // 获取单条Api消息
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)   // 更新api
		apiRouter.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有api
		apiRouter.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表
	}
}
