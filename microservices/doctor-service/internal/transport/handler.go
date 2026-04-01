package transport

import (
	"context"

	"github.com/fallendwn/appointment/doctor-service/internal/model"
	"github.com/gin-gonic/gin"
)

type DoctorUseCase interface {
	GetDoctorsInfo(ctx context.Context) ([]model.Doctor, error)
	GetDoctorInfo(ctx context.Context, id string) (model.Doctor, error)
	CreateDoctor(ctx context.Context, doc model.Doctor) error
}
type DoctorHandler struct {
	uc DoctorUseCase
}

func NewDoctorHandler(uc DoctorUseCase) *DoctorHandler {
	return &DoctorHandler{
		uc: uc,
	}
}
func (h *DoctorHandler) Register(c *gin.Context) {
	var doc model.Doctor

	if err := c.ShouldBindJSON(&doc); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.uc.CreateDoctor(c.Request.Context(), doc)
	if err != nil {
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

	doctors, err := h.uc.GetDoctorInfo(c.Request.Context(), id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, doctors)

}
