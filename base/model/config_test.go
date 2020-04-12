package model

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewConfig(t *testing.T) {
	assertion := assert.New(t)

	appConfigStr := `
	  {
		"host": "localhost:27017",
		"user": "root",
		"password": "",
		"database": "testing_model",
		"mode": "Strong",
		"pool": 5,
		"timeout": 5,
		"replica": "mgset-500148149"
	  }
	`

	modelConfig, err := NewConfig([]byte(appConfigStr))
	assertion.Nil(err)
	assertion.Equal("localhost:27017", modelConfig.Host)
	assertion.Equal("root", modelConfig.User)
	assertion.Equal("", modelConfig.Passwd)
	assertion.Equal("testing_model", modelConfig.Database)
	assertion.Equal("Strong", modelConfig.Mode)
	assertion.Equal(5, modelConfig.Pool)
	assertion.Equal(5, modelConfig.Timeout)
	assertion.Equal("mongodb://localhost:27017/testing_model?replicaSet=mgset-500148149", modelConfig.DSN())
}

func Test_ConfigCopy(t *testing.T) {
	config := new(Config)
	copiedConfig := config.Copy()

	assertion := assert.New(t)
	assertion.NotEqual(fmt.Sprintf("%p", config), fmt.Sprintf("%p", copiedConfig))
}

func Test_Config_GetUser(t *testing.T) {
	assertion := assert.New(t)

	cfg := &Config{
		User: "user",
	}

	assertion.Equal(cfg.GetUser(), "user")

	cfg.User = ""
	os.Setenv(MONGODB_USER, "MONGODB_USER")
	assertion.Equal(cfg.GetUser(), MONGODB_USER)
	os.Setenv(MONGODB_USER, "")
}

func Test_Config_GetPasswd(t *testing.T) {
	assertion := assert.New(t)

	cfg := &Config{
		Passwd: "password",
	}

	assertion.Equal(cfg.GetPasswd(), "password")

	cfg.Passwd = ""
	os.Setenv(MONGODB_PASSWORD, "MONGODB_PASSWORD")
	assertion.Equal(cfg.GetPasswd(), MONGODB_PASSWORD)
	os.Setenv(MONGODB_PASSWORD, "")
}
