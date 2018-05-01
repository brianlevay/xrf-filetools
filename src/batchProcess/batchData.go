package batchProcess

import (
	conf "configureSpe"
	speprocess "processAvaatechSpe"
	"sync"
)

type BatchSpectra struct {
	Mtx      sync.Mutex             `json:"-"`
	RootPath string                 `json:"-"`
	Config   *conf.Configuration    `json:"-"`
	Existing map[string]bool        `json:-`
	List     []*speprocess.Spectrum `json:"List"`
	Error    error                  `json:"-"`
}
