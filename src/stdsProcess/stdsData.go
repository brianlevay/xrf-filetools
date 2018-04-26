package stdsProcess

import (
	conf "configureSpe"
	speprocess "processAvaatechSpe"
	"sync"
)

type Standards struct {
	Mtx      sync.Mutex             `json:"-"`
	Config   *conf.Configuration    `json:"-"`
	Existing map[string]bool        `json:-`
	List     []*speprocess.Spectrum `json:"List"`
	Error    error                  `json:"-"`
}
