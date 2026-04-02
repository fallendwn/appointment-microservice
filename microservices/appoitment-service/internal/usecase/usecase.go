package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/fallendwn/appointment/appointment-service/internal/client"
	"github.com/fallendwn/appointment/appointment-service/internal/dto"
	"github.com/fallendwn/appointment/appointment-service/internal/model"
)

type AppointmentRepository interface {
	CreateAppointment(ctx context.Context, appoint model.Appointment) error
	GetAppointments(ctx context.Context) ([]model.Appointment, error)
	GetAppointment(ctx context.Context, id string) (model.Appointment, error)
	PatchAppointment(ctx context.Context, input dto.PatchDTO) (model.Appointment, error)
}
type AppointmentUseCase struct {
	repo AppointmentRepository
}

func NewAppointmentUseCase(repo AppointmentRepository) *AppointmentUseCase {
	return &AppointmentUseCase{
		repo: repo,
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
	ok, err := client.CheckDoctorExists(ctx, appointment.DoctorID.Hex())
	if err != nil {
		return fmt.Errorf("network error: cannot reach doctor-service: %v", err)
	}
	if !ok {
		return errors.New("doctor not found in doctor-service")
	}
	if appointment.Title == "" {
		return errors.New("title is empty")
	}
	if appointment.Description == "" {
		return errors.New("description is empty")
	}
	if !appointment.Status.IsValid() {
		return fmt.Errorf("invalid status: received '%s'", appointment.Status)
	}
	return u.repo.CreateAppointment(ctx, appointment)
}

func (u *AppointmentUseCase) PatchAppointment(ctx context.Context, input dto.PatchDTO) (model.Appointment, error) {

	appointment, err := u.repo.GetAppointment(ctx, input.Id)
	if err != nil {
		return appointment, errors.New("appointment d.n.e.")
	}

	var status model.Status = appointment.Status
	switch {
	case status == "done":
		return appointment, errors.New("cannot change appointment that is already done")
	case status == "in_progress" && input.Status != nil && *input.Status == model.StatusNew:
		return appointment, errors.New("cannot change appointment from progress to new")
	}
	return u.repo.PatchAppointment(ctx, input)
}
