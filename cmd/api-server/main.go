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

	router := gin.Default()

	route.SetupUserRoute(router, userHandler)

	router.Run(fmt.Sprintf(":%s", cfg.App.Port))
}
