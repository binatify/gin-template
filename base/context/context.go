package context

import (
	"github.com/binatify/gin-template/base/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	ContextLoggerKey = "_contextLoggerKey"
)

type Context struct {
	*gin.Context
	logger *logrus.Entry
}

func NewLogger(ctx *gin.Context) *logrus.Entry {
	ctxEntry, ok := ctx.Get(ContextLoggerKey)
	if ok {
		return (ctxEntry).(*logrus.Entry)
	}

	return nil
}

func (ctx *Context) Logger() *logrus.Entry {
	if ctx.logger != nil {
		return ctx.logger
	}

	if ctx.Context == nil {
		return nil
	}

	ctx.logger = NewLogger(ctx.Context)

	return ctx.logger
}

func NewHandler(fn func(*Context)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fn(&Context{
			Context: ctx,
		})
	}
}

func NewLoggerMiddleware(l *logrus.Logger, reqIDKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(ContextLoggerKey, logger.NewAppLogger(l, ctx.Request.Header.Get(reqIDKey)))
	}
}
