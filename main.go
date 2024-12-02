package main

import (
	"net/http"

	"github.com/BorisMustakimov/12-13-todolist/nextdate"
	sqltable "github.com/BorisMustakimov/12-13-todolist/sql_table"
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
	sqltable.Sql_table()
	handleRequest()

}
