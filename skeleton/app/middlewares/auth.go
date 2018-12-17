package middlewares

import (
	"github.com/gin-gonic/gin"
)

var (
	Auth *_Auth
)

type _Auth struct{}

func (_ *_Auth) AuthRequired(ctx *gin.Context) {
}
