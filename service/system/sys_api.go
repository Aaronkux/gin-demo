package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gandi.icu/demo/utils"

	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

type ApiService struct{}

var ApiServiceApp = new(ApiService)

func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	api.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	if !errors.Is(global.AM_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return global.AM_DB.Create(&api).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApi
//@description: 删除基础api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	if err = global.AM_DB.Where("id = ?", api.ID).First(&oldA).Error; err != nil {
		return err
	}
	err = global.AM_DB.Delete(&oldA).Error
	println("delete", oldA.Path, oldA.Method)
	CasbinServiceApp.ClearCasbin(1, oldA.Path, oldA.Method)
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAPIInfoList
//@description: 分页获取数据,
//@param: api model.SysApi, info request.PageInfo, order string, desc bool
//@return: err error

func (apiService *ApiService) GetAPIInfoList(r systemReq.SearchApiParams) (list interface{}, total int64, err error) {
	order := r.OrderKey
	desc := r.Desc
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysApi{})
	var apiList []system.SysApi

	if r.Path != "" {
		db = db.Where("path LIKE ?", "%"+r.Path+"%")
	}

	if r.Description != "" {
		db = db.Where("description LIKE ?", "%"+r.Description+"%")
	}

	if r.Method != "" {
		db = db.Where("method = ?", r.Method)
	}

	if r.ApiGroup != "" {
		db = db.Where("api_group = ?", r.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			order := utils.CamelToSnake(order)
			var OrderStr string
			// 设置有效排序key 防止sql注入
			orderMap := make(map[string]bool, 4)
			orderMap["path"] = true
			orderMap["api_group"] = true
			orderMap["description"] = true
			orderMap["method"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return apiList, total, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAllApis
//@description: 获取所有的api
//@return: err error, apis []model.SysApi

func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.AM_DB.Order("api_group").Find(&apis).Error
	return apis, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApiById
//@description: 根据id获取api
//@param: id float64
//@return: err error, api model.SysApi

func (apiService *ApiService) GetApiById(id global.SnowflakeID) (api system.SysApi, err error) {
	err = global.AM_DB.Where("id = ?", id).First(&api).Error
	return api, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApi
//@description: 根据id更新api
//@param: api model.SysApi
//@return: err error

func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.AM_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(global.AM_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
			if err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method); err != nil {
				return err
			}
			if err = tx.Save(&api).Error; err != nil {
				return err
			}

			return nil
		})
	}
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApis
//@description: 删除选中API
//@param: apis []model.SysApi
//@return: err error

func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	err = global.AM_DB.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
	return err
}
