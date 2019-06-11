package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"

	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/models"
)

type UserRepositoryMongo struct {
	db *mongo.Database
	collection string
}

//NewProfileRepositoryMongo
func NewUserRepository(db *mongo.Database, collection string) *UserRepositoryMongo{
	return &UserRepositoryMongo{
		db: db,
		collection: collection,
	}
}

//Save
func (r *UserRepositoryMongo) Save(user *models.User) error{
	res, err := r.db.Collection(r.collection).InsertOne(ctx, user)
	fmt.Println(res)
	return  err
}
//func (r *UserRepositoryMongo) SaveSubject(subject *models.Subject, user *models.User) error{
//	err := r.db.Collection(r.collection).Insert(user)
//	err2 := r.db.Collection(r.collection).Update(bson.M{"subject": subject}, user)
//	fmt.Println(err2)
//	return err
//}
//Update
func (r *UserRepositoryMongo) Update(user *models.User) error{
	res, err := r.db.Collection(r.collection).UpdateOne(ctx, bson.M{"_id": user.ID}, user)
	fmt.Println(res)
	return err
}

//Delete
func (r *UserRepositoryMongo) Delete(id bson.ObjectId) error{
	res, err := r.db.Collection(r.collection).DeleteOne(ctx, bson.M{"_id": id})
	fmt.Println(res)
	return err
}

//FindByID
func (r *UserRepositoryMongo) FindByID(id bson.ObjectId) (*models.User, error){
	var user models.User
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

//FindAll
func (r *UserRepositoryMongo) FindAll() (models.Users, error){
	var user models.Users

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		// decode the document
		if err := cursor.Decode(&user); err != nil {
		}
	}
	return user, nil
}


//FindByName
func (r *UserRepositoryMongo) FindByName(name string) (*models.User, error){
	var user models.User
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"firstname": name}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryMongo) FindByEmail(Email string) (*models.User, error){
	var user models.User
	err := r.db.Collection(r.collection).FindOne(ctx, bson.M{"email": Email}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
