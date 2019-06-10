package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
"net/http"
"CPEProject/src/models"
"CPEProject/config"
"CPEProject/src/repository"
)


func AddFavoriteSubject(w http.ResponseWriter, req *http.Request)  {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	favoriteSubjectRepository := repository.NewFavoriteSubjectRepository(db, "FavoriteSubject")
	userRepository := repository.NewUserRepository(db, "User")
	subjectRepository := repository.NewSubjectRepository(db, "Subject")

	params := mux.Vars(req)
	var email = string(params["email"])
	var code = string(params["code"])
	user, err2 := userRepository.FindByEmail(email)
	if err2 != nil {
		fmt.Println(err2)
	}

	subject, err3 := subjectRepository.FindByCode(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	var p models.FavoriteSubject
	p.User = user
	p.Subject = subject
	favoriteSubjectRepository.Save(&p)

}


//Default data FavoriteSubject
func AddFavoriteSubjectDefault(email string, code string)  {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	favoriteSubjectRepository := repository.NewFavoriteSubjectRepository(db, "FavoriteSubject")
	subjectRepository := repository.NewSubjectRepository(db, "Subject")
	userRepository := repository.NewUserRepository(db, "User")

	user, err2 := userRepository.FindByEmail(email)
	if err2 != nil {
		fmt.Println(err2)
	}

	subject, err3 := subjectRepository.FindByCode(code)
	if err3 != nil {
		fmt.Println(err3)
	}
	var p models.FavoriteSubject
	p.User = user
	p.Subject = subject
	favoriteSubjectRepository.Save(&p)
}
func GetFavoriteSubjectByUserEmail(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	favoriteSubjectRepository := repository.NewFavoriteSubjectRepository(db, "FavoriteSubject")

	params := mux.Vars(req)
	var email = string(params["email"])
	user, err2 := favoriteSubjectRepository.FindByEmail(email)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(user)

}