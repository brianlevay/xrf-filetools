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
	var errArr []error
	errArr = make([]error, 8)

	cleanName := strings.Replace(spe.FileName, ",", ".", -1)
	namePts := strings.Split(cleanName, "!")
	spe.Sample = namePts[0]
	spe.Date, errArr[0] = convertNameDate(namePts[4], namePts[5])
	spe.Live, errArr[1] = strconv.ParseUint(namePts[6], 10, 64)
	spe.Voltage, errArr[2] = strconv.ParseFloat(namePts[7], 64)
	spe.Current, errArr[3] = strconv.ParseFloat(namePts[8], 64)
	if errArr[3] == nil {
		spe.Current = spe.Current / 1000
	}
	spe.X, errArr[4] = strconv.ParseFloat(namePts[9], 64)
	spe.Y, errArr[5] = strconv.ParseFloat(namePts[10], 64)
	spe.Filter = namePts[11]
	spe.DC, errArr[6] = strconv.ParseFloat(namePts[12], 64)
	spe.CC, errArr[7] = strconv.ParseFloat(namePts[13], 64)
	for e := 0; e < len(errArr); e++ {
		if errArr[e] != nil {
			return errArr[e]
		}
	}
	return nil
}

func (spe *SPE) ParseNameOld() error {
	var lastLetter, lastTwo, firstTwo, firstLetter, subStr string
	var partLength int
	var xFound bool
	xFound = false
	var errArr []error
	errArr = make([]error, 7)

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
				spe.Live, errArr[0] = strconv.ParseUint(subStr, 10, 64)
			} else if lastTwo == "mm" {
				if firstTwo == "DC" {
					subStr = namePts[i][2 : partLength-2]
					spe.DC, errArr[1] = strconv.ParseFloat(subStr, 64)
				} else if firstTwo == "CC" {
					subStr = namePts[i][2 : partLength-2]
					spe.CC, errArr[2] = strconv.ParseFloat(subStr, 64)
				} else if firstLetter == "X" {
					subStr = namePts[i][1 : partLength-2]
					spe.X, errArr[3] = strconv.ParseFloat(subStr, 64)
				} else if firstLetter == "Y" {
					subStr = namePts[i][1 : partLength-2]
					spe.Y, errArr[4] = strconv.ParseFloat(subStr, 64)
				} else if xFound == false {
					subStr = namePts[i][:partLength-2]
					spe.X, errArr[3] = strconv.ParseFloat(subStr, 64)
					xFound = true
				} else {
					subStr = namePts[i][:partLength-2]
					spe.Y, errArr[4] = strconv.ParseFloat(subStr, 64)
				}
			} else if lastTwo == "kV" {
				subStr = namePts[i][:partLength-2]
				spe.Voltage, errArr[5] = strconv.ParseFloat(subStr, 64)
			} else if lastTwo == "uA" {
				subStr = namePts[i][:partLength-2]
				spe.Current, errArr[6] = strconv.ParseFloat(subStr, 64)
				if errArr[6] == nil {
					spe.Current = spe.Current / 1000
				}
				spe.Filter = namePts[i+1]
			}
		}
	}
	for e := 0; e < len(errArr); e++ {
		if errArr[e] != nil {
			return errArr[e]
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
