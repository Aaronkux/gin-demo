package system

import (
	"errors"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/system"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var existingUser system.SysUser
	if !errors.Is(global.AM_DB.Where("email = ?", u.Email).First(&existingUser).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("邮箱已被注册")
	}
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return userInter, err
	}
	u.Password = string(encryptedPassword)
	err = global.AM_DB.Create(&u).Error
	return u, err
}
