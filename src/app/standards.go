package main

import (
	"log"
	"net/http"
	stds "speThroughput"
	"strings"
)

func setStandardsHandler(through *stds.Throughput) {
	http.HandleFunc("/update_stds", func(w http.ResponseWriter, r *http.Request) {
		through.Mtx.Lock()
		defer through.Mtx.Unlock()
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
		if strings.Compare(sourcePath, through.SourcePath) != 0 {
			through.Initialize(sourcePath)
		}
		through.RecursiveSearch()
		w.Write(through.JSON())
	})
}
