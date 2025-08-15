package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stephenZ22/mini_dash/internal/service"
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

		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  err,
			"data": nil,
		})
		return
	}

	_, err := lg.svc.LoginByPassword(login_request.Username, login_request.Password)
	if err != nil {
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
		"data": "token", // use jwt token
	})
}
