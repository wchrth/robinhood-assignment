package main

import (
	"log"
	"os"

	"robinhood-assignment/domain/entity"
	"robinhood-assignment/infrastructure/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	postgresConfig := postgres.Config{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDBName:   os.Getenv("POSTGRES_DB_NAME"),
	}

	db, err := postgres.NewDB(postgresConfig)
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(&entity.Appointment{}, &entity.Comment{}, &entity.User{})

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
