package avaatechSpe

import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadSPE(path string, info os.FileInfo) *SPE {
	spe := new(SPE)
	spe.Path = path
	spe.Name = getName(info.Name())
	spe.Folder = getFolder(path)
	spe.Filter = getFilter(info.Name())
	spe.ParseFile()
	return spe
}

func (spe *SPE) ParseFile() {
	var cps, voltage, current, dc, cc, nextRow string
	var measureTimes []string
	var convErr error
	fileBytes, errRead := ioutil.ReadFile(spe.Path)
	if errRead != nil {
		log.Print(errRead)
		return nil
	}
	fileString := string(fileBytes)
	fileRows := strings.Split(fileContents, "\n")
	nRows := len(fileRows)
	for i := 0; i < nRows-1; i++ {
		nextRow = strings.Replace(fileRows[i+1], "\r", "", -1)
		nextRow = strings.Replace(nextRow, ",", ".", -1)
		nextRow = strings.Trim(nextRow)
		if strings.Contains(fileRows[i], "$X_Position:") == true {
			spe.X, convErr = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Y_Position:") == true {
			spe.Y, convErr = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Slit_DC:") == true {
			spe.DC, convErr = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$Slit_CC:") == true {
			spe.CC, convErr = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$TotalCPS:") == true {
			spe.CPS, convErr = strconv.ParseUint(nextRow, 10, 64)
		} else if strings.Contains(fileRows[i], "$ACC_VOLT:") == true {
			spe.Voltage, convErr = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$TUBE_CUR:") == true {
			spe.Current, convErr = strconv.ParseFloat(nextRow, 64)
		} else if strings.Contains(fileRows[i], "$MEAS_TIM:") == true {
			measureTimes = strings.Split(nextRow, " ")
			spe.Live, convErr = strconv.ParseUint(measureTimes[0], 64)
		} else if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			spe.Date = nextRow
		} else if strings.Contains(fileRows[i], "$DATA:") == true {
			spe.Counts = getChannelCounts(fileRows, (i + 2))
		}
	}
}

func getName(fileName string) string {
	var name string
	var namePts []string
	if strings.Contains(fileName, "!") == true {
		namePts = strings.Split(fileName, "!")
	} else {
		namePts = strings.Split(fileName, " ")
	}
	name = namePts[0]
	return name
}

func getFolder(path string) string {

}

func getFilter(fileName string) string {
	return "n/a"
}

func getChannelCounts(fileRows []string, start int) []uint64 {

}
