package main

import (
	"fmt"
	"log"
	"robinhood-assignment/internal/config"
	"robinhood-assignment/internal/database"
	"robinhood-assignment/internal/handler"
	"robinhood-assignment/internal/repository"
	"robinhood-assignment/internal/route"
	"robinhood-assignment/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserServiceImpl(userRepository)
	userHandler := handler.NewUserHandler(userService)

	appointmentRepository := repository.NewAppointmentRepositoryDB(db)
	appointmentService := service.NewAppointmentServiceImpl(appointmentRepository, userRepository)
	appointmentHandler := handler.NewAppointmentHandler(appointmentService)

	commentRepository := repository.NewCommentRepositoryDB(db)
	commentService := service.NewCommentServiceImpl(commentRepository, userRepository)
	commentHandler := handler.NewCommentHandler(commentService)

	jwtService := service.NewJWTServiceImpl(&cfg.JWT)
	authService := service.NewAuthServiceImpl(userRepository, jwtService)
	authHandler := handler.NewAuthHandler(authService)

	router := gin.Default()

	route.SetupUserRoute(router, userHandler, jwtService)
	route.SetupAppointmentRoute(router, appointmentHandler, jwtService)
	route.SetupCommentRoute(router, commentHandler, jwtService)
	route.SetupAuthRoute(router, authHandler)

	router.Run(fmt.Sprintf(":%s", cfg.App.Port))
}
