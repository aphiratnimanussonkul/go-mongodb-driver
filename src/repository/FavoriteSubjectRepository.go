package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"CPEProject/src/models"
)

type FavoriteSubjectRepositoryMongo struct {
	db *mongo.Database
	collection string
}

//NewProfileRepositoryMongo
func NewFavoriteSubjectRepository(db *mongo.Database, collection string) *FavoriteSubjectRepositoryMongo{
	return &FavoriteSubjectRepositoryMongo{
		db: db,
		collection: collection,
	}
}

//Save
func (r *FavoriteSubjectRepositoryMongo) Save(favoriteSubject *models.FavoriteSubject) error{
	res, err := r.db.Collection(r.collection).InsertOne(ctx, favoriteSubject)
	fmt.Println(res)
	return  err
}

//Update
func (r *FavoriteSubjectRepositoryMongo) Update(favoriteSubject *models.FavoriteSubject) error{
	res, err := r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"_id": favoriteSubject.ID}, favoriteSubject)
	fmt.Println(res)
	return err
}

//Delete
func (r *FavoriteSubjectRepositoryMongo) Delete(id primitive.ObjectID) error{
	res, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(res)
	return err
}

//FindByID
func (r *FavoriteSubjectRepositoryMongo) FindByID(id primitive.ObjectID) (*models.FavoriteSubject, error){
	var favoriteSubject models.FavoriteSubject
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&favoriteSubject)

	if err != nil {
		return nil, err
	}
	return &favoriteSubject, nil
}

//FindAll
func (r *FavoriteSubjectRepositoryMongo) FindAll() (models.FavoriteSubjects, error){
	var favoriteSubjects models.FavoriteSubjects

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var f models.FavoriteSubject
		if err := cursor.Decode(&f); err != nil {
		}
		favoriteSubjects = append(favoriteSubjects,f)
	}

	return favoriteSubjects, nil
}

//FindByEmail
func (r *FavoriteSubjectRepositoryMongo) FindByEmail(useremail string) (models.FavoriteSubjects, error){
	var favs models.FavoriteSubjects
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{"user.email": useremail})

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var f models.FavoriteSubject
		if err := cursor.Decode(&f); err != nil {
		}
		favs = append(favs, f)
	}
	return favs, nil
}

//FindBySubject
func (r *FavoriteSubjectRepositoryMongo) FindBySubject(subjectName string) (models.FavoriteSubjects, error){
	var name models.FavoriteSubjects
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{"user.email": subjectName})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var f models.FavoriteSubject
		if err := cursor.Decode(&f); err != nil {
		}
		name = append(name, f)
	}
	return name, nil
}
