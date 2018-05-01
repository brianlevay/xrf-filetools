package main

import (
	batch "batchProcess"
	"encoding/json"
	"log"
)

func batchResponse(batch *batch.BatchSpectra) []byte {
	var errStart []byte = []byte("{\"Error\":")
	var dataStart []byte = []byte("{\"Data\":")
	var end []byte = []byte("}")
	var resp []byte
	if batch.Error != nil {
		log.Println("BATCH PROCESSING ERROR:", batch.Error)
		resp = append(errStart, []byte(batch.Error.Error())...)
	} else {
		dataJSON, errJSON := json.Marshal(batch.List)
		if errJSON != nil {
			log.Println("JSON CONVERSION ERROR:", errJSON)
			resp = append(errStart, []byte(errJSON.Error())...)
		} else {
			resp = append(dataStart, dataJSON...)
		}
	}
	resp = append(resp, end...)
	return resp
}

func generalErrorResponse(context string, err error) []byte {
	var errStart []byte = []byte("{\"Error\":")
	var end []byte = []byte("}")
	var resp []byte
	log.Println(context, ":", err)
	resp = append(errStart, []byte(err.Error())...)
	resp = append(resp, end...)
	return resp
}
