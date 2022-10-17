package router

import (
	"mygram/controller"
	"mygram/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controller.RegisterUser)
		userRouter.POST("/login", controller.LoginUser)
		userRouter.PUT("/:userId", middleware.Authentication(), middleware.UserAuthorization(), controller.UpdateUserData)
		userRouter.DELETE("/:userId", middleware.Authentication(), middleware.UserAuthorization(), controller.DeleteUserAccount)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Authentication())
		photoRouter.POST("/", controller.PostPhoto)
		photoRouter.GET("/", controller.GetPhotos)
		photoRouter.PUT("/:photoId", middleware.PhotoAuthorization(), controller.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middleware.PhotoAuthorization(), controller.DeletePhoto)
	}
	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Authentication())
		commentRouter.POST("/", controller.PostComment)
		commentRouter.GET("/", controller.GetComments)
		commentRouter.PUT("/:commentId", middleware.CommentAuthorization(), controller.UpdateComment)
		commentRouter.DELETE("/:commentId", middleware.CommentAuthorization(), controller.DeleteComment)
	}
	sosmedRouter := r.Group("/socialmedias")
	{
		sosmedRouter.Use(middleware.Authentication())
		sosmedRouter.POST("/", controller.PostSocialMedia)
		sosmedRouter.GET("/", controller.GetSocialMedias)
		sosmedRouter.PUT("/:socialMediaId", middleware.SosmedAuthorization(), controller.UpdateSocialMedia)
		sosmedRouter.DELETE("/:socialMediaId", middleware.SosmedAuthorization(), controller.DeleteSocialMedia)
	}

	return r
}
