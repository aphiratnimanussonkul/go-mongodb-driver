package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Major struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name		string 		`json:"name"`
	Faculty     *Faculty      `bson:"faculty" json:"faculty"`
}

type Majors []Major
type MajorPointer []*Major
