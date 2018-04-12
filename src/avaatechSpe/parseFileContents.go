package avaatechSpe

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const defaultChannelN int = 2048

func (spe *SPE) ParseFileContents() error {
	var nextRow string
	var measureTimes []string
	fileBytes, errRead := ioutil.ReadFile(spe.FilePath)
	if errRead != nil {
		return errRead
	}
	fileString := string(fileBytes)
	fileRows := strings.Split(fileString, "\n")
	nRows := len(fileRows)
	for i := 0; i < nRows-1; i++ {
		nextRow = strings.Replace(fileRows[i+1], "\r", "", -1)
		nextRow = strings.Replace(nextRow, ",", ".", -1)
		nextRow = strings.Trim(nextRow, " ")
		if strings.Contains(fileRows[i], "$X_Position:") == true {
			spe.X, _ = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Y_Position:") == true {
			spe.Y, _ = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Slit_DC:") == true {
			spe.DC, _ = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Slit_CC:") == true {
			spe.CC, _ = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$TotalCPS:") == true {
			spe.CPS, _ = strconv.ParseUint(nextRow, 10, 64)
		} else if strings.Contains(fileRows[i], "$ACC_VOLT:") == true {
			spe.Voltage, _ = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$TUBE_CUR:") == true {
			spe.Current, _ = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$MEAS_TIM:") == true {
			measureTimes = strings.Split(nextRow, " ")
			spe.Live, _ = strconv.ParseUint(measureTimes[0], 10, 64)
		} else if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			spe.Date, _ = convertContentDate(nextRow)
		} else if strings.Contains(fileRows[i], "$DATA:") == true {
			spe.Counts = getChannelCounts(nextRow, fileRows[(i+2):])
		}
	}
	return nil
}

func convertContentDate(dateStr string) (time.Time, error) {
	var timeObj time.Time
	var errTime error
	timeObj = time.Now()
	layout := "01-02-2006 15:04:05"
	loc, _ := time.LoadLocation("Local")
	timeObj, errTime = time.ParseInLocation(layout, dateStr, loc)
	if errTime != nil {
		return timeObj, errTime
	}
	return timeObj, nil
}

func getChannelCounts(channelBounds string, channelData []string) []uint64 {
	var counts []uint64
	var i, j, n, nrows, ncols int
	var row string
	var rowPts []string
	var errCts error
	channelBoundPts := strings.Split(channelBounds, " ")
	maxCh, errCh := strconv.ParseUint(channelBoundPts[len(channelBoundPts)-1], 10, 64)
	if errCh != nil {
		counts = make([]uint64, (maxCh + 1))
	} else {
		counts = make([]uint64, defaultChannelN)
	}
	n = 0
	nrows = len(channelData)
	for i = 0; i < nrows; i++ {
		row = strings.Replace(channelData[i], "\r", "", -1)
		row = strings.Trim(row, " ")
		rowPts = strings.Split(row, " ")
		ncols = len(rowPts)
		for j = 0; j < ncols; j++ {
			if strings.Contains(rowPts[j], " ") == false {
				counts[n], errCts = strconv.ParseUint(rowPts[j], 10, 64)
				if errCts != nil {
					counts[n] = 0
				}
				n++
			}
		}
	}
	return counts
}
