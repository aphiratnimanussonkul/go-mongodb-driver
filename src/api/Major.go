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


func AddMajor(w http.ResponseWriter, req *http.Request)  {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	majorRepository := repository.NewMajorRepository(db, "Major")
	facultyRepository := repository.NewFacultyRepositoryMongo(db, "Faculty")

	params := mux.Vars(req)
	var facultyName string
	facultyName = string(params["facultyName"])
	faculty, err2 := facultyRepository.FindByName(facultyName)
	if err2 != nil {
		fmt.Println(err)
	}
	var name string
	name = string(params["name"])

	var p models.Major
	p.Name = name
	p.Faculty = faculty
	majorRepository.Save(&p)

}


func GetMajorByFaculty(w http.ResponseWriter, req *http.Request)  {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	majorRepository := repository.NewMajorRepository(db, "Major")

	params := mux.Vars(req)
	var facultyName = string(params["facultyName"])
	major, err2 := majorRepository.FindByFaculty(facultyName)
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(major)


}

func GetMajorByFacultyEmail(w http.ResponseWriter, req *http.Request)  {
	//
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	majorRepository := repository.NewMajorRepository(db, "Major")
	userRepository := repository.NewUserRepository(db, "User")
	params := mux.Vars(req)
	var facultyName = string(params["facultyName"])
	var email = string(params["email"])
	user, err := userRepository.FindByEmail(email)
	var majors models.MajorPointer
	var isHave = false
	var tempMajor []string
	for i := 0; i < len(user.Subject); i++ {
		major, err := majorRepository.FindByName(user.Subject[i].Major.Name)
		if err != nil {
		}
		for j := 0; j < len(tempMajor); j++ {
			if majors[j].Name == major.Name {
				isHave = true
			} else {
				isHave = false
				continue
			}
		}
		if major.Faculty.Name == facultyName && !isHave{
			tempMajor = append(tempMajor, major.Name)
			majors = append(majors, major)
		}
	}
	json.NewEncoder(w).Encode(majors)
}
func DeleteMajor(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Go Mongo Db")
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	majorRepository := repository.NewMajorRepository(db, "Major")
	params := mux.Vars(req)
	var majorName = string(params["majorname"])
	majorRepository.DeleteByName(majorName)
}

func GetMajorAll(w http.ResponseWriter, req *http.Request) {

	//
	fmt.Println("Go Mongo Db")
	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}
	majorRepository := repository.NewMajorRepository(db, "Major")
	major, err2 := majorRepository.FindAll()
	if err2 != nil {
		fmt.Println(err2)
	}
	json.NewEncoder(w).Encode(major)
}