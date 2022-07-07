package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	Id        primitive.ObjectID `bson:"id" json:"id,omitempty"`
	ArticleId primitive.ObjectID `bson:"article_id" json:"article_id,omitempty" validate:"required"`
	Username  string             `bson:"username" json:"username,omitempty" validate:"required"`
	Content   string             `bson:"content" json:"content,omitempty" validate:"required"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at,omitempty" validate:"required"`
}
