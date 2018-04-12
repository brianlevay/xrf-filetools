package avaatechSpe

import (
	"strconv"
	"strings"
	"time"
)

func (spe *SPE) ParseFileName() error {
	if strings.Contains(spe.FileName, "!") == true {
		errNew := spe.ParseNameNew()
		if errNew != nil {
			return errNew
		}
	} else {
		errOld := spe.ParseNameOld()
		if errOld != nil {
			return errOld
		}
	}
	return nil
}

func (spe *SPE) ParseNameNew() error {
	var err error
	cleanName := strings.Replace(spe.FileName, ",", ".", -1)
	namePts := strings.Split(cleanName, "!")
	spe.Sample = namePts[0]
	spe.Date, _ = convertNameDate(namePts[4], namePts[5])
	spe.Live, _ = strconv.ParseUint(namePts[6], 10, 64)
	spe.Voltage, _ = strconv.ParseFloat(namePts[7], 64)
	spe.Current, err = strconv.ParseFloat(namePts[8], 64)
	if err == nil {
		spe.Current = spe.Current / 1000
	}
	spe.X, _ = strconv.ParseFloat(namePts[9], 64)
	spe.Y, _ = strconv.ParseFloat(namePts[10], 64)
	spe.Filter = namePts[11]
	spe.DC, _ = strconv.ParseFloat(namePts[12], 64)
	spe.CC, _ = strconv.ParseFloat(namePts[13], 64)
	return nil
}

func (spe *SPE) ParseNameOld() error {
	var err error
	var lastLetter, lastTwo, firstTwo, firstLetter, subStr string
	var partLength int
	var xFound bool
	xFound = false

	cleanName := strings.Replace(spe.FileName, ",", ".", -1)
	namePts := strings.Split(cleanName, " ")
	nPts := len(namePts)
	spe.Sample = namePts[0]

	for i := 1; i < nPts; i++ {
		partLength = len(namePts[i])
		if partLength >= 2 {
			lastLetter = namePts[i][partLength-1:]
			lastTwo = namePts[i][partLength-2:]
			firstTwo = namePts[i][:2]
			firstLetter = namePts[i][:1]
			if lastLetter == "s" {
				subStr = namePts[i][:partLength-1]
				spe.Live, _ = strconv.ParseUint(subStr, 10, 64)
			} else if lastTwo == "mm" {
				if firstTwo == "DC" {
					subStr = namePts[i][2 : partLength-2]
					spe.DC, _ = strconv.ParseFloat(subStr, 64)
				} else if firstTwo == "CC" {
					subStr = namePts[i][2 : partLength-2]
					spe.CC, _ = strconv.ParseFloat(subStr, 64)
				} else if firstLetter == "X" {
					subStr = namePts[i][1 : partLength-2]
					spe.X, _ = strconv.ParseFloat(subStr, 64)
				} else if firstLetter == "Y" {
					subStr = namePts[i][1 : partLength-2]
					spe.Y, _ = strconv.ParseFloat(subStr, 64)
				} else if xFound == false {
					subStr = namePts[i][:partLength-2]
					spe.X, _ = strconv.ParseFloat(subStr, 64)
					xFound = true
				} else {
					subStr = namePts[i][:partLength-2]
					spe.Y, _ = strconv.ParseFloat(subStr, 64)
				}
			} else if lastTwo == "kV" {
				subStr = namePts[i][:partLength-2]
				spe.Voltage, _ = strconv.ParseFloat(subStr, 64)
			} else if lastTwo == "uA" {
				subStr = namePts[i][:partLength-2]
				spe.Current, err = strconv.ParseFloat(subStr, 64)
				if err == nil {
					spe.Current = spe.Current / 1000
				}
				spe.Filter = namePts[i+1]
			}
		}
	}
	return nil
}

func convertNameDate(dayStr string, timeStr string) (time.Time, error) {
	var timeObj time.Time
	var errTime error
	timeObj = time.Now()
	dateStr := dayStr + " " + timeStr
	layout := "2006-01-02 15-04-05"
	loc, _ := time.LoadLocation("Local")
	timeObj, errTime = time.ParseInLocation(layout, dateStr, loc)
	if errTime != nil {
		return timeObj, errTime
	}
	return timeObj, nil
}
