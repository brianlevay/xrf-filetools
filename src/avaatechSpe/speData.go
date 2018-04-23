package avaatechSpe

import (
	"time"
)

const defaultChannelN int = 2048

type SPE struct {
	FilePath string    `json:"-"`
	FileName string    `json:"-"`
	Opened   bool      `json:"-"`
	Folder   string    `json:"Folder"`
	Sample   string    `json:"Sample"`
	Date     time.Time `json:"Date"`
	Voltage  float64   `json:"Voltage"`
	Filter   string    `json:"Filter"`
	Current  float64   `json:"Current"`
	Live     int64     `json:"Live"`
	DC       float64   `json:"DC"`
	CC       float64   `json:"CC"`
	X        float64   `json:"X"`
	Y        float64   `json:"Y"`
	CPS      int64     `json:"CPS"`
	Counts   []float64 `json:"-"`
}

func (spe *SPE) Initialize() {
	spe.FilePath = "n/a"
	spe.FileName = "n/a"
	spe.Opened = false
	spe.Folder = "n/a"
	spe.Sample = "n/a"
	spe.Date = time.Date(2000, time.January, 01, 01, 0, 0, 0, time.UTC)
	spe.Voltage = 0.0
	spe.Filter = "n/a"
	spe.Current = 0.0
	spe.Live = 0
	spe.DC = 0.0
	spe.CC = 0.0
	spe.X = 0.0
	spe.Y = 0.0
	spe.CPS = 0
	spe.Counts = make([]float64, defaultChannelN)
}
