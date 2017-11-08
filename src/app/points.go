package main

import (
	"log"
	"net/http"
	pts "sectionPoints"
)

func setSectionPointsHandler() {
	http.HandleFunc("/section_points", func(w http.ResponseWriter, r *http.Request) {
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

		points := new(pts.UniquePoints)
		points.Initialize(sourcePath)
		points.RecursiveSearch()

		w.Write(points.JSON())
	})
}
