package main

import (
	"github.com/aphiratnimanussonkul/go-mongodb-driver/src/api"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	//Subject
	router.HandleFunc("/subject/{name}/{code}/{majorName}", api.AddSubject).Methods("GET")
	router.HandleFunc("/subject/{majorName}", api.GetSubjectByMajor).Methods("GET")
	router.HandleFunc("/subjectbycode/{code}", api.GetSubjectByCode).Methods("GET")
	router.HandleFunc("/subjects", api.GetSubjectAll).Methods("GET")
	router.HandleFunc("/subjectbyemail/{major}/{email}", api.GetSubjectByMajorEmail).Methods("GET")
	router.HandleFunc("/subjectfromuser/{email}", api.GetSubjectFromUser).Methods("GET")
	router.HandleFunc("/deletesubject/{code}", api.DeleteSubject).Methods("GET")
	router.HandleFunc("/subject", api.CreateSubject).Methods("POST")
	//Faculty
	router.HandleFunc("/faculty/{name}", api.AddFaculty).Methods("GET")
	router.HandleFunc("/facultyemail/{email}", api.GetFacultyByEmail).Methods("GET")
	router.HandleFunc("/faculties", api.GetFacultyAll).Methods("GET")
	router.HandleFunc("/deletefaculty/{facultyname}", api.DeleteFaculty).Methods("GET")
	//Major
	router.HandleFunc("/major/{name}/{facultyName}", api.AddMajor).Methods("GET")
	router.HandleFunc("/major/{facultyName}", api.GetMajorByFaculty).Methods("GET")
	router.HandleFunc("/majorbyemail/{facultyName}/{email}", api.GetMajorByFacultyEmail).Methods("GET")
	router.HandleFunc("/deletemajor/{majorname}", api.DeleteMajor).Methods("GET")
	router.HandleFunc("/majors", api.GetMajorAll).Methods("GET")
	// User Ping
	router.HandleFunc("/createuser", api.CreateUser).Methods("POST")
	//User
	router.HandleFunc("/users", api.GetUserAll).Methods("GET")
	router.HandleFunc("/deleteuser/{id}", api.DeleteUserById).Methods("GET")
	router.HandleFunc("/user/{Email}", api.GetUserByEmail).Methods("GET")
	router.HandleFunc("/follow/{email}/{code}", api.FollowSubject).Methods("GET")
	router.HandleFunc("/user/{firstName}/{lastName}/{Email}", api.AddUser).Methods("GET")
	router.HandleFunc("/unfollow/{email}/{code}", api.UnfollowSubject).Methods("GET")
	//Post
	router.HandleFunc("/postvdo/{text}/{email}/{code}/{vdoLink}", api.AddPost).Methods("GET")
	router.HandleFunc("/postfile/{text}/{email}/{code}/{name}/{token}", api.AddPost).Methods("GET")
	router.HandleFunc("/postfull/{text}/{email}/{code}/{vdoLink}/{name}/{token}", api.AddPost).Methods("GET")
	router.HandleFunc("/post/{text}/{email}/{code}", api.AddPost).Methods("POST")
	router.HandleFunc("/post", api.AddPost).Methods("POST")
	router.HandleFunc("/posts", api.GetPostAll).Methods("GET")
	router.HandleFunc("/post/{code}", api.GetPostByCode).Methods("GET")
	router.HandleFunc("/deletepost/{postid}", api.DeletePost).Methods("GET")
	router.HandleFunc("/getpost/{postid}", api.GetPostById).Methods("GET")
	//comment
	router.HandleFunc("/comments", api.GetCommentAll).Methods("GET")
	router.HandleFunc("/commenttext/{text}", api.AddComment).Methods("GET")
	router.HandleFunc("/comment/{postID}", api.AddComment).Methods("POST")
	router.HandleFunc("/deletecomment/{id}", api.DeleteCommentById).Methods("GET")
	router.HandleFunc("/deletecomment/{id}/{postid}", api.DeleteCommentByIdANDPostId).Methods("GET")
	//feedback
	router.HandleFunc("/feedback", api.AddFeedback).Methods("POST")
	router.HandleFunc("/deletefeedback/{id}", api.DeleteFeedback).Methods("GET")
	router.HandleFunc("/feedbacks", api.GetFeedbackAll).Methods("GET")
	//request
	router.HandleFunc("/request", api.AddRequest).Methods("POST")
	router.HandleFunc("/deleterequest/{id}", api.DeleteRequest).Methods("GET")
	router.HandleFunc("/requests", api.GetRequestAll).Methods("GET")
	log.Fatal(http.ListenAndServe(getPort(), handlers.CORS(handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "12345"
	}
	return ":" + port
}
