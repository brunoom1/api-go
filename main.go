package main

import (
	"github.com/GA-Marketing/service-viability/db"
	"github.com/GA-Marketing/service-viability/middlewares"
	"github.com/GA-Marketing/service-viability/migration"
	"github.com/GA-Marketing/service-viability/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	db.InitDatabase()
	app.Use(middlewares.PassDbMiddleware(db.GetCurrentConnection()))

	migration.Run()

	routes.InitRoutes(app)
	app.Run()
}
