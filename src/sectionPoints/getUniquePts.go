package sectionPoints

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Highest level function //
func (points *UniquePoints) RecursiveSearch() {
	err := filepath.Walk(points.SourcePath, addSample(points))
	if err != nil {
		log.Fatal(err)
	}
	points.FlattenToArr()
}

// Called on each file during recursiveSearch, uses closure to bind pointer to struct //
func addSample(points *UniquePoints) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Print(err)
			return nil
		}
		if (info.IsDir() == false) && (strings.Contains(info.Name(), ".spe") == true) {
			getPoint(info.Name(), path, points)
		}
		return nil
	}
}

// Called on each file inside WalkFunc closure //
func getPoint(fileName string, path string, points *UniquePoints) {
	var point Point
	if strings.Contains(fileName, "!") == true {
		point = pointFromNew(fileName, path)
	} else {
		point = pointFromOld(fileName, path)
	}
	points.DataMap[point]++
}

func pointFromNew(fileName string, path string) Point {
	cleanName := strings.Replace(fileName, ",", ".", -1)
	filePts := strings.Split(cleanName, "!")

	point := new(Point)
	point.Sample = filePts[0]
	point.X = filePts[9]
	point.Y = filePts[10]
	point.DC = filePts[12]
	point.CC = filePts[13]
	return *point
}

func pointFromOld(fileName string, path string) Point {
	var xFound bool
	cleanName := strings.Replace(fileName, ",", ".", -1)
	filePts := strings.Split(cleanName, " ")
	nPts := len(filePts)

	point := new(Point)
	point.Sample = filePts[0]
	point.DC = "10.0"
	point.CC = "12.0"
	for i := 1; i < nPts; i++ {
		if strings.Contains(filePts[i], "mm") == true {
			rootStr := strings.Replace(filePts[i], "mm", "", -1)
			if strings.Contains(rootStr, "DC") == true {
				point.DC = strings.Replace(rootStr, "DC", "", -1)
			} else if strings.Contains(rootStr, "CC") == true {
				point.CC = strings.Replace(rootStr, "CC", "", -1)
			} else if xFound == true {
				point.Y = rootStr
			} else {
				point.X = rootStr
				xFound = true
			}
		}
	}
	return *point
}
