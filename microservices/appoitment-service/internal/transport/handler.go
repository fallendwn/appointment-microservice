package transport

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/fallendwn/appointment/appointment-service/internal/dto"
	"github.com/fallendwn/appointment/appointment-service/internal/model"
	"github.com/gin-gonic/gin"
)

type AppointmentUseCase interface {
	GetAppointmentsInfo(ctx context.Context) ([]model.Appointment, error)
	GetAppointmentInfo(ctx context.Context, id string) (model.Appointment, error)
	PatchAppointment(ctx context.Context, input dto.PatchDTO) (model.Appointment, error)
	CreateAppointment(ctx context.Context, appointment model.Appointment) error
}

type AppointmentHandler struct {
	uc AppointmentUseCase
}

func NewAppointmentHandler(uc AppointmentUseCase) *AppointmentHandler {
	return &AppointmentHandler{uc: uc}
}

func (h *AppointmentHandler) RegisterAppointment(c *gin.Context) {
	var input dto.CreateAppointmentDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	appointment := model.Appointment{
		Title:       input.Title,
		Description: input.Description,
		DoctorID:    input.DoctorID,
	}

	if err := h.uc.CreateAppointment(c.Request.Context(), appointment); err != nil {
		if err.Error() == "doctor not found in doctor-service" {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}
		if strings.Contains(err.Error(), "network error") {
			log.Printf("[ERROR] doctor-service unavailable: %v", err)
			c.JSON(503, gin.H{"error": "doctor-service unavailable"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, appointment)
}

func (h *AppointmentHandler) RetrieveAppointments(c *gin.Context) {
	appointments, err := h.uc.GetAppointmentsInfo(c.Request.Context())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, appointments)
}

func (h *AppointmentHandler) RetrieveAppointmentByID(c *gin.Context) {
	id := c.Param("id")

	appointment, err := h.uc.GetAppointmentInfo(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, model.ErrNotFound) {
			c.JSON(404, gin.H{"error": "appointment not found"})
			return
		}
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, appointment)
}

func (h *AppointmentHandler) PatchAppointmentStatus(c *gin.Context) {
	id := c.Param("id")

	var req struct {
		Status model.Status `json:"status"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	input := dto.PatchDTO{Id: id, Status: &req.Status}
	appointment, err := h.uc.PatchAppointment(c.Request.Context(), input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, appointment)
}
