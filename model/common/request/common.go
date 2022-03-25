package request

import "gandi.icu/demo/global"

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	OrderKey string `json:"orderKey"`                 // 排序
	Desc     bool   `json:"desc"`                     // 排序方式:升序false(默认)|降序true
}

type IdsReq struct {
	Ids []global.SnowflakeID `json:"ids" form:"ids"`
}

type GetById struct {
	ID global.SnowflakeID `json:"id" form:"id"` // 主键ID
}
