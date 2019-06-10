package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Request struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SubjectCode		string 		`json:"subjectcode"`
	SubjectName		string 		`json:"subjectname"`
	Timestamp 	string  	`json:"timestamp"`
	Date		string 		`json:"date"`
	User		*User		`bson:"user" json:"user"`
}

type Requests []Request
