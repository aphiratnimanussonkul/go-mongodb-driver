package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	// "time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/models"
)

type FeedbackRepository struct {
	db *mongo.Database
	collection string
}

//NewProfileRepositoryMongo
func NewFeedbackRepository(db *mongo.Database, collection string) *FeedbackRepository{
	return &FeedbackRepository{
		db: db,
		collection: collection,
	}
}

//Save
func (r *FeedbackRepository) Save(feedback *models.Feedback) error{
	res, err := r.db.Collection(r.collection).InsertOne(ctx, feedback)
	fmt.Println(res)
	return  err
}

//Update
func (r *FeedbackRepository) Update( feedback *models.Feedback) error{
	res, err := r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"_id": feedback.ID}, feedback)
	fmt.Println(res)
	return err
}

//Delete
func (r *FeedbackRepository) Delete(feedback *models.Feedback) error{
	res, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": feedback.ID})
	fmt.Println(res)
	return err
}

//FindByID
func (r *FeedbackRepository) FindByID(id primitive.ObjectID) (*models.Feedback, error){
	var feedback models.Feedback
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&feedback)
	if err != nil {
		return nil, err
	}
	return &feedback, nil
}

//FindAll
func (r *FeedbackRepository) FindAll() (models.Feedbacks, error){
	var feedback models.Feedbacks

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var f models.Feedback
		if err := cursor.Decode(&f); err != nil {
		}
		feedback = append(feedback, f)
	}

	return feedback, nil
}


//FindByName
func (r *FeedbackRepository) FindByName(name string) (*models.Feedback, error){
	var feedback models.Feedback
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"text": name}).Decode(&feedback)

	if err != nil {
		return nil, err
	}

	return &feedback, nil
}
func (r *FeedbackRepository) FindByCode(code string) (models.Feedbacks, error){
	var feedback models.Feedbacks
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{"subject.code": code})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var f models.Feedback
		if err := cursor.Decode(&f); err != nil {
		}
		feedback = append(feedback, f)
	}
	return feedback, nil
}
