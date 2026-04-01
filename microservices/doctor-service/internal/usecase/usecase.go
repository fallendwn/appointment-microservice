package usecase

import (
	"context"
	"errors"

	"github.com/fallendwn/appointment/doctor-service/internal/model"
)

type DoctorRepository interface {
	CreateDoctor(ctx context.Context, doc model.Doctor) error
	GetDoctors(ctx context.Context) ([]model.Doctor, error)
	GetDoctor(ctx context.Context, id string) (model.Doctor, error)
}
type DoctorUseCase struct {
	repo DoctorRepository
}

func NewDoctorUseCase(repo DoctorRepository) *DoctorUseCase {
	return &DoctorUseCase{
		repo: repo,
	}

}

func (u *DoctorUseCase) GetDoctorsInfo(ctx context.Context) ([]model.Doctor, error) {
	var doctors []model.Doctor
	doctors, err := u.repo.GetDoctors(ctx)
	if err != nil {
		return doctors, err
	}
	return doctors, nil

}
func (u *DoctorUseCase) GetDoctorInfo(ctx context.Context, id string) (model.Doctor, error) {
	var doctor model.Doctor
	doctor, err := u.repo.GetDoctor(ctx, id)
	if err != nil {
		return doctor, err
	}
	return doctor, nil

}

func (u *DoctorUseCase) CreateDoctor(ctx context.Context, doc model.Doctor) error {
	if doc.Email == "" || doc.FullName == "" || doc.Specialization == "" {
		return errors.New("Some value is empty")
	}
	u.repo.CreateDoctor(ctx, doc)
	return nil
}
