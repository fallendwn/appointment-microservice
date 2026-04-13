package app

import (
	"log"
	"net"

	appclient "github.com/fallendwn/appointment/appointment-service/internal/client"
	pb "github.com/fallendwn/appointment/appointment-service/internal/proto"
	"github.com/fallendwn/appointment/appointment-service/internal/repository"
	"github.com/fallendwn/appointment/appointment-service/internal/transport"
	"github.com/fallendwn/appointment/appointment-service/internal/usecase"
	"google.golang.org/grpc"
)

func Run(cfg *Config) {
	mongoClient, err := repository.InitMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("Failed to connect to mongo %v", err)
	}
	db := mongoClient.Database("appointment_db")

	appointmentRepo := repository.NewAppointmentRepo(db.Collection("appointments"))
	doctorClient, err := appclient.NewDoctorClient(cfg.DoctorServiceAddr)
	if err != nil {
		log.Fatalf("Failed to connect to doctor-service: %v", err)
	}
	appointmentService := usecase.NewAppointmentUseCase(appointmentRepo, doctorClient)
	appointmentHandler := transport.NewAppointmentHandler(appointmentService)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAppointmentServiceServer(s, appointmentHandler)

	log.Printf("Appointment gRPC server starting on %s", cfg.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
