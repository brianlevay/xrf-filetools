package uniqueNames

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Highest level function //
func (unique *UniqueNames) RecursiveSearch() {
	err := filepath.Walk(unique.SourcePath, readSampleName(unique))
	if err != nil {
		log.Fatal(err)
	}
}

// Called on each file during recursiveSearch, uses closure to bind pointer to struct //
func readSampleName(unique *UniqueNames) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			getUniqueName(info.Name(), path, unique)
		}
		return nil
	}
}

// Called on each file inside WalkFunc closure //
func getUniqueName(fileName string, path string, unique *UniqueNames) {
	var Name, Folder, Combined string
	var namePts, pathPts []string
	var pathLength int

	if strings.Contains(fileName, "!") == true {
		namePts = strings.Split(fileName, "!")
	} else {
		namePts = strings.Split(fileName, " ")
	}
	Name = namePts[0]

	pathPts = filepath.SplitList(path)
	pathLength = len(pathPts)
	switch pathLength {
	case 0:
		{
			Folder = "Root"
		}
	case 1:
		{
			Folder = "Root"
		}
	case 2:
		{
			if strings.Contains(pathPts[0], "kV") {
				Folder = "Root"
			} else {
				Folder = pathPts[0]
			}
		}
	default:
		{
			if strings.Contains(pathPts[pathLength-2], "kV") {
				Folder = pathPts[pathLength-3]
			} else {
				Folder = pathPts[pathLength-2]
			}
		}
	}
	Combined = Folder + "," + Name
	_, ok := unique.Data[Combined]
	if ok == false {
		unique.Data[Combined] = true
	}
}
