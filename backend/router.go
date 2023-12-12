package main

import (
	"backend/controller"
	"backend/middleware/jwt"
	"backend/middleware/rabbitmq"
	"backend/middleware/redis"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// 基础接口
	apiRouter.GET("/video/recommended", jwt.AuthWithoutLogin(), controller.Feed)
	//apiRouter.GET("/feed/", jwt.AuthWithoutLogin(), controller.Feed)
	apiRouter.GET("/me/my", jwt.AuthWithoutLogin(), controller.UserInfo)
	apiRouter.GET("/user/profile/other", jwt.AuthWithoutLogin(), controller.UserInfo)
	apiRouter.GET("/user/number_existence", controller.UserExistence)
	apiRouter.POST("/user/getcode", controller.GenerateCode)
	apiRouter.POST("/user/password_change", controller.PasswordChange)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/password_login/", controller.PasswordLogin)
	apiRouter.POST("/user/login", controller.Login)
	// apiRouter.POST("/publish/action/", jwt.AuthBody(), controller.Publish)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/me/mypost", jwt.AuthWithoutLogin(), controller.PublishList)

	// 互动接口
	apiRouter.POST("/favorite/action/", jwt.AuthWithoutLogin(), controller.FavoriteAction)
	apiRouter.GET("/me/like", jwt.AuthWithoutLogin(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", jwt.Auth(), controller.CommentAction)
	apiRouter.GET("/comment/list/", jwt.AuthWithoutLogin(), controller.CommentList)

	// 社交接口
	apiRouter.POST("/relation/action/", jwt.Auth(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", jwt.AuthWithoutLogin(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", jwt.AuthWithoutLogin(), controller.FollowerList)
	apiRouter.GET("/relation/friend/list/", jwt.Auth(), controller.FriendList)
	apiRouter.GET("/message/chat/", jwt.Auth(), controller.MessageChat)
	apiRouter.POST("/message/action/", jwt.Auth(), controller.MessageAction)

	// 测试接口
	apiRouter.POST("/test/", jwt.Auth(), controller.Test)
}

func initMiddleware() {
	redis.InitRedis()
	rabbitmq.InitRabbitMQ()
	rabbitmq.InitLikeRabbitMQ()
	rabbitmq.InitFollowRabbitMQ()
	rabbitmq.InitCommentRabbitMQ()
}
