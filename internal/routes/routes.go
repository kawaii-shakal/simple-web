package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	// Роуты для аутентификации
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", auth.LoginHandler)
		authRouter.POST("/signup", auth.SignupHandler)
	}
}
