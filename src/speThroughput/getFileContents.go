package speThroughput

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Highest level function //
func (through *Throughput) RecursiveSearch() {
	err := filepath.Walk(through.SourcePath, readFileContents(through))
	if err != nil {
		log.Fatal(err)
	}
	through.LastDate = through.NextDate
}

// Called on each file during recursiveSearch, uses closure to bind pointer to struct //
func readFileContents(through *Throughput) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			if info.ModTime().After(through.LastDate) == true {
				if info.ModTime().After(through.NextDate) == true {
					through.NextDate = info.ModTime()
				}
				fileBytes, errRead := ioutil.ReadFile(path)
				if errRead != nil {
					log.Print(errRead)
					return nil
				}
				fileString := string(fileBytes)
				getContents(info.Name(), fileString, through)
			}
		}
		return nil
	}
}

// "Name","X","Date","CPS","kVp","mA","DC Slit","CC Slit"
// Called on each file inside WalkFunc closure //
func getContents(fileName string, fileContents string, through *Throughput) {
	var Name, X, Date, CPS, KVp, Curr, DC, CC string
	var namePts []string
	var nextRow string

	if strings.Contains(fileName, "!") == true {
		namePts = strings.Split(fileName, "!")
	} else {
		namePts = strings.Split(fileName, " ")
	}
	Name = namePts[0]

	fileRows := strings.Split(fileContents, "\n")
	nRows := len(fileRows)
	for i := 0; i < nRows-1; i++ {
		nextRow = strings.Replace(fileRows[i+1], "\r", "", -1)
		nextRow = strings.Replace(nextRow, ",", ".", -1)
		if strings.Contains(fileRows[i], "$X_Position:") == true {
			X = nextRow
		} else if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			Date = nextRow
		} else if strings.Contains(fileRows[i], "$TotalCPS:") == true {
			CPS = nextRow
		} else if strings.Contains(fileRows[i], "$ACC_VOLT:") == true {
			KVp = nextRow
		} else if strings.Contains(fileRows[i], "$TUBE_CUR:") == true {
			Curr = nextRow
		} else if strings.Contains(fileRows[i], "$Slit_DC:") == true {
			DC = nextRow
		} else if strings.Contains(fileRows[i], "$Slit_CC:") == true {
			CC = nextRow
		}
	}

	row := &Row{Name: Name, X: X, Date: Date, CPS: CPS, KVp: KVp, Curr: Curr, DC: DC, CC: CC}
	rowJSON, _ := json.Marshal(row)
	through.Data = append(through.Data, string(rowJSON))
}
