package authservice

import (
	"time"
)

type claims struct {
	IsAdmin  bool   `json:"is_admin"`
	ID       string `json:"id"`
	ExpireAt int64  `json:"expire_at"`
}

func (c *claims) Valid() error {
	return nil
}

func (c *claims) expired() bool {
	return time.Now().Unix() >= c.ExpireAt
}

func expireTime() int64 {
	return time.Now().Unix() + expireDuration
}
