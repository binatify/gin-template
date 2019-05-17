package app

import (
	"fmt"
	"github.com/binatify/gin-template/base/appconfig"
	"github.com/binatify/gin-template/base/runmode"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"
)

var (
	mockApp     *Application
	mockSrcPath string
	appCfg      *appconfig.AppConfig
)

func TestMain(m *testing.M) {
	rand.Seed(time.Now().UnixNano())

	mockApp = NewApplication(runmode.Test, path.Clean("./"), &appCfg)
	mockApp.cfg = appCfg

	code := m.Run()
	os.Exit(code)
}

func TestNewApplication(t *testing.T) {
	assertion := assert.New(t)

	// panic test
	{
		func() {
			defer func() {
				if r := recover(); r != nil {
					assertion.Contains(r, "no such file")
				}
			}()
			NewApplication(runmode.Test, ".", &appCfg)
		}()
	}

	// logger file test
	{
		{
			fileName := path.Join(mockSrcPath, "config", "application.logfile.json")
			f, err := os.Create(fileName)
			assertion.Nil(err)
			f.WriteString(`{
  "logger": {
    "output": "logfile.log",
    "level": "debug"
  }
}`)
			f.Close()
			defer os.Remove(fileName)
			defer os.Remove("./logfile.log")

			NewApplication(runmode.RunMode("logfile"), mockSrcPath, &appCfg)
			assertion.Nil(err)

			// check logger file exist
			_, err = os.Stat("./logfile.log")
			assertion.Nil(err)
		}
	}

	// logger file with error
	func() {
		defer func() {
			if r := recover(); r != nil {
				assertion.Contains(fmt.Sprintf("%s", r), "open logs/logfile.log")
			}
		}()

		{
			fileName := path.Join(mockSrcPath, "config", "application.logfile.json")
			f, err := os.Create(fileName)
			assertion.Nil(err)
			f.WriteString(`{
  "logger": {
    "output": "logs/logfile.log",
    "level": "debug"
  }
}`)
			f.Close()
			defer os.Remove(fileName)
			defer os.RemoveAll("./logs")

			NewApplication(runmode.RunMode("logfile"), mockSrcPath, &appCfg)
		}
	}()

	// logger file with error
	func() {
		defer func() {
			if r := recover(); r != nil {
				assertion.Contains(fmt.Sprintf("%s", r), "invalidlevel")
			}
		}()

		{
			fileName := path.Join(mockSrcPath, "config", "application.logfile.json")
			f, err := os.Create(fileName)
			assertion.Nil(err)
			f.WriteString(`{
  "logger": {
    "output": "logfile.log",
    "level": "invalidlevel"
  }
}`)
			NewApplication(runmode.RunMode("logfile"), mockSrcPath, &appCfg)
		}
	}()
}

func TestApplicationRun(t *testing.T) {
	assertion := assert.New(t)

	func() {
		host := appCfg.Server.Host

		defer func() {
			if r := recover(); r != nil {
				assertion.Contains(fmt.Sprintf("%s", r), "missing port in")
			}
			appCfg.Server.Host = host
		}()

		appCfg.Server.Host = "////" // make server start with error
		mockApp.Run()
	}()

	appCfg.Server.Host = "////"
	// no meaningful, a trickery for test cover
	go mockApp.Run()
}

func TestApplicationLogger(t *testing.T) {
	assertion := assert.New(t)

	assertion.NotNil(mockApp.Logger())
	assertion.Equal(mockApp.logger, mockApp.Logger())
}
