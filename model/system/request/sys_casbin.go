package request

import "gandi.icu/demo/global"

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId global.SnowflakeID `json:"authorityId"` // 权限id
	CasbinInfos []CasbinInfo       `json:"casbinInfos"`
}

func DefaultCasbin() []CasbinInfo {
	return []CasbinInfo{
		{Path: "/base/login", Method: "POST"},
		{Path: "/user/register", Method: "POST"},
		{Path: "/user/changePassword", Method: "POST"},
	}
}
