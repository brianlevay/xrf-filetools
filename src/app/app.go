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
	setSectionPointsHandler()
	setTimeUsageHandler()

	port := ":8080"
	htmlPath := "http://localhost" + port
	log.Println("Listening at " + htmlPath)
	http.ListenAndServe(port, nil)
}
