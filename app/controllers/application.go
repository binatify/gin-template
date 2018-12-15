package controllers

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/binatify/gin-template/app/models"
	"github.com/binatify/gin-template/base/runmode"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	APP    *Application
	Config *AppConfig
)

type Application struct {
	*gin.Engine
	Mode runmode.RunMode

	appLogger *logrus.Logger
	v1        *gin.RouterGroup
	admin     *gin.RouterGroup
}

func NewApplication(runMode, srcPath string) *Application {
	if err := NewAppConfig(runMode, srcPath); err != nil {
		panic(err.Error())
	}

	// set gin with logger
	{
		if Config.Logger.IsFileOutput() {
			gin.DisableConsoleColor()

			f, err := os.Create(Config.Logger.Output)
			if err != nil {
				panic(err)
			}

			gin.DefaultWriter = io.MultiWriter(f)
		}

		gin.SetMode(Config.Logger.Level)
	}

	engine := gin.Default()

	// set logger
	appLogger := logrus.New()
	appLogger.SetOutput(gin.DefaultWriter)

	// setup model
	models.SetupModelWithConfig(Config.Mongo, appLogger)


	APP = &Application{
		Engine:    engine,
		Mode:      runmode.RunMode(runMode),
		appLogger: appLogger,

		v1:    engine.Group("/v1"),
		admin: engine.Group("/admin"),
	}

	return APP
}

func (app *Application) Run() {
	// http server config
	s := http.Server{
		Addr:           Config.Server.Host,
		Handler:        app.Engine,
		ReadTimeout:    time.Duration(Config.Server.RequestTimeout) * time.Second,
		WriteTimeout:   time.Duration(Config.Server.ResponseTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app.appLogger.Infof("Run app at: %v\n", s.Addr)

	s.ListenAndServe()
}

func (app *Application) Logger() *logrus.Logger {
	return app.appLogger
}

// use for app middlewares inject
func (app *Application) Use(route string, middlewares ...gin.HandlerFunc) {
	switch route {
	case "*":
		app.Engine.Use(middlewares...)

	case "admin":
		app.admin.Use(middlewares...)

	default:
		panic("Unkown route of " + route)
	}
}

// Resources for routes inject
func (app *Application) Resource() {
	app.v1.POST("/examples", Example.Create)
	app.v1.PUT("/examples/:id", Example.Update)
	app.v1.GET("/examples/:id", Example.Show)
	app.v1.GET("/examples", Example.All)
}
