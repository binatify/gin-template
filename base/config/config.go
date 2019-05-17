package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func Load(runMode, srcPath string, cfg interface{}) error {
	configFileName := fmt.Sprintf("application.%s.json", runMode)

	configFilePath := path.Join(srcPath, "config", configFileName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		configFilePath = path.Join(srcPath, "config", "application.json")
	}

	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &cfg)
}
