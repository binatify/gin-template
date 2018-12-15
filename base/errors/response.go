package errors

type ErrorResponse struct {
	RequestID string `json:"requestId"`
	Err       Error  `json:"error"`
}

func NewErrorResponse(requestID string, err Error) *ErrorResponse {
	return &ErrorResponse{
		RequestID: requestID,
		Err:       err,
	}
}

func (er *ErrorResponse) StatusCode() int {
	return er.Err.Code
}

func (er *ErrorResponse) Error() string {
	return "[" + er.Err.Status + "] " + er.Err.Message
}
