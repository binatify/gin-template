package model

import (
	"encoding/json"
	"os"
)

type Config struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Passwd   string `json:"password"`
	Database string `json:"database"`
	Mode     string `json:"mode"`
	Pool     int    `json:"pool"`
	Timeout  int    `json:"timeout"`
}

const (
	MONGODB_USER     = "MONGODB_USER"
	MONGODB_PASSWORD = "MONGODB_PASSWORD"
)

func NewConfig(buff []byte) (*Config, error) {
	var c Config
	err := json.Unmarshal(buff, &c)
	return &c, err
}

func (c *Config) Copy() *Config {
	config := *c

	return &config
}

func (c *Config) GetUser() string {
	if c.User != "" {
		return c.User
	}

	return os.Getenv(MONGODB_USER)
}

func (c *Config) GetPasswd() string {
	if c.Passwd != "" {
		return c.Passwd
	}

	return os.Getenv(MONGODB_PASSWORD)
}
