package sectionPoints

import (
	"encoding/json"
	"sort"
)

type Point struct {
	X  float64
	Y  float64
	DC float64
	CC float64
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
	points.Headers = []string{"X", "Y", "DC", "CC"}
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
		if sortedPts[i].X < sortedPts[j].X {
			return true
		}
		if sortedPts[i].X > sortedPts[j].X {
			return false
		}
		return sortedPts[i].Y < sortedPts[j].Y
	})
	return sortedPts
}
