package repository

import (
	"time"

	"github.com/fallendwn/appointment/appointment-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type appointmentDAO struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	DoctorID    primitive.ObjectID `bson:"doctor_id"`
	Status      string             `bson:"status"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
}

func toAppointmentDAO(a model.Appointment) appointmentDAO {
	return appointmentDAO{
		ID:          a.ID,
		Title:       a.Title,
		Description: a.Description,
		DoctorID:    a.DoctorID,
		Status:      string(a.Status),
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}

func fromAppointmentDAO(a appointmentDAO) model.Appointment {
	return model.Appointment{
		ID:          a.ID,
		Title:       a.Title,
		Description: a.Description,
		DoctorID:    a.DoctorID,
		Status:      model.Status(a.Status),
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
	}
}
