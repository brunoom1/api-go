package routes

import (
	"github.com/GA-Marketing/service-viability/controllers"
	"github.com/GA-Marketing/service-viability/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	routeSecurity := app.Group("")
	routeSecurity.Use(middlewares.CheckToken())

	var userController controllers.UserController
	routeSecurity.GET("/users", userController.List)
	routeSecurity.GET("/users/types", userController.UserTypes)
	routeSecurity.POST("/users", userController.Create)
	routeSecurity.GET("/users/me", userController.Me)

	routesPublics := app.Group("")
	routesPublics.POST("/users/auth/login", userController.Login)

}
