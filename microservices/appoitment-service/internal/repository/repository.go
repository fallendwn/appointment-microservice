package repository

import (
	"context"
	"time"

	"github.com/fallendwn/appointment/appointment-service/internal/dto"
	"github.com/fallendwn/appointment/appointment-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AppointmentRepo struct {
	col *mongo.Collection
}

func NewAppointmentRepo(col *mongo.Collection) *AppointmentRepo {
	return &AppointmentRepo{
		col: col,
	}
}

func (r *AppointmentRepo) CreateAppointment(ctx context.Context, appoint model.Appointment) error {

	_, err := r.col.InsertOne(ctx, appoint)
	return err

}

func (r *AppointmentRepo) GetAppointments(ctx context.Context) ([]model.Appointment, error) {
	var appointmentsList []model.Appointment
	filter := bson.M{}
	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var appoints model.Appointment
		if err := cursor.Decode(&appoints); err != nil {
			return nil, err
		}
		appointmentsList = append(appointmentsList, appoints)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return appointmentsList, nil
}

func (r *AppointmentRepo) GetAppointment(ctx context.Context, id string) (model.Appointment, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Appointment{}, err
	}
	filter := bson.M{"_id": objectId}
	var result model.Appointment
	err = r.col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return model.Appointment{}, err
	}

	return result, nil
}

func (r *AppointmentRepo) PatchAppointment(ctx context.Context, dto dto.PatchDTO) (model.Appointment, error) {
	objectId, err := primitive.ObjectIDFromHex(dto.Id)
	if err != nil {
		return model.Appointment{}, err
	}

	filterId := bson.M{"_id": objectId}
	filterStatus := bson.M{"$set": bson.M{"status": dto.Status, "updated_at": time.Now()}}

	var updatedDoc model.Appointment

	err = r.col.FindOneAndUpdate(ctx, filterId, filterStatus, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&updatedDoc)
	if err != nil {
		return model.Appointment{}, err
	}
	return updatedDoc, nil
}
