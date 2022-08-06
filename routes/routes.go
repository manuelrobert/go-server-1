package routes

import (
	"github.com/gorilla/mux"
	"github.com/manuelrobert/test_1_server/controllers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/students", controllers.GetAllStudents).Methods("GET")
	router.HandleFunc("/api/student", controllers.CreateOneStudent).Methods("POST")
	router.HandleFunc("/api/student/{id}", controllers.UpdateOneStudent).Methods("PUT")
	router.HandleFunc("/api/student/{id}", controllers.GetOneStudent).Methods("GET")
	router.HandleFunc("/api/student/{id}", controllers.DeleteOneStudent).Methods("DELETE")

	return router
}
