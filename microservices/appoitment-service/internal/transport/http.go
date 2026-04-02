package transport

import "github.com/gin-gonic/gin"

func RegisterRouter(r *gin.Engine, h *AppointmentHandler) {

	api := r.Group("/api/v1")
	{
		api.POST("/appointments", h.RegisterAppointment)
		api.GET("/appointments/", h.RetrieveAppointments)
		api.GET("/appointments/:id", h.RetrieveAppointmentByID)
		api.PATCH("/appointments/:id/status", h.PatchAppointmentStatus)
	}

}
