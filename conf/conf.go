package conf

import (
	"cofmgr/logger"
	"cofmgr/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ServerAddress string

func Init() {
	if err := godotenv.Load(); err != nil {
		logger.Info("Not detect .env file")
	}

	logger.Init()

	ServerAddress = util.Env("SERVER_ADDRESS", ":8888")

	gin.SetMode(util.Env("GIN_MODE", "debug"))

	// fileservice.Init()

	// authservice.Init()

	// serializer.InitSerializer()

	// model.InitDatabase()
}
