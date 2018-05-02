package processAvaatechSpe

import (
	conf "configureSpe"
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
	spect.AssignPrimaryLines(config.GainMinKeV, config.GainMaxKeV)
	spect.CalculateEnergyScale(config.GainMidKeV)
	spect.AssignSecondaryLines()
	return spect, nil
}
