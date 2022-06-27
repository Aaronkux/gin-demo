package system

import (
	"errors"
	"fmt"
	"mime/multipart"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SaleService struct{}

func (s *SaleService) CreateSale(c *gin.Context, r systemReq.CreateSale) (err error) {
	var saleExist system.SysSale
	var branchExist system.SysBranch
	if !errors.Is(global.AM_DB.Where("email = ?", r.Email).First(&saleExist).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "已存销售使用该邮箱"}
	}
	if errors.Is(global.AM_DB.Where("id = ?", r.BranchId).First(&branchExist).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "该部门不存在"}
	}
	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		fileRes, err := FileServiceApp.UploadFile(c, "avatar", r.Avatar)
		if err != nil {
			return err
		}

		newSale := system.SysSale{Name: r.Name, Email: r.Email, AvatarID: fileRes.ID, BranchID: r.BranchId, IsActive: true}
		newSale.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
		if err := global.AM_DB.Create(&newSale).Error; err != nil {
			return err
		}
		return nil
	})

	return err
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
	updateSale := system.SysSale{Name: r.Name, Email: r.Email, BranchID: r.BranchId, IsActive: *r.IsActive}
	err = global.AM_DB.Model(&saleExist).Select("Name", "Email", "BranchID", "IsActive").Updates(&updateSale).Error
	return err
}

func (s *SaleService) UpdateSaleAvatar(c *gin.Context, id global.SnowflakeID, file *multipart.FileHeader) (err error) {
	var saleExist system.SysUser
	if errors.Is(global.AM_DB.Where("id = ?", id).First(&saleExist).Error, gorm.ErrRecordNotFound) {
		return &response.CusError{Msg: "用户不存在"}
	}

	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		// 查看文件是否存在，存在则中删除
		if saleExist.AvatarID != nil {
			if err = FileServiceApp.DeleteFileById(c, *saleExist.AvatarID); err != nil {
				return err
			}
		}
		// 上传文件到minio
		fileRes, err := FileServiceApp.UploadFile(c, "avatar", file)
		if err != nil {
			return err
		}
		// 更新用户头像
		saleExist.AvatarID = &fileRes.ID
		if err = tx.Save(&saleExist).Error; err != nil {
			return err
		}
		return nil
	})
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
	err = db.Preload("Branch").Preload("Avatar").Limit(limit).Offset(offset).Find(&saleList).Error
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
	err = db.Preload("Branch").Preload("Avatar").Find(&saleList).Error
	return saleList, total, err
}

func (s *SaleService) DeleteSale(id global.SnowflakeID) (err error) {
	// TODO
	var deletedSale system.SysSale
	err = global.AM_DB.Delete(&deletedSale, id).Error
	fmt.Println(deletedSale.AvatarID)
	return err
}
