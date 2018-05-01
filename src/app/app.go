package main

import (
	batch "batchProcess"
	conf "configureSpe"
	"log"
	"net/http"
)

func main() {
	port := ":8080"
	startupDisplay(port)

	stdsPath, config := conf.ReadConfig()
	log.Println("Beginning standards processing...")
	stds := batch.GetBatchSpectra(stdsPath, config)
	if stds.Error != nil {
		log.Println(stds.Error)
	} else {
		log.Println("Standards processing complete.", len(stds.List), "files processed.")
	}

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	setStandardsHandler(stds)
	setSectionHandler(config)
	http.ListenAndServe(port, nil)
}
