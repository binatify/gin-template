package app

import (
	"github.com/atarantini/ginrequestid"
	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/runmode"
	"github.com/binatify/gin-template/example/app/controllers"
	"github.com/binatify/gin-template/example/app/middlewares"
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
	app.Use("*", ginrequestid.RequestId(), context.NewLoggerMiddleware(app.Logger()))
	app.Use("admin", middlewares.Auth.AuthRequired)
}

func (app *Application) Run() {
	app.Middlewares()

	app.Resource()

	app.Application.Run()
}
