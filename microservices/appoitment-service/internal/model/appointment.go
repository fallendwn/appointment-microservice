package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
)

type Status string

func (s Status) Error(internal codes.Code, param2 string) error {
	panic("unimplemented")
}

const (
	StatusNew       Status = "new"
	StatusInProgres Status = "in_progress"
	StatusDone      Status = "done"
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
	case StatusNew, StatusInProgres, StatusDone:
		return true
	}
	return false
}
