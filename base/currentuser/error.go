package currentuser

import (
	"github.com/binatify/gin-template/base/errors"
)

var (
	NotLoginError = errors.Error{401, "Unauthorized", "login please"}
)
