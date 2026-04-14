package transport

import (
	"context"
	"errors"

	"github.com/fallendwn/appointment/appointment-service/internal/dto"
	"github.com/fallendwn/appointment/appointment-service/internal/model"
	pb "github.com/fallendwn/appointment/appointment-service/internal/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AppointmentUseCase interface {
	GetAppointmentsInfo(ctx context.Context) ([]model.Appointment, error)
	GetAppointmentInfo(ctx context.Context, id string) (model.Appointment, error)
	PatchAppointment(ctx context.Context, input dto.PatchDTO) (model.Appointment, error)
	CreateAppointment(ctx context.Context, appointment model.Appointment) (model.Appointment, error)
}

type AppointmentHandler struct {
	pb.UnimplementedAppointmentServiceServer
	uc AppointmentUseCase
}

func NewAppointmentHandler(uc AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{uc: uc}
}

func (h *AppointmentHandler) CreateAppointment(ctx context.Context, req *pb.CreateAppointmentRequest) (*pb.AppointmentResponse, error) {
	if req.Title == "" {
		return nil, status.Error(codes.InvalidArgument, "title is required")
	}
	if req.DoctorId == "" {
		return nil, status.Error(codes.InvalidArgument, "doctor_id is required")
	}

	doctorID, err := primitive.ObjectIDFromHex(req.DoctorId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid doctor_id format")
	}

	appoint, err := h.uc.CreateAppointment(ctx, model.Appointment{
		Title:       req.Title,
		Description: req.Description,
		DoctorID:    doctorID,
	})
	if err != nil {
		if errors.Is(err, model.ErrDoctorUnavailable) {
			return nil, status.Error(codes.Unavailable, "doctor service is unavailable")
		}
		if errors.Is(err, model.ErrDoctorNotFound) {
			return nil, status.Error(codes.FailedPrecondition, "doctor does not exist")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return appointmentToProto(&appoint), nil
}

func (h *AppointmentHandler) ListAppointments(ctx context.Context, req *pb.ListAppointmentsRequest) (*pb.ListAppointmentsResponse, error) {
	appointments, err := h.uc.GetAppointmentsInfo(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	var result []*pb.AppointmentResponse
	for _, a := range appointments {
		result = append(result, appointmentToProto(&a))
	}
	return &pb.ListAppointmentsResponse{Appointments: result}, nil
}

func (h *AppointmentHandler) GetAppointment(ctx context.Context, req *pb.GetAppointmentRequest) (*pb.AppointmentResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	res, err := h.uc.GetAppointmentInfo(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "appointment not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return appointmentToProto(&res), nil
}

func (h *AppointmentHandler) UpdateAppointmentStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.AppointmentResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	s := protoToStatus(req.Status)
	input := dto.PatchDTO{Id: req.Id, Status: &s}
	appointment, err := h.uc.PatchAppointment(ctx, input)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "appointment not found")
		}
		if err.Error() == "cannot transition from done to new" {
			return nil, status.Error(codes.InvalidArgument, "cannot transition status from done to new")
		}
		if err.Error() == "invalid status value" {
			return nil, status.Error(codes.InvalidArgument, "invalid status value")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return appointmentToProto(&appointment), nil
}
