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
	FileApi
	UserApi
	ReferralApi
	ClientApi
	ViewApi
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
	referralService  = service.ServiceGroupApp.SystemServiceGroup.ReferralService
	clientService    = service.ServiceGroupApp.SystemServiceGroup.ClientService
	viewService      = service.ServiceGroupApp.SystemServiceGroup.ViewService
)
