package avaatechSpe

import (
	"time"
)

type SPE struct {
	FilePath string
	FileName string
	Opened   bool
	Folder   string    // From FilePath
	Sample   string    // From FileName
	Date     time.Time // From FileName(New), FileContents
	Voltage  float64   // From FileName, FileContents
	Filter   string    // From FileName
	Current  float64   // From FileName, FileContents
	Live     uint64    // From FileName, FileContents
	DC       float64   // From FileName, FileContents
	CC       float64   // From FileName, FileContents
	X        float64   // From FileName, FileContents
	Y        float64   // From FileName, FileContents
	CPS      uint64    // From FileContents
	Counts   []uint64  // From FileContents
}

func (spe *SPE) Initialize() {
	spe.Opened = false
}
