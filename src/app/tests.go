package main

import (
	"fmt"
	speprocesser "processAvaatechSpe"
)

func testPackages() {
	spePath := `./_misc/testData/standards/Standard_1 X  50.0mm   6s   9kV 250uA No-Filter DC10.0mm CC12.0mm Y  0.0mm Run1 Rep0 Sett Low.spe`
	config := &speprocesser.Configuration{Threshold: 1000.0, GainMinKeV: 0.02000, GainMaxKeV: 0.02050}
	spect, err := speprocesser.Process(spePath, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\nPROCESS SPECTRUM TEST\n")
	fmt.Println("File:", spect.SPE.FileName)
	fmt.Println("Sample:", spect.SPE.Sample)
	for _, peak := range spect.Peaks {
		fmt.Println("Channel:", peak.Channel, " Area:", peak.Area, " FWHM:", peak.FWHM, "ChiSq:", peak.ChiSq)
	}
	for line, peakAssigned := range spect.Lines {
		fmt.Println("Line:", line, " Channel:", peakAssigned.Channel)
	}
	fmt.Println("Gain:", spect.Gain, " Offset:", spect.Offset, " R2:", spect.R2)
	return
}
