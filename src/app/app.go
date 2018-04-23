package main

import (
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	port := ":8080"
	startupDisplay(port)
	testPackages()
	http.ListenAndServe(port, nil)
}
