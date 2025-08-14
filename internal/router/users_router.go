package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/handler"
)

func RegisterUserRoutes(router *gin.Engine, handler *handler.UserHandler) *gin.Engine {
	// User routes
	router.POST("/users", handler.CreateUser)
	router.GET("/users/:id", handler.GetUser)
	router.PUT("/users/:id", handler.UpdateUser)
	router.DELETE("/users/:id", handler.DeleteUser)
	return router
}
