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
