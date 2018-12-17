package controllers

import (
	"encoding/json"
	"net/http"

	"{{.Module}}/app/common"
	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/errors"
)

func ErrHandler(ctx *context.Context, err errors.Error, msg ...string) {
	requestId := ctx.MustGet(common.REQUEST_ID).(string)

	if len(msg) > 0 {
		err.Message = msg[0]
	}

	ctx.JSON(err.Code, errors.NewErrorResponse(requestId, err))
}

type CommonResponse struct {
	RequestId string      `json:"requestId"`
	Result    interface{} `json:"result"`
}

func NewCommonResponse(requestId string, result interface{}) *CommonResponse {
	return &CommonResponse{
		RequestId: requestId,
		Result:    result,
	}
}

func ResponseJSON( ctx *context.Context, result interface{}, code ...int){
	StatusCode := http.StatusOK
	if len(code) > 0{
		StatusCode = code[0]
	}


	ctx.JSON(StatusCode, NewCommonResponse(ctx.RequestID(), result))
}

func (resp *CommonResponse) BindResultData(obj interface{}) error {
	data, err := json.Marshal(resp.Result)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &obj)
}
