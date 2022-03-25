package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gorm.io/gorm"
)

type SaleService struct{}

func (b *SaleService) CreateSale(r systemReq.CreateSale) (saleRes system.SysSale, err error) {
	var saleExist system.SysSale
	var branchExist system.SysBranch
	if !errors.Is(global.AM_DB.Where("email = ?", r.Email).First(&saleExist).Error, gorm.ErrRecordNotFound) {
		return saleRes, &response.CusError{Msg: "已存销售使用该邮箱"}
	}
	if errors.Is(global.AM_DB.Where("id = ?", r.BranchId).First(&branchExist).Error, gorm.ErrRecordNotFound) {
		return saleRes, &response.CusError{Msg: "该部门不存在"}
	}
	newSale := system.SysSale{Name: r.Name, Email: r.Email, Avatar: r.Avatar, IsActive: true}
	newSale.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&newSale).Error; err != nil {
			return err
		}
		if err := tx.Model(&branchExist).Association("Sales").Append(&newSale); err != nil {
			return err
		}
		return nil
	})
	return newSale, err
}
