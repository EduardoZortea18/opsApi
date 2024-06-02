package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DelteOpeningHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete route",
	})
}
