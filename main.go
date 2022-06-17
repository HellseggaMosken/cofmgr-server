package main

import (
	"cofmgr/conf"
	_ "cofmgr/docs"
	"cofmgr/logger"
	"cofmgr/router"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           会议投稿系统
// @version         1.0
// @description     会议投稿系统的后端接口文档。
// @basePath        /api/v1

func main() {
	conf.Init()

	r := router.NewRouter()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	if err := r.Run(conf.ServerAddress); err != nil {
		logger.Fatal(err)
	}
}
