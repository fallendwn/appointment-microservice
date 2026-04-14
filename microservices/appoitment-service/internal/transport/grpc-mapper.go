package transport

import (
	"github.com/fallendwn/appointment/appointment-service/internal/model"
	pb "github.com/fallendwn/appointment/appointment-service/proto"
)

func appointmentToProto(a *model.Appointment) *pb.AppointmentResponse {
	return &pb.AppointmentResponse{

		Id:          a.ID.Hex(),
		Title:       a.Title,
		Description: a.Description,
		DoctorId:    a.DoctorID.Hex(),
		Status:      statusToProto(a.Status),
		CreatedAt:   a.CreatedAt.String(),
		UpdatedAt:   a.UpdatedAt.String(),
	}
}

func statusToProto(s model.Status) pb.AppointmentStatus {

	switch s {
	case model.StatusInProgress:
		return pb.AppointmentStatus_IN_PROGRESS
	case model.StatusDone:
		return pb.AppointmentStatus_DONE

	default:
		return pb.AppointmentStatus_NEW
	}
}

func protoToStatus(s pb.AppointmentStatus) model.Status {

	switch s {
	case pb.AppointmentStatus_DONE:
		return model.StatusDone
	case pb.AppointmentStatus_IN_PROGRESS:
		return model.StatusInProgress

	default:
		return model.StatusNew
	}

}
