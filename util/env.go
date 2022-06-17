package util

import "os"

func Env(key string, defaultValue string) string {
	e := os.Getenv(key)
	if e == "" {
		e = defaultValue
	}
	return e
}
