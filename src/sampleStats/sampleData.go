package sampleStats

import (
	"sort"
)

type SampleStats struct {
	SourcePath string          `json:"-"`
	Error      string          `json:"Error"`
	DataMap       map[string]map[string]int `json:"Data"`
	Conditions map[string]int `json:Conditions"`
}

func (samples *SampleStats) Initialize(sourcePath string) {
	samples.SourcePath = sourcePath
	samples.Error = "none"
	samples.DataMap = make(map[string]map[string]int)
	samples.Conditions = make(map[string]int)
}

func (samples *SampleStats) SortedConditions() []string {
	var conds []string
	for k := range samples.Conditions {
		conds = append(conds, k)
	}
	sort.Strings(conds)
	return conds
}