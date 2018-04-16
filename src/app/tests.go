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
	spect := spectrum.GetSpectrum(spe)
	fmt.Println(spect.SPE.CPS)
	for _, peak := range spect.Peaks {
		fmt.Println(peak)
	}
	return
}
