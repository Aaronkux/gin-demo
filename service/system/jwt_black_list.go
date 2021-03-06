package system

import (
	"context"
	"time"

	"go.uber.org/zap"

	"gandi.icu/demo/global"
	"gandi.icu/demo/model/system"
	"github.com/go-redis/redis/v8"
)

type JwtService struct{}

var JwtServiceApp = new(JwtService)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	jwtList.ID = global.SnowflakeID(global.AM_SNOWFLAKE.Generate().Int64())
	err = global.AM_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func (jwtService *JwtService) IsBlacklist(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
	// err := global.AM_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
	// return !isNotFound
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: email string
//@return: err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(email string) (redisJWT string, err error) {
	redisJWT, err = global.AM_REDIS.Get(context.Background(), email).Result()
	return redisJWT, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, email string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, email string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.AM_CONFIG.JWT.ExpiresTime) * time.Second
	err = global.AM_REDIS.Set(context.Background(), email, jwt, timer).Err()
	return err
}

func (jwtService *JwtService) SetEmailBlackList(email string) (err error) {
	if jwtStr, err := JwtServiceApp.GetRedisJWT(email); err != redis.Nil {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := JwtServiceApp.JsonInBlacklist(blackJWT); err != nil {
			return err
		}
	}
	return nil
}

func LoadAll() {
	var data []string
	err := global.AM_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.AM_LOG.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
