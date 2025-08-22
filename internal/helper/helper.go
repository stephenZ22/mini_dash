package helper

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) (uint, error) {
	userIDVal, ok := c.Get("user_id")
	if !ok {
		return 0, fmt.Errorf("user_id not found in context")
	}
	userIDStr, ok := userIDVal.(string)
	if !ok {
		return 0, fmt.Errorf("user_id is not a string")
	}
	id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func GetUserName(c *gin.Context) (string, error) {
	if val, ok := c.Get("username"); ok {
		if username, ok := val.(string); ok {
			return username, nil
		}
		return "", fmt.Errorf("username is not a string")
	}
	return "", fmt.Errorf("username not found in context")
}

func GetUserEmail(c *gin.Context) (string, error) {
	if val, ok := c.Get("user_email"); ok {
		if email, ok := val.(string); ok {
			return email, nil
		}
		return "", fmt.Errorf("user_email is not a string")
	}
	return "", fmt.Errorf("user_email not found in context")
}
