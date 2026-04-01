package app

import (
	"log"

	"github.com/fallendwn/appointment/doctor-service/internal/repository"
	"github.com/fallendwn/appointment/doctor-service/internal/transport"
	"github.com/fallendwn/appointment/doctor-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Run(cfg *Config) {
	client, err := repository.InitMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to mongo %v", err)
	}
	db := client.Database("doctor_db")

	doctorRepo := repository.NewDoctorRepo(db.Collection("doctors"))
	doctorService := usecase.NewDoctorUseCase(doctorRepo)
	doctorHandler := transport.NewDoctorHandler(doctorService)
	r := gin.Default()
	transport.RegisterRouter(r, doctorHandler)
	r.Run(cfg.Port)
}
