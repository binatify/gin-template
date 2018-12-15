package app

import (
	"github.com/binatify/gin-template/app/controllers"
	"github.com/binatify/gin-template/app/middlewares"
	"github.com/gin-gonic/gin"
)

type Application struct {
	*controllers.Application
}

func New(runMode, srcPath string) *Application {
	app := &Application{
		Application: controllers.NewApplication(runMode, srcPath),
	}

	return app
}

func (app *Application) Middlewares() {
	app.Use("*", gin.Recovery())

	app.Use("admin", middlewares.Auth.AuthRequired)
}

func (app *Application) Run() {
	app.Middlewares()

	app.Resource()

	app.Application.Run()
}
