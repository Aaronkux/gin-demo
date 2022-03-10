package request

import "gandi.icu/demo/global"

type CreateAuthority struct {
	AuthorityName string             `json:"authorityName"`
	ParentId      global.SnowflakeID `json:"parentId"`
}
