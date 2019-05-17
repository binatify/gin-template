package app

import (
	"github.com/binatify/gin-template/base/appconfig"
	"github.com/binatify/gin-template/base/config"
	"github.com/binatify/gin-template/base/logger"
	"github.com/binatify/gin-template/base/runmode"
	"github.com/binatify/gin-template/base/runmodegin"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type Application struct {
	*gin.Engine
	Mode runmode.RunMode

	cfg    *appconfig.AppConfig
	logger *logrus.Logger
}

func NewApplication(runMode runmode.RunMode, srcPath string, cfg interface{}) *Application {
	if err := config.Load(string(runMode), srcPath, &cfg); err != nil {
		panic(err.Error())
	}

	// resolve to app config
	var appCfg *appconfig.AppConfig
	if err := config.Load(string(runMode), srcPath, &appCfg); err != nil {
		panic(err.Error())
	}

	// init logger
	appLogger, err := logger.NewLogger(appCfg.Logger)
	if err != nil {
		panic(err)
	}

	// set gin with logger
	{
		if !appCfg.Logger.IsStdout() {
			gin.DisableConsoleColor()
		}

		gin.DefaultWriter = appLogger.Out
		gin.SetMode(runmodegin.ParseMode(runMode))
	}

	engine := gin.Default()
	appLogger.Infof("Initialized %s in %s mode", appCfg.GetAppName(), runMode)

	return &Application{
		Engine: engine,
		Mode:   runmode.RunMode(runMode),

		cfg:    appCfg,
		logger: appLogger,
	}

}

func (app *Application) Run() {
	//	http sever config
	s := http.Server{
		Addr:           app.cfg.Server.Host,
		Handler:        app.Engine,
		ReadTimeout:    time.Duration(app.cfg.Server.RequestTimeout) * time.Second,
		WriteTimeout:   time.Duration(app.cfg.Server.ResponseTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app.Logger().Infof("Listening on %s", app.cfg.Server.Host)

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

// Logger for app logger getter
func (app *Application) Logger() *logrus.Logger {
	return app.logger
}
