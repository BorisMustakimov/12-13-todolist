package main

import (
	"net/http"

	"github.com/BorisMustakimov/12-13-todolist/nextdate"
	sqltable "github.com/BorisMustakimov/12-13-todolist/sql_table"
)

func handleRequest() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":7540", nil)
}

func main() {
	sqltable.Sql_table()
	handleRequest()
	nextdate.NextDate()

}
