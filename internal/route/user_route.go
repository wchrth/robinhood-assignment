package route

import (
	"robinhood-assignment/internal/handler"

	"github.com/gin-gonic/gin"
)

func SetupUserRoute(router *gin.Engine, userHandler *handler.UserHandler) {

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", userHandler.GetAllUsers)
		v1.GET("/users/:id", userHandler.GetUserByID)
		// v1.GET("/users/:email", userHandler.GetUserByEmail)
		v1.POST("/users/register", userHandler.Register)
	}

}
