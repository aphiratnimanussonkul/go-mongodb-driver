package api

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/config"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/models"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)
func AddComment(w http.ResponseWriter, req *http.Request)  {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg models.Comment
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	commentRepository := repository.NewCommentRepository(db, "Comment")
	postRepository := repository.NewPostRepository(db, "Post")
	userRepository := repository.NewUserRepository(db, "User")
	currentTime := time.Now()
	//
	params := mux.Vars(req)
	var postID string
	postID = string(params["postID"])
	user, err := userRepository.FindByEmail(msg.User.Email)
	if err != nil {
		fmt.Println(err)
	}
	var p models.Comment
	p.Text = msg.Text
	p.ID = primitive.NewObjectID()
	p.Timestamp = currentTime.Format("3:4:5")
	p.Date = currentTime.Format("2006-01-02")
	p.User = user
	fmt.Println("bson:",p.ID)
	commentRepository.Save(&p)
	postIDHex, err := primitive.ObjectIDFromHex(postID)
	comment, err := commentRepository.FindByID(p.ID)
	post, err := postRepository.FindByID(postIDHex)
	post.Comment = append(post.Comment,comment)
	postRepository.Update(post)


}

// func GetFacultyById(w http.ResponseWriter, req *http.Request) {

// 	//
// 	fmt.Println("Go Mongo Db")
// 	db, err := config.GetMongoDB()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Profile")
// 	//
// 	params := mux.Vars(req)
// 	fmt.Println(params["id"])


// 	profile, err := facultyRepository.FindByID(params["id"])

// 	fmt.Println("===== 1 =====")
// 	fmt.Println(profile)
// 	fmt.Println("===== 2 =====")
// 	fmt.Println(profile.ID)


// }

// func GetFacultyAll(w http.ResponseWriter, req *http.Request) {

// 	//
// 	fmt.Println("Go Mongo Db")
// 	db, err := config.GetMongoDB()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Faculty")
// 	post, err2 := facultyRepository.FindAll()
// 	if err2 != nil {
// 		fmt.Println(err2)
// 	}
// 	json.NewEncoder(w).Encode(post)
// }

// //Defualt add data
// func AddFacultyDefualt(facultyname string)  {
// 	db, err := config.GetMongoDB()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Faculty")
// 	var p models.Faculty
// 	p.Name = facultyname
// 	facultyRepository.Save(&p)
// }
