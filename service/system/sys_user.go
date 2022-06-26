package system

import (
	"errors"
	"mime/multipart"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
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
	var userExist system.SysUser
	if !errors.Is(global.AM_DB.Where("email = ?", r.Email).First(&userExist).Error, gorm.ErrRecordNotFound) {
		return userRes, &response.CusError{Msg: "邮箱已被注册"}
	}

	newUser := system.SysUser{Email: r.Email, NickName: r.NickName, Password: r.Password, AvatarID: r.AvatarID, Phone: r.Phone, Authorities: authorities}
	var encryptedPassword []byte
	if encryptedPassword, err = bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost); err != nil {
		return userRes, err
	}
	newUser.Password = string(encryptedPassword)
	err = global.AM_DB.Create(&newUser).Error
	return newUser, err
}

func (userService *UserService) Login(r systemReq.Login) (userRes system.SysUser, err error) {
	var user system.SysUser
	if err := global.AM_DB.Where("email = ? and is_active = ?", r.Email, true).Preload("Avatar").Preload("Authorities").First(&user).Error; err != nil {
		return userRes, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password)); err != nil {
		return userRes, err
	}
	return user, err
}

func (userService *UserService) UpdateSelf(r systemReq.UpdateSelf, id global.SnowflakeID) (err error) {
	var userExist system.SysUser
	user := system.SysUser{NickName: r.NickName, Phone: r.Phone}
	err = global.AM_DB.Where("id = ?", id).First(&userExist).Select("NickName", "Phone").Updates(&user).Error
	return err
}

func (userService *UserService) UpdateSelfAvatar(c *gin.Context, id global.SnowflakeID, file *multipart.FileHeader) (fileRes system.SysFile, err error) {
	var userExist system.SysUser
	if errors.Is(global.AM_DB.Where("id = ?", id).First(&userExist).Error, gorm.ErrRecordNotFound) {
		return fileRes, &response.CusError{Msg: "用户不存在"}
	}

	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		// 查看文件是否存在，存在则中删除
		if err = FileServiceApp.DeleteFileById(c, *userExist.AvatarID); err != nil {
			return err
		}
		// 上传文件到minio
		fileRes, err = FileServiceApp.UploadFile(c, "avatar", file)
		if err != nil {
			return err
		}
		// 更新用户头像
		userExist.AvatarID = &fileRes.ID
		if err = tx.Save(&userExist).Error; err != nil {
			return err
		}
		return nil
	})
	return fileRes, err
}

func (userService *UserService) UpdateUser(r systemReq.UpdateUser) (err error) {
	var oldUser system.SysUser
	user := system.SysUser{NickName: r.NickName, Phone: r.Phone, IsActive: *r.IsActive}
	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Preload("Authorities").Where("id = ?", r.ID).First(&oldUser).Error; err != nil {
			return err
		}
		oldAuthorities := oldUser.Authorities
		if err = tx.Model(&oldUser).Select("NickName", "Phone", "IsActive").Updates(&user).Error; err != nil {
			return err
		}
		var authorities []system.SysAuthority
		if err = tx.Where("id In ?", r.AuthorityIds).Find(&authorities).Error; err != nil {
			return err
		}
		if err = tx.Model(&oldUser).Association("Authorities").Replace(authorities); err != nil {
			return err
		}
		var oldAuthorityIds []string
		for _, authority := range oldAuthorities {
			oldAuthorityIds = append(oldAuthorityIds, authority.ID.String())
		}
		if !utils.SameStringSlice(oldAuthorityIds, r.AuthorityIds) || !user.IsActive {
			if err = JwtServiceApp.SetEmailBlackList(oldUser.Email); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

func (userService *UserService) UpdateUserAvatar(c *gin.Context, id global.SnowflakeID, file *multipart.FileHeader) (fileRes system.SysFile, err error) {
	var userExist system.SysUser
	if errors.Is(global.AM_DB.Where("id = ?", id).First(&userExist).Error, gorm.ErrRecordNotFound) {
		return fileRes, &response.CusError{Msg: "用户不存在"}
	}

	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {
		// 查看文件是否存在，存在则中删除
		if err = FileServiceApp.DeleteFileById(c, *userExist.AvatarID); err != nil {
			return err
		}
		// 上传文件到minio
		fileRes, err = FileServiceApp.UploadFile(c, "avatar", file)
		if err != nil {
			return err
		}
		// 更新用户头像
		userExist.AvatarID = &fileRes.ID
		if err = tx.Save(&userExist).Error; err != nil {
			return err
		}
		return nil
	})
	return fileRes, err
}

func (userService *UserService) GetUserList(r systemReq.SearchUserParams) (list interface{}, total int64, err error) {
	var userList []system.SysUser
	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.AM_DB.Model(&system.SysUser{})
	db.Where("is_active = ?", *r.IsActive)
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
	err = db.Preload("Authorities").Limit(limit).Offset(offset).Find(&userList).Error
	return userList, total, err
}

func (userService *UserService) GetUserById(id global.SnowflakeID) (userRes system.SysUser, err error) {
	err = global.AM_DB.Preload("Authorities").Where("id = ?", id).First(&userRes).Error
	return userRes, err
}

func (userService *UserService) DeleteUser(id global.SnowflakeID) (err error) {
	var user system.SysUser
	err = global.AM_DB.Transaction(func(tx *gorm.DB) error {

		if err = tx.Where("id = ?", id).First(&user).Error; err != nil {
			return err
		}
		if err = tx.Delete(&user).Error; err != nil {
			return err
		}
		if err = tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
			return err
		}
		if err := JwtServiceApp.SetEmailBlackList(user.Email); err != nil {
			return err
		}
		return nil
	})

	return err
}
