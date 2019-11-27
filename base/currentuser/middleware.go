package currentuser

import (
	"github.com/binatify/gin-template/base/context"
	"github.com/binatify/gin-template/base/errors"
	"github.com/binatify/gin-template/base/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	Middleware *_Middleware
)

type _Middleware struct{}

func (*_Middleware) ShouldLogin(ctx *gin.Context) {
	session := sessions.Default(ctx)
	sessionStore := session.Get(SessionKey)

	reqID := ctx.GetString(context.RequestId)
	currentUser, ok := sessionStore.(*User)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, errors.NewErrorResponse(reqID, NotLoginError))
		ctx.Abort()
		return
	}

	ctx.Set(SessionKey, currentUser)
	ctx.Next()
}

func (*_Middleware) ShouldBe(roles ...UserRole) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		cUser := ctx.MustGet(SessionKey).(*User)

		if cUser.InRoles(roles) {
			ctx.Next()
			return
		}

		reqID := ctx.GetString(context.RequestId)
		ctx.JSON(http.StatusForbidden, errors.NewErrorResponse(reqID, errors.AccessDenied))
		ctx.Abort()
	}
}
