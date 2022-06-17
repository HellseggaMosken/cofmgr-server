package main

import (
	"cofmgr/conf"
	"cofmgr/logger"
	"cofmgr/router"
)

func main() {
	conf.Init()

	r := router.NewRouter()
	if err := r.Run(conf.ServerAddress); err != nil {
		logger.Fatal(err)
	}
}
