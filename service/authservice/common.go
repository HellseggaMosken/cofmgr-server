package authservice

import (
	"os"
)

var secret []byte

const expireDuration = 2 * 24 * 60 * 60 // 2 days in seconds

func Init() {
	secret = []byte(os.Getenv("AUTH_SECRET"))
}
