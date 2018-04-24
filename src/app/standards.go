package main

import (
	"encoding/json"
	"log"
	"net/http"
	stds "stdsProcess"
)

func setStandardsHandler(stds *stds.Standards) {
	http.HandleFunc("/get_stds", func(w http.ResponseWriter, r *http.Request) {
		resp := createResponse(stds)
		w.Write(resp)
	})

	http.HandleFunc("/update_stds", func(w http.ResponseWriter, r *http.Request) {
		stds.UpdateStds()
		resp := createResponse(stds)
		w.Write(resp)
	})
}

func createResponse(stds *stds.Standards) []byte {
	var errStart []byte = []byte("{\"Error\":\"")
	var dataStart []byte = []byte("{\"Data\":\"")
	var end []byte = []byte("\"}")
	var resp []byte
	if stds.Error != nil {
		resp = append(errStart, []byte(stds.Error.Error())...)
	} else {
		dataJSON, errJSON := json.Marshal(stds.List)
		if errJSON != nil {
			log.Println(errJSON)
			resp = append(errStart, []byte(errJSON.Error())...)
		} else {
			resp = append(dataStart, dataJSON...)
		}
	}
	resp = append(resp, end...)
	return resp
}
