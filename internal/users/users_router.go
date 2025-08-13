package users

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine, handler *UserHandler) *gin.Engine {
	// User routes
	router.POST("/users", handler.CreateUser)
	router.GET("/users/:id", handler.GetUser)
	router.PUT("/users/:id", handler.UpdateUser)
	router.DELETE("/users/:id", handler.DeleteUser)
	return router
}
