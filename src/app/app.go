package main

import (
	"log"
	"net/http"
	stds "speThroughput"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	through := new(stds.Throughput)
	through.Initialize("")
	setStandardsHandler(through)

	setSampleStatsHandler()

	log.Println("Listening at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
