package request

import "gandi.icu/demo/global"

type CreateAuthority struct {
	AuthorityName string             `json:"authorityName"`
	ParentId      global.SnowflakeID `json:"parentId"`
}

type SetAuthorityMenu struct {
	AuthorityId global.SnowflakeID   `json:"authorityId"`
	MenuIds     []global.SnowflakeID `json:"menuIds"`
}
