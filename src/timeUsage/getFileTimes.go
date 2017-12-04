package timeUsage

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Highest level function //
func (times *TimeStats) RecursiveSearch() {
	err := filepath.Walk(times.SourcePath, readFileContents(times))
	if err != nil {
		log.Fatal(err)
	}
}

// Called on each file during recursiveSearch, uses closure to bind pointer to struct //
func readFileContents(times *TimeStats) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			fileBytes, errRead := ioutil.ReadFile(path)
			if errRead != nil {
				log.Print(errRead)
				return nil
			}
			fileString := string(fileBytes)
			getContents(info.Name(), fileString, times)
		}
		return nil
	}
}

// Called on each file inside WalkFunc closure //
func getContents(fileName string, fileContents string, times *TimeStats) {
	var Name, Timestamp string
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
		if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			Timestamp = nextRow
			break
		}
	}
	updateData(Name, Timestamp, times)
}

// Performs the necessary date conversions and updates the data structure //
func updateData(Name string, Timestamp string, times *TimeStats) {
	timestampPts := strings.Split(Timestamp, " ")
	dayMeas := timestampPts[0]
	layout := "01-02-2006 15:04:05"
	loc, _ := time.LoadLocation("Local")
	timeObj, err1 := time.ParseInLocation(layout, Timestamp, loc)
	if err1 != nil {
		log.Print(err1)
		return
	}
	daySummary, ok := times.DataMap[dayMeas]
	if ok == false {
		daySummary = DaySummary{}
		daySummary.Start = timeObj
		daySummary.Finish = timeObj
		daySummary.PtsSect = make(map[string]int)
		daySummary.PtsSect[Name]++
		times.DataMap[dayMeas] = daySummary
	} else {
		if timeObj.Before(daySummary.Start) {
			daySummary.Start = timeObj
		}
		if timeObj.After(daySummary.Finish) {
			daySummary.Finish = timeObj
		}
		daySummary.PtsSect[Name]++
	}
}