package apiret

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mockReqID = "mockReqID"
)

func TestNewRet(t *testing.T) {
	assertion := assert.New(t)

	data := []string{
		"a", "b",
	}

	ret := NewRet(data, mockReqID)

	assertion.NotNil(ret)
	assertion.NotNil(ret.Result)
	assertion.Equal(ret.RequestID, mockReqID)
}

func TestRetBindData(t *testing.T) {
	assertion := assert.New(t)

	data := []string{
		"a", "b",
	}

	var (
		unmarshalRet Ret
		bindData     []string
		bindListData []string
	)

	{
		ret := NewRet(data, mockReqID)
		buf, err := json.Marshal(ret)
		assertion.Nil(err)

		assertion.Nil(json.Unmarshal(buf, &unmarshalRet))
		assertion.Equal(unmarshalRet.RequestID, mockReqID)

		assertion.Nil(unmarshalRet.BindData(&bindData))
		assertion.NotEmpty(bindData)
	}

	{
		ret := NewRet(NewList(data), mockReqID)
		buf, err := json.Marshal(ret)
		assertion.Nil(err)

		assertion.Nil(json.Unmarshal(buf, &unmarshalRet))
		assertion.Equal(unmarshalRet.RequestID, mockReqID)

		assertion.Nil(unmarshalRet.BindListData(&bindListData))
		assertion.NotEmpty(bindData)
	}
}
