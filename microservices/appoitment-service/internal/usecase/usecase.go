package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/fallendwn/appointment/appointment-service/internal/dto"
	"github.com/fallendwn/appointment/appointment-service/internal/model"
)

type AppointmentRepository interface {
	CreateAppointment(ctx context.Context, appoint model.Appointment) error
	GetAppointments(ctx context.Context) ([]model.Appointment, error)
	GetAppointment(ctx context.Context, id string) (model.Appointment, error)
	PatchAppointment(ctx context.Context, input dto.PatchDTO) (model.Appointment, error)
}
type DoctorClient interface {
	CheckDoctorExists(ctx context.Context, doctorID string) (bool, error)
}
type AppointmentUseCase struct {
	repo         AppointmentRepository
	doctorClient DoctorClient
}

func NewAppointmentUseCase(repo AppointmentRepository, dc DoctorClient) *AppointmentUseCase {
	return &AppointmentUseCase{
		repo:         repo,
		doctorClient: dc,
	}

}

func (u *AppointmentUseCase) GetAppointmentsInfo(ctx context.Context) ([]model.Appointment, error) {

	appointments, err := u.repo.GetAppointments(ctx)
	if err != nil {
		return nil, err
	}

	return appointments, nil

}

func (u *AppointmentUseCase) GetAppointmentInfo(ctx context.Context, id string) (model.Appointment, error) {

	var appointment model.Appointment
	appointment, err := u.repo.GetAppointment(ctx, id)
	if err != nil {
		return appointment, err
	}
	return appointment, nil

}

func (u *AppointmentUseCase) CreateAppointment(ctx context.Context, appointment model.Appointment) error {
	if appointment.DoctorID.IsZero() {
		return errors.New("doctor_id is required")
	}
	ok, err := u.doctorClient.CheckDoctorExists(ctx, appointment.DoctorID.Hex())
	if err != nil {
		return fmt.Errorf("network error: cannot reach doctor-service: %v", err)
	}
	if !ok {
		return errors.New("doctor not found in doctor-service")
	}
	if appointment.Title == "" {
		return errors.New("title is empty")
	}
	appointment.Status = model.StatusNew
	now := time.Now()
	appointment.CreatedAt = now
	appointment.UpdatedAt = now
	return u.repo.CreateAppointment(ctx, appointment)
}

func (u *AppointmentUseCase) PatchAppointment(ctx context.Context, input dto.PatchDTO) (model.Appointment, error) {
	if input.Status != nil && !input.Status.IsValid() {
		return model.Appointment{}, errors.New("invalid status value")
	}
	appointment, err := u.repo.GetAppointment(ctx, input.Id)
	if err != nil {
		return appointment, errors.New("appointment d.n.e.")
	}

	var status model.Status = appointment.Status
	if status == model.StatusDone && input.Status != nil && *input.Status == model.StatusNew {
		return appointment, errors.New("cannot transition from done to new")
	}
	appointment.UpdatedAt = time.Now()
	return u.repo.PatchAppointment(ctx, input)
}
