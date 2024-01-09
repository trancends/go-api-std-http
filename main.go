package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type responseSucces struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"Data,omitempty"`
}

type responseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

var tasks []Task

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(responseError{
				Status:  "error",
				Message: "Method Not Allowed",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequestMethodDelete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			json.NewEncoder(w).Encode(responseError{
				Status:  "error",
				Message: "Method Not Allowed",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method Client: ", r.Method)
		w.Write([]byte("TODO"))
	})

	http.Handle("/task", RequestMethodGet(http.HandlerFunc(getTask)))
	http.Handle("/task/", RequestMethodGet(http.HandlerFunc(getTaskById)))

	http.HandleFunc("/task/add", addTask)
	http.HandleFunc("/task/update", updateTask)

	http.Handle("/task/delete", RequestMethodDelete(http.HandlerFunc(deleteTask)))
	http.Handle("/task/delete/all", RequestMethodDelete(http.HandlerFunc(deleteAllTask)))

	fmt.Println("Server Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Hello")
}

func addTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{
			Status:  "Error",
			Message: "Failed to decode json",
		})
		return
	}

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(responseSucces{
		Status:  "OK",
		Message: "Sucesfully Added New Task",
		Data:    tasks,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func getTaskById(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/task/")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{
			Status:  "error",
			Message: "id needs to be number",
		})
		return
	}
	for _, task := range tasks {
		if task.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(responseSucces{
				Status:  "OK",
				Message: "Sucesfully get data",
				Data:    task,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(responseSucces{
		Status:  "error",
		Message: "task not found",
	})
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(responseSucces{
		Status:  "OK",
		Message: "Data",
		Data:    tasks,
	}); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func deleteAllTask(w http.ResponseWriter, r *http.Request) {
	tasks = []Task{}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseSucces{
		Status:  "Succes",
		Message: "All Task Deleted",
		Data:    tasks,
	})
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{
			Status:  "Error",
			Message: "Failed to decode json",
		})
		return
	}

	for index, task := range tasks {
		if task.ID == newTask.ID {
			tasks = append(tasks[:index], tasks[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(responseSucces{
				Status:  "Succes",
				Message: "Task Deleted",
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(responseError{
		Status:  "error",
		Message: "Task Not Found",
	})
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(responseError{Message: "Metod Not Allowed", Status: "error"})
		return
	}

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{
			Status:  "Error",
			Message: "Failed to decode json",
		})
		return
	}

	for index, task := range tasks {
		if task.ID == newTask.ID {
			tasks[index].Title = newTask.Title
			tasks[index].Description = newTask.Description
			tasks[index].Status = newTask.Status
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(responseSucces{
				Status:  "success",
				Message: "Task Updated",
				Data:    tasks,
			})
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(responseSucces{
		Status:  "error",
		Message: "Task Not Found",
	})
}
