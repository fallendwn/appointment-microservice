package transport

import (
	"context"

	"github.com/fallendwn/appointment/doctor-service/internal/model"
	pb "github.com/fallendwn/appointment/doctor-service/internal/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DoctorUseCase interface {
	GetDoctorsInfo(ctx context.Context) ([]model.Doctor, error)
	GetDoctorInfo(ctx context.Context, id string) (model.Doctor, error)
	CreateDoctor(ctx context.Context, doc model.Doctor) (model.Doctor, error)
}

type DoctorHandler struct {
	uc DoctorUseCase
	pb.UnimplementedDoctorServiceServer
}

func NewDoctorHandler(uc DoctorUseCase) *DoctorHandler {
	return &DoctorHandler{uc: uc}
}

func (h *DoctorHandler) CreateDoctor(ctx context.Context, req *pb.CreateDoctorRequest) (*pb.DoctorResponse, error) {
	doc := model.Doctor{
		FullName:       req.FullName,
		Specialization: req.Specialization,
		Email:          req.Email,
	}

	if doctor, err := h.uc.CreateDoctor(ctx, doc); err != nil {
		if err.Error() == "full_name and email are required" {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}
		return nil, status.Error(codes.Internal, "internal error")
	} else {
		return doctorToProto(&doctor), nil
	}

}

func (h *DoctorHandler) GetDoctor(ctx context.Context, req *pb.GetDoctorRequest) (*pb.DoctorResponse, error) {

	doctor, err := h.uc.GetDoctorInfo(ctx, req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "doctor not found")
	}
	return doctorToProto(&doctor), nil

}

func (h *DoctorHandler) ListDoctors(ctx context.Context, req *pb.ListDoctorsRequest) (*pb.ListDoctorsResponse, error) {
	doc, err := h.uc.GetDoctorsInfo(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal server error")
	}
	var res []*pb.DoctorResponse
	for _, d := range doc {
		res = append(res, doctorToProto(&d))
	}
	return &pb.ListDoctorsResponse{Doctors: res}, nil
}
