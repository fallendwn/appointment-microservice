package usecase

import (
	"context"
	"errors"

	"github.com/fallendwn/appointment/doctor-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DoctorRepository interface {
	CreateDoctor(ctx context.Context, doc model.Doctor) (model.Doctor, error)
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
	doc, err := u.repo.GetDoctors(ctx)
	if err != nil {
		return []model.Doctor{}, errors.New("internal server error")
	}
	return doc, nil
}

func (u *DoctorUseCase) GetDoctorInfo(ctx context.Context, id string) (model.Doctor, error) {
	doc, err := u.repo.GetDoctor(ctx, id)
	if err != nil {
		return model.Doctor{}, err
	}
	return doc, nil
}

func (u *DoctorUseCase) CreateDoctor(ctx context.Context, input model.Doctor) (model.Doctor, error) {
	if input.Email == "" || input.FullName == "" {
		return model.Doctor{}, errors.New("full_name and email are required")
	}

	doc := model.Doctor{
		ID:             primitive.NewObjectID(),
		FullName:       input.FullName,
		Email:          input.Email,
		Specialization: input.Specialization,
	}
	res, err := u.repo.CreateDoctor(ctx, doc)
	if err != nil {
		return model.Doctor{}, err
	}
	return res, nil
}
