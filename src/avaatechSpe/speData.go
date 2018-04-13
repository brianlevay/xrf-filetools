package avaatechSpe

import (
	"fmt"
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
	spe.Opened = false
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
