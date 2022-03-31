package request

import (
	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/request"
)

type Register struct {
	Email        string   `json:"email"`
	Phone        string   `json:"phone"`
	Password     string   `json:"password"`
	NickName     string   `json:"nickname"`
	Avatar       string   `json:"avatar"`
	AuthorityIds []string `json:"authorityIds"`
}

type Login struct {
	Email     string `json:"email"`     // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type UpdateSelf struct {
	NickName string `json:"nickname"`
	Avatar   string `json:"avatar"`
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
	Phone        string             `json:"phone"`
	NickName     string             `json:"nickname"`
	Avatar       string             `json:"avatar"`
	IsActive     *bool              `json:"isActive"`
	AuthorityIds []string           `json:"authorityIds"`
}
