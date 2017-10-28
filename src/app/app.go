package main

import (
	"log"
	"net/http"
	"os"
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

func dirExists(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	if stat.IsDir() == false {
		return false
	}
	return true
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
