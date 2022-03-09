package request

import "encoding/json"

type CreateAuthority struct {
	AuthorityName string      `json:"authorityName"`
	ParentId      json.Number `json:"parentId"`
}
