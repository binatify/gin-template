package httplimiter

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
)

func NewHttpLimiterMiddleware(r rate.Limit, b int) gin.HandlerFunc {
	limiter := rate.NewLimiter(r, b)

	return func(ctx *gin.Context) {
		if !limiter.Allow() {
			ctx.String(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
