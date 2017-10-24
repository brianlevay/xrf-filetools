package speSummaries

import (
	"strings"
	"time"
)

type Summary struct {
	LastDate   time.Time
	NextDate   time.Time
	SourcePath string
	Data       []string
}

func NewSummary(sourcePath string) *Summary {
	summ := new(Summary)
	defaultTime := time.Date(2000, time.January, 01, 01, 0, 0, 0, time.UTC)
	summ.LastDate = defaultTime
	summ.NextDate = defaultTime
	summ.SourcePath = sourcePath
	return summ
}

func (summ *Summary) JSON() string {
	JSONstr := "{" + strings.Join(summ.Data, ",") + "}"
	return JSONstr
}

func (summ *Summary) Reset(sourcePath string) {
	defaultTime := time.Date(2000, time.January, 01, 01, 0, 0, 0, time.UTC)
	summ.LastDate = defaultTime
	summ.NextDate = defaultTime
	summ.SourcePath = sourcePath
	summ.Data = nil
}
