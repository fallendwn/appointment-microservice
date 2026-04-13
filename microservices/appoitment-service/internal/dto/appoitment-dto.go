package dto

import (
	"github.com/fallendwn/appointment/appointment-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateAppointmentDTO struct {
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DoctorID    primitive.ObjectID `json:"doctor_id"`
}

type PatchDTO struct {
	Id     string        `json:"id"`
	Status *model.Status `json:"status"`
}
