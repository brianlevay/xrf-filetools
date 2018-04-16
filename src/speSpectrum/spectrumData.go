package speSpectrum

import (
    avaatech "avaatechSpe"
    "time"
)

type Spectrum struct {
	Sample  string
	Date    time.Time
	Voltage float64
	Filter  string
	Current float64
	Live    uint64
	DC      float64
	CC      float64
	X       float64
	Y       float64
	CPS     uint64
}

func (spect *Spectrum) MetaData(spe *avaatech.SPE) {
	spect.Sample = spe.Sample
	spect.Date = spe.Date
	spect.Voltage = spe.Voltage
	spect.Filter = spe.Filter
	spect.Current = spe.Current
	spect.Live = spe.Live
	spect.DC = spe.DC
	spect.CC = spe.CC
	spect.X = spe.X
	spect.Y = spe.Y
	spect.CPS = spe.CPS
}