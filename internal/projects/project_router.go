package projects

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handler *ProjectHandler) *gin.Engine {

	// Project routes
	router.POST("/projects", handler.CreateProject)
	router.GET("/projects/:id", handler.GetProject)
	router.PUT("/projects/:id", handler.UpdateProject)
	router.DELETE("/projects/:id", handler.DeleteProject)
	router.GET("/projects", handler.ListProjects)
	return router
}
