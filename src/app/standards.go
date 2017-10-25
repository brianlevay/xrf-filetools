package main

import (
	"log"
	"net/http"
	stds "speSummaries"
	"strings"
)

func setStandardsHandler(summ *stds.Summary) {
	http.HandleFunc("/update_stds", func(w http.ResponseWriter, r *http.Request) {
		summ.Mtx.Lock()
		defer summ.Mtx.Unlock()
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			w.Write([]byte("{\"Error\":\"Unable to parse request\"}"))
			return
		}
		var sourcePath string
		sourcePath = r.Form["stdsPath"][0]
		if sourcePath == "" {
			log.Println("No path provided for standards.")
			w.Write([]byte("{\"Error\":\"No path provided\"}"))
			return
		}
		isDir := dirExists(sourcePath)
		if isDir == false {
			log.Println("Invalid directory.")
			w.Write([]byte("{\"Error\":\"Invalid directory\"}"))
			return
		}
		if strings.Compare(sourcePath, summ.SourcePath) == 0 {
			summ.Initialize(sourcePath)
		}
		summ.RecursiveSearch()
		w.Write(summ.JSON())
	})
}
