package main

import (
	"log"
	"net/http"
	"timeUsage"
)

func setTimeUsageHandler() {
	http.HandleFunc("/time_usage", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		if errP != nil {
			log.Println(errP)
			w.Write([]byte("{\"Error\":\"Unable to parse request\"}"))
			return
		}
		var sourcePath string
		sourcePath = r.Form["srcPath"][0]
		if sourcePath == "" {
			log.Println("No source path provided")
			w.Write([]byte("{\"Error\":\"Missing source path\"}"))
			return
		}
		isDirSrc := dirExists(sourcePath)
		if isDirSrc == false {
			log.Println("Invalid source directory")
			w.Write([]byte("{\"Error\":\"Invalid source directory\"}"))
			return
		}

		timeuse := new(timeUsage.TimeStats)
		timeuse.Initialize(sourcePath)
		timeuse.RecursiveSearch()

		w.Write(timeuse.JSON())
	})
}
