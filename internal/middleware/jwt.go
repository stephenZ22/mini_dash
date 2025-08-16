package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stephenZ22/mini_dash/pkg/logger"
)

var secret = "minidash"

func GenerateJWTToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"user_name": username,
		"exp":       time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":       time.Now().Unix(),
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
		c.Set("user_name", claims["user_name"])

		// logger.MiniLogger().Infof("jwt token authorization successfully current_username is %s", claims["user_name"].(string))
		if username, ok := claims["user_name"].(string); ok {
			logger.MiniLogger().Infof("jwt token authorization successfully current_username is %s", username)
		} else {
			logger.MiniLogger().Warn("user_name claim is missing or not a string")
		}

		c.Next()
	}

}
