package app

import (
	"log"

	"github.com/fallendwn/appointment/appointment-service/internal/repository"
	"github.com/fallendwn/appointment/appointment-service/internal/transport"
	"github.com/fallendwn/appointment/appointment-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Run(cfg *Config) {
	client, err := repository.InitMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to mongo %v", err)
	}
	db := client.Database("appointment_db")

	appointmentRepo := repository.NewAppointmentRepo(db.Collection("appointments"))
	appointmentService := usecase.NewAppointmentUseCase(appointmentRepo)
	appointmentHandler := transport.NewAppointmentHandler(appointmentService)
	r := gin.Default()
	transport.RegisterRouter(r, appointmentHandler)
	r.Run(cfg.Port)
}
