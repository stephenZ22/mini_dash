package projects

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handler *ProjectHandler) *gin.Engine {

	// Project routes
	router.POST("/projects", handler.CreateProject)
	return router
}
