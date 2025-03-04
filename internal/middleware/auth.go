package middleware

import (
	"robinhood-assignment/internal/api"
	"robinhood-assignment/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.HasPrefix(authHeader, "Bearer ") {
			api.RespondError(c, api.ErrUnauthorized)
			c.Abort()
			return
		}

		tokenStr := authHeader[len("Bearer "):]

		claims, err := jwtService.ValidateToken(tokenStr, "access")
		if err != nil {
			api.RespondError(c, err)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
}
