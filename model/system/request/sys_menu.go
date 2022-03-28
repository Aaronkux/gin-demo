package request

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
)

type CreateMenu struct {
	MenuName string              `json:"menuName"`
	MenuKey  string              `json:"menuKey"`
	Path     string              `json:"path"`
	Hidden   *bool               `json:"hidden"`
	ParentId *global.SnowflakeID `json:"parentId"`
}

type UpdateMenu struct {
	ID       global.SnowflakeID `json:"id"`
	MenuName string             `json:"menuName"`
	MenuKey  string             `json:"menuKey"`
	Path     string             `json:"path"`
	Hidden   *bool              `json:"hidden"`
}

type SearchMenuParams struct {
	request.PageInfo
	MenuName string `json:"menuName"`
}
