package routes

import (
	"github.com/GA-Marketing/service-viability/controllers"
	"github.com/GA-Marketing/service-viability/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {

	routeSecurity := app.Group("")
	routeSecurity.Use(middlewares.CheckToken())

	routeSecurity.GET("/users", controllers.UsersList)
	routeSecurity.GET("/users/types", controllers.UsersTypes)
	routeSecurity.POST("/users", controllers.UsersCreate)
	routeSecurity.GET("/users/me", controllers.UserMe)

	routesPublics := app.Group("")
	routesPublics.POST("/users/auth/login", controllers.UsersAuthLogin)

}
