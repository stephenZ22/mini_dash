package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stephenZ22/mini_dash/internal/model"
	"github.com/stephenZ22/mini_dash/pkg/logger"
)

var secret = "minidash"

func GenerateJWTToken(user *model.User) (string, error) {
	logger.MiniLogger().Infof("generate jwt token by id: %d, name: %s. email: %s", user.ID, user.Username, user.Email)
	claims := jwt.MapClaims{
		"user_id":  fmt.Sprintf("%d", user.ID),
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.MiniLogger().Info("JWT auth")
		// 解析 Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":  0,
				"error": "missing or invalid token"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 解析并验证 token
		var claims jwt.MapClaims
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 0,
				"msg":  "invalid token",
			})
			c.Abort()
			return
		}

		// 将解析出来的 claims 存到 gin.Context
		c.Set("username", claims["username"])
		c.Set("user_id", claims["user_id"])
		c.Set("user_email", claims["email"])

		c.Next()
	}

}
