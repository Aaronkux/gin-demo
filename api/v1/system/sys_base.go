package system

import (
	"strconv"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/common/response"
	"gandi.icu/demo/model/system"
	systemReq "gandi.icu/demo/model/system/request"
	systemRes "gandi.icu/demo/model/system/response"
	"gandi.icu/demo/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var r systemReq.Login
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if store.Verify(r.CaptchaId, r.Captcha, true) {
		if user, err := userService.Login(r); err != nil {
			global.AM_LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			response.FailWithMessage("用户名不存在或者密码错误", c)
		} else {
			b.tokenNext(c, user)
		}
	} else {
		response.FailWithMessage("验证码错误", c)
	}
}

// 登录以后签发jwt
func (b *BaseApi) tokenNext(c *gin.Context, user system.SysUser) {
	j := utils.NewJWT()
	var authorityIds []string
	for _, authority := range user.Authorities {
		authorityIds = append(authorityIds, authority.ID.String())
	}
	claims := j.CreateClaims(systemReq.BaseClaims{
		ID:           user.ID,
		NickName:     user.NickName,
		Email:        user.Email,
		AuthorityIds: authorityIds,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.AM_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.AM_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: strconv.FormatInt(claims.StandardClaims.ExpiresAt*1000, 10),
		}, "登录成功", c)
		return
	}
	if jwtStr, err := jwtService.GetRedisJWT(user.Email); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Email); err != nil {
			global.AM_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: strconv.FormatInt(claims.StandardClaims.ExpiresAt*1000, 10),
		}, "登录成功", c)
	} else if err != nil {
		global.AM_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Email); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: strconv.FormatInt(claims.StandardClaims.ExpiresAt*1000, 10),
		}, "登录成功", c)
	}
}
