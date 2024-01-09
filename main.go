package main

import (
	"fmt"
	"log"
	"net/http"
)

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

	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok, starting..."))
	})

	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method Client: ", r.Method)
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/task/add", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method Client: ", r.Method)
		if r.Method != "POST" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/task/update", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method Client: ", r.Method)
		if r.Method != "UPDATE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Hello World!"))
	})

	http.HandleFunc("/task/delete", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Method Client: ", r.Method)
		if r.Method != "DELETE" {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		w.Write([]byte("Hello World!"))
	})

	fmt.Println("Server Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	fmt.Println("Hello")
}
