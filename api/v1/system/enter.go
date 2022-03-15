package system

import "gandi.icu/demo/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
	MenuApi
	CasbinApi
	SystemApiApi
}

var (
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	menuService      = service.ServiceGroupApp.SystemServiceGroup.MenuService
	casbinService    = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	apiService       = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService       = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
