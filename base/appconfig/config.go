package appconfig

import "github.com/binatify/gin-template/base/logger"

type AppConfig struct {
	Name   string         `json:"name"`
	Server *ServerConfig  `json:"server"`
	Logger *logger.Config `json:"logger"`
}

func (c *AppConfig) Copy() *AppConfig {
	cfg := *c

	return &cfg
}

func (c *AppConfig) GetAppName() string {
	if c.Name == "" {
		return "app"
	}

	return c.Name
}

type ServerConfig struct {
	Host            string `json:"host"`
	RequestTimeout  int    `json:"request_timeout"`
	ResponseTimeout int    `json:"response_timeout"`

	Throttle   int `json:"throttle"`
	RequestMax int `json:"request_max"`

	RequestID  string `json:"request_id"`
	RequestPin string `json:"request_pin"`
}

func (cfg *ServerConfig) GetThrottle() int {
	if cfg.Throttle == 0 {
		return 100
	}

	return cfg.Throttle
}

func (cfg *ServerConfig) GetRequestMax() int {
	if cfg.RequestMax == 0 {
		return 5 * cfg.GetThrottle()
	}

	return cfg.RequestMax
}
