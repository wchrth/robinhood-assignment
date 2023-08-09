package main

import (
	"log"
	"os"

	"robinhood-assignment/api"
	"robinhood-assignment/application/service"
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

	ar := postgres.NewAppointmentRepository(db)
	as := service.NewAppointmentService(ar)
	ah := api.NewAppointmentHandler(as)

	cr := postgres.NewCommentRepository(db)
	cs := service.NewCommentService(cr)
	ch := api.NewCommentHandler(cs)

	ur := postgres.NewUserRepository(db)
	us := service.NewUserService(ur)
	uh := api.NewUserHandler(us)

	r := gin.New()
	r.GET("/appointments/:id", ah.GetByID)
	r.GET("/appointments", ah.GetAll)
	r.POST("/appointments", ah.Create)
	r.GET("/comments/:id", ch.GetByID)
	r.GET("/comments", ch.GetAll)
	r.POST("/comments", ch.Create)
	r.GET("/users/:id", uh.GetByID)
	r.GET("/users", uh.GetAll)
	r.POST("/users", uh.Create)
	r.Run()
}
