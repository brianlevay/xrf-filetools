package batchProcessSpe

import (
	conf "configureSpe"
	speprocess "processAvaatechSpe"
)

type Batch struct {
	Root     string                 `json:-`
	Config   *conf.Configuration    `json:-`
	Existing map[string]bool        `json:-`
	List     []*speprocess.Spectrum `json:List`
}
