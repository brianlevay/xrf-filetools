package speThroughput

import (
	"encoding/json"
	"sync"
	"time"
)

type Throughput struct {
	Mtx        sync.Mutex `json:"-"`
	LastDate   time.Time  `json:"-"`
	NextDate   time.Time  `json:"-"`
	SourcePath string     `json:"-"`
	Error      string     `json:"Error"`
	Data       []string   `json:"Data"`
}

type Row struct {
	Name string `json:"Name"`
	X    string `json:"X"`
	Date string `json:"Date"`
	CPS  string `json:"CPS"`
	KVp  string `json:"KVp"`
	Curr string `json:"Curr"`
	DC   string `json:"DC"`
	CC   string `json:"CC"`
}

func (through *Throughput) Initialize(sourcePath string) {
	defaultTime := time.Date(2000, time.January, 01, 01, 0, 0, 0, time.UTC)
	through.LastDate = defaultTime
	through.NextDate = defaultTime
	through.SourcePath = sourcePath
	through.Error = "none"
	through.Data = nil
}

func (through *Throughput) JSON() []byte {
	var JSONbytes []byte
	var err error
	JSONbytes, err = json.Marshal(through)
	if err != nil {
		return []byte("{\"Error\":\"Unable to marshal JSON\"}")
	}
	return JSONbytes
}
