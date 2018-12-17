package models

import (
	"github.com/binatify/gin-template/base/model"
	"github.com/sirupsen/logrus"
)

var (
	mongo *model.Model
)

func SetupModel(model *model.Model) {
	mongo = model
}

func SetupModelWithConfig(config *model.Config, logger *logrus.Logger) {
	mongo = model.NewModel(config, logger)
}

func Model() *model.Model {
	return mongo
}
