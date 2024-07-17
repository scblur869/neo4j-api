package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, "service up")
}
