package conf

import (
	"os"

	"cofmgr/logger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ServerAddress = ":8888"

func Init() {
	if err := godotenv.Load(); err != nil {
		logger.Info("Not detect .env file")
	}

	logger.Init()

	if addr := os.Getenv("SERVER_ADDRESS"); addr != "" {
		ServerAddress = addr
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	// fileservice.Init()

	// authservice.Init()

	// serializer.InitSerializer()

	// model.InitDatabase()
}
