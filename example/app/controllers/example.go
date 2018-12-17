package controllers

import (
	"github.com/binatify/gin-template/example/app/models"
	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/errors"
)

var (
	Example *_Example
)

type _Example struct{}

func (_ *_Example) Create(ctx *context.Context) {
	var input *CreateExampleInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Logger().Errorf("Unmarshal json error: %v", err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	example := models.NewExampleModel(input.Name)
	example.Phone = input.Phone
	if err := example.Save(); err != nil {
		ctx.Logger().Errorf("exmaple.Save(): %v", err)
		ErrHandler(ctx, errors.InternalError)
		return
	}

	res := NewShowExampleOutput(example)
	ResponseJSON(ctx, res)
}

func (_ *_Example) Update(ctx *context.Context) {
	id := ctx.Param("id")

	example, err := models.Example.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.Example.Find(%v): %v", id, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	var input *UpdateExampleInput
	if err := ctx.BindJSON(&input); err != nil {
		ctx.Logger().Errorf("Unmarshal json error: %v", err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	example.Name = input.Name
	if err := example.Save(); err != nil {
		ctx.Logger().Errorf("example.Save(): %v", err)
		ErrHandler(ctx, errors.InternalError)
		return
	}

	ResponseJSON(ctx, input)
}

func (_ *_Example) Show(ctx *context.Context) {
	id := ctx.Param("id")

	example, err := models.Example.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.Example.Find(%v): %v", id, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	res := NewShowExampleOutput(example)
	ResponseJSON(ctx, res)
}

func (_ *_Example) All(ctx *context.Context) {
	var input ListExamplesInput // DO NOT use pointer

	if err := ctx.ShouldBindQuery(&input); err != nil {
		ctx.Logger().Errorf("ctx.BindQuery(%v): %v", input, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	filter := models.FilterExample{}

	examples, err := models.Example.List(100, &filter)
	if err != nil {
		ctx.Logger().Errorf("models.Example.List(%v, %v): %v", 100, filter, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	res := make([]*ShowExampleOutput, 0)
	for _, v := range examples {
		res = append(res, NewShowExampleOutput(v))
	}

	ResponseJSON(ctx, res)
}

func (_ *_Example) Remove(ctx *context.Context) {
	id := ctx.Param("id")

	example, err := models.Example.Find(id)
	if err != nil {
		ctx.Logger().Errorf("models.Example.Find(%v): %v", id, err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	if err := example.Delete(); err != nil {
		ctx.Logger().Errorf("example.Delete(): %v", err)
		ErrHandler(ctx, errors.InvalidParameter)
		return
	}

	ResponseJSON(ctx, "Success")
}
