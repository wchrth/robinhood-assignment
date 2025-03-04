package route

import (
	"robinhood-assignment/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoute(router *gin.Engine, authHandler *handler.AuthHandler) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/auth/login", authHandler.Login)
		v1.GET("/auth/refresh", authHandler.RefreshToken)
	}

}
