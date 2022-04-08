package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gorm.io/gorm"
)

type ClientService struct{}

func (clientService *ClientService) CreateClient(r system.SysClient) (clientRes system.SysClient, err error) {
	r.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	return r, global.AM_DB.Create(&r).Error
}

func (clientService *ClientService) UpdateClient(r system.SysClient) (err error) {
	if !errors.Is(global.AM_DB.Where("id = ?", r.ID).First(&system.SysClient{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("未找到该客户")
	}
	return global.AM_DB.Save(&r).Error
}

func (clientService *ClientService) GetClientList(r systemReq.SearchClientParams) (list interface{}, total int64, err error) {
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysClient{})
	var clientList []system.SysClient
	db = db.Where("client_type = ?", r.ClientType)
	if r.Name != "" {
		db = db.Where("name LIKE ?", "%"+r.Name+"%")
	}
	if r.RegistrationId != 0 {
		db = db.Where("registration_Id = ?", r.RegistrationId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Preload("Sale").Limit(limit).Offset(offset).Order("id desc").Find(&clientList).Error
	return clientList, total, err
}

func (clientService *ClientService) GetClientById(id global.SnowflakeID) (clientRes system.SysClient, err error) {
	return clientRes, global.AM_DB.Preload("Beneficiaries").Where("id = ?", id).First(&clientRes).Error
}

func (clientService *ClientService) DeleteClient(id global.SnowflakeID) (err error) {
	return global.AM_DB.Where("id = ?", id).Delete(&system.SysClient{}).Error
}
