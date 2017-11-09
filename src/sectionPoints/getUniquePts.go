package sectionPoints

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
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
	var err error
	if strings.Contains(fileName, "!") == true {
		point, err = pointFromNew(fileName, path)
	} else {
		point, err = pointFromOld(fileName, path)
	}
	if err != nil {
		return
	}
	points.DataMap[point]++
}

func pointFromNew(fileName string, path string) (Point, error) {
	var errs [4]error
	cleanName := strings.Replace(fileName, ",", ".", -1)
	filePts := strings.Split(cleanName, "!")

	point := new(Point)
	point.X, errs[0] = strconv.ParseFloat(filePts[9], 64)
	point.Y, errs[1] = strconv.ParseFloat(filePts[10], 64)
	point.DC, errs[2] = strconv.ParseFloat(filePts[12], 64)
	point.CC, errs[3] = strconv.ParseFloat(filePts[13], 64)
	for i := 0; i < 4; i++ {
		if errs[i] != nil {
			return *point, errs[i]
		}
	}
	return *point, nil
}

func pointFromOld(fileName string, path string) (Point, error) {
	var errs [4]error
	var xFound bool
	cleanName := strings.Replace(fileName, ",", ".", -1)
	filePts := strings.Split(cleanName, " ")
	nPts := len(filePts)

	point := new(Point)
	point.DC = 10
	point.CC = 12
	for i := 1; i < nPts; i++ {
		if strings.Contains(filePts[i], "mm") == true {
			rootStr := strings.Replace(filePts[i], "mm", "", -1)
			if strings.Contains(rootStr, "DC") == true {
				point.DC, errs[0] = strconv.ParseFloat(strings.Replace(rootStr, "DC", "", -1), 64)
			} else if strings.Contains(rootStr, "CC") == true {
				point.CC, errs[1] = strconv.ParseFloat(strings.Replace(rootStr, "CC", "", -1), 64)
			} else if xFound == true {
				point.Y, errs[2] = strconv.ParseFloat(rootStr, 64)
			} else {
				point.X, errs[3] = strconv.ParseFloat(rootStr, 64)
				xFound = true
			}
		}
	}
	for i := 0; i < 4; i++ {
		if errs[i] != nil {
			return *point, errs[i]
		}
	}
	return *point, nil
}
