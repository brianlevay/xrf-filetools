package speSummaries

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Highest level function //
func (summ *Summary) RecursiveSearch() {
	err := filepath.Walk(summ.SourcePath, readFileContents(summ))
	if err != nil {
		log.Fatal(err)
	}
}

// Called on each file during recursiveSearch, uses closure to bind pointer to struct //
func readFileContents(summ *Summary) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			if info.ModTime().After(summ.LastDate) == true {
				if info.ModTime().After(summ.NextDate) == true {
					summ.NextDate = info.ModTime()
				}
				fileBytes, errRead := ioutil.ReadFile(path)
				if errRead != nil {
					log.Print(errRead)
					return nil
				}
				fileString := string(fileBytes)
				getContents(info.Name(), fileString, summ)
			}
		}
		return nil
	}
}

// "Name","X","Date","CPS","kVp","mA","DC Slit","CC Slit"
// Called on each file inside WalkFunc closure //
func getContents(fileName string, fileContents string, summ *Summary) {
	var namePts []string
	var nextRow string

	if strings.Contains(fileName, "!") == true {
		namePts = strings.Split(fileName, "!")
	} else {
		namePts = strings.Split(fileName, " ")
	}
	summ.Name = append(summ.Name, namePts[0])

	fileRows := strings.Split(fileContents, "\n")
	nRows := len(fileRows)
	for i := 0; i < nRows-1; i++ {
		nextRow = strings.Replace(fileRows[i+1], "\r", "", -1)
		nextRow = strings.Replace(nextRow, ",", ".", -1)
		if strings.Contains(fileRows[i], "$X_Position:") == true {
			summ.X = append(summ.X, nextRow)
		} else if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			summ.Date = append(summ.Date, nextRow)
		} else if strings.Contains(fileRows[i], "$TotalCPS:") == true {
			summ.CPS = append(summ.CPS, nextRow)
		} else if strings.Contains(fileRows[i], "$ACC_VOLT:") == true {
			summ.KVp = append(summ.KVp, nextRow)
		} else if strings.Contains(fileRows[i], "$TUBE_CUR:") == true {
			summ.Curr = append(summ.Curr, nextRow)
		} else if strings.Contains(fileRows[i], "$Slit_DC:") == true {
			summ.DC = append(summ.DC, nextRow)
		} else if strings.Contains(fileRows[i], "$Slit_CC:") == true {
			summ.CC = append(summ.CC, nextRow)
		}
	}
}
