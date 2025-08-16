package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/handler"
	"github.com/stephenZ22/mini_dash/internal/middleware"
)

func RegisterProjectRoutes(router *gin.Engine, handler *handler.ProjectHandler) {

	// Project routes
	projects_router := router.Group("/projects", middleware.JWTAuth())
	{
		projects_router.POST("", handler.CreateProject)
		projects_router.GET("/:id", handler.GetProject)
		projects_router.PUT("/:id", handler.UpdateProject)
		projects_router.DELETE("/:id", handler.DeleteProject)
		projects_router.GET("", handler.ListProjects)
	}
}
