package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Article struct {
	Id             primitive.ObjectID   `bson:"_id" json:"id,omitempty"`
	Authors        []primitive.ObjectID `bson:"authors" json:"authors,omitempty" validate:"required"`
	Title          string               `bson:"title " json:"title,omitempty" validate:"required"`
	Tags           []string             `bson:"tags" json:"tags,omitempty"`
	TitlePictureId primitive.ObjectID   `bson:"title_picture_id" json:"title_picture_id,omitempty"`
	PreviewText    string               `bson:"preview_text" json:"preview_text,omitempty"`
	CreatedAt      primitive.DateTime   `bson:"created_at" json:"created_at,omitempty" validate:"required"`
	ModifiedAt     primitive.DateTime   `bson:"modified_at" json:"modified_at,omitempty" validate:"required"`
	Views          int64                `bson:"views" json:"views,omitempty"`
	IsVisible      bool                 `bson:"is_visible" json:"is_visible,omitempty" validate:"required"`
}
