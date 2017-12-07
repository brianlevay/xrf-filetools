package timeUsage

import (
	"encoding/json"
	"sort"
	"strconv"
	"time"
)

type TimeStats struct {
	SourcePath string                `json:"-"`
	Error      string                `json:"Error"`
	DataMap    map[string]DaySummary `json:"-"`
	Stats      []DayDisplay          `json:"Stats"`
}

type DaySummary struct {
	Start   time.Time `json:"-"`
	Finish  time.Time `json:"-"`
	Seconds int64     `json:"-"`
	Points  int64     `json:"-"`
}

type DayDisplay struct {
	Day     string `json:"Day"`
	Start   string `json:"Start"`
	Finish  string `json:"Finish"`
	Elapsed string `json:"Elapsed"`
	Runtime string `json:"Runtime"`
	Points  string `json:"Points"`
}

func (timeuse *TimeStats) Initialize(sourcePath string) {
	timeuse.SourcePath = sourcePath
	timeuse.Error = "none"
	timeuse.DataMap = make(map[string]DaySummary)
}

func (timeuse *TimeStats) FlattenToArr() {
	sortedDays := SortedDays(timeuse.DataMap)
	timeuse.Stats = sortedDays
}

func (timeuse *TimeStats) JSON() []byte {
	var JSONbytes []byte
	var err error
	JSONbytes, err = json.Marshal(timeuse)
	if err != nil {
		return []byte("{\"Error\":\"Unable to marshal JSON\"}")
	}
	return JSONbytes
}

func SortedDays(m map[string]DaySummary) []DayDisplay {
	var sortedRaw []DaySummary
	for _, val := range m {
		sortedRaw = append(sortedRaw, val)
	}
	sort.Slice(sortedRaw, func(i, j int) bool {
		if sortedRaw[i].Start.Before(sortedRaw[j].Start) {
			return true
		}
		return false
	})

	var sortedDays []DayDisplay
	var year, month, day string
	var startHr, startMin, finishHr, finishMin string
	var elapsedHrs, runHrs string
	nRows := len(sortedRaw)
	for i := 0; i < nRows; i++ {
		year = strconv.Itoa(sortedRaw[i].Start.Year())
		month = sortedRaw[i].Start.Month().String()
		day = strconv.Itoa(sortedRaw[i].Start.Day())
		startHr = intToStringPadded(sortedRaw[i].Start.Hour())
		startMin = intToStringPadded(sortedRaw[i].Start.Minute())
		finishHr = intToStringPadded(sortedRaw[i].Finish.Hour())
		finishMin = intToStringPadded(sortedRaw[i].Finish.Minute())
		elapsedHrs = strconv.FormatFloat(sortedRaw[i].Finish.Sub(sortedRaw[i].Start).Hours(), 'f', 2, 64)
		runHrs = strconv.FormatFloat(float64(sortedRaw[i].Seconds)/(60*60), 'f', 2, 64)

		dayDisplay := DayDisplay{}
		dayDisplay.Day = month + " " + day + ", " + year
		dayDisplay.Start = startHr + ":" + startMin
		dayDisplay.Finish = finishHr + ":" + finishMin
		dayDisplay.Elapsed = elapsedHrs
		dayDisplay.Runtime = runHrs
		dayDisplay.Points = strconv.FormatInt(sortedRaw[i].Points, 10)

		sortedDays = append(sortedDays, dayDisplay)
	}
	return sortedDays
}

func intToStringPadded(value int) string {
	if value < 10 {
		return ("0" + strconv.Itoa(value))
	} else {
		return strconv.Itoa(value)
	}
}
