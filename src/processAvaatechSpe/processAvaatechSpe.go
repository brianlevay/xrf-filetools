package processAvaatechSpe

import (
	conf "configureSpe"
	"math"
	spereader "readAvaatechSpe"
	"snip"
)

func Process(spePath string, config *conf.Configuration) (*Spectrum, error) {
	spect := new(Spectrum)
	spe, err := spereader.ReadSPE(spePath, config.UTCoffset, true)
	if err != nil {
		return spect, err
	}
	spect.SPE = spe
	spect.MaxChannel = int(math.Min(float64(len(spe.Counts)-1), float64(config.MaxChannel)))
	spect.Signal = snip.RemoveBackground(spe.Counts[0:spect.MaxChannel+1], config.SNIPwidth)
	spect.ModelPeaks(config.MinPeakHeight)
	spect.AssignPrimaryLines(config.GainMinKeV, config.GainMaxKeV)
	spect.CalculateEnergyScale(config.GainMidKeV)
	spect.AssignSecondaryLines()
	return spect, nil
}
