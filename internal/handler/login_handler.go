package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/middleware"
	"github.com/stephenZ22/mini_dash/internal/service"
	"github.com/stephenZ22/mini_dash/pkg/logger"
)

type LoginHandler struct {
	svc service.LoginService
}

func NewLoginHandler(svc service.LoginService) *LoginHandler {
	return &LoginHandler{
		svc: svc,
	}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (lg *LoginHandler) LoginByPassword(c *gin.Context) {
	var login_request LoginRequest

	if err := c.ShouldBindJSON(&login_request); err != nil {
		logger.MiniLogger().Error("invalid request data")
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	_, err := lg.svc.LoginByPassword(login_request.Username, login_request.Password)
	if err != nil {
		logger.MiniLogger().Error("login by password error\n")
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err,
			"data": nil,
		})

		return
	}

	token, err := middleware.GenerateJWTToken(login_request.Username)
	if err != nil {
		logger.MiniLogger().Errorf("generate jwt token error: %s", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err,
			"data": nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  "login successfully",
		"data": token, // use jwt token
	})
}
