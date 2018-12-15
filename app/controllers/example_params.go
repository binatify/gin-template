package controllers

import "github.com/binatify/gin-template/app/models"

type CreateExampleInput struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (in CreateExampleInput) IsValid() bool {
	return true
}

type ListExamplesInput struct {
	Total *int `form:"total"` // NOTE: have to use tag 'form' if this param is in query string, eg: /v1/examples?total=1
}

type ShowExampleOutput struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewShowExampleOutput(example *models.ExampleModel) *ShowExampleOutput {
	return &ShowExampleOutput{
		ID:   example.ID.Hex(),
		Name: example.Name,
	}
}

type UpdateExampleInput struct {
	Name string `json:"name"`
}

func (in UpdateExampleInput) IsValid() bool {
	return true
}
