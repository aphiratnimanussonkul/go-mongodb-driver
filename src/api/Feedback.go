package api

import (
	"github.com/aphiratnimanussonkul/go-mongodb-driver/config"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/models"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net/http"
	//"strings"
	"time"
)

func AddFeedback(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg models.Feedback
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	feedbackRepository := repository.NewFeedbackRepository(db, "Feedback")
	userRepository := repository.NewUserRepository(db, "User")
	currentTime := time.Now()

	user, err2 := userRepository.FindByEmail(msg.User.Email)
	if err2 != nil {
		fmt.Println(err2)
	}
	var p models.Feedback
	p.Text = msg.Text
	p.Timestamp = currentTime.Format("3:4:5")
	p.Date = currentTime.Format("2006-01-02")
	p.User = user

	feedbackRepository.Save(&p)
}
func GetFeedbackAll(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	feedbackRepository := repository.NewFeedbackRepository(db, "Feedback")
	feedback, err2 := feedbackRepository.FindAll()
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(feedback)

}
func DeleteFeedback(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	feedbackRepository := repository.NewFeedbackRepository(db, "Feedback")
	params := mux.Vars(req)
	var feedbackId = string(params["feedbackid"])
	fmt.Println(feedbackId)
	feedIdHex, err := primitive.ObjectIDFromHex(feedbackId)
	feedback , err := feedbackRepository.FindByID(feedIdHex)
	fmt.Println(feedback)
	err = feedbackRepository.Delete(feedback)
	if err != nil {
	}
}
func GetFeedbackById(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	feedbackRepository := repository.NewFeedbackRepository(db, "Feedback")

	params := mux.Vars(req)
	var feedbackid = string(params["feedbackid"])
	feedIdHex, err := primitive.ObjectIDFromHex(feedbackid)
	feedback, err2 := feedbackRepository.FindByID(feedIdHex)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(feedback)

}
