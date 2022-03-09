package system

import "gandi.icu/demo/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
}

var (
	userService      = service.ServiceGroupApp.SystemServiceGroup.UserService
	authorityService = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
)
