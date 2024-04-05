package v1

import (
	v1Controller "boilerplate/app/http/controllers/v1"
	"boilerplate/app/http/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		userController := v1Controller.NewUserController()

		authRoute := v1.Group("/auth")
		{
			authRoute.POST("/register", userController.RegisterUser)
			authRoute.POST("/login", userController.Login)
			authRoute.POST("/refresh_token", userController.Login)
		}

		usersRoute := v1.Group("/users")
		{
			usersRoute.POST("/email_checkers", userController.CheckEmailAvailability)
			usersRoute.POST("/avatars", middlewares.AuthMiddleware(), userController.UploadAvatar)
			usersRoute.GET("/users/fetch", middlewares.AuthMiddleware(), userController.FetchUser)
		}
	}
}
