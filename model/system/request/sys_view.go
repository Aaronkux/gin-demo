package request

import "gandi.icu/demo/global"

type SearchViewParams struct {
	Type   string             `json:"type"`
	UserID global.SnowflakeID `json:"userId"`
}
