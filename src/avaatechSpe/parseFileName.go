package avaatechSpe

import (
	"strconv"
	"strings"
)

func (spe *SPE) ParseFileName() error {
	errFolder := spe.GetFolder()
	if errFolder != nil {
		return errFolder
	}
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

func (spe *SPE) GetFolder() error {

}

func (spe *SPE) ParseNameNew() error {
	cleanName := strings.Replace(spe.FileName, ",", ".", -1)
	namePts := strings.Split(cleanName, "!")
	spe.Sample = namePts[0]
	spe.Date, _ = convertNameDate(namePts[4], namePts[5])
	spe.Live, _ = strconv.ParseUint(namePts[6], 10, 64)
	spe.Voltage, _ = strconv.ParseFloat(namePts[7], 64)
	spe.Current, _ = strconv.ParseFloat(namePts[8], 64)
	spe.X, _ = strconv.ParseFloat(namePts[9], 64)
	spe.Y, _ = strconv.ParseFloat(namePts[10], 64)
	spe.Filter = namePts[11]
	spe.DC, _ = strconv.ParseFloat(namePts[12], 64)
	spe.CC, _ = strconv.ParseFloat(namePts[13], 64)
	return nil
}

func (spe *SPE) ParseNameOld() error {
	cleanName := strings.Replace(spe.FileName, ",", ".", -1)
	namePts := strings.Split(cleanName, " ")
	nPts := len(namePts)
	spe.Sample = namePts[0]
	return nil
}

func convertNameDate(dayStr string, timeStr string) (time.Time, error) {
	dateStr := dayStr + " " + timeStr
	layout := "2006-01-02 15-04-05"
	loc, _ := time.LoadLocation("Local")
	timeObj, err := time.ParseInLocation(layout, dateStr, loc)
	if err != nil {
		return nil, err
	}
	return timeObj, nil
}
