package timeUsage

import (
	"time"
)

type TimeStats struct {
	SourcePath string                `json:"-"`
	Error      string                `json:"Error"`
	DataMap    map[string]DaySummary `json:"-"`
}

type DaySummary struct {
	Start   time.Time      `json:"-"`
	Finish  time.Time      `json:"-"`
	PtsSect map[string]int `json:"-"`
}

func (times *TimeStats) Initialize(sourcePath string) {
	times.SourcePath = sourcePath
	times.Error = "none"
	times.DataMap = make(map[string]DaySummary)
}
