package sampleStats

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Highest level function //
func (samples *SampleStats) WriteToCSV(outPath string, outName string) error {
	f, errC := os.Create(filepath.Join(outPath, outName))
	if errC != nil {
		log.Println(errC)
		return errC
	}
	defer f.Close()

	_, errW := f.WriteString(strings.Join(samples.Headers, ",") + "\n")
	if errW != nil {
		log.Println(errW)
		return errW
	}
	for _, m := range samples.Stats {
		for j, key := range samples.Headers {
			_, errW = f.WriteString(m[key])
			if errW != nil {
				log.Println(errW)
				return errW
			}
			if j < (len(samples.Headers) - 1) {
				f.WriteString(",")
			}
		}
		f.WriteString("\n")
	}
	f.Sync()

	return nil
}
