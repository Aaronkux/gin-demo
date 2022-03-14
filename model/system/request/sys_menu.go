package request

import "gandi.icu/demo/global"

type CreateMenu struct {
	MenuName string              `json:"menuName"`
	Path     string              `json:"path"`
	Hidden   *bool               `json:"hidden"`
	Icon     string              `json:"icon"`
	Order    *uint               `json:"order"`
	ParentId *global.SnowflakeID `json:"parentId"`
}
