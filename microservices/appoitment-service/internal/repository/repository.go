package repository

import (
	"context"
	"time"

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
	return &AppointmentRepo{col: col}
}

func (r *AppointmentRepo) CreateAppointment(ctx context.Context, appoint model.Appointment) (model.Appointment, error) {
	result, err := r.col.InsertOne(ctx, toAppointmentDAO(appoint))
	if err != nil {
		return model.Appointment{}, err
	}
	appoint.ID = result.InsertedID.(primitive.ObjectID)
	return appoint, nil
}

func (r *AppointmentRepo) GetAppointments(ctx context.Context) ([]model.Appointment, error) {
	cursor, err := r.col.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []model.Appointment
	for cursor.Next(ctx) {
		var a appointmentDAO
		if err := cursor.Decode(&a); err != nil {
			return nil, err
		}
		result = append(result, fromAppointmentDAO(a))
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *AppointmentRepo) GetAppointment(ctx context.Context, id string) (model.Appointment, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Appointment{}, err
	}
	var a appointmentDAO
	err = r.col.FindOne(ctx, bson.M{"_id": objectId}).Decode(&a)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Appointment{}, model.ErrNotFound
		}
		return model.Appointment{}, err
	}
	return fromAppointmentDAO(a), nil
}

func (r *AppointmentRepo) PatchAppointment(ctx context.Context, id string, status model.Status) (model.Appointment, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Appointment{}, err
	}

	var a appointmentDAO
	err = r.col.FindOneAndUpdate(
		ctx,
		bson.M{"_id": objectId},
		bson.M{"$set": bson.M{"status": string(status), "updated_at": time.Now()}},
		options.FindOneAndUpdate().SetReturnDocument(options.After),
	).Decode(&a)
	if err != nil {
		return model.Appointment{}, err
	}
	return fromAppointmentDAO(a), nil
}
