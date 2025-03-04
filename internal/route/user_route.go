package route

import (
	"robinhood-assignment/internal/handler"
	"robinhood-assignment/internal/middleware"
	"robinhood-assignment/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupUserRoute(router *gin.Engine, userHandler *handler.UserHandler, jwtService service.JWTService) {

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users/register", userHandler.Register)

		v1.Use(middleware.AuthMiddleware(jwtService))
		v1.GET("/users", userHandler.GetAllUsers)
		v1.GET("/users/:id", userHandler.GetUserByID)

	}

}
