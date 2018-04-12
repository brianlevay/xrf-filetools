package avaatechSpe

import (
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func (spe *SPE) ParseFileContents() error {
	var cps, voltage, current, dc, cc, nextRow string
	var measureTimes []string
	fileBytes, errRead := ioutil.ReadFile(spe.Path)
	if errRead != nil {
		return errRead
	}
	fileString := string(fileBytes)
	fileRows := strings.Split(fileContents, "\n")
	nRows := len(fileRows)
	for i := 0; i < nRows-1; i++ {
		nextRow = strings.Replace(fileRows[i+1], "\r", "", -1)
		nextRow = strings.Replace(nextRow, ",", ".", -1)
		nextRow = strings.Trim(nextRow)
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
			spe.Live, _ = strconv.ParseUint(measureTimes[0], 64)
		} else if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			spe.Date, _ = convertContentDate(nextRow)
		} else if strings.Contains(fileRows[i], "$DATA:") == true {
			spe.Counts = getChannelCounts(fileRows, (i + 2))
		}
	}
	return nil
}

func convertContentDate(dateStr string) (time.Time, error) {
	layout := "01-02-2006 15:04:05"
	loc, _ := time.LoadLocation("Local")
	timeObj, err := time.ParseInLocation(layout, dateStr, loc)
	if err != nil {
		return nil, err
	}
	return timeObj, nil
}

func getChannelCounts(fileRows []string, start int) []uint64 {

}
