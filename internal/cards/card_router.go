package cards

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handler *CardHandler) *gin.Engine {
	// Card routes
	router.POST("/card", handler.CreateCard)
	router.GET("/card/:id", handler.GetCard)
	router.PUT("/card/:id", handler.UpdateCard)
	router.DELETE("/card/:id", handler.DeleteCard)
	return router
}
