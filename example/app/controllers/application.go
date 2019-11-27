package controllers

import (
	"github.com/binatify/gin-template/base/app"
	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/ping"
	"github.com/binatify/gin-template/base/runmode"
	"github.com/binatify/gin-template/example/app/models"

	"github.com/gin-gonic/gin"
)

var (
	APP    *Application
	Config *AppConfig
)

type Application struct {
	*app.Application
	v1    *gin.RouterGroup
	admin *gin.RouterGroup
}

func NewApplication(runMode runmode.RunMode, srcPath string) *Application {
	application := app.NewApplication(runMode, srcPath, &Config)
	// setup model
	models.SetupModelWithConfig(Config.Mongo, application.Logger())

	APP = &Application{
		Application: application,

		v1:    application.Group("/v1"),
		admin: application.Group("/admin"),
	}

	return APP
}

// use for app middlewares inject
func (app *Application) Use(route string, middlewares ...gin.HandlerFunc) {
	switch route {
	case "*":
		app.Engine.Use(middlewares...)
		app.v1.Use(middlewares...)
		app.admin.Use(middlewares...)

	case "v1":
		app.v1.Use(middlewares...)

	case "admin":
		app.admin.Use(middlewares...)

	default:
		panic("Unkown route of " + route)
	}
}

// Resources for routes inject
func (app *Application) Resource() {
	app.GET("/ping", ping.PongHandler)
	app.v1.POST("/examples", context.NewHandler(Example.Create))
	app.v1.PUT("/examples/:id", context.NewHandler(Example.Update))
	app.v1.GET("/examples/:id", context.NewHandler(Example.Show))
	app.v1.GET("/examples", context.NewHandler(Example.All))
}
