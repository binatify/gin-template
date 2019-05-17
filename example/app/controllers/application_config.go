package controllers

import (
	"github.com/binatify/gin-template/base/appconfig"
	"github.com/binatify/gin-template/base/model"
)

type AppConfig struct {
	*appconfig.AppConfig
	Mongo *model.Config `json:"mongo"`
}

func (c *AppConfig) Copy() *AppConfig {
	cfg := *c

	return &cfg
}
