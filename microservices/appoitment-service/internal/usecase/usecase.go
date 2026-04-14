package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/fallendwn/appointment/appointment-service/internal/model"
)

type AppointmentRepository interface {
	CreateAppointment(ctx context.Context, appoint model.Appointment) (model.Appointment, error)
	GetAppointments(ctx context.Context) ([]model.Appointment, error)
	GetAppointment(ctx context.Context, id string) (model.Appointment, error)
	PatchAppointment(ctx context.Context, id string, status model.Status) (model.Appointment, error)
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
	return u.repo.GetAppointment(ctx, id)
}

func (u *AppointmentUseCase) CreateAppointment(ctx context.Context, appointment model.Appointment) (model.Appointment, error) {
	if appointment.Title == "" {
		return model.Appointment{}, errors.New("title is required")
	}
	if appointment.DoctorID.IsZero() {
		return model.Appointment{}, errors.New("doctor_id is required")
	}

	ok, err := u.doctorClient.CheckDoctorExists(ctx, appointment.DoctorID.Hex())
	if err != nil {
		return model.Appointment{}, err
	}
	if !ok {
		return model.Appointment{}, model.ErrDoctorNotFound
	}

	appointment.Status = model.StatusNew
	now := time.Now()
	appointment.CreatedAt = now
	appointment.UpdatedAt = now

	return u.repo.CreateAppointment(ctx, appointment)
}

func (u *AppointmentUseCase) PatchAppointment(ctx context.Context, id string, newStatus model.Status) (model.Appointment, error) {
	if !newStatus.IsValid() {
		return model.Appointment{}, model.ErrInvalidStatus
	}

	appointment, err := u.repo.GetAppointment(ctx, id)
	if err != nil {
		return model.Appointment{}, err
	}

	if appointment.Status == model.StatusDone && newStatus == model.StatusNew {
		return model.Appointment{}, model.ErrInvalidStatusTransition
	}

	appointment.UpdatedAt = time.Now()
	return u.repo.PatchAppointment(ctx, id, newStatus)
}
