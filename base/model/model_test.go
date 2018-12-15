package model

import (
	"fmt"
	"testing"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"

	"os"

	"github.com/golib/assert"
	"github.com/sirupsen/logrus"
)

var (
	newMockConfig = func() *Config {
		return &Config{
			Host:     "localhost:27017",
			User:     "root",
			Passwd:   "",
			Database: "testing_model",
		}
	}

	mockLogger = func() *logrus.Logger {
		logger := logrus.New()
		logger.Out = os.Stdout
		return logger
	}()

	mockModel = func() *Model {
		return NewModel(newMockConfig(), mockLogger)
	}()

	mockTestModelIndexes = []mgo.Index{
		{
			Key:    []string{"name"},
			Unique: false,
		},
	}
)

type (
	mockTestModel struct {
		ID   bson.ObjectId `bson:"_id" json:"id"`
		Name string        `bson:"name" json:"name"`
	}
)

func Test_NewModel(t *testing.T) {
	assertion := assert.New(t)

	model := NewModel(newMockConfig(), mockLogger)
	assertion.NotNil(model)
	assertion.NotNil(model.session)
	assertion.Nil(model.collection)
	assertion.NotNil(model.config)
	assertion.NotNil(model.logger)
	assertion.Condition(func() bool {
		return fmt.Sprintf("%v", mockLogger) == fmt.Sprintf("%v", model.logger)
	})
	assertion.Empty(model.indexes)
}

func Test_ModelUse(t *testing.T) {
	assertion := assert.New(t)

	model := NewModel(newMockConfig(), mockLogger)
	assertion.Equal("testing_model", model.Database())

	model.Use("testing_database")
	assertion.Equal("testing_database", model.Database())
}

func Test_ModelCopy(t *testing.T) {
	assertion := assert.New(t)
	model := NewModel(newMockConfig(), mockLogger)

	copiedModel := model.Copy()
	assertion.Condition(func() bool {
		return fmt.Sprintf("%p", model) != fmt.Sprintf("%p", copiedModel)
	})
	assertion.Condition(func() bool {
		return fmt.Sprintf("%p", model.session) != fmt.Sprintf("%p", copiedModel.session)
	})
	assertion.Condition(func() bool {
		return fmt.Sprintf("%p", model.config) != fmt.Sprintf("%p", copiedModel.config)
	})
	assertion.Condition(func() bool {
		return fmt.Sprintf("%v", model.logger) == fmt.Sprintf("%v", copiedModel.logger)
	})

	copiedModel.Use("testing_database")
	assertion.NotEqual(model.Database(), copiedModel.Database())
}

func Test_ModelC(t *testing.T) {
	assertion := assert.New(t)
	model := NewModel(newMockConfig(), mockLogger)

	db := model.C("testing_collection")
	assertion.NotNil(db.collection)
	assertion.Equal(model.Database(), db.Database())
}

func Test_ModelQuery(t *testing.T) {
	assertion := assert.New(t)
	model := NewModel(newMockConfig(), mockLogger)
	test := &mockTestModel{bson.NewObjectId(), "testing"}

	model.Query("testing_collection", mockTestModelIndexes, func(c *mgo.Collection) {
		err := c.Insert(test)
		assertion.Nil(err)
	})
}

func Benchmark_ModelQuery(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testModel := &mockTestModel{bson.NewObjectId(), "testing"}

		mockModel.Query("testing_collection", mockTestModelIndexes, func(c *mgo.Collection) {
			c.Insert(testModel)
		})
	}
}
