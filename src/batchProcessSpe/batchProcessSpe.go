package batchProcessSpe

import (
	conf "configureSpe"
	"fmt"
	"os"
	"path/filepath"
	speprocess "processAvaatechSpe"
	"strings"
)

func NewBatch(rootPath string, config *conf.Configuration) (*Batch, error) {
	batch := new(Batch)
	batch.Root = rootPath
	batch.Config = config
	batch.Existing = make(map[string]bool)
	errWalk := filepath.Walk(rootPath, processFile(batch))
	if errWalk != nil {
		fmt.Println("Error walking the path:", errWalk, "\n")
		return nil, errWalk
	}
	return batch, nil
}

func (batch *Batch) Update() {
	errWalk := filepath.Walk(batch.Root, processFile(batch))
	if errWalk != nil {
		fmt.Println("Error walking the path:", errWalk, "\n")
	}
}

func processFile(batch *Batch) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
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
				}
			}
		}
		return nil
	}
}
