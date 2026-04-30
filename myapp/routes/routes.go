package routes

import (
	"fmt"
	"log"
	"myapp/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	// creating a new router
	router := mux.NewRouter()

	// Student routes
	router.HandleFunc("/student/all", controller.GetAllStuds).Methods("GET")
	router.HandleFunc("/student/{sid}", controller.GetStud).Methods("GET")
	router.HandleFunc("/student/add", controller.AddStudent).Methods("POST")
	router.HandleFunc("/student/{sid}", controller.UpdateStud).Methods("PUT")
	router.HandleFunc("/student/{sid}", controller.DeleteStud).Methods("DELETE")

	// Course routes
	router.HandleFunc("/course/all", controller.GetAllCourses).Methods("GET")
	router.HandleFunc("/course/add", controller.AddCourse).Methods("POST")
	router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")

	// load static files
	fHandler := http.FileServer(http.Dir("./view"))
	// serve static files as a route by registering all static files on the mux router
	router.PathPrefix("/").Handler(fHandler)

	fmt.Println("Server started successfully")
	// start the http server
	log.Fatal(http.ListenAndServe(":8080", router))
}
