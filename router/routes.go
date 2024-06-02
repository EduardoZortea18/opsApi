package router

import (
	"github.com/EduardoZortea18/opsApi/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/openings/:id", handler.GetByIdOpeningHandler)
		v1.POST("/openings", handler.CreateOpeningHandler)
		v1.DELETE("/openings", handler.DelteOpeningHandler)
		v1.PUT("/openings", handler.EditOpeningHandler)
		v1.GET("/openings", handler.GetOpeningHandler)
	}
}
