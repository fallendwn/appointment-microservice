package transport

import (
	"context"
	"errors"
	"strings"

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
	//doctor exists?
	doctorID, err := primitive.ObjectIDFromHex(req.DoctorId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid doctor_id")
	}
	appointment := model.Appointment{
		Title:       req.Title,
		Description: req.Description,
		DoctorID:    doctorID,
	}

	if appoint, err := h.uc.CreateAppointment(ctx, appointment); err != nil {
		if err.Error() == "doctor not found in doctor-service" {
			return nil, status.Error(codes.NotFound, "invalid doctor_id")
		}
		if strings.Contains(err.Error(), "network error") {
			return nil, status.Error(codes.Unavailable, "doctor service is unavailable")
		}
		return nil, status.Error(codes.Internal, "internal error")
	} else {
		return appointmentToProto(&appoint), nil
	}
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

	res, err := h.uc.GetAppointmentInfo(ctx, req.Id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			return nil, status.Error(codes.NotFound, "not found")
		}
		return nil, status.Error(codes.Internal, "internal error")
	}
	return appointmentToProto(&res), nil

}

func (h *AppointmentHandler) UpdateAppointmentStatus(ctx context.Context, req *pb.UpdateStatusRequest) (*pb.AppointmentResponse, error) {
	var s model.Status = protoToStatus(req.Status)
	input := dto.PatchDTO{Id: req.Id, Status: &s}
	appointment, err := h.uc.PatchAppointment(ctx, input)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return appointmentToProto(&appointment), nil
}
