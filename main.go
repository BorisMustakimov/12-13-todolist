package main

import (
	"net/http"

	handlers "github.com/BorisMustakimov/12-13-todolist/handlers"
	nextdate "github.com/BorisMustakimov/12-13-todolist/nextdate"
	service "github.com/BorisMustakimov/12-13-todolist/service"
	sqltable "github.com/BorisMustakimov/12-13-todolist/sql_table"
	"github.com/go-chi/chi/v5"
)

func handleRequest() {

	var x service.TaskService

	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./web"))
	r.Handle("/*", http.StripPrefix("/", fileServer))

	taskHandler := handlers.NewTaskHandler(x)

	r.Get("/api/nextdate", nextdate.HandlerNextDate)

	r.Post("/api/task", http.HandlerFunc(taskHandler.TaskHandler))

	http.ListenAndServe(":7540", r)
}

func main() {

	handleRequest()
	sqltable.Sql_table()
}
