package sampleSummaries

import ()

type SampleSummaries struct {
	SourcePath string          `json:"-"`
	Error      string          `json:"Error"`
	Data       map[string]map[string]int `json:"Data"`
}

func (samples *SampleSummaries) Initialize(sourcePath string) {
	samples.SourcePath = sourcePath
	samples.Error = "none"
	samples.Data = make(map[string]map[string]int)
}
