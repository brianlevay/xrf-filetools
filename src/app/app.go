package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := ":8080"
	htmlPath := "http://localhost" + port
	log.Println("Listening at " + htmlPath)
	http.ListenAndServe(port, nil)
}
