package transport

import (
	"github.com/fallendwn/appointment/doctor-service/internal/model"
	pb "github.com/fallendwn/appointment/doctor-service/proto"
)

func doctorToProto(doc *model.Doctor) *pb.DoctorResponse {

	return &pb.DoctorResponse{

		Id:             doc.ID.Hex(),
		FullName:       doc.FullName,
		Specialization: doc.Specialization,
		Email:          doc.Email,
	}

}
