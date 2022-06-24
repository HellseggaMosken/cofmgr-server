package router

import (
	"cofmgr/api"
	md "cofmgr/middleware"
	"cofmgr/util"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(md.Cors([]string{util.Env("ALLOW_ORIGIN", "*")}))
	v1 := r.Group("/api/v1")

	v1.GET("/ping", api.Ping)

	{
		user := v1.Group("/users")
		user.POST("", api.UserRegister)
		{
			user := user.Group("")
			user.Use(md.AuthRequired(md.AuthRequiredUser))
			user.PUT("", api.UserUpdate)
		}

		{
			user := user.Group("")
			user.Use(md.AuthRequired(md.AuthRequiredAdmin))
			user.GET(":id", api.UserShow)
			user.GET("", api.UserList)
		}
	}

	{
		password := v1.Group("/password")
		{
			password := password.Group("")
			password.Use(md.AuthRequired(md.AuthRequiredUser))
			password.PUT("user", api.PasswordChangeForUser)
		}
		{
			password := password.Group("")
			password.Use(md.AuthRequired(md.AuthRequiredAdmin))
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
			cof.Use(md.AuthRequired(md.AuthRequiredAdmin))
			cof.POST("", api.ConferenceCreate)
			cof.PUT("", api.ConferenceUpdate)
		}
	}

	{
		ctb := v1.Group("/contributions")
		ctb.GET("common/:id", api.ContributionShow)
		{
			ctb.Use(md.AuthRequired(md.AuthRequiredUser))
			ctb.POST("", api.ContributionCreate)
			ctb.PUT("", api.ContributionUpdate)
			ctb.GET("user", api.ContributionListWithUser)
			ctb.GET("conference/:cof_id", api.ContributionListWithConference)
		}
	}

	{
		referee := v1.Group("/referees")
		{
			referee := referee.Group("")
			referee.Use(md.AuthRequired(md.AuthRequiredUser))
			referee.GET("conferences", api.RefereeConferenceListForCurrent)
			referee.GET("contributions", api.RefereeContributionListForCurrent)
		}
		{
			referee := referee.Group("")
			referee.Use(md.AuthRequired(md.AuthRequiredAdmin))
			referee.GET("conference/:cof_id", api.RefereeListForConference)
			referee.GET("conferences/:referee_user_id", api.RefereeConferenceList)
			referee.GET("contributions/:referee_user_id", api.RefereeContributionList)
			referee.POST("conference/:cof_id/user/:user_id", api.RefereeAddForConference)
			referee.DELETE("conference/:cof_id/user/:user_id", api.RefereeRemoveForConference)
			referee.POST("contribution/:ctb_id/user/:user_id", api.RefereeAddForContribution)
			referee.DELETE("contribution/:ctb_id/user/:user_id", api.RefereeRemoveForContribution)
		}
	}

	{
		oss := v1.Group("/oss")
		oss.Use(md.AuthRequired(md.AuthRequiredAny))
		oss.POST("upload", api.OSSGetUploadURL)
		oss.POST("download", api.OSSGetUDownloaddURL)
	}

	return r
}
