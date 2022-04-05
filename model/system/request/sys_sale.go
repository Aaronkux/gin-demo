package request

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
)

type CreateSale struct {
	Name     string             `json:"name"`
	Avatar   string             `json:"avatar"`
	Email    string             `json:"email"`
	BranchId global.SnowflakeID `json:"branchId"`
}

type UpdateSale struct {
	ID       global.SnowflakeID `json:"id"`
	Name     string             `json:"name"`
	Avatar   string             `json:"avatar"`
	Email    string             `json:"email"`
	BranchId global.SnowflakeID `json:"branchId"`
	IsActive *bool              `json:"isActive"`
}

type SearchSaleParams struct {
	request.PageInfo
	Name     string `json:"name"`
	IsActive *bool  `json:"isActive"`
}
