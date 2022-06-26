package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gorm.io/gorm"
)

type BranchService struct{}

func (b *BranchService) CreateBranch(r systemReq.CreateBranch) (branchRes system.SysBranch, err error) {
	var branchExist system.SysBranch
	if !errors.Is(global.AM_DB.Where("name = ?", r.Name).First(&branchExist).Error, gorm.ErrRecordNotFound) {
		return branchRes, &response.CusError{Msg: "已存在同名Branch"}
	}
	newBranch := system.SysBranch{Name: r.Name}
	return newBranch, global.AM_DB.Create(&newBranch).Error
}

func (b *BranchService) GetBranchList(r systemReq.SearchBranchParams) (list interface{}, total int64, err error) {
	var branchList []system.SysBranch
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysBranch{})

	if r.Name != "" {
		db = db.Where("name like ?", "%"+r.Name+"%")
	}

	err = db.Count(&total).Error
	if err != nil {
		return branchList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if r.OrderKey != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 1)
			orderMap["name"] = true
			if orderMap[r.OrderKey] {
				if r.Desc {
					OrderStr = r.OrderKey + " desc"
				} else {
					OrderStr = r.OrderKey
				}
			}

			err = db.Order(OrderStr).Find(&branchList).Error
		} else {
			err = db.Find(&branchList).Error
		}
	}
	return branchList, total, err
}

func (b *BranchService) GetBranchById(id global.SnowflakeID) (branchRes system.SysBranch, err error) {
	var branch system.SysBranch
	err = global.AM_DB.Where("id = ?", id).First(&branch).Error
	return branch, err
}

func (b *BranchService) UpdateBranch(r system.SysBranch) (err error) {
	var branchExist system.SysBranch
	if !errors.Is(global.AM_DB.Where("name = ?", r.Name).First(&branchExist).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "已存在同名Branch"}
	}

	updatedBranch := system.SysBranch{Name: r.Name}
	err = global.AM_DB.Where("id = ?", r.ID).First(&branchExist).Updates(&updatedBranch).Error
	return err
}

func (b *BranchService) DeleteBranch(id global.SnowflakeID) error {
	return global.AM_DB.Delete(&[]system.SysBranch{}, "id = ?", id).Error
}
