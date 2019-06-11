package api

import (
	"github.com/aphiratnimanussonkul/go-mongodb-driver/config"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/models"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)


func AddUser(w http.ResponseWriter, req *http.Request)  {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(db, "User")

	//get variable by path
	params := mux.Vars(req)
	var firstName, lastName, Email string
	firstName = string(params["firstName"])
	lastName = string(params["lastName"])
	Email = string(params["Email"])


	var p models.User
	p.Name = firstName + " " + lastName
	p.Email = Email
	userRepository.Save(&p)

}

func GetUserByEmail(w http.ResponseWriter, req *http.Request)  {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(db, "User")

	//get variable by path
	params := mux.Vars(req)
	var Email string
	Email = string(params["Email"])

	user, err2 := userRepository.FindByEmail(Email)
	if err2 != nil{
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(user)

}

//func AddUserSubject(w http.ResponseWriter, req *http.Request)  {
//	//
//	db, err := config.GetMongoDB()
//	if err != nil {
//		fmt.Println(err)
//	}
//	userRepository := repository.NewUserRepository(db, "User")
//	subjectRepository := repository.NewSubjectRepository(db, "Subject")
//
//	//get variable by path
//	params := mux.Vars(req)
//	var firstName, lastName, Email string
//	firstName = string(params["firstName"])
//	lastName = string(params["lastName"])
//	Email = string(params["Email"])
//	var codeSubject = string(params["code"])
//	subject, err := subjectRepository.FindByCode(codeSubject)
//	var p models.User
//	p.Name = firstName + " " + lastName
//	p.Email = Email
//	userRepository.SaveSubject(subject, &p)
//
//}
func FollowSubject(w http.ResponseWriter, req *http.Request)  {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(db, "User")
	subjectRepository := repository.NewSubjectRepository(db, "Subject")

	//get variable by path
	params := mux.Vars(req)
	var email = string(params["email"])
	var codeSubject = string(params["code"])
	user, err := userRepository.FindByEmail(email)
	subject, err := subjectRepository.FindByCode(codeSubject)
	var status = 1
	for i := 0; i < len(user.Subject); i++ {
		if user.Subject[i].Code == codeSubject {
			json.NewEncoder(w).Encode("false")
			status = 0
			break
		}
	}
	if status == 1 {
		user.Subject = append(user.Subject, subject)
		userRepository.Update(user)
	}
}
func CreateUser(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	// Unmarshal
	var msg models.User
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(db, "User")

	userTemp, err := userRepository.FindByEmail(msg.Email)
	if userTemp == nil {
		var user models.User
		user.Email = msg.Email
		user.Name = msg.Name
		user.Picture = msg.Picture
		user.StudentId = msg.StudentId
		user.Major = msg.Major
		userRepository.Save(&user)
	} else {
		json.NewEncoder(w).Encode("Can not confirm account this email is already used!")
	}

}
func UnfollowSubject(w http.ResponseWriter, req *http.Request)  {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	userRepository := repository.NewUserRepository(db, "User")

	//get variable by path
	params := mux.Vars(req)
	var email = string(params["email"])
	var codeSubject = string(params["code"])
	user, err := userRepository.FindByEmail(email)
	var subjects models.SubjectPointer
	for i := 0; i < len(user.Subject); i++ {
		if user.Subject[i].Code != codeSubject {
			subjects = append(subjects, user.Subject[i])
		}
	}
	user.Subject = subjects
	userRepository.Update(user);
}
