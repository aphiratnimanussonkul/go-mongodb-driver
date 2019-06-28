package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/models"
)

type SubjectRepositoryMongo struct {
	db *mongo.Database
	collection string
}

//NewProfileRepositoryMongo
func NewSubjectRepository(db *mongo.Database, collection string) *SubjectRepositoryMongo{
	return &SubjectRepositoryMongo{
		db: db,
		collection: collection,
	}
}

//Save
func (r *SubjectRepositoryMongo) Save(subject *models.Subject) error{
	res, err := r.db.Collection(r.collection).InsertOne(ctx, subject)
	fmt.Println(res)
	return  err
}

//Update
func (r *SubjectRepositoryMongo) Update(subject *models.Subject) error{
	res, err := r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"_id": subject.ID}, bson.M{"$set": subject})
	fmt.Println(res)
	return err
}

//Delete
func (r *SubjectRepositoryMongo) Delete(id primitive.ObjectID) error{
	res, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(res)
	return err
}

//FindByID
func (r *SubjectRepositoryMongo) FindByID(id primitive.ObjectID) (*models.Subject, error){
	var subject models.Subject
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&subject)

	if err != nil {
		return nil, err
	}

	return &subject, nil
}

//FindAll
func (r *SubjectRepositoryMongo) FindAll() (models.Subjects, error){
	var subject models.Subjects

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var s models.Subject
		if err := cursor.Decode(&s); err != nil {
		}
		subject = append(subject, s)
	}
	return subject, nil
}


//FindByName
func (r *SubjectRepositoryMongo) FindByName(name string) (*models.Subject, error){
	var subject models.Subject
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"name": name}).Decode(&subject)

	if err != nil {
		return nil, err
	}

	return &subject, nil
}

func (r *SubjectRepositoryMongo) FindByCode(code string) (*models.Subject, error){
	var subject models.Subject
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"code": code}).Decode(&subject)

	if err != nil {
		return nil, err
	}

	return &subject, nil
}


func (r *SubjectRepositoryMongo) FindByMajor (majorName string) (models.Subjects, error){
	var subject models.Subjects
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{"major.name": majorName})
	if err != nil {

	}
	for cursor.Next(ctx) {
		var s models.Subject
		if err := cursor.Decode(&s); err != nil {
		}
		subject = append(subject, s)
	}
	return subject, nil
}

func (r *SubjectRepositoryMongo) FindByCodeEx(code string) (models.Subjects, error){
	var subject models.Subjects
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{"code": primitive.Regex{code, ""}})

	if err != nil {
	
	}
	for cursor.Next(ctx) {
		var s models.Subject
		if err := cursor.Decode(&s); err != nil {
		}
		subject = append(subject, s)
	}
	return subject, nil
}

func (r *SubjectRepositoryMongo) FindByNameEx(name string) (models.Subjects, error){
	var subject models.Subjects
	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{"name": primitive.Regex{name, ""}})

	if err != nil {

	}
	for cursor.Next(ctx) {
		var s models.Subject
		if err := cursor.Decode(&s); err != nil {
		}
		subject = append(subject, s)
	}
	return subject, nil
}
func (r *SubjectRepositoryMongo) DeleteByCode(code string) error{
	res, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"code": code})
	fmt.Println(res)
	return err
}
