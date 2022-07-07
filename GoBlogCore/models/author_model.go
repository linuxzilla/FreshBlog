package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	Id        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username  string             `bson:"username" json:"username,omitempty" validate:"required"`
	Password  string             `bson:"password" json:"password,omitempty" validate:"required"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at,omitempty" validate:"required"`
}
