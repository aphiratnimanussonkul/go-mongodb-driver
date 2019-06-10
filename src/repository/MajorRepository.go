package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"CPEProject/src/models"
)

type MajorRepositoryMongo struct {
	db *mongo.Database
	collection string
}

//NewProfileRepositoryMongo
func NewMajorRepository(db *mongo.Database, collection string) *MajorRepositoryMongo{
	return &MajorRepositoryMongo{
		db: db,
		collection: collection,
	}
}

//Save
func (r *MajorRepositoryMongo) Save(major *models.Major) error{
	res, err := r.db.Collection(r.collection).InsertOne(ctx, major)
	fmt.Println(res)
	return  err
}

//Update
func (r *MajorRepositoryMongo) Update(major *models.Major) error{
	res, err := r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"_id": major.ID}, major)
	fmt.Println(res)
	return err
}

//Delete
func (r *MajorRepositoryMongo) Delete(id primitive.ObjectID) error{
	res, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(res)
	return err
}

//FindByID
func (r *MajorRepositoryMongo) FindByID(id primitive.ObjectID) (*models.Major, error){
	var major models.Major
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&major)

	if err != nil {
		return nil, err
	}

	return &major, nil
}

//FindAll
func (r *MajorRepositoryMongo) FindAll() (models.Majors, error){
	var major models.Majors

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var m models.Major
		if err := cursor.Decode(&m); err != nil {
		}
		major = append(major, m)
	}
	return major, nil
}


//FindByName
func (r *MajorRepositoryMongo) FindByName(name string) (*models.Major, error){
	var major models.Major
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"name": name}).Decode(&major)

	if err != nil {
		return nil, err
	}

	return &major, nil
}

func (r *MajorRepositoryMongo) FindByFaculty(facultyName string) (models.Majors, error){
	var major models.Majors
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{"faculty.name": facultyName})
	if err != nil {

	}
	for cursor.Next(ctx) {
		var m models.Major
		if err := cursor.Decode(&m); err != nil {
		}
		major = append(major, m)
	}
	return major, nil
}

func (r *MajorRepositoryMongo) DeleteByName(name string) error{
	res, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"name": name})
	fmt.Println(res)
	return err
}
