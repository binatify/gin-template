package app

import (
	"github.com/binatify/gin-template/example/app/controllers"
	"github.com/binatify/gin-template/example/app/middlewares"
	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/runmode"
	"github.com/gin-gonic/gin"
)

type Application struct {
	*controllers.Application
}

func New(runMode runmode.RunMode, srcPath string) *Application {
	app := &Application{
		Application: controllers.NewApplication(runMode, srcPath),
	}

	return app
}

func (app *Application) Middlewares() {
	app.Use("*",context.NewLoggerMiddleware(app.Logger()), gin.Recovery())

	app.Use("v1",context.NewLoggerMiddleware(app.Logger()))

	app.Use("admin", middlewares.Auth.AuthRequired)
}

func (app *Application) Run() {
	app.Middlewares()

	app.Resource()

	app.Application.Run()
}
