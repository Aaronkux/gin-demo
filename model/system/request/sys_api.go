package request

import (
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/system"
)

type SearchApiParams struct {
	system.SysApi
	request.PageInfo
}
