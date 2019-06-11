package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/models"

)

type PostRepositoryMongo struct {
	db *mongo.Database
	collection string
}

//NewProfileRepositoryMongo
func NewPostRepository(db *mongo.Database, collection string) *PostRepositoryMongo{
	return &PostRepositoryMongo{
		db: db,
		collection: collection,
	}
}

//Save
func (r *PostRepositoryMongo) Save(post *models.Post) error{
	res, err := r.db.Collection(r.collection).InsertOne(ctx, post)
	fmt.Println(res)
	return  err
}

//Update
func (r *PostRepositoryMongo) Update(post *models.Post) error{
	res, err := r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"_id": post.ID}, post)
	fmt.Println(res)
	return err
}

//Delete
func (r *PostRepositoryMongo) Delete(post *models.Post) error{
	res, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": post.ID})
	fmt.Println(res)
	return err
}

//FindByID
func (r *PostRepositoryMongo) FindByID(id primitive.ObjectID) (*models.Post, error){
	var post models.Post
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	if err != nil {
	}
	return &post, nil
}

//FindAll
func (r *PostRepositoryMongo) FindAll() (models.Posts, error){
	var post models.Posts

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var p models.Post
		if err := cursor.Decode(&p); err != nil {
		}
		post = append(post, p)
	}
	return post, nil
}


//FindByName
func (r *PostRepositoryMongo) FindByName(name string) (*models.Post, error){
	var post models.Post
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"text": name}).Decode(&post)

	if err != nil {
		return nil, err
	}

	return &post, nil
}
func (r *PostRepositoryMongo) FindByCode(code string) (models.Posts, error){
	var post models.Posts
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{"subject.code": code})

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var p models.Post
		if err := cursor.Decode(&post); err != nil {
		}
		post = append(post, p)
	}
	return post, nil
}
