package main

import (
	avaatech "avaatechSpe"
	"fmt"
	spectrum "speSpectrum"
)

func testReader() {
	path := `./_misc/testData/standards/Standard_1 X  50.0mm   6s   9kV 250uA No-Filter DC10.0mm CC12.0mm Y  0.0mm Run1 Rep0 Sett Low.spe`
	name := `Standard_1 X  50.0mm   6s   9kV 250uA No-Filter DC10.0mm CC12.0mm Y  0.0mm Run1 Rep0 Sett Low.spe`
	spe, err := avaatech.ReadSPE(path, name, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	spe.Print(80)

	config := &spectrum.Configuration{Threshold: 1000.0, GainMinKeV: 0.02000, GainMaxKeV: 0.02050}
	spect := spectrum.Process(spe, config)
	fmt.Println("CPS:", spect.SPE.CPS)
	for _, peak := range spect.Peaks {
		fmt.Println("Channel:", peak.Channel, " Area:", peak.Area, " FWHM:", peak.FWHM, "ChiSq:", peak.ChiSq)
	}
	for line, peakAssigned := range spect.Lines {
		fmt.Println("Line:", line, " Channel:", peakAssigned.Channel)
	}
	fmt.Println("Gain:", spect.Gain, " Offset:", spect.Offset, " R2:", spect.R2)
	return
}
