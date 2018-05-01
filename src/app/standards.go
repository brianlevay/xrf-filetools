package main

import (
	batch "batchProcess"
	"net/http"
)

func setStandardsHandler(stds *batch.BatchSpectra) {
	http.HandleFunc("/standards", func(w http.ResponseWriter, r *http.Request) {
		stds.UpdateBatch()
		resp := batchResponse(stds)
		w.Write(resp)
	})
}
