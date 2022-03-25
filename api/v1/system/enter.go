package system

import "gandi.icu/demo/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
	MenuApi
	CasbinApi
	SystemApiApi
	BranchApi
	SaleApi
}

var (
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	menuService      = service.ServiceGroupApp.SystemServiceGroup.MenuService
	casbinService    = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	apiService       = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService       = service.ServiceGroupApp.SystemServiceGroup.JwtService
	fileService      = service.ServiceGroupApp.SystemServiceGroup.FileService
	branchService    = service.ServiceGroupApp.SystemServiceGroup.BranchService
	saleService      = service.ServiceGroupApp.SystemServiceGroup.SaleService
)
