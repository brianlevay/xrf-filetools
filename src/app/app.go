package main

import (
	"log"
	"net/http"
	"os"
	stds "speSummaries"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("Default path to standards not set")
		return
	}
	defaultPath := os.Args[1]
	summ := new(stds.Summary)
	summ.Initialize(defaultPath)
	summ.RecursiveSearch()

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	setStandardsHandlers(summ)

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
