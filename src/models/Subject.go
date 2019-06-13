package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Subject struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name		string 		`json:"name"`
	Code		string			`json:"code"`
	Major       *Major       `bson:"major" json:"major"`
	Picture		string		`json:"picture"`
}
type Subjects []Subject
type SubjectPointer []*Subject 
