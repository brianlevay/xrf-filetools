package timeUsage

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Highest level function //
func (timeuse *TimeStats) RecursiveSearch() {
	err := filepath.Walk(timeuse.SourcePath, readFileContents(timeuse))
	if err != nil {
		log.Fatal(err)
	}
	timeuse.FlattenToArr()
}

// Called on each file during recursiveSearch, uses closure to bind pointer to struct //
func readFileContents(timeuse *TimeStats) filepath.WalkFunc {
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
			getContents(info.Name(), fileString, timeuse)
		}
		return nil
	}
}

// Called on each file inside WalkFunc closure //
func getContents(fileName string, fileContents string, timeuse *TimeStats) {
	var Name, Seconds, Timestamp string
	var nextRow string

	var namePts []string
	if strings.Contains(fileName, "!") {
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
		if strings.Contains(fileRows[i], "$MEAS_TIM:") == true {
			Seconds = nextRow
		}
		if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			Timestamp = nextRow
			break
		}
	}
	updateData(Name, Seconds, Timestamp, timeuse)
}

// Performs the necessary date conversions and updates the data structure //
func updateData(Name string, Seconds string, Timestamp string, timeuse *TimeStats) {
	secondsPts := strings.Split(Seconds, " ")
	actualSec, err1 := strconv.ParseInt(secondsPts[len(secondsPts)-1], 10, 64)
	if err1 != nil {
		log.Print(err1)
		return
	}

	timestampPts := strings.Split(Timestamp, " ")
	dayMeas := timestampPts[0]
	layout := "01-02-2006 15:04:05"
	loc, _ := time.LoadLocation("Local")
	timeObj, err2 := time.ParseInLocation(layout, Timestamp, loc)
	if err2 != nil {
		log.Print(err2)
		return
	}

	daySummary, ok := timeuse.DataMap[dayMeas]
	if ok == false {
		daySummary = DaySummary{}
		daySummary.Start = timeObj
		daySummary.Finish = timeObj
		daySummary.Seconds = actualSec
		daySummary.Points = 1
		daySummary.Sections = make(map[string]bool)
		daySummary.Sections[Name] = true
		timeuse.DataMap[dayMeas] = daySummary
	} else {
		if timeObj.Before(daySummary.Start) {
			daySummary.Start = timeObj
		}
		if timeObj.After(daySummary.Finish) {
			daySummary.Finish = timeObj
		}
		daySummary.Seconds += actualSec
		daySummary.Points += 1
		daySummary.Sections[Name] = true
		timeuse.DataMap[dayMeas] = daySummary
	}
}
