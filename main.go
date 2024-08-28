package main

import (
	"log"
	"net/http"
	"todoApp/todocrud"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/todo/tasks", todocrud.GetTasks).Methods("GET")
	router.HandleFunc("/todo/tasks/{id}", todocrud.GetTask).Methods("GET")
	router.HandleFunc("/todo/tasks", todocrud.CreateTask).Methods("POST")
	// router.HandleFunc("/todo/tasks/{id}", todocrud.UpdateTaskContext).Methods("PUT")
	router.HandleFunc("/todo/tasks/{id}", todocrud.DeleteTask).Methods("DELETE")
	log.Println("SERVER IS RUNNING! :)")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("[WARNING] Error running server")
	}
}
