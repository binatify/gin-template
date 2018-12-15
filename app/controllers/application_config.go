package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/binatify/gin-template/base/model"
)

type AppConfig struct {
	Server  *ServerConfig `json:"server"`
	Logger  *LoggerConfig `json:"logger"`
	Mongo   *model.Config `json:"mongo"`
}

type ServerConfig struct {
	Host            string `json:"host"`
	RequestTimeout  int    `json:"request_timeout"`
	ResponseTimeout int    `json:"response_timeout"`
}

type LoggerConfig struct {
	Output string `json:"output"`
	Level  string `json:"level"`
}

func (logger *LoggerConfig) IsFileOutput() bool {
	return logger.Output != "stdout"
}

func NewAppConfig(runMode, srcPath string) error {
	configFileName := fmt.Sprintf("application.%s.json", runMode)

	configFilePath := path.Join(srcPath, "config", configFileName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		configFilePath = path.Join(srcPath, "config", "application.json")
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(data, &Config); err != nil {
		return err
	}

	return nil
}

