package request

import (
	"mime/multipart"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
)

type Register struct {
	Email        string                `form:"email"`
	Phone        string                `form:"phone"`
	Password     string                `form:"password"`
	NickName     string                `form:"nickname"`
	Avatar       *multipart.FileHeader `form:"avatar"`
	AuthorityIds []string              `form:"authorityIds"`
}

type Login struct {
	Email     string `json:"email"`     // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type UpdateSelf struct {
	NickName string `json:"nickname"`
	Phone    string `json:"phone"`
}

type SearchUserParams struct {
	request.PageInfo
	NickName string `json:"nickname"`
	Email    string `json:"email"`
	IsActive *bool  `json:"isActive"`
}

type UpdateUser struct {
	ID           global.SnowflakeID `json:"id"`
	Email        string             `json:"email"`
	Phone        string             `json:"phone"`
	NickName     string             `json:"nickname"`
	IsActive     *bool              `json:"isActive"`
	AuthorityIds []string           `json:"authorityIds"`
}

type UpdateUserAvatar struct {
	ID   global.SnowflakeID    `form:"id"`
	File *multipart.FileHeader `form:"file"`
}
