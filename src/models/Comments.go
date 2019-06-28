package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Text		string 		`json:"text"`
	Timestamp 	string   `json:"timestamp"`
	Date		string 		`json:"date"`
	User		*User		`bson:"user" json:"user"`

}
type Comments []Comment
type CommentPointer []*Comment