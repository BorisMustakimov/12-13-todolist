package main

import (
	"net/http"
	
	)

func handleRequest() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.ListenAndServe(":7540", nil)
}

func main() {
	handleRequest()
}	
