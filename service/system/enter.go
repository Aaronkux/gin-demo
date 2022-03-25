package system

type ServiceGroup struct {
	JwtService
	UserService
	CasbinService
	AuthorityService
	MenuService
	ApiService
	FileService
	BranchService
}
