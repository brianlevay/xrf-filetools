package main

import (
	"fmt"
	"log"
	"net/http"
	stds "speSummaries"
)

func main() {
	sourcePath := ".misc/testData"
	summ := stds.NewSummary(sourcePath)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	http.HandleFunc("/all_stds", func(w http.ResponseWriter, r *http.Request) {
		summ.Reset(sourcePath)
		summ.RecursiveSearch()
		fmt.Fprintf(w, summ.JSON())
	})

	http.HandleFunc("/update_stds", func(w http.ResponseWriter, r *http.Request) {
		summ.RecursiveSearch()
		fmt.Fprintf(w, summ.JSON())
	})

	log.Println("Listening at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
