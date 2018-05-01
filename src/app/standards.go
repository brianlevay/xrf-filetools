package main

import (
	batch "batchProcess"
	"net/http"
)

func setStandardsHandler(stds *batch.BatchSpectra) {
	http.HandleFunc("/get_stds", func(w http.ResponseWriter, r *http.Request) {
		resp := batchResponse(stds)
		w.Write(resp)
	})

	http.HandleFunc("/update_stds", func(w http.ResponseWriter, r *http.Request) {
		stds.UpdateBatch()
		resp := batchResponse(stds)
		w.Write(resp)
	})
}
