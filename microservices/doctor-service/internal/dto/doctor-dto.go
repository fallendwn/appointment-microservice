package dto

type CreateDoctorDTO struct {
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	Specialization string `json:"specialization"`
}
