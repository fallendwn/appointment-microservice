package repository

import (
	"github.com/fallendwn/appointment/doctor-service/internal/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type doctorDAO struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	FullName       string             `bson:"full_name"`
	Specialization string             `bson:"specialization"`
	Email          string             `bson:"email"`
}

func toDoctorDAO(d model.Doctor) doctorDAO {
	return doctorDAO{
		ID:             d.ID,
		FullName:       d.FullName,
		Specialization: d.Specialization,
		Email:          d.Email,
	}
}

func fromDoctorDAO(d doctorDAO) model.Doctor {
	return model.Doctor{
		ID:             d.ID,
		FullName:       d.FullName,
		Specialization: d.Specialization,
		Email:          d.Email,
	}
}
