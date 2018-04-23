package speSpectrum

// Tests with bAxil on spectra from both instruments indicate that offset is
// always << 20 eV, which is < 1 channel width. Thus, it can be
// approximated as 0 eV for all fitting in this module.
// Gain varies from ~20.0 eV/ch to ~20.4 eV/ch, which leads to channel differences of
// ~2 at Si Ka and ~7 out at Fe Kb. Thus, gain needs to be a variable.

const gain_keV float64 = 0.0200
const offset_keV float64 = 0.0000
const gain_delta float64 = 0.0050

type LineID struct {
	Name   string
	Energy float64
}

// Source for Energies: Kaye and Laby, Table of Physical and Chemical Constants //
var line_keV = []LineID{
	LineID{Name: "Al_Ka", Energy: 1.487},
	LineID{Name: "Si_Ka", Energy: 1.740},
	LineID{Name: "Rh_La", Energy: 2.698},
	LineID{Name: "K_Ka", Energy: 3.314},
	LineID{Name: "Ca_Ka", Energy: 3.692},
	LineID{Name: "Ca_Kb", Energy: 4.013},
	LineID{Name: "Fe_Ka", Energy: 6.404},
	LineID{Name: "Fe_Kb", Energy: 7.058},
}

func channelFromEnergy(energy float64, gain float64, offset float64) float64 {
	channel := (energy - offset) / gain
	return channel
}

func energyFromChannel(channel float64, gain float64, offset float64) float64 {
	energy := (gain * channel) + offset
	return energy
}
