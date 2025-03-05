package route

import (
	"robinhood-assignment/internal/handler"
	"robinhood-assignment/internal/middleware"
	"robinhood-assignment/internal/service"

	"github.com/gin-gonic/gin"
)

func SetupAppointmentRoute(router *gin.Engine, appointmentHandler *handler.AppointmentHandler, jwtService service.JWTService) {

	v1 := router.Group("/api/v1")
	{

		v1.Use(middleware.AuthMiddleware(jwtService))
		v1.GET("/appointments/statuses", appointmentHandler.GetStatuses)
		v1.GET("/appointments", appointmentHandler.GetAllAppointments)
		v1.GET("/appointments/:id", appointmentHandler.GetAppointmentByID)
		v1.POST("/appointments", appointmentHandler.CreateAppointment)
		v1.PUT("/appointments/:id", appointmentHandler.UpdateAppointment)
		v1.DELETE("/appointments/:id", appointmentHandler.DeleteAppointment)
		v1.POST("/appointments/:id/archive", appointmentHandler.ArchiveAppointment)
		v1.GET("/appointments/:id/histories", appointmentHandler.GetAppointmentHistories)

	}

}
