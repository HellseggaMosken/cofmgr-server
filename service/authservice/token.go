package authservice

import (
	"cofmgr/logger"

	"github.com/golang-jwt/jwt/v4"
)

func CreateToken(id string, isAdmin bool) (token string) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		IsAdmin:  isAdmin,
		ID:       id,
		ExpireAt: expireTime(),
	})

	token, err := jwtToken.SignedString(secret)
	if err != nil {
		logger.Panic("when create token:", err)
	}
	return
}

func ParseToken(token string) (isAdmin bool, id string, expired bool, valid bool) {
	jwtToken, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		logger.Warning(err)
		return false, "", false, false
	}

	if c, ok := jwtToken.Claims.(*claims); ok && jwtToken.Valid {
		return c.IsAdmin, c.ID, c.expired(), true
	}
	return false, "", false, false
}
