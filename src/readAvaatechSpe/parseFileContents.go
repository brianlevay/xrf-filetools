package readAvaatechSpe

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func (spe *SPE) ParseFileContents(UTCoffset string) error {
	var nextRow string
	var measureTimes []string
	var errArr []error
	errArr = make([]error, 9)

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
			spe.X, errArr[0] = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Y_Position:") == true {
			spe.Y, errArr[1] = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Slit_DC:") == true {
			spe.DC, errArr[2] = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Slit_CC:") == true {
			spe.CC, errArr[3] = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$TotalCPS:") == true {
			spe.CPS, errArr[4] = strconv.ParseInt(nextRow, 10, 64)
		} else if strings.Contains(fileRows[i], "$ACC_VOLT:") == true {
			spe.Voltage, errArr[5] = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$TUBE_CUR:") == true {
			spe.Current, errArr[6] = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$MEAS_TIM:") == true {
			measureTimes = strings.Split(nextRow, " ")
			spe.Live, errArr[7] = strconv.ParseInt(measureTimes[0], 10, 64)
		} else if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			spe.Date, errArr[8] = convertContentDate(nextRow, UTCoffset)
		} else if strings.Contains(fileRows[i], "$DATA:") == true {
			spe.Counts = getChannelCounts(nextRow, fileRows[(i+2):])
		}
	}
	for e := 0; e < len(errArr); e++ {
		if errArr[e] != nil {
			return errArr[e]
		}
	}
	return nil
}

func convertContentDate(dateStr string, UTCoffset string) (time.Time, error) {
	var timeObj time.Time
	var errTime error
	timeObj = time.Now()
	layout := "01-02-2006 15:04:05 -07"
	dateStr = dateStr + " " + UTCoffset
	timeObj, errTime = time.Parse(layout, dateStr)
	if errTime != nil {
		return timeObj, errTime
	}
	return timeObj, nil
}

func getChannelCounts(channelBounds string, channelRows []string) []float64 {
	var counts []float64
	var i, j, n, nrows, ncols int
	var row, channelVal string
	var rowPts []string
	var errCts error

	channelBoundPts := strings.Split(channelBounds, " ")
	maxCh, errCh := strconv.ParseInt(channelBoundPts[len(channelBoundPts)-1], 10, 64)
	if errCh == nil {
		counts = make([]float64, (maxCh + 1))
	}
	n = 0
	nrows = len(channelRows)
	for i = 0; i < nrows; i++ {
		row = strings.Replace(channelRows[i], "\r", "", -1)
		rowPts = strings.Split(row, " ")
		ncols = len(rowPts)
		for j = 0; j < ncols; j++ {
			channelVal = strings.Replace(rowPts[j], " ", "", -1)
			if channelVal != "" {
				counts[n], errCts = strconv.ParseFloat(rowPts[j], 64)
				if errCts != nil {
					counts[n] = 0
				}
				n++
			}
		}
	}
	return counts
}
