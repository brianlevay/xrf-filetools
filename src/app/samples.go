package main

import (
	"log"
	"net/http"
	stats "sampleStats"
	"strconv"
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
		var toSave bool
		var errB error
		sourcePath = r.Form["srcPath"][0]
		toSave, errB = strconv.ParseBool(r.Form["toSave"][0])
		if errB != nil {
			toSave = false
		}
		outPath = r.Form["outPath"][0]
		outName = r.Form["outName"][0]
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

		if toSave == true {
			isDirOut := dirExists(outPath)
			if isDirOut == false {
				log.Println("Invalid output directory")
				w.Write([]byte("{\"Error\":\"Invalid output directory\"}"))
				return
			}
			if outName == "" {
				outName = "sample_stats"
			}
			outName = strings.Split(outName, ".")[0] + ".csv"
		}

		samples := new(stats.SampleStats)
		samples.Initialize(sourcePath)
		samples.RecursiveSearch()

		if toSave == true {
			errW := samples.WriteToCSV(outPath, outName)
			if errW != nil {
				log.Println(errW)
				w.Write([]byte("{\"Error\":\"Unable to save csv file\"}"))
				return
			}
		}

		w.Write(samples.JSON())
	})
}
