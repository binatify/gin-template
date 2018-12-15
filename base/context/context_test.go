package context

import (
	"fmt"
	"github.com/binatify/gin-template/base/logger"
	"github.com/gin-gonic/gin"
	"github.com/golib/assert"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestNewLogger(t *testing.T) {
	assertion := assert.New(t)

	ctx := &gin.Context{}

	log := NewLogger(ctx)
	assertion.Nil(log)

	log = NewLogger(NewMockCtxWithLogger())
	assertion.NotNil(log)
}

func TestContext_Logger(t *testing.T) {
	assertion := assert.New(t)

	mockCtx := NewMockCtxWithLogger()

	ctx := Context{}
	assertion.Nil(ctx.Logger())

	ctx.Context = mockCtx
	assertion.NotNil(ctx.Logger())

	ctx.Context = nil
	assertion.NotNil(ctx.Logger())
}

func TestNewHandler(t *testing.T) {
	assertion := assert.New(t)

	fn := func(*Context) {}
	handler := NewHandler(fn)

	handler(&gin.Context{})

	assertion.NotNil(handler)
	assertion.Equal(fmt.Sprintf("%T", handler), "gin.HandlerFunc")
}

func TestNewLoggerMiddleware(t *testing.T) {
	assertion := assert.New(t)

	l := NewMockLogger()

	handler := NewLoggerMiddleware(l, "reqIdKey")
	assertion.NotNil(handler)
	assertion.Equal(fmt.Sprintf("%T", handler), "gin.HandlerFunc")

	ctx := &gin.Context{
		Request: &http.Request{
			Header: http.Header{},
		},
	}
	handler(ctx)

	assertion.NotNil(ctx.Get(ContextLoggerKey))
}

func NewMockCtxWithLogger() *gin.Context {
	l := NewMockLogger()

	ctx := &gin.Context{}
	ctx.Set(ContextLoggerKey, logger.NewAppLogger(l, "reqId"))

	return ctx
}

func NewMockLogger() *logrus.Logger {
	l, _ := logger.NewLogger(&logger.Config{
		Output: "stdout",
		Level:  "info",
	})

	return l
}
