package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)
type Post struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Text		string 		`json:"text"`
	Timestamp 	string  	`json:"timestamp"`
	Date		string 		`json:"date"`
	User		*User		`bson:"user" json:"user"`
	VdoLink		[]string   	`json:"vdoLink"`
	File	    []string 	`json:"file"`
	FileName	[]string	`json:"filename"`
	Picture	    []string 	`json:"picture"`
	Subject		*Subject	`bson:"subject" json:"subject"`
	Comment		[]*Comment	`bson:"comment" json:"comment"`
}

type Posts []Post
