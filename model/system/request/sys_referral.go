package request

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
)

type CreateReferral struct {
	Name string `json:"name"`
}

type UpdateReferral struct {
	ID   global.SnowflakeID `json:"id"`
	Name string             `json:"name"`
}

type SearchReferralParams struct {
	request.PageInfo
	Name string `json:"name"`
}
