package main

import (
	conf "configureSpe"
	"log"
	"net/http"
	standards "stdsProcess"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)
	port := ":8080"
	startupDisplay(port)

	config := conf.ReadConfig()
	log.Println("Beginning standards processing...")
	stds := standards.GetStds(config)
	if stds.Error != nil {
		log.Println(stds.Error)
	} else {
		log.Println("Standards processing complete.", len(stds.List), "files processed.")
	}
	setStandardsHandler(stds)
	http.ListenAndServe(port, nil)
}
