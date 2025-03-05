package route

import (
	"robinhood-assignment/internal/handler"
	"robinhood-assignment/internal/middleware"
	"robinhood-assignment/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupCommentRoute(router *gin.Engine, commentHandler *handler.CommentHandler, jwtService service.JWTService) {

	v1 := router.Group("/api/v1")
	{

		v1.Use(middleware.AuthMiddleware(jwtService))
		v1.POST("/appointments/:appointment_id/comments", commentHandler.CreateComment)
		v1.GET("/appointments/:appointment_id/comments", commentHandler.GetCommentsByAppointment)
		v1.PUT("/comments/:id", commentHandler.UpdateComment)
		v1.DELETE("/comments/:id", commentHandler.DeleteComment)

	}

}
