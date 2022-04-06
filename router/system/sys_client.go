package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type ClientRouter struct{}

func (s *ClientRouter) InitClientRouter(Router *gin.RouterGroup) {
	clientRouter := Router.Group("client")
	clientApi := v1.ApiGroupApp.SystemApiGroup.ClientApi
	{
		clientRouter.POST("createClient", clientApi.CreateClient)
		clientRouter.POST("getClientList", clientApi.GetClientList)
		clientRouter.POST("getClientById", clientApi.GetClientById)
		clientRouter.POST("updateClient", clientApi.UpdateClient)
		clientRouter.POST("deleteClient", clientApi.DeleteClient)
	}
}
