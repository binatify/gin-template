package controllers

import (
	"encoding/json"

	"github.com/binatify/gin-template/app/common"
	"github.com/binatify/gin-template/base/errors"
	"github.com/gin-gonic/gin"
)

func ErrHandler(ctx *gin.Context, err errors.Error, msg ...string) {
	requestId := ctx.GetHeader(common.REQUEST_ID)

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

func (resp *CommonResponse) BindResultData(obj interface{}) error {
	data, err := json.Marshal(resp.Result)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &obj)
}
