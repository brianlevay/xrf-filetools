package main

import (
	"log"
	"net/http"
	stds "speSummaries"
)

func setStandardsHandlers(summ *stds.Summary) {
	http.HandleFunc("/refresh_stds", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			w.Write([]byte("{'Error':Unable to parse request}"))
			return
		}
		var sourcePath string
		sourcePath = r.Form["stdsPath"][0]
		if sourcePath == "" {
			log.Println("No path provided for standards.")
			w.Write([]byte("{'Error':No path provided}"))
			return
		}
		isDir := dirExists(sourcePath)
		if isDir == false {
			log.Println("Invalid directory.")
			w.Write([]byte("{'Error':Invalid directory}"))
			return
		}
		summ.Initialize(sourcePath)
		summ.RecursiveSearch()
		w.Write(summ.JSON())
	})

	http.HandleFunc("/update_stds", func(w http.ResponseWriter, r *http.Request) {
		summ.RecursiveSearch()
		w.Write(summ.JSON())
	})

	http.HandleFunc("/current_stds", func(w http.ResponseWriter, r *http.Request) {
		w.Write(summ.JSON())
	})
}
