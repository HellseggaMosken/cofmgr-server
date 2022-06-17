package model

import (
	"cofmgr/logger"

	"golang.org/x/crypto/bcrypt"
)

func PasswordDigest(password string) string {
	const passwordCost = 12 // 密码加密难度
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), passwordCost)
	if err != nil {
		logger.Panic("when encrypt user password:", password, err)
	}
	return string(bytes)
}

func PasswordCompare(password string, digest string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(password))
	return err == nil
}
