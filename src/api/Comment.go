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
	fmt.Println(post.ID)
	post.Comment = append(post.Comment,comment)
	postRepository.Update(post)


}

func GetCommentAll(w http.ResponseWriter, req *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	commentRepository := repository.NewCommentRepository(db, "Comment")
	comment, err2 := commentRepository.FindAll()
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(comment)
}

func DeleteCommentById(w http.ResponseWriter, req *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	commentRepository := repository.NewCommentRepository(db, "Comment")
	params := mux.Vars(req)
	var commentId = string(params["id"])
	commentIdHex, err := primitive.ObjectIDFromHex(commentId)
	commentRepository.Delete(commentIdHex)
}
