package main

import (
	"net/http"

	nextdate "github.com/BorisMustakimov/12-13-todolist/nextdate"
	sqltable "github.com/BorisMustakimov/12-13-todolist/sql_table"
	//handlers "github.com/BorisMustakimov/12-13-todolist/handlers"
	//service "github.com/BorisMustakimov/12-13-todolist/service"
	"github.com/go-chi/chi/v5"
)

func handleRequest() {

	r := chi.NewRouter()

	fileServer := http.FileServer(http.Dir("./web"))
	r.Handle("/*", http.StripPrefix("/", fileServer))

	r.Get("/api/nextdate", nextdate.HandlerNextDate)

	http.ListenAndServe(":7540", r)
}

func main() {

	handleRequest()
	sqltable.Sql_table()
}
