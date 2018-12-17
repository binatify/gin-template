package models

import (
	db "github.com/binatify/gin-template/base/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type _Example struct{}

var (
	Example *_Example

	exampleCollection = "example"
	exampleIndexes    = []mgo.Index{
		{
			Key:    []string{"name"},
			Unique: false,
		},
		{
			Key:    []string{"phone"},
			Unique: false,
		},
	}
)

type ExampleModel struct {
	Name  string `bson:"name"`
	Phone string `bson:"phone"`

	db.BaseModel `bson:",inline"`
}

func NewExampleModel(name string) *ExampleModel {
	return &ExampleModel{
		Name: name,

		BaseModel: db.NewBaseModel(),
	}
}

func (model *ExampleModel) IsValid() bool {
	return true
}

func (model *ExampleModel) Save() (err error) {
	query := bson.M{
		"name":  model.Name,
		"phone": model.Phone,
	}

	return db.Save(model, query)
}

func (example *_Example) Find(id string) (r *ExampleModel, err error) {
	err = db.Find(example, id, &r)
	return
}

func (example *_Example) BatchInsert(examples []*ExampleModel) (err error) {
	t := time.Now()

	res := make([]interface{}, len(examples))
	for i, v := range examples {
		v.CreatedAt = t
		res[i] = v
	}

	return db.BatchInsert(example, res...)
}

type FilterExample struct {
}

func (filter *FilterExample) Resolve() bson.M {
	query := bson.M{}
	return query
}

func (example *_Example) List(total int, filter *FilterExample) (r []*ExampleModel, err error) {
	query := filter.Resolve()
	err = db.Where(example, query, EnsureWithinMaxItems(total), &r)
	return
}

func (model *ExampleModel) Delete() (err error) {
	return db.Destroy(model)
}

func (model *ExampleModel) C() db.Collection {
	return Example
}

func (_ *_Example) Query(query func(c *mgo.Collection)) {
	Model().Query(exampleCollection, exampleIndexes, query)
}
