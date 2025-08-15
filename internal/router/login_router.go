package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/handler"
)

func RegisterLoginRoutes(engin *gin.Engine, h *handler.LoginHandler) *gin.Engine {
	engin.POST("/login", h.LoginByPassword)
	return engin
}
