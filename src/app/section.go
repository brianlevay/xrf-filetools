package main

import (
	batch "batchProcess"
	conf "configureSpe"
	"errors"
	"net/http"
)

func setSectionHandler(config *conf.Configuration) {
	http.HandleFunc("/get_section", func(w http.ResponseWriter, r *http.Request) {
		var resp []byte
		errP := r.ParseForm()
		if errP != nil {
			resp = generalErrorResponse("PARSING FORM", errP)
			w.Write(resp)
			return
		}
		sectionPath := r.Form["sectionPath"][0]
		if sectionPath == "" {
			resp = generalErrorResponse("SECT PATH", errors.New("Missing source path"))
			w.Write(resp)
			return
		}
		isDir := dirExists(sectionPath)
		if isDir == false {
			resp = generalErrorResponse("SECT PATH", errors.New("Invalid source directory"))
			w.Write(resp)
			return
		}
		sect := batch.GetBatchSpectra(sectionPath, config)
		resp = batchResponse(sect)
		w.Write(resp)
	})
}
