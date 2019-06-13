package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Picture	string `json:"picture"`
	StudentId	string `json:"studentId"`
	Name	string `json:"name"`
	Major 	string `json:"major"`
	Email		string `json:"email"`
	Subject 	[]*Subject `json:"subject" bson:"subject"`
}
type Users []User
