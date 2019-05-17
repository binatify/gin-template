package ping

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PongHandler(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
