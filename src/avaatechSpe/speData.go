package avaatechSpe

import (
	"fmt"
	"time"
)

const defaultChannelN int = 2048

type SPE struct {
	FilePath string
	FileName string
	Opened   bool
	Folder   string    // From FilePath
	Sample   string    // From FileName
	Date     time.Time // From FileName(New), FileContents
	Voltage  float64   // From FileName, FileContents
	Filter   string    // From FileName
	Current  float64   // From FileName, FileContents [in mA]
	Live     uint64    // From FileName, FileContents
	DC       float64   // From FileName, FileContents
	CC       float64   // From FileName, FileContents
	X        float64   // From FileName, FileContents
	Y        float64   // From FileName, FileContents
	CPS      uint64    // From FileContents
	Counts   []uint64  // From FileContents
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
	spe.Counts = make([]uint64, defaultChannelN)
}

func (spe *SPE) Print(index int) {
	fmt.Println("\nFilePath: ", spe.FilePath)
	fmt.Println("FileName: ", spe.FileName)
	fmt.Println("Opened: ", spe.Opened)
	fmt.Println("Folder: ", spe.Folder)
	fmt.Println("Sample: ", spe.Sample)
	fmt.Println("Date: ", spe.Date)
	fmt.Println("Voltage: ", spe.Voltage)
	fmt.Println("Filter: ", spe.Filter)
	fmt.Println("Current: ", spe.Current)
	fmt.Println("Live: ", spe.Live)
	fmt.Println("DC: ", spe.DC)
	fmt.Println("CC: ", spe.CC)
	fmt.Println("X: ", spe.X)
	fmt.Println("Y: ", spe.Y)
	fmt.Println("CPS :", spe.CPS)
	fmt.Println("Counts[", index, "]: ", spe.Counts[index], "\n")
}
