package system

import (
	v1 "gandi.icu/demo/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthorityRouter struct{}

func (s *AuthorityRouter) InitAuthorityRouter(Router *gin.RouterGroup) {
	authorityRouter := Router.Group("authority")
	authorityApi := v1.ApiGroupApp.SystemApiGroup.AuthorityApi
	{
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority)
		authorityRouter.POST("getAuthorityList", authorityApi.GetAuthorityList)
		authorityRouter.POST("getAuthorityById", authorityApi.GetAuthorityById)
		authorityRouter.POST("updateAuthority", authorityApi.UpdateAuthority)
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority)
		authorityRouter.POST("setAuthorityMenu", authorityApi.SetAuthorityMenu)
		authorityRouter.POST("getAuthorityMenu", authorityApi.GetAuthorityMenu)
	}
}
