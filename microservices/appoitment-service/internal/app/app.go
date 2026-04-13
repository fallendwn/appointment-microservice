package app

import (
	"log"
	"os"

	appclient "github.com/fallendwn/appointment/appointment-service/internal/client"
	"github.com/fallendwn/appointment/appointment-service/internal/repository"
	"github.com/fallendwn/appointment/appointment-service/internal/transport"
	"github.com/fallendwn/appointment/appointment-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Run(cfg *Config) {
	mongoClient, err := repository.InitMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to mongo %v", err)
	}
	db := mongoClient.Database("appointment_db")

	appointmentRepo := repository.NewAppointmentRepo(db.Collection("appointments"))
	doctorClient := appclient.NewDoctorClient(os.Getenv("DOCTOR_SERVICE_URL"))
	appointmentService := usecase.NewAppointmentUseCase(appointmentRepo, doctorClient)
	appointmentHandler := transport.NewAppointmentHandler(appointmentService)
	r := gin.Default()
	transport.RegisterRouter(r, appointmentHandler)
	r.Run(cfg.Port)
}
