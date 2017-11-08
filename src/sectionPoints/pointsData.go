package sectionPoints

import (
	"encoding/json"
	"sort"
	"strconv"
)

type Point struct {
	Sample string
	X      string
	Y      string
	DC     string
	CC     string
}

type UniquePoints struct {
	SourcePath string        `json:"-"`
	Error      string        `json:"Error"`
	DataMap    map[Point]int `json:"-"`
	Points     []Point       `json:"Points"`
	Headers    []string      `json:"Headers"`
}

func (points *UniquePoints) Initialize(sourcePath string) {
	points.SourcePath = sourcePath
	points.Error = "none"
	points.DataMap = make(map[Point]int)
	points.Points = nil
	points.Headers = []string{"Sample", "X", "Y", "DC", "CC"}
}

func (points *UniquePoints) FlattenToArr() {
	sortedPoints := SortedPtsXY(points.DataMap)
	points.Points = sortedPoints
}

func (points *UniquePoints) JSON() []byte {
	var JSONbytes []byte
	var err error
	JSONbytes, err = json.Marshal(points)
	if err != nil {
		return []byte("{\"Error\":\"Unable to marshal JSON\"}")
	}
	return JSONbytes
}

func SortedPtsXY(m map[Point]int) []Point {
	var sortedPts []Point
	for key := range m {
		sortedPts = append(sortedPts, key)
	}
	sort.Slice(sortedPts, func(i, j int) bool {
		Xi, _ := strconv.Atoi(sortedPts[i].X)
		Xj, _ := strconv.Atoi(sortedPts[j].X)
		Yi, _ := strconv.Atoi(sortedPts[i].Y)
		Yj, _ := strconv.Atoi(sortedPts[j].Y)
		if Xi < Xj {
			return true
		}
		if Xi > Xj {
			return false
		}
		return Yi < Yj
	})
	return sortedPts
}
