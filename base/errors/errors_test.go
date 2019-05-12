package errors

import (
	"github.com/golib/assert"
	"testing"
)

func Test_IsEmptyError(t *testing.T) {
	assertion := assert.New(t)

	assertion.True(Error{}.IsEmptyError())
	assertion.False(BadRequest.IsEmptyError())
}

func Test_Error(t *testing.T) {
	assertion := assert.New(t)
	assertion.Equal(BadRequest.Error(), "400:BadRequest")
}

func Test_Error_String(t *testing.T) {
	assertion := assert.New(t)
	assertion.Equal(BadRequest.String(), "BadRequest: Bad Request")
}

func Test_Error_SetMsg(t *testing.T) {
	assertion := assert.New(t)
	err := InvalidParameter
	msg := "Customerd Msg"
	err.SetMsg(msg)
	assertion.Equal(err.Message ,msg)
}
