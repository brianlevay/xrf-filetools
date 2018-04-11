package avaatechSpe

import (
	"time"
)

type SPE struct {
	Path    string
	Name    string
	Date    string
	Voltage float64
	Filter  string
	Current float64
	Live    uint64
	DC      float64
	CC      float64
	X       float64
	Y       float64
	CPS     uint64
	Counts  []uint64
}
