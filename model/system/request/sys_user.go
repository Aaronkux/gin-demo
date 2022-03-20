package request

import "encoding/json"

type Register struct {
	Email        string        `json:"email"`
	Phone        *string       `json:"phone"`
	Password     string        `json:"password"`
	NickName     string        `json:"nickname"`
	Avatar       string        `json:"avatar"`
	AuthorityIds []json.Number `json:"authorityIds"`
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
