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

func (s *SaleService) CreateSale(r systemReq.CreateSale) (saleRes system.SysSale, err error) {
	var saleExist system.SysSale
	var branchExist system.SysBranch
	if !errors.Is(global.AM_DB.Where("email = ?", r.Email).First(&saleExist).Error, gorm.ErrRecordNotFound) {
		return saleRes, &response.CusError{Msg: "已存销售使用该邮箱"}
	}
	if errors.Is(global.AM_DB.Where("id = ?", r.BranchId).First(&branchExist).Error, gorm.ErrRecordNotFound) {
		return saleRes, &response.CusError{Msg: "该部门不存在"}
	}
	newSale := system.SysSale{Name: r.Name, Email: r.Email, Avatar: r.Avatar, BranchID: r.BranchId, IsActive: true}
	if err := global.AM_DB.Create(&newSale).Error; err != nil {
		return saleRes, err
	}
	return newSale, err
}

func (s *SaleService) UpdateSale(r systemReq.UpdateSale) (err error) {
	var saleExist system.SysSale
	var branchExist system.SysBranch
	if errors.Is(global.AM_DB.Where("id = ?", r.ID).First(&saleExist).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "该销售不存在"}
	}
	branchChanged := saleExist.BranchID != r.BranchId
	if branchChanged && errors.Is(global.AM_DB.Where("id = ?", r.BranchId).First(&branchExist).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "该部门不存在"}
	}
	updateSale := system.SysSale{Name: r.Name, Email: r.Email, Avatar: r.Avatar, BranchID: r.BranchId, IsActive: *r.IsActive}
	err = global.AM_DB.Model(&saleExist).Select("*").Updates(&updateSale).Error
	return err
}

func (s *SaleService) GetSaleList(r systemReq.SearchSaleParams) (list interface{}, total int64, err error) {
	var saleList []system.SysSale
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysSale{})
	db.Where("is_active = ?", *r.IsActive)
	if r.Name != "" {
		db = db.Where("name like ?", "%"+r.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return saleList, total, err
	}
	err = db.Preload("Branch").Limit(limit).Offset(offset).Find(&saleList).Error
	return saleList, total, err
}

func (s *SaleService) GetAllSales() (list interface{}, total int64, err error) {
	var saleList []system.SysSale
	db := global.AM_DB.Model(&system.SysSale{})
	db.Where("is_active = ?", true)
	err = db.Count(&total).Error
	if err != nil {
		return saleList, total, err
	}
	err = db.Preload("Branch").Find(&saleList).Error
	return saleList, total, err
}

func (s *SaleService) DeleteSale(id global.SnowflakeID) (err error) {
	err = global.AM_DB.Delete(&system.SysSale{}, id).Error
	return err
}
