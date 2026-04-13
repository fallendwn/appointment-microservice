package usecase

import (
	"context"
	"errors"

	"github.com/fallendwn/appointment/doctor-service/internal/dto"
	"github.com/fallendwn/appointment/doctor-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	return &DoctorUseCase{repo: repo}
}

func (u *DoctorUseCase) GetDoctorsInfo(ctx context.Context) ([]model.Doctor, error) {
	return u.repo.GetDoctors(ctx)
}

func (u *DoctorUseCase) GetDoctorInfo(ctx context.Context, id string) (model.Doctor, error) {
	return u.repo.GetDoctor(ctx, id)
}

func (u *DoctorUseCase) CreateDoctor(ctx context.Context, input dto.CreateDoctorDTO) error {
	if input.Email == "" || input.FullName == "" {
		return errors.New("full_name and email are required")
	}

	doc := model.Doctor{
		ID:             primitive.NewObjectID(),
		FullName:       input.FullName,
		Email:          input.Email,
		Specialization: input.Specialization,
	}

	return u.repo.CreateDoctor(ctx, doc)
}
