package app

import (
	"log"
	"net"

	"github.com/fallendwn/appointment/doctor-service/internal/repository"
	"github.com/fallendwn/appointment/doctor-service/internal/transport"
	"github.com/fallendwn/appointment/doctor-service/internal/usecase"
	pb "github.com/fallendwn/appointment/doctor-service/proto"
	"google.golang.org/grpc"
)

func Run(cfg *Config) {
	client, err := repository.InitMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatalf("fail to connect to mongo %v", err)
	}
	db := client.Database("doctor_db")

	doctorRepo := repository.NewDoctorRepo(db.Collection("doctors"))
	doctorService := usecase.NewDoctorUseCase(doctorRepo)
	doctorHandler := transport.NewDoctorHandler(doctorService)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterDoctorServiceServer(s, doctorHandler)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
