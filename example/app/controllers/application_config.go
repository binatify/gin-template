package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/binatify/gin-template/base/logger"
	"github.com/binatify/gin-template/base/model"
	"github.com/binatify/gin-template/base/runmode"
	"io/ioutil"
	"os"
	"path"
)

type AppConfig struct {
	Name   string         `json:"name"`
	Server *ServerConfig  `json:"server"`
	Logger *logger.Config `json:"logger"`
	Mongo  *model.Config  `json:"mongo"`
}

type ServerConfig struct {
	Host            string `json:"host"`
	RequestTimeout  int    `json:"request_timeout"`
	ResponseTimeout int    `json:"response_timeout"`
}

func NewAppConfig(runMode runmode.RunMode, srcPath string) error {
	configFileName := fmt.Sprintf("application.%s.json", runMode)

	configFilePath := path.Join(srcPath, "config", configFileName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		configFilePath = path.Join(srcPath, "config", "application.json")
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &Config)
}
