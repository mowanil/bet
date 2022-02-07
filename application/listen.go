package application

import (
	"github.com/gin-gonic/gin"
	"github.com/mowahaeser/bet/controllers"
	"github.com/mowahaeser/bet/inputs"
	"github.com/mowahaeser/bet/middlewares"
)

func Listen() {
	router := gin.Default()

	v1 := router.Group("v1")
	{
		auth := v1.Group("auth")
		{
			auth.POST("register", middlewares.Body(inputs.Register{}), controllers.Account.Register)
			auth.POST("login", middlewares.Body(inputs.Login{}), controllers.Account.Login)
			auth.POST("logout", controllers.Account.Logout)
		}
	}

	router.Run(":42069")
}
