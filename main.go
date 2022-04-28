package main

import (
	"fmt"
	"go-mongo/db"
	"go-mongo/handlers"
	"log"
	"net/http"
	"text/template"
)

var templates, ErrTmpl = template.ParseFiles("index.html")

func main() {
	err := db.CreateConnection()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected...")

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/task/", handlers.CreateTask)
	mux.HandleFunc("/gettask/", handlers.GetTask)
	mux.HandleFunc("/update/", handlers.UpdateTask)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
