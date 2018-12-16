package runmodegin

import (
	"github.com/binatify/gin-template/base/runmode"
	"github.com/gin-gonic/gin"
	"github.com/golib/assert"
	"testing"
)

func TestParseMode(t *testing.T) {
	assertion := assert.New(t)

	assertion.Equal(ParseMode(runmode.Production), gin.ReleaseMode)
	assertion.Equal(ParseMode(runmode.Development), gin.DebugMode)
	assertion.Equal(ParseMode(runmode.Test), gin.TestMode)
	assertion.Equal(ParseMode(runmode.RunMode("invalid mode")), "")
}
