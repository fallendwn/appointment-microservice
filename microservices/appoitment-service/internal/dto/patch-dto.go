package dto

import "github.com/fallendwn/appointment/appointment-service/internal/model"

type PatchDTO struct {
	Id     string        `json:"id"`
	Status *model.Status `json:"status"`
}
