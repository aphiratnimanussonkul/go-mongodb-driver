package api

import (
	"github.com/aphiratnimanussonkul/go-mongodb-driver/config"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/models"
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AddFaculty(w http.ResponseWriter, req *http.Request) {

	//
	db, err := config.GetMongoDB()
	fmt.Println(db)
	if err != nil {
		fmt.Println(err)
	}
	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Faculty")
	//

	params := mux.Vars(req)
	var name string
	name = string(params["name"])

	var p models.Faculty
	p.Name = name
	facultyRepository.Save(&p)
}

func GetFacultyById(w http.ResponseWriter, req *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Faculty")
	params := mux.Vars(req)
	idHex, err := primitive.ObjectIDFromHex(params["id"])
	faculty, err := facultyRepository.FindByID(idHex)
	json.NewEncoder(w).Encode(faculty)
}

func GetFacultyAll(w http.ResponseWriter, req *http.Request) {

	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Faculty")
	faculty, err2 := facultyRepository.FindAll()
	if err2 != nil {
	}
	json.NewEncoder(w).Encode(faculty)
}

func GetFacultyByEmail(w http.ResponseWriter, req *http.Request) {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Faculty")
	userRepository := repository.NewUserRepository(db, "User")
	majorRepository := repository.NewMajorRepository(db, "Major")
	params := mux.Vars(req)
	var email = string(params["email"])
	var facultyAll models.FacultiesPointer
	var tempFacultyName []string
	var isHave = false
	user, err := userRepository.FindByEmail(email)
	for i := 0; i < len(user.Subject); i++ {
		major, err := majorRepository.FindByName(user.Subject[i].Major.Name)
		if err != nil {
		}
		faculty, err := facultyRepository.FindByName(major.Faculty.Name)

		for j := 0; j < len(tempFacultyName); j++ {
			if tempFacultyName[j] == faculty.Name {
				isHave = true
				break
			} else {
				isHave = false
			}
		}
		if !isHave {
			facultyAll = append(facultyAll, faculty)
			tempFacultyName = append(tempFacultyName, faculty.Name)
		}

	}
	json.NewEncoder(w).Encode(facultyAll)
}
func DeleteFaculty(w http.ResponseWriter, req *http.Request) {
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Faculty")
	params := mux.Vars(req)
	var facultyName = string(params["facultyname"])
	facultyRepository.DeleteByName(facultyName)
}
