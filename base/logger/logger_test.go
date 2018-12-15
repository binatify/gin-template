package logger

import (
	"github.com/golib/assert"
	"testing"
)

func TestNewLogger(t *testing.T) {
	assertion := assert.New(t)

	cfg := &Config{
		Level:  "info",
		Output: "stdout",
	}

	logger, err := NewLogger(cfg)
	assertion.Nil(err)
	assertion.NotNil(logger)

	cfg.Level = "xxxx"

	logger, err = NewLogger(cfg)
	assertion.NotNil(err)
	assertion.Nil(logger)
	cfg.Level = "info"

	cfg.Output = "./bbb/a.log"

	logger, err = NewLogger(cfg)
	assertion.NotNil(err)
	assertion.Nil(logger)
}

func TestNewAppLogger(t *testing.T) {
	assertion := assert.New(t)

	cfg := &Config{
		Level:  "info",
		Output: "stdout",
	}

	logger, _ := NewLogger(cfg)

	applogger := NewAppLogger(logger, "reqId")

	assertion.NotNil(applogger)
	assertion.NotNil(applogger.Data)
	assertion.Equal(applogger.Data["ReqID"], "reqId")
}
