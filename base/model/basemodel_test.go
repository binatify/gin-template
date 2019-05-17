package model

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	_mockTestCollection *_mockTestBaseModelCollection

	mockTestCollection        = "examples"
	mockTestCollectionIndexes = []mgo.Index{}
)

type (
	mockTestBaseModel struct {
		Name      string `bson:"name" json:"name"`
		BaseModel `bson:",inline"`
	}

	_mockTestBaseModelCollection struct{}
)

func (model *mockTestBaseModel) IsValid() bool {
	return true
}

func (model *mockTestBaseModel) C() Collection {
	return _mockTestCollection
}

func (model *mockTestBaseModel) Save() (err error) {
	query := bson.M{
		"name": model.Name,
	}

	return Save(model, query)
}

func (*_mockTestBaseModelCollection) Query(query func(c *mgo.Collection)) {
	mockModel.Query(mockTestCollection, mockTestCollectionIndexes, query)
}

func (*_mockTestBaseModelCollection) Find(id string) (ret *mockTestBaseModel, err error) {
	err = Find(_mockTestCollection, id, &ret)
	return
}

func Test_NewMockTestBaseModel(t *testing.T) {
	assertion := assert.New(t)

	testBaseModel := &mockTestBaseModel{
		Name:      "testname",
		BaseModel: NewBaseModel(),
	}

	err := testBaseModel.Save()
	assertion.Nil(err)
	assertion.NotNil(testBaseModel.ID)

	// test found model
	foundModel, err := _mockTestCollection.Find(testBaseModel.ID.Hex())
	assertion.Nil(err)
	assertion.Equal(foundModel.ID.Hex(), testBaseModel.ID.Hex())
}
