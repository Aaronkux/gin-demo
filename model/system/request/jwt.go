package request

import (
	"gandi.icu/demo/global"
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type BaseClaims struct {
	ID           global.SnowflakeID
	Email        string
	NickName     string
	AuthorityIds []string
}
