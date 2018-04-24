package stdsProcess

import (
	conf "configureSpe"
	"os"
	"path/filepath"
	speprocess "processAvaatechSpe"
	"strings"
)

func GetStds(config *conf.Configuration) *Standards {
	stds := new(Standards)
	stds.Config = config
	stds.Existing = make(map[string]bool)
	errWalk := filepath.Walk(stds.Config.StdsPath, processFile(stds))
	if errWalk != nil {
		stds.Error = errWalk
	}
	return stds
}

func (stds *Standards) UpdateStds() {
	errWalk := filepath.Walk(stds.Config.StdsPath, processFile(stds))
	if errWalk != nil {
		stds.Error = errWalk
	}
}

func processFile(stds *Standards) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		stds.Mtx.Lock()
		defer stds.Mtx.Unlock()
		if err != nil {
			return err
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			_, ok := stds.Existing[path]
			if ok == false {
				spect, errProcess := speprocess.Process(path, stds.Config)
				if errProcess == nil {
					stds.List = append(stds.List, spect)
					stds.Existing[path] = true
				} else {
					stds.Error = errProcess
				}
			}
		}
		return nil
	}
}
