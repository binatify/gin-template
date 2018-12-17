package controllers

import (
	"html/template"
	"os"
	"testing"

	"github.com/golib/assert"
)

func Test_AutoGenerate(t *testing.T) {
	assertion := assert.New(t)

	input := &TemplateModel{
		Name:          "Order",
		LowerCaseName: "order",
		Project:       "github.com/binatify/gin-template",
	}

	// generate controller file
	controllerFileName := input.LowerCaseName + ".go"
	controllerTemplate := template.Must(template.New("controller").Parse(ctlStr))
	controllerFile, err := os.OpenFile(controllerFileName, os.O_CREATE|os.O_WRONLY, 0644)
	assertion.Nil(err)
	err = controllerTemplate.Execute(controllerFile, input)
	assertion.Nil(err)

	// generate controller params file
	ctlParamsFileName := input.LowerCaseName + "_params.go"
	ctlParamsTemplate := template.Must(template.New("controller params").Parse(ctlParamsStr))
	ctlParamsFile, err := os.OpenFile(ctlParamsFileName, os.O_CREATE|os.O_WRONLY, 0644)
	assertion.Nil(err)
	err = ctlParamsTemplate.Execute(ctlParamsFile, input)
	assertion.Nil(err)

	// generate models file
	modelFileName := "../models/" + input.LowerCaseName + ".go"
	modelTemplate := template.Must(template.New("models").Parse(modelStr))
	modelFile, err := os.OpenFile(modelFileName, os.O_CREATE|os.O_WRONLY, 0644)
	assertion.Nil(err)
	err = modelTemplate.Execute(modelFile, input)
	assertion.Nil(err)
}

type TemplateModel struct {
	Name          string
	LowerCaseName string
	Project       string
}

var ctlStr = `package controllers

import (
	"{{.Project}}/app/models"
	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/errors"
)

var (
	{{.Name}} *_{{.Name}} 
)

type _{{.Name}} struct{}

func (_ *_{{.Name}}) Create(ctx *context.Context) {
	var input *Create{{.Name}}Input
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Logger().Errorf("Unmarshal json error: %v", err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	{{.LowerCaseName}} := models.New{{.Name}}Model()
	if err := {{.LowerCaseName}}.Save(); err != nil {
		ctx.Logger().Errorf("{{.LowerCaseName}}.Save(): %v", err)
		ErrHandler(ctx, errors.InternalError)
		return
	}

	res := NewShow{{.Name}}Output({{.LowerCaseName}})
	ResponseJSON(ctx, res)
}

func (_ *_{{.Name}}) Update(ctx *context.Context) {
	id := ctx.Param("id")

	{{.LowerCaseName}}, err := models.{{.Name}}.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.{{.Name}}.Find(%v): %v", id, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	var input *Update{{.Name}}Input
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Logger().Errorf("Unmarshal json error: %v", err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	//TODO: update the fields
	if err := {{.LowerCaseName}}.Save(); err != nil {
		ctx.Logger().Errorf("{{.LowerCaseName}}.Save(): %v", err)
		ErrHandler(ctx, errors.InternalError)
		return
	}

	ResponseJSON(ctx, input)
}

func (_ _{{.Name}}) Show(ctx *context.Context) {
	id := ctx.Param("id")

	{{.LowerCaseName}}, err := models.{{.Name}}.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.{{.Name}}.Find(%v): %v", id, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	res := NewShow{{.Name}}Output({{.LowerCaseName}})
	ResponseJSON(ctx, res)
}

func (_ _{{.Name}}) All(ctx *context.Context) {
	var input List{{.Name}}sInput // DO NOT use pointer
	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.Logger().Errorf("ctx.BindQuery(%v): %v", input, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	filter := models.Filter{{.Name}}{}

	{{.LowerCaseName}}s, err := models.{{.Name}}.List(100, &filter)
	if err != nil {
		ctx.Logger().Errorf("models.{{.Name}}.List(%v, %v): %v", 100, filter, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	res := make([]*Show{{.Name}}Output, 0)
	for _, v := range {{.LowerCaseName}}s {
		res = append(res, NewShow{{.Name}}Output(v))
	}

	ResponseJSON(ctx, res)
}

func (_ _{{.Name}}) Remove(ctx *context.Context) {
	id := ctx.Param("id")

	{{.LowerCaseName}}, err := models.{{.Name}}.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.{{.Name}}.Find(%v): %v", id, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	if err := {{.LowerCaseName}}.Delete(); err != nil {
		ctx.Logger().Errorf("{{.LowerCaseName}}.Delete(): %v", err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	ResponseJSON(ctx, "Success")
}
`

var ctlParamsStr = `package controllers

import "{{.Project}}/app/models"

type Create{{.Name}}Input struct {

}

func (in Create{{.Name}}Input) IsValid() bool {
	return true
}

type List{{.Name}}sInput struct {

}

type Show{{.Name}}Output struct {

}

func NewShow{{.Name}}Output({{.LowerCaseName}} *models.{{.Name}}Model) *Show{{.Name}}Output {
	return &Show{{.Name}}Output{

	}
}

type Update{{.Name}}Input struct {

}

func (in Update{{.Name}}Input) IsValid() bool {
	return true
}
`

var modelStr = `package models

import (
	db "github.com/binatify/gin-template/base/model"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type _{{.Name}} struct{}

var (
	{{.Name}} *_{{.Name}}

	{{.LowerCaseName}}Collection = "{{.LowerCaseName}}"
	{{.LowerCaseName}}Indexes    = []mgo.Index{

	}
)

type {{.Name}}Model struct {
	db.BaseModel ` + "`" + `bson:",inline"` + "`" + `
}

func New{{.Name}}Model() *{{.Name}}Model {
	return &{{.Name}}Model{
		BaseModel: db.NewBaseModel(),
	}
}

func (model *{{.Name}}Model) IsValid() bool {
	return true
}

func (model *{{.Name}}Model) Save() (err error) {
	query := bson.M{}

	return db.Save(model, query)
}

func ({{.LowerCaseName}} *_{{.Name}}) Find(id string) (r *{{.Name}}Model, err error) {
	err = db.Find({{.LowerCaseName}}, id, &r)
	return
}

func ({{.LowerCaseName}} *_{{.Name}}) BatchInsert({{.LowerCaseName}}s []*{{.Name}}Model)(err error){
	t := time.Now()

	res := make([]interface{}, len({{.LowerCaseName}}s))
	for i, v := range {{.LowerCaseName}}s{
		v.CreatedAt = t
		res[i] = v
	}

	return db.BatchInsert({{.LowerCaseName}}, res...)
}

type Filter{{.Name}} struct{}

func (filter *Filter{{.Name}}) Resolve() bson.M{
	query := bson.M{}
	return query
}

func ({{.LowerCaseName}} *_{{.Name}}) List(total int, filter *Filter{{.Name}}) (r []*{{.Name}}Model, err error) {
	query := filter.Resolve()
	err = db.Where({{.LowerCaseName}}, query, EnsureWithinMaxItems(total), &r)
	return
}

func (model *{{.Name}}Model) Delete() (err error) {
	return db.Destroy(model)
}

func (model *{{.Name}}Model) C() db.Collection {
	return {{.Name}}
}

func (_ *_{{.Name}}) Query(query func(c *mgo.Collection)) {
	Model().Query({{.LowerCaseName}}Collection, {{.LowerCaseName}}Indexes, query)
}
`

