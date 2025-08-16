package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/handler"
	"github.com/stephenZ22/mini_dash/internal/middleware"
)

func RegisterUserRoutes(router *gin.Engine, handler *handler.UserHandler) {
	// User routes
	user_routers := router.Group("/users", middleware.JWTAuth())
	{
		user_routers.POST("", handler.CreateUser)
		user_routers.GET("/:id", handler.GetUser)
		user_routers.PUT("/:id", handler.UpdateUser)
		user_routers.DELETE("/:id", handler.DeleteUser)
	}
}
