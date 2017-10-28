package sampleStats

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"strconv"
)

// Highest level function //
func (samples *SampleStats) WriteToCSV(outPath string, outName string) error {
	f, errC := os.Create(filepath.Join(outPath, outName))
	if errC != nil {
		log.Println(errC)
		return errC
	}
	defer f.Close()

	conditions := samples.SortedConditions()
	_, errW := f.WriteString("FolderName,FileName," + strings.Join(conditions, ",") + "\n")
	if errW != nil {
		log.Println(errW)
		return errW
	}
	for key, m := range samples.DataMap {
		keyPts := strings.Split(key, "/")
		condStr := keyValsToString(m, conditions)
		_, errW = f.WriteString(keyPts[0] + "," + keyPts[1] + "," + condStr + "\n")
		if errW != nil {
			log.Println(errW)
			return errW
		}
	}
	f.Sync()

	return nil
}

func keyValsToString(m map[string]int, orderArray []string) string {
	var valArr []string
	var valStr string
	for _, item := range orderArray {
		valArr = append(valArr, strconv.Itoa(m[item]))
	}
	valStr = strings.Join(valArr, ",")
	return valStr
}
