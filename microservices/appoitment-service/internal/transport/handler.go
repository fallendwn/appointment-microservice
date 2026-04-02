package transport

import (
	"context"

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
	return &AppointmentHandler{
		uc: uc,
	}
}
func (h *AppointmentHandler) RegisterAppointment(c *gin.Context) {
	var appointment model.Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.uc.CreateAppointment(context.Background(), appointment)
	if err != nil {
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

	doctors, err := h.uc.GetAppointmentInfo(c.Request.Context(), id)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, doctors)

}

func (h *AppointmentHandler) PatchAppointmentStatus(c *gin.Context) {

	id := c.Param("id")
	var req struct {
		Status model.Status `json:"status"`
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	input := dto.PatchDTO{Id: id, Status: &req.Status}
	appointment, err := h.uc.PatchAppointment(context.Background(), input)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, appointment)

}
