package appconfig

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppConfig_Copy(t *testing.T) {
	assertion := assert.New(t)

	cfg1 := &AppConfig{}
	cfg2 := cfg1.Copy()

	assertion.NotEqual(fmt.Sprintf("%p", cfg1), fmt.Sprintf("%p", cfg2))
}

func TestAppConfig_GetAppName(t *testing.T) {
	assertion := assert.New(t)

	cfg := &AppConfig{
		Name: "",
	}

	{
		assertion.Equal(cfg.GetAppName(), "app")
	}

	{
		cfg.Name = "epnc-config"
		assertion.Equal(cfg.GetAppName(), cfg.Name)
	}
}

func TestServerConfig(t *testing.T) {
	assertion := assert.New(t)

	jsonStr := `
{
    "host": "0.0.0.0:4002",
    "request_timeout": 30,
    "response_timeout": 30,
	"request_max": 30,
	"throttle": 15,
    "request_id": "x-jdcloud-request-id",
    "request_pin": "x-jdcloud-pin"
}`

	var cfg ServerConfig
	err := json.Unmarshal([]byte(jsonStr), &cfg)
	assertion.Nil(err)

	{
		assertion.Equal(cfg.Host, "0.0.0.0:4002")
		assertion.Equal(cfg.RequestTimeout, 30)
		assertion.Equal(cfg.ResponseTimeout, 30)
		assertion.Equal(cfg.RequestID, "x-jdcloud-request-id")
		assertion.Equal(cfg.RequestPin, "x-jdcloud-pin")

		assertion.Equal(cfg.GetThrottle(), 15)
		assertion.Equal(cfg.GetRequestMax(), 30)

		cfg.Throttle = 0
		cfg.RequestMax = 0

		assertion.NotEqual(cfg.GetThrottle(), 0)
		assertion.NotEqual(cfg.GetRequestMax(), 0)
	}
}
