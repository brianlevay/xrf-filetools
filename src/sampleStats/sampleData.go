package sampleStats

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"
)

type SampleStats struct {
	SourcePath string                    `json:"-"`
	Error      string                    `json:"Error"`
	DataMap    map[string]map[string]int `json:"-"`
	Conditions map[string]int            `json:"-"`
	Stats      []map[string]string       `json:"Stats"`
	Headers    []string                  `json:"Headers"`
}

func (samples *SampleStats) Initialize(sourcePath string) {
	samples.SourcePath = sourcePath
	samples.Error = "none"
	samples.DataMap = make(map[string]map[string]int)
	samples.Conditions = make(map[string]int)
	samples.Stats = nil
	samples.Headers = nil
}

func (samples *SampleStats) GenerateStatsArr() {
	var nameKey string
	sortedConds := SortedConditions(samples.Conditions)
	sortedSamples := SortedSamples(samples.DataMap)
	samples.Headers = GenerateHeaders(sortedConds)
	for _, arr := range sortedSamples {
		rowMap := make(map[string]string)
		nameKey = arr[0] + "/" + arr[1]
		excMap, _ := samples.DataMap[nameKey]
		rowMap["Folder"] = arr[0]
		rowMap["Sample"] = arr[1]
		for _, cond := range sortedConds {
			rowMap[cond] = strconv.Itoa(excMap[cond])
		}
		samples.Stats = append(samples.Stats, rowMap)
	}
}

func (samples *SampleStats) JSON() []byte {
	var JSONbytes []byte
	var err error
	JSONbytes, err = json.Marshal(samples)
	if err != nil {
		return []byte("{\"Error\":\"Unable to marshal JSON\"}")
	}
	return JSONbytes
}

func SortedConditions(m map[string]int) []string {
	var sortedArr []string
	for key := range m {
		sortedArr = append(sortedArr, key)
	}
	sort.Strings(sortedArr)
	return sortedArr
}

func SortedSamples(m map[string]map[string]int) [][]string {
	var sortedArr [][]string
	for key := range m {
		keyPts := strings.Split(key, "/")
		sortedArr = append(sortedArr, keyPts)
	}
	sort.Slice(sortedArr, func(i, j int) bool {
		folderLess := sortedArr[i][0] < sortedArr[j][0]
		nameLess := sortedArr[i][1] < sortedArr[j][1]
		return folderLess && nameLess
	})
	return sortedArr
}

func GenerateHeaders(sortedConds []string) []string {
	headers := make([]string, (len(sortedConds) + 2))
	headers[0] = "Folder"
	headers[1] = "Sample"
	for i, cond := range sortedConds {
		headers[i+2] = cond
	}
	return headers
}
