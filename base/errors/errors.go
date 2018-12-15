package errors

import "strconv"

type Error struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (e Error) IsEmptyError() bool {
	return e.Code == 0 && e.Status == "" && e.Message == ""
}

func (e Error) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Status
}

func (e Error) String() string {
	return e.Status + ": " + e.Message
}
