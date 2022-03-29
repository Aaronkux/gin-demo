package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Register(r systemReq.Register) (userRes system.SysUser, err error) {
	// 检查权限ID是否存在
	var authorities []system.SysAuthority
	if err := global.AM_DB.Where("id In ?", r.AuthorityIds).Find(&authorities).Error; err != nil {
		return userRes, err
	}
	if len(authorities) != len(r.AuthorityIds) {
		return userRes, &response.CusError{Msg: "权限ID不存在"}
	}
	// 检查邮箱是否存在
	var existingUser system.SysUser
	if !errors.Is(global.AM_DB.Where("email = ?", r.Email).First(&existingUser).Error, gorm.ErrRecordNotFound) {
		return userRes, &response.CusError{Msg: "邮箱已被注册"}
	}

	newUser := system.SysUser{Email: r.Email, NickName: r.NickName, Password: r.Password, Avatar: r.Avatar, Authorities: authorities}
	var encryptedPassword []byte
	if encryptedPassword, err = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost); err != nil {
		return userRes, err
	}
	newUser.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	newUser.Password = string(encryptedPassword)
	err = global.AM_DB.Create(&newUser).Error
	return newUser, err
}

func (userService *UserService) Login(r systemReq.Login) (userRes system.SysUser, err error) {
	var user system.SysUser
	if err := global.AM_DB.Where("email = ?", r.Email).Preload("Authorities").First(&user).Error; err != nil {
		return userRes, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
		return userRes, err
	}
	return user, err
}

func (userService *UserService) UpdateSelf(r systemReq.UpdateSelf, id global.SnowflakeID) (userRes system.SysUser, err error) {
	user := system.SysUser{NickName: r.NickName, Avatar: r.Avatar, Phone: r.Phone}
	if err := global.AM_DB.Where("id = ?", id).Preload("Authorities").First(&userRes).Select("NickName", "Avatar", "Phone").Updates(&user).Error; err != nil {
		return userRes, err
	}
	return userRes, err
}
func (userService *UserService) UpdateUser(r systemReq.UpdateUser) (userRes system.SysUser, err error) {

	user := system.SysUser{NickName: r.NickName, Avatar: r.Avatar, Phone: r.Phone}
	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", r.ID).First(&userRes).Select("NickName", "Avatar", "Phone").Updates(&user).Error; err != nil {
			return err
		}

		var authorities []system.SysAuthority
		if err := tx.Where("id In ?", r.AuthorityIds).Find(&authorities).Error; err != nil {
			return err
		}
		if err := tx.Model(&userRes).Association("Authorities").Replace(authorities); err != nil {
			return err
		}
		return nil
	})
	return userRes, err
}

func (userService *UserService) GetUserList(r systemReq.SearchUserParams) (list interface{}, total int64, err error) {
	var userList []system.SysUser
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysUser{})
	if r.NickName != "" {
		db = db.Where("nick_name like ?", "%"+r.NickName+"%")
	}
	if r.Email != "" {
		db = db.Where("email like ?", "%"+r.Email+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

func (userService *UserService) GetUserById(id global.SnowflakeID) (userRes system.SysUser, err error) {
	err = global.AM_DB.Where("id = ?", id).First(&userRes).Error
	return userRes, err
}
