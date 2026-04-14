package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status string

const (
	StatusNew        Status = "new"
	StatusInProgress Status = "in_progress"
	StatusDone       Status = "done"
)

type Appointment struct {
	ID          primitive.ObjectID
	Title       string
	Description string
	DoctorID    primitive.ObjectID
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s Status) IsValid() bool {
	switch s {
	case StatusNew, StatusInProgress, StatusDone:
		return true
	}
	return false
}
