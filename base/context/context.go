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

func NewLoggerMiddleware(l *logrus.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqid, exist := ctx.Get("RqeustId")

		logrus.Infof(">>>>>>> requestid is : %v, exist is %v", reqid, exist)

		// ctx.Set(ContextLoggerKey, logger.NewAppLogger(l, ctx.MustGet("RequestId").(string)))
		ctx.Set(ContextLoggerKey, logger.NewAppLogger(l, reqid.(string)))
	}
}
