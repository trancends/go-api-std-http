package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type responseSucces struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type responseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type todo struct {
	id   int
	task string
}

var todos []todo

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method Client: ", r.Method)
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/task", getTask)
	http.HandleFunc("/task/add", addTask)
	http.HandleFunc("/task/update", updateTask)
	http.HandleFunc("/task/delete", deleteTask)

	fmt.Println("Server Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Hello")
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Task added",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(responseError{Message: "Metod Not Allowed", Status: "error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(responseSucces{
		Status:  "succes",
		Message: "Data",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(responseError{Message: "Metod Not Allowed", Status: "error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(responseSucces{
		Status:  "succes",
		Message: "Task Deleted",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "UPDATE" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(responseError{Message: "Metod Not Allowed", Status: "error"})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(responseSucces{
		Status:  "succes",
		Message: "Task Updated",
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
