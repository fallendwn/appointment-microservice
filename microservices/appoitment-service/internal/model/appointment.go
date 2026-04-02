package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	StatusNew       Status = "new"
	StatusInProgres Status = "in_progress"
	StatusDone      Status = "done"
)

type Appointment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	DoctorID    primitive.ObjectID `bson:"doctor_id" json:"doctor_id"`
	Status      Status             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

func (s Status) IsValid() bool {
	switch s {
	case StatusNew, StatusInProgres, StatusDone:
		return true
	}
	return false
}
