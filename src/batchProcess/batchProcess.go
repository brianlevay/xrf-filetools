package batchProcess

import (
	conf "configureSpe"
	"os"
	"path/filepath"
	speprocess "processAvaatechSpe"
	"strings"
)

func GetBatchSpectra(rootPath string, config *conf.Configuration) *BatchSpectra {
	batch := new(BatchSpectra)
	batch.RootPath = rootPath
	batch.Config = config
	batch.Existing = make(map[string]bool)
	errWalk := filepath.Walk(batch.RootPath, processFile(batch))
	if errWalk != nil {
		batch.Error = errWalk
	}
	return batch
}

func (batch *BatchSpectra) UpdateBatch() {
	errWalk := filepath.Walk(batch.RootPath, processFile(batch))
	if errWalk != nil {
		batch.Error = errWalk
	}
}

func processFile(batch *BatchSpectra) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		batch.Mtx.Lock()
		defer batch.Mtx.Unlock()
		if err != nil {
			return err
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			_, ok := batch.Existing[path]
			if ok == false {
				spect, errProcess := speprocess.Process(path, batch.Config)
				if errProcess == nil {
					batch.List = append(batch.List, spect)
					batch.Existing[path] = true
				} else {
					batch.Error = errProcess
				}
			}
		}
		return nil
	}
}
