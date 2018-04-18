package speSpectrum

// Tests with bAxil on spectra from both instruments indicate that offset is
// always << 20 eV, which is < 1 channel width. Thus, it can be
// approximated as 0 eV for all fitting in this module.
// Gain varies from ~20.0 eV/ch to ~20.4 eV/ch, which leads to channel differences of
// ~2 at Si Ka and ~7 out at Fe Kb. Thus, gain needs to be a variable.

const gain_keV float64 = 0.0200
const offset_keV float64 = 0.0000

// Source for Energies: Kaye and Laby, Table of Physical and Chemical Constants //
var lineEnergies = map[string]float64{
	"Al_Ka": 1.487,
	"Si_Ka": 1.740,
	"Rh_La": 2.698,
	"K_Ka":  3.314,
	"Ca_Ka": 3.692,
	"Ca_Kb": 4.013,
	"Fe_Ka": 6.404,
	"Fe_Kb": 7.058,
}

type Elements struct {
	Al_Ka *Peak
	Si_Ka *Peak
	Rh_La *Peak
	K_Ka  *Peak
	Ca_Ka *Peak
	Ca_Kb *Peak
	Fe_Ka *Peak
	Fe_Kb *Peak
}

func channelFromEnergy(energy float64, gain float64, offset float64) float64 {
	channel := (energy - offset) / gain
	return channel
}

func energyFromChannel(channel float64, gain float64, offset float64) float64 {
	energy := (gain * channel) + offset
	return energy
}
