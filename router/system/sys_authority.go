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
		authorityRouter.POST("createAuthority", authorityApi.CreateAuthority)   // 创建角色
		authorityRouter.POST("getAuthorityList", authorityApi.GetAuthorityList) // 获取角色列表
		authorityRouter.POST("updateAuthority", authorityApi.UpdateAuthority)   // 获取角色列表
		authorityRouter.POST("deleteAuthority", authorityApi.DeleteAuthority)   // 获取角色列表
		authorityRouter.POST("setAuthorityMenu", authorityApi.SetAuthorityMenu) // 获取角色列表
	}
}
