package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Doctor struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FullName       string             `bson:"full_name" json:"name"`
	Specialization string             `bson:"specialization" json:"specialization"`
	Email          string             `bson:"email" json:"email"`
}
