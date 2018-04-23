package processAvaatechSpe

import (
	spereader "readAvaatechSpe"
)

func Process(spePath string, config *Configuration) (*Spectrum, error) {
	spect := new(Spectrum)
	spe, err := spereader.ReadSPE(spePath, true)
	if err != nil {
		return spect, err
	}
	spect.SPE = spe
	spect.Config = config
	spect.ModelPeaks()
	spect.AssignPeaks()
	spect.CalculateEnergyScale()
	return spect, nil
}
