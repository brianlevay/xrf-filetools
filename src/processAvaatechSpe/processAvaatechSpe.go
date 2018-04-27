package processAvaatechSpe

import (
	conf "configureSpe"
	"fmt"
	spereader "readAvaatechSpe"
)

func Process(spePath string, config *conf.Configuration) (*Spectrum, error) {
	spect := new(Spectrum)
	spe, err := spereader.ReadSPE(spePath, config.UTCoffset, true)
	if err != nil {
		return spect, err
	}
	spect.SPE = spe
	spect.ModelPeaks(config.Threshold)
	spect.AssignPeaks(config.GainMinKeV, config.GainMaxKeV)
	_, ok := spect.Lines["Si_Ka"] ////// FOR DEBUGGING ONLY
	if ok == false {              ////// FOR DEBUGGING ONLY
		fmt.Println(spect.SPE.FilePath)    ////// FOR DEBUGGING ONLY
		for _, peak := range spect.Peaks { ////// FOR DEBUGGING ONLY
			fmt.Println(peak.Channel) ////// FOR DEBUGGING ONLY
		} ////// FOR DEBUGGING ONLY
	} //////
	return spect, nil
}
