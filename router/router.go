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

	v1.GET("/ping", api.Ping)

	{
		user := v1.Group("/users")
		user.POST("", api.UserRegister)
		{
			user := user.Group("")
			user.Use(middleware.AuthRequired(false))
			user.PUT("", api.UserUpdate)
		}

		{
			user := user.Group("")
			user.Use(middleware.AuthRequired(true))
			user.GET(":id", api.UserShow)
			user.GET("", api.UserList)
		}
	}

	{
		password := v1.Group("/password")
		{
			password := password.Group("")
			password.Use(middleware.AuthRequired(false))
			password.PUT("user", api.PasswordChangeForUser)
		}
		{
			password := password.Group("")
			password.Use(middleware.AuthRequired(true))
			password.PUT("admin", api.PasswordChangeForAdmin)
		}
	}

	{
		login := v1.Group("/login")
		login.POST("user", api.LoginUser)
		login.POST("admin", api.LoginAdmin)
	}

	{
		cof := v1.Group("/conferences")
		cof.GET(":id", api.ConferenceShow)
		cof.GET("", api.ConferenceList)
		{
			cof := cof.Group("")
			cof.Use(middleware.AuthRequired(true))
			cof.POST("", api.ConferenceCreate)
			cof.PUT("", api.ConferenceUpdate)
		}
	}

	{
		ctb := v1.Group("/contributions")
		ctb.GET("common/:id", api.ContributionShow)
		{
			ctb.Use(middleware.AuthRequired(false))
			ctb.POST("", api.ContributionCreate)
			ctb.PUT("", api.ContributionUpdate)
			ctb.GET("user", api.ContributionListWithUser)
			ctb.GET("conference/:cof_id", api.ContributionListWithConference)
		}
	}

	{
		referee := v1.Group("/referees")
		referee.Use(middleware.AuthRequired(true))
		referee.POST("conference/:cof_id/user/user_id", api.RefereeAddForConference)
		referee.DELETE("conference/:cof_id/user/user_id", api.RefereeRemoveForConference)
		referee.POST("contribution/:ctb_id/user/user_id", api.RefereeAddForContribution)
		referee.DELETE("contribution/:ctb_id/user/user_id", api.RefereeRemoveForContribution)
	}

	return r
}
