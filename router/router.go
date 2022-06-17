package router

import (
	"os"

	"cofmgr/api"
	"cofmgr/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors([]string{os.Getenv("ALLOW_ORIGIN")}))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", api.Ping)
	}
	// {
	// 	v1.GET("/ping", api.Ping)

	// 	auth := v1.Group("/auth")
	// 	{
	// 		auth.POST("/getToken", api.GetToken)
	// 		auth.POST("/refreshToken", api.RefreshToken)
	// 	}

	// 	user := v1.Group("/users")
	// 	{
	// 		user.POST("", api.UserRegister)

	// 		user.Use(middleware.AuthRequired())

	// 		user.GET("self", api.UserSelf)
	// 		user.GET("show/:id", api.UserShow)
	// 		user.GET("", api.UserList)
	// 		user.PUT("password", api.UserChangePassword)
	// 		user.PUT("", api.UserUpdate)
	// 	}

	// 	friend := v1.Group("/friends")
	// 	{
	// 		friend.Use(middleware.AuthRequired())

	// 		friend.GET("", api.FriendList)

	// 		friend.Use(middleware.IsFriend("fid"))

	// 		friend.GET(":fid", api.FriendShow)
	// 		friend.PUT(":fid", api.FriendUpdate)
	// 		friend.DELETE(":fid", api.FriendRemove)
	// 	}

	// 	group := v1.Group("/groups")
	// 	{
	// 		group.Use(middleware.AuthRequired())

	// 		group.GET("", api.GroupList)
	// 		group.GET(":id", api.GroupShow)
	// 		group.POST("", api.GroupCreate)

	// 		owner := group.Group("")
	// 		owner.Use(middleware.IsGroupOwner("gid"))
	// 		{
	// 			owner.PUT(":gid", api.GroupUpdate)
	// 			owner.DELETE(":gid", api.GroupRemove)
	// 		}

	// 	}

	// 	member := v1.Group("/members/:gid")
	// 	{
	// 		member.Use(middleware.AuthRequired())

	// 		member1 := member.Group("")
	// 		member1.Use(middleware.IsMember("gid"))
	// 		{
	// 			// member1.GET("", api.MemberList)
	// 			member1.PUT("", api.MemberUpdate)
	// 			member1.GET(":mid", api.MemberShow)
	// 		}
	// 		// 此处是个临时的写法，建议改进
	// 		member2 := member.Group(":mid")
	// 		member2.Use(middleware.IsMemberWithMemberID("gid", "mid"))
	// 		member2.DELETE("", api.MemberRemove)
	// 	}

	// 	action := v1.Group("/actions")
	// 	{
	// 		action.Use(middleware.AuthRequired())

	// 		friendAction := action.Group("/friends")
	// 		{
	// 			friendAction.PUT("/reject", api.FriendRequestReject)
	// 			friendAction.PUT("/accept", api.FriendRequestAccept)
	// 			friendAction.DELETE("", api.FriendRequestRemove)

	// 			friendAction.Use(middleware.ValidReceiver("userid"))
	// 			friendAction.POST("/:userid", api.FriendRequestSend)
	// 		}

	// 		memberAction := action.Group("/members")
	// 		{
	// 			memberAction.PUT("/reject", api.MemberRequestReject)
	// 			memberAction.PUT("/accept", api.MemberRequestAccept)
	// 			memberAction.DELETE("", api.MemberRequestRemove)

	// 			memberAction.Use(middleware.ValidReceiver("userid"))
	// 			memberAction.POST("/:userid", api.MemberRequestSend)
	// 		}

	// 		groupAction := action.Group("/groups")
	// 		{
	// 			groupAction.PUT("/reject", api.GroupInvitationReject)
	// 			groupAction.PUT("/accept", api.GroupInvitationAccept)
	// 			groupAction.DELETE("", api.GroupInvitationRemove)

	// 			groupAction.Use(middleware.ValidReceiver("userid"))
	// 			groupAction.POST("/:userid", api.GroupInvitationSend)
	// 		}
	// 	}

	// 	message := v1.Group("/messages")
	// 	{
	// 		message.Use(middleware.AuthRequired())

	// 		friendMessage := message.Group("/friends")
	// 		friendMessage.Use(middleware.IsFriend("fid"))
	// 		friendMessage.POST(":fid", api.FriendMessageSend)

	// 		groupMessage := message.Group("/groups")
	// 		groupMessage.Use(middleware.IsMember("gid"))
	// 		groupMessage.POST(":gid", api.GroupMessageSend)
	// 	}

	// 	push := v1.Group("/push")
	// 	{
	// 		push.Use(middleware.AuthRequiredFromQuery())
	// 		push.GET("/register/websocket", api.PushRegisterWebsocket)
	// 	}

	// 	file := v1.Group("/files")
	// 	{
	// 		file.Use(middleware.AuthRequired())
	// 		file.POST("/url/download", api.FileGetDownloadURL)
	// 		file.POST("/url/upload", api.FileGetUploadURL)
	// 	}
	// }

	return r
}
