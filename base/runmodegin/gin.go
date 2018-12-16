package runmodegin

import (
	"github.com/binatify/gin-template/base/runmode"
	"github.com/gin-gonic/gin"
)

func ParseMode(mode runmode.RunMode) string {
	switch mode {
	case runmode.Production:
		return gin.ReleaseMode
	case runmode.Development:
		return gin.DebugMode
	case runmode.Test:
		return gin.TestMode
	default:
		return ""
	}
}
