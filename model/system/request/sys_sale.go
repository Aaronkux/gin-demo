package request

import "gandi.icu/demo/global"

type CreateSale struct {
	Name     string              `json:"name"`
	Avatar   string              `json:"avatar"`
	Email    string              `json:"email"`
	BranchId *global.SnowflakeID `json:"BranchId"`
}
