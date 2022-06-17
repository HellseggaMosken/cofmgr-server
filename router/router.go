package router

import (
	"cofmgr/api"
	"cofmgr/middleware"
	"cofmgr/util"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors([]string{util.Env("ALLOW_ORIGIN", "*")}))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", api.Ping)
	}

	return r
}
