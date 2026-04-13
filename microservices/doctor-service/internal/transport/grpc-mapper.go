package transport

import (
	"github.com/fallendwn/appointment/doctor-service/internal/model"
	pb "github.com/fallendwn/appointment/doctor-service/internal/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func doctorToProto(doc *model.Doctor) *pb.DoctorResponse {

	return &pb.DoctorResponse{

		Id:             doc.ID.Hex(),
		FullName:       doc.FullName,
		Specialization: doc.Specialization,
		Email:          doc.Email,
	}

}

func protoToObjectID(hex string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(hex)
}
