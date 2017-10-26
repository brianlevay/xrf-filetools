package main

import (
	"log"
	"net/http"
)

func setUniqueNamesHandler() {
	http.HandleFunc("/unique_names", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			w.Write([]byte("{\"Error\":\"Unable to parse request\"}"))
			return
		}
		var sourcePath, outPath, outName string
		sourcePath = r.Form["srcPath"][0]
		outPath = r.Form["outPath"][0]
		outName = r.Form["outName"][0]
		if (sourcePath == "") || (outPath == "") {
			log.Println("No path provided for standards.")
			w.Write([]byte("{\"Error\":\"Missing source or output path\"}"))
			return
		}
		isDirSrc := dirExists(sourcePath)
		if isDirSrc == false {
			log.Println("Invalid source directory.")
			w.Write([]byte("{\"Error\":\"Invalid source directory\"}"))
			return
		}
		isDirOut := dirExists(outPath)
		if isDirOut == false {
			log.Println("Invalid output directory.")
			w.Write([]byte("{\"Error\":\"Invalid output directory\"}"))
			return
		}
		if outName == "" {
			outName = "unique_names"
		}
		outName = outName + ".csv"
		//// DO THE WORK HERE ////
		w.Write([]byte("{\"Error\":\"Calculation steps not yet programmed\"}"))
	})
}
