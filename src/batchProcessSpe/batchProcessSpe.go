package batchProcessSpe

import (
	conf "configureSpe"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	speprocess "processAvaatechSpe"
	"strings"
)

func BatchProcess(rootPath string, config *conf.Configuration) ([]byte, error) {
	var batchList []*speprocess.Spectrum
	errWalk := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			spect, errProcess := speprocess.Process(path, config)
			if errProcess == nil {
				batchList = append(batchList, spect)
			}
		}
		return nil
	})
	if errWalk != nil {
		fmt.Println("Error walking the path:", errWalk, "\n")
		return nil, errWalk
	}
	jsonBytes, errJSON := json.Marshal(batchList)
	if errJSON != nil {
		fmt.Println("Unable to convert results to JSON:", errJSON, "\n")
		return nil, errJSON
	}
	return jsonBytes, nil
}
