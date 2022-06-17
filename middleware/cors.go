package middleware

import (
	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors(origins []string) gin.HandlerFunc {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Cookie", "token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowWebSockets:  true,
		AllowOrigins:     origins,
	}
	// 测试环境下模糊匹配本地开头的请求
	if gin.Mode() == gin.DebugMode {
		config.AllowOriginFunc = func(origin string) bool {
			if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
				return true
			}
			if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
				return true
			}
			return false
		}
	}
	return cors.New(config)
}
