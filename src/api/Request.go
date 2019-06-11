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

func AddRequest(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg models.Request
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	requestRepository := repository.NewRequestRepository(db, "Request")
	userRepository := repository.NewUserRepository(db, "User")
	currentTime := time.Now()

	user, err2 := userRepository.FindByEmail(msg.User.Email)
	if err2 != nil {
		fmt.Println(err2)
	}
	var p models.Request
	p.SubjectCode = msg.SubjectCode
	p.SubjectName = msg.SubjectName
	p.Timestamp = currentTime.Format("3:4:5")
	p.Date = currentTime.Format("2006-01-02")
	p.User = user

	requestRepository.Save(&p)
}
func GetRequestAll(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	requestRepository := repository.NewRequestRepository(db, "Request")
	feedback, err2 := requestRepository.FindAll()
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(feedback)

}
func DeleteRequest(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	requestRepository := repository.NewRequestRepository(db, "Request")
	params := mux.Vars(req)
	var requestId = string(params["requestid"])
	requestIdHex, err := primitive.ObjectIDFromHex(requestId)
	request , err := requestRepository.FindByID(requestIdHex)
	err = requestRepository.Delete(request)
	if err != nil {
	}
}
func GetRequestById(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	requestRepository := repository.NewRequestRepository(db, "Request")

	params := mux.Vars(req)
	var requestid = string(params["requestid"])
	requestIdHex, err := primitive.ObjectIDFromHex(requestid)
	request, err2 := requestRepository.FindByID(requestIdHex)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(request)

}
