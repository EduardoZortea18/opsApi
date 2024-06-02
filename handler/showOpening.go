package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetByIdOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get by id route",
	})
}
