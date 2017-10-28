package sampleSummaries

import (
	"log"
	"os"
	"path/filepath"
)

// Highest level function //
func (samples *SampleSummaries) WriteToCSV(outPath string, outName string) error {
	f, errC := os.Create(filepath.Join(outPath, outName))
	if errC != nil {
		log.Println(errC)
		return errC
	}
	defer f.Close()

	_, errW := f.WriteString("FolderName,FileName\n")
	if errW != nil {
		log.Println(errW)
		return errW
	}
	for key, _ := range samples.Data {
		_, errW = f.WriteString(key + "\n")
		if errW != nil {
			log.Println(errW)
			return errW
		}
	}
	f.Sync()

	return nil
}
