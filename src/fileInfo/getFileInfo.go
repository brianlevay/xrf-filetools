package fileInfo

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"strconv"
)

func ReadContents(path string, info os.FileInfo) *FileContents {
	fileBytes, errRead := ioutil.ReadFile(path)
	if errRead != nil {
		log.Print(errRead)
		return nil
	}
	fileString := string(fileBytes)
	contents := new(FileContents)
	contents.Path = path
	parseFile(info.Name(), fileString, contents)
	return contents
}

func parseFile(fileName string, fileContents string, contents *FileContents) {
	var cps, voltage, current, dc, cc, nextRow string
	var measureTimes []string
	var convErr error
	contents.Name = getName(fileName)
	contents.Filter = getFilter(fileName)
	fileRows := strings.Split(fileContents, "\n")
	nRows := len(fileRows)
	for i := 0; i < nRows-1; i++ {
		nextRow = strings.Replace(fileRows[i+1], "\r", "", -1)
		nextRow = strings.Replace(nextRow, ",", ".", -1)
		nextRow = strings.Trim(nextRow)
		if strings.Contains(fileRows[i], "$X_Position:") == true {
			contents.X, convErr = strconv.ParseFloat(nextRow,64)
		} else if strings.Contains(fileRows[i], "$Y_Position:") == true {
			contents.Y, convErr = strconv.ParseFloat(nextRow,64)
		} else if strings.Contains(fileRows[i], "$Slit_DC:") == true {
			contents.DC, convErr = strconv.ParseFloat(nextRow,64)
		} else if strings.Contains(fileRows[i], "$Slit_CC:") == true {
			contents.CC, convErr = strconv.ParseFloat(nextRow,64)
		} else if strings.Contains(fileRows[i], "$TotalCPS:") == true {
			contents.CPS, convErr = strconv.ParseUint(nextRow,10,64)
		} else if strings.Contains(fileRows[i], "$ACC_VOLT:") == true {
			contents.Voltage, convErr = strconv.ParseFloat(nextRow,64)
		} else if strings.Contains(fileRows[i], "$TUBE_CUR:") == true {
			contents.Current, convErr = strconv.ParseFloat(nextRow,64)
		} else if strings.Contains(fileRows[i], "$MEAS_TIM:") == true {
			measureTimes = strings.Split(nextRow, " ")
			contents.Live, convErr = strconv.ParseUint(measureTimes[0],64)
		} else if strings.Contains(fileRows[i], "$DATE_MEA:") == true {
			contents.Date = nextRow
		} else if strings.Contains(fileRows[i], "$DATA:") == true {
			contents.Counts = getChannelCounts(fileRows, (i+2))
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

func getFilter(fileName string) string {
	return "n/a"
}

func getChannelCounts(fileRows []string, start int) {
	
}