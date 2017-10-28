package sampleStats

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Highest level function //
func (samples *SampleStats) RecursiveSearch() {
	err := filepath.Walk(samples.SourcePath, addSample(samples))
	if err != nil {
		log.Fatal(err)
	}
	samples.GenerateStatsArr()
}

// Called on each file during recursiveSearch, uses closure to bind pointer to struct //
func addSample(samples *SampleStats) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			getDetails(info.Name(), path, samples)
		}
		return nil
	}
}

// Called on each file inside WalkFunc closure //
func getDetails(fileName string, path string, samples *SampleStats) {
	var uniqueName, excitation string

	if strings.Contains(fileName, "!") == true {
		uniqueName, excitation = newInfo(fileName, path)
	} else {
		uniqueName, excitation = oldInfo(fileName, path)
	}
	excMap, ok := samples.DataMap[uniqueName]
	if ok == false {
		excMap = make(map[string]int)
		samples.DataMap[uniqueName] = excMap
	}
	excMap[excitation]++
	samples.Conditions[excitation]++
}

func newInfo(fileName string, path string) (string, string) {
	var cleanName string
	var sample, folder, voltage, filter string
	var uniqueName, excitation string
	folder = getFolderName(path)
	cleanName = strings.Replace(fileName, ",", ".", -1)
	filePts := strings.Split(cleanName, "!")
	sample = filePts[0]
	voltage = strings.Replace(filePts[7], " ", "", -1) + "kV"
	voltage = leftPadVoltage(voltage)
	filter = filePts[11]
	uniqueName = folder + "/" + sample
	excitation = voltage + "_" + filter
	return uniqueName, excitation
}

func oldInfo(fileName string, path string) (string, string) {
	var cleanName string
	var sample, folder, voltage, filter string
	var uniqueName, excitation string
	folder = getFolderName(path)
	cleanName = strings.Replace(fileName, ",", ".", -1)
	filePts := strings.Split(cleanName, " ")
	sample = filePts[0]
	for i, _ := range filePts {
		if strings.Contains(filePts[i], "kV") == true {
			voltage = strings.Replace(filePts[i], " ", "", -1)
			voltage = leftPadVoltage(voltage)
		}
		if strings.Contains(filePts[i], "uA") == true {
			filter = filePts[i+1]
		}
	}
	uniqueName = folder + "/" + sample
	excitation = voltage + "_" + filter
	return uniqueName, excitation
}

func leftPadVoltage(voltage string) string {
	var padded string
	if len(voltage) == 3 {
		padded = "0" + voltage
	} else {
		padded = voltage
	}
	return padded
}

func getFolderName(path string) string {
	var folder string
	var pathPts []string
	var pathLength int
	pathPts = strings.Split(path, string(os.PathSeparator))
	pathLength = len(pathPts)
	switch pathLength {
	case 0:
		{
			folder = "Root"
		}
	case 1:
		{
			folder = "Root"
		}
	case 2:
		{
			if strings.Contains(pathPts[pathLength-2], "Run") {
				folder = "Root"
			} else {
				folder = pathPts[pathLength-2]
			}
		}
	default:
		{
			if strings.Contains(pathPts[pathLength-2], "Run") {
				folder = pathPts[pathLength-3]
			} else {
				folder = pathPts[pathLength-2]
			}
		}
	}
	return folder
}
