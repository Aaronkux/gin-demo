package request

import (
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/system"
)

type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式:升序false(默认)|降序true
}
