package conf

import (
	"cofmgr/logger"
	"cofmgr/model"
	"cofmgr/service/authservice"
	"cofmgr/util"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var ServerAddress string

func Init() {
	if err := godotenv.Load(); err != nil {
		logger.Info("Not detect .env file")
	}

	logger.Debug("inti logger")
	logger.Init()

	logger.Debug("inti server address")
	ServerAddress = util.Env("SERVER_ADDRESS", ":8888")

	logger.Debug("inti gin mode")
	gin.SetMode(util.Env("GIN_MODE", "debug"))

	logger.Debug("inti authservice")
	authservice.Init()

	// logger.Debug("inti fileservice")
	// fileservice.Init()

	logger.Debug("inti database")
	model.InitDatabase()
}
