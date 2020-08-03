package token

import (
	"github.com/dgrijalva/jwt-go"
	"errors"
	"time"
)

const (
	COOKIE_NAME      = "wgsession"
	TOKEN_CLAIMS_UID = "uid"

	//系统CLAIMS
	TOKEN_CLAIMS_EXP = "exp"
	TOKEN_CLAIMS_IAT = "iat"

	LONG_TOKEN_EXPIRE_DAYS = 30

	SigningKey = "AEIOU12345"
)

type TokenClaims struct {
	Uid int
}

func GenerateToken(uid int) string {
	var claims = jwt.MapClaims{
		TOKEN_CLAIMS_EXP: time.Now().Add(24 * time.Hour * time.Duration(LONG_TOKEN_EXPIRE_DAYS)).Unix(),
		TOKEN_CLAIMS_IAT: time.Now().Unix(),
		TOKEN_CLAIMS_UID: uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tkStr, err := token.SignedString([]byte(SigningKey))
	if err != nil {
		return ""
	}
	return tkStr
}

func ParseTokenClaims(claims jwt.MapClaims) *TokenClaims {
	var tokenClaims TokenClaims
	if value, ok := claims[TOKEN_CLAIMS_UID]; ok {
		tokenClaims.Uid = int(value.(float64))
	}
	return &tokenClaims
}

func ParseTokenUid(tkStr string) (int, error) {
	token, err := jwt.Parse(tkStr, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SigningKey), nil
	})
	if err != nil {
		return 0, err
	}
	if token.Valid {
		tokenClaims := ParseTokenClaims(token.Claims.(jwt.MapClaims))
		if tokenClaims != nil {
			return tokenClaims.Uid, nil
		}
	}
	return 0, errors.New("token不合法")
}
