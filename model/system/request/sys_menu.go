package request

import "gandi.icu/demo/global"

type CreateMenu struct {
	MenuName string              `json:"menuName"`
	Path     string              `json:"path"`
	Hidden   *bool               `json:"hidden"`
	ParentId *global.SnowflakeID `json:"parentId"`
}
