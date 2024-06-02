package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			ctx.JSON(http.StatusOK, gin.H{
				"message": id,
			})
		})

		v1.POST("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "second route",
			})
		})

		v1.DELETE("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "third route",
			})
		})

		v1.PUT("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "fourth route",
			})
		})

		v1.GET("/openings", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "fifth route",
			})
		})
	}
}
