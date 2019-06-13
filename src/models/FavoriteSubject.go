package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type FavoriteSubject struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	User		*User		`bson:"user" json:"user"`
	Subject		*Subject	`bson:"subject" json:"subject"`
}
type FavoriteSubjects []FavoriteSubject
