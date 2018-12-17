package models

import (
	"encoding/json"
	"github.com/binatify/gin-template/base/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestMain(m *testing.M) {
	srcPath := path.Clean("../../")

	var modelConfig struct {
		Mongo *model.Config `json:"mongo"`
	}

	cfg, _ := ioutil.ReadFile(srcPath + "/config/application.test.json")
	json.Unmarshal(cfg, &modelConfig)

	// setup logger
	logger := logrus.New()
	logger.Out = os.Stdout

	// start up MongoDB
	SetupModelWithConfig(modelConfig.Mongo, logger)

	// database clear up
	code := m.Run()
	cleanDatabase()
	os.Exit(code)
}

// clean database
func cleanDatabase() {
	mongo := Model()
	mongo.Session().DB(mongo.Database()).DropDatabase()
}

func Test_SetupModel(t *testing.T) {
	model := &model.Model{}
	tmp := mongo

	mongo = model
	SetupModel(model)

	// restore variable mongo
	mongo = tmp
}
