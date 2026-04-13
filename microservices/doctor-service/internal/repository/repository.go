package repository

import (
	"context"
	"fmt"

	"github.com/fallendwn/appointment/doctor-service/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DoctorRepo struct {
	col *mongo.Collection
}

func NewDoctorRepo(col *mongo.Collection) *DoctorRepo {
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{Key: "email", Value: 1}},
		Options: options.Index().SetUnique(true),
	}
	_, err := col.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		fmt.Print(err)
	}
	return &DoctorRepo{col: col}
}

func (r *DoctorRepo) CreateDoctor(ctx context.Context, doc model.Doctor) error {
	_, err := r.col.InsertOne(ctx, toDoctorDAO(doc))
	return err
}

func (r *DoctorRepo) GetDoctors(ctx context.Context) ([]model.Doctor, error) {
	filter := bson.M{}
	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var result []model.Doctor
	for cursor.Next(ctx) {
		var d doctorDAO
		if err := cursor.Decode(&d); err != nil {
			return nil, err
		}
		result = append(result, fromDoctorDAO(d))
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *DoctorRepo) GetDoctor(ctx context.Context, id string) (model.Doctor, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Doctor{}, err
	}
	filter := bson.M{"_id": objectId}
	var d doctorDAO
	err = r.col.FindOne(ctx, filter).Decode(&d)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Doctor{}, model.ErrNotFound
		}
		return model.Doctor{}, err
	}
	return fromDoctorDAO(d), nil
}
