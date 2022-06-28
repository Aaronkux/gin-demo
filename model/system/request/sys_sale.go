package request

import (
	"mime/multipart"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
)

type CreateSale struct {
	Name     string                `form:"name"`
	Avatar   *multipart.FileHeader `form:"avatar"`
	Email    string                `form:"email"`
	BranchId global.SnowflakeID    `form:"branchId"`
}

type UpdateSale struct {
	ID       global.SnowflakeID `json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	BranchId global.SnowflakeID `json:"branchId"`
	IsActive *bool              `json:"isActive"`
}

type UpdateSaleAvatar struct {
	ID   global.SnowflakeID    `form:"id"`
	File *multipart.FileHeader `form:"file"`
}

type SearchSaleParams struct {
	request.PageInfo
	Name     string `json:"name"`
	IsActive *bool  `json:"isActive"`
}
