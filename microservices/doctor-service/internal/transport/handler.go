package transport

import (
	"context"
	"errors"

	"github.com/fallendwn/appointment/doctor-service/internal/dto"
	"github.com/fallendwn/appointment/doctor-service/internal/model"
	"github.com/gin-gonic/gin"
)

type DoctorUseCase interface {
	GetDoctorsInfo(ctx context.Context) ([]model.Doctor, error)
	GetDoctorInfo(ctx context.Context, id string) (model.Doctor, error)
	CreateDoctor(ctx context.Context, doc dto.CreateDoctorDTO) error
}

type DoctorHandler struct {
	uc DoctorUseCase
}

func NewDoctorHandler(uc DoctorUseCase) *DoctorHandler {
	return &DoctorHandler{uc: uc}
}

func (h *DoctorHandler) Register(c *gin.Context) {
	var doc dto.CreateDoctorDTO
	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.CreateDoctor(c.Request.Context(), doc); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, doc)
}

func (h *DoctorHandler) RetrieveDoctors(c *gin.Context) {
	doctors, err := h.uc.GetDoctorsInfo(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, doctors)
}

func (h *DoctorHandler) RetrieveDoctorByID(c *gin.Context) {
	id := c.Param("id")

	doctor, err := h.uc.GetDoctorInfo(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(404, gin.H{"error": "doctor not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, doctor)
}
