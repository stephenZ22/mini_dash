package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/handler"
)

func RegisterCardRoutes(router *gin.Engine, handler *handler.CardHandler) *gin.Engine {
	// Card routes
	router.POST("/card", handler.CreateCard)
	router.GET("/card/:id", handler.GetCard)
	router.PUT("/card/:id", handler.UpdateCard)
	router.DELETE("/card/:id", handler.DeleteCard)
	return router
}
