package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/handler"
)

func RegisterProjectRoutes(router *gin.Engine, handler *handler.ProjectHandler) *gin.Engine {

	// Project routes
	router.POST("/projects", handler.CreateProject)
	router.GET("/projects/:id", handler.GetProject)
	router.PUT("/projects/:id", handler.UpdateProject)
	router.DELETE("/projects/:id", handler.DeleteProject)
	router.GET("/projects", handler.ListProjects)
	return router
}
