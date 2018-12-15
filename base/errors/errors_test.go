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
