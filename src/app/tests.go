package main

import (
	avaatech "avaatechSpe"
	"fmt"
	spectrum "speSpectrum"
)

func testReader() {
	spePath := `./_misc/testData/standards/Standard_1 X  50.0mm   6s   9kV 250uA No-Filter DC10.0mm CC12.0mm Y  0.0mm Run1 Rep0 Sett Low.spe`
	spe, err := avaatech.ReadSPE(spePath, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("LOAD SPE FILE TEST\n")
	fmt.Println("File:", spe.FileName)
	fmt.Println("Sample:", spe.Sample)
	fmt.Println("Voltage:", spe.Voltage, " Filter:", spe.Filter, " Current:", spe.Current)
	fmt.Println("CPS:", spe.CPS)

	config := &spectrum.Configuration{Threshold: 1000.0, GainMinKeV: 0.02000, GainMaxKeV: 0.02050}
	spect := spectrum.Process(spe, config)
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
