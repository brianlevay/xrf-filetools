package main

import (
	"log"
	"net/http"
	stats "sampleStats"
	"strings"
)

func setSampleStatsHandler() {
	http.HandleFunc("/sample_stats", func(w http.ResponseWriter, r *http.Request) {
		errP := r.ParseForm()
		if errP != nil {
			log.Println(errP)
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
			outName = "sample_stats"
		}
		outName = strings.Split(outName, ".")[0] + ".csv"

		samples := new(stats.SampleStats)
		samples.Initialize(sourcePath)
		samples.RecursiveSearch()
		errW := samples.WriteToCSV(outPath, outName)
		if errW != nil {
			log.Println(errW)
			w.Write([]byte("{\"Error\":\"Unable to save csv file\"}"))
			return
		}
		w.Write([]byte("{\"Error\":\"none\"}"))
	})
}
