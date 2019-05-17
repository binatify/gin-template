package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewErrorResponse(t *testing.T) {
	assertion := assert.New(t)
	responseErr := NewErrorResponse("mockReqId", BadRequest)
	assertion.NotNil(responseErr)

	assertion.Equal(responseErr.RequestID, "mockReqId")
	assertion.Equal(responseErr.Error(), "["+BadRequest.Status+"] "+BadRequest.Message)
	assertion.Equal(responseErr.StatusCode(), BadRequest.Code)
}
